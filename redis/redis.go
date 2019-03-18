package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/panjf2000/ants"
	"strconv"
	"sync"
	"time"
)

func client() {
	//链接redis
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "root",
		DB:       1,
	})
	//i := 1
	//pong, err := client.Ping().Result()
	//fmt.Println(pong, err)
	for i := 1; i <= 10000; i++ {
		str := strconv.Itoa(i)
		err := client.Set("key"+str, "value"+str, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}
func main() {

	t1 := time.Now().Unix()
	runTimes := 10000

	// 使用普通池
	var wg sync.WaitGroup
	//使用协程池
	p, _ := ants.NewPoolWithFunc(8, func(i interface{}) {
		client()
		wg.Done()
	})
	//最后销毁协程池
	defer p.Release()
	//循环提交任务
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	t2 := time.Now().Unix()
	fmt.Printf("运行协程: %d\n", p.Running())
	fmt.Printf("结果")
	fmt.Printf("使用时间 %d\n", t2-t1)
}
