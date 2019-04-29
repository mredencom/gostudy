package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

//pool管理一组可以安全的在多个协程间共享资源
//实现io.Closer
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

//ErrPoolClosed标识请求关闭了一个pool
var ErrPoolClosed = errors.New("资源池已经关闭")

// New 创建一个用来管理资源的池。
// 这个池需要一个可以分配新资源的函数，
// 并规定池的大小
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("资源池太小")
	}
	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

//从池子中取得一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire:", "使用资源")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	//没有资源使用时,提供一个新的资源使用, 这样不会导致阻塞
	default:
		log.Println("Acquire:", "使用新资源")
		return p.factory()
	}
}

//将一个使用后的资源放回资源回池中
func (p *Pool) Release(r io.Closer) {
	//使用互斥锁
	p.m.Lock()
	//使用解锁
	defer p.m.Unlock()
	//判断池是否被关闭, 销毁这个资源
	if p.closed {
		r.Close()
		return
	}
	select {
	case p.resources <- r:
		log.Println("Release:", "进入队列")
	default:
		log.Println("Release:", "关闭")
		_ = r.Close()
	}
}

//close 会让资源池停止工作, 并让关闭现有资源
func (p *Pool) Close() {
	//使用互斥锁
	p.m.Lock()
	//使用解锁
	defer p.m.Unlock()
	if p.closed {
		return
	}
	//关闭
	p.closed = true
	//在清空通道你资源前,将通道关闭,不然发生死锁
	close(p.resources)
	for r := range p.resources {
		r.Close()
	}
}
