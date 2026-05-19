package main

import (
	"fmt"
	"sync"
)

var sum1 int

var wait1 sync.WaitGroup
var lock sync.Mutex

func add1() {
	lock.Lock() //上锁
	for i := 0; i < 100000; i++ {
		sum1++
	}
	lock.Unlock() //解锁
	wait1.Done()
}

func sub1() {
	lock.Lock()
	for i := 0; i < 100000; i++ {
		sum1--
	}
	lock.Unlock()
	/**
	Lock()
	拿不到锁就阻塞等待
	TryLock()
	拿不到锁立刻返回 false
	*/
	//lock.TryLock() //非阻塞地尝试获取锁
	wait1.Done()
}

// 线程安全了
func main() {
	wait1.Add(2)
	go add1()
	go sub1()
	wait1.Wait()
	if lock.TryLock() {
		fmt.Println("success")
	} else {
		fmt.Println("failed")
	}

	fmt.Println(sum1)
}
