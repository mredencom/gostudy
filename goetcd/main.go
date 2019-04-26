package main

import (
	"fmt"
	"gostudy/goetcd/etcd"
	"time"
)

func main() {
	etcdR, err := etcd.CreateCli([]string{"127.0.0.1:2379"}, 2)

	defer etcdR.Close()

	if err != nil {
		fmt.Println(err)
	}
	//塞入kv
	ptr := etcd.PutValue(etcdR, "abc", "123")
	//查看kv是否存入成功
	fmt.Println("存放是否成功", ptr)
	//取kv值
	gtr := etcd.GetValue(etcdR, "abc")
	//打印kv
	fmt.Println("查询结果", gtr)
	//删除kv
	gi := etcd.DelValue(etcdR, "abc")
	//查看删除结果
	fmt.Println("删除kv结果:", gi)
	//删除kv后在取
	gtr = etcd.GetValue(etcdR, "abc")
	//查询结果
	fmt.Println("查询结果", gtr)

	ptr = etcd.PutValue(etcdR, "abcdef", "1234")

	gtbP := etcd.GetValueByPrefix(etcdR, "abcd")

	fmt.Println("查询结果", gtbP)

	dt := etcd.DelValueByPrefix(etcdR, "abc")

	fmt.Println("删除是否成功", dt)

	gtr = etcd.GetValue(etcdR, "abcdef")

	fmt.Println("查询结果", gtr)

	etcd.PutValue(etcdR, "abc", "123")

	ofDeal := etcd.OfficeDeal(etcdR, "abc", "123", "456")

	fmt.Println("事务是否成功", ofDeal)

	gtr = etcd.GetValue(etcdR, "abc")

	fmt.Println("查询结果", gtr)

	ofDeal = etcd.OfficeDeal(etcdR, "abc", "456", "345")

	fmt.Println("事务是否成功", ofDeal)

	gtr = etcd.GetValue(etcdR, "abc")

	fmt.Println("查询结果", gtr)

	var cha chan bool
	var chan1 chan bool

	go func() {
		m := etcd.WatchKey(etcdR, "abc")
		fmt.Println("监看键值", string(m))
		cha <- true
	}()

	go func() {
		i := etcd.WatchKeyWithPrefix(etcdR, "a")
		fmt.Println("监看带有前缀的键", string(i))
		chan1 <- true
	}()

	etcd.PutValue(etcdR, "abc", "1234")

	etcd.PutValue(etcdR, "abc", "121")

	cySetKv := etcd.CycleSetKeyValueWithTime(etcdR, 10, "haha", "balbal")

	fmt.Println("周期是否成功", cySetKv)

	gv := etcd.GetValue(etcdR, "haha")

	fmt.Println("周期结束之前", gv)

	time.Sleep(11 * time.Second)

	gv = etcd.GetValue(etcdR, "haha")

	fmt.Println("周期结束之后", gv)
	<-cha
	<-chan1
}
