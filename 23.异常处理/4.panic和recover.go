package main

import (
	"fmt"
	"log"
	"runtime/debug"
)

// 业务层函数 - 可能发生 panic
func processData() {
	var list []int = []int{1, 2}

	// 故意制造 panic
	fmt.Println(list[2]) // 数组越界
}

// 安全执行器 - 处理 panic 并记录
func safeExecute(fn func()) (panicErr error) {
	defer func() {
		if r := recover(); r != nil {
			// 1. 记录 panic 信息
			panicErr = fmt.Errorf("panic recovered: %v", r)

			// 2. 记录堆栈跟踪
			stack := debug.Stack()
			log.Printf("PANIC DETAILS:\nError: %v\nStack Trace:\n%s", r, stack)

			// 3. 可以选择发送到监控系统
			// sendToMonitoring(r, stack)
		}
	}()

	fn() // 执行可能 panic 的函数
	return nil
}

// 业务包装函数
func read1() error {
	return safeExecute(processData)
}

func main() {
	// 执行可能 panic 的操作
	if err := read1(); err != nil {
		log.Printf("操作失败: %v", err)

		// 根据错误类型决定后续逻辑
		// 可以继续执行，或返回，或重试
	}

	// 正常逻辑继续执行
	fmt.Println("hello - 程序继续正常运行")
}
