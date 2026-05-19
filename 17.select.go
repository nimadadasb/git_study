package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan1 = make(chan int)    //声明初始化一个长度为0的信道
var nameChan1 = make(chan string)  // 声明初始化一个长度为0的信道
var doneChan = make(chan struct{}) //声明一个用于关闭select的信道

// 必须用指针，不能传值
func send(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s start shopping\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s finish shopping\n", name)

	//把money这个值发送到moneyChan这个信道
	/**
	这是一个 阻塞操作
	如果：
	没有人正在接收 → 当前 goroutine(协程) 会卡在这里
	直到：
	有另一个 goroutine 执行 <-moneyChan
	在你程序中：
	发送发生在 send的 goroutine 里
	接收发生在 main的select里
	*/
	moneyChan1 <- money
	nameChan1 <- name

	wait.Done()
}

// 协程
func main() {
	//Add、Done、Wait三者一定是配套使用的，
	//它们共同构成了 sync.WaitGroup的完整生命周期。
	//Add(3)写在启动 goroutine 之前
	//Done()写在每个具体的 goroutine 内部
	//Wait()写在单独的 goroutine 里，也就是专门负责关闭 channel 的 goroutine

	//使用局部变量而非全局变量，避免并发问题
	//声明一个WaitGroup，其内部维护一个计时器，初始值为0
	var wait sync.WaitGroup
	startTime := time.Now()
	//设置WaitGroup的等待数量，现在需要等3个任务
	wait.Add(3)

	//主线程结束，协程函数跟着结束
	go send("dkw1", 2, &wait)
	go send("dkw2", 3, &wait)
	go send("dkw3", 5, &wait)

	//启动一个“专门负责关闭 channel 的 goroutine”
	go func() {

		defer close(moneyChan1)
		defer close(nameChan1)
		defer close(doneChan)
		//阻塞当前 goroutine(即go func()这个负责关闭channel的goroutine)，直到wait.Done() 被调用 3 次，也就是等所有 pay goroutine 都干完活
		wait.Wait()
		//关闭channel，之后不能再发送，但可以接受剩余数据，这是 range channel 能退出的关键
		//close(moneyChan)

	}()

	//准备接收 channel 中的数据
	//声明一个切片
	//用来保存从 channel 中收到的数据
	var moneyList []int
	var nameList []string

	//select处理多个channel，类似于switch
	var event = func() {
		//for循环无限监听，因为channel不会立刻关闭，这是一个事件循环模型
		for {
			select {
			//select里不能用range
			case money := <-moneyChan1:
				moneyList = append(moneyList, money)
			case name := <-nameChan1:
				nameList = append(nameList, name)
			case <-doneChan:
				return
			}
		}
	}
	event()
	fmt.Println("结束")
	fmt.Println("shopping cost:", time.Since(startTime))
	fmt.Println("moneyList", moneyList)
	fmt.Println("nameList", nameList)

	//用 range 从 channel 消费数据（核心）
	/**
	range moneyChan
	含义：
	从 moneyChan中 不断接收数据
	每次接收：
	把值赋给 money
	进入循环体
	*/
	//for money := range moneyChan1 {
	//	fmt.Println(money)
	//	moneyList = append(moneyList, money)
	//}

	//下面的写法等同于上面的
	//for {
	//	money, ok := <-moneyChan
	//	fmt.Println(money, ok)
	//	if !ok {
	//		break
	//	}
	//}

}
