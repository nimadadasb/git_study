package main

import (
	"fmt"
	"sync"
)

var sum int

var wait sync.WaitGroup

func add() {
	for i := 0; i < 100000; i++ {
		sum++
	}
	wait.Done()
}

func sub() {
	for i := 0; i < 100000; i++ {
		sum--
	}
	wait.Done()
}

// 并发问题
func main() {
	wait.Add(2)
	go add()
	go sub()
	wait.Wait()

	fmt.Println(sum)
}
