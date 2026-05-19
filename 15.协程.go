package main

import (
	"fmt"
	"sync"
	"time"
)

func shopping(name string, wait *sync.WaitGroup) {
	fmt.Printf("%s start shopping\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s finish shopping\n", name)
	wait.Done()
}

// 协程
func main() {
	//使用局部变量而非全局变量，避免并发问题
	var wait sync.WaitGroup
	startTime := time.Now()
	//现在的模式，就是购物接力
	//shopping("dkw1")
	//shopping("dkw2")
	//shopping("dkw3")
	wait.Add(3)

	//主线程结束，协程函数跟着结束
	go shopping("dkw1", &wait)
	go shopping("dkw2", &wait)
	go shopping("dkw3", &wait)

	wait.Wait()

	fmt.Println("shopping cost:", time.Since(startTime))

}
