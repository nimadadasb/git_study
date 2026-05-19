package main

import (
	"fmt"
	"os"
)

func init() {
	_, err := os.ReadFile("12345")
	if err != nil {
		//log.Fatal(err)

		panic("错误") //用于初始化

	}

}

func main() {
	fmt.Println("main")

}
