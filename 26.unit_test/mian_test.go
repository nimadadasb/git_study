package main

import (
	"fmt"
	"os"
	"testing"
)

func setup() {
	fmt.Println("setup", "测试前")
}

func teardown() {
	fmt.Println("teardown", "测试后")
}

func TestAdd2(t *testing.T) {
	fmt.Println("TestAdd2", "测试中")
	t.Errorf("测试不通过")
}

// 必须是这个名字，测试主入口
func TestMain(m *testing.M) {
	fmt.Println("TestMain")
	//测试前执行
	setup()
	code := m.Run()
	//测试后执行，不管测试通没通过
	teardown()
	os.Exit(code)
}
