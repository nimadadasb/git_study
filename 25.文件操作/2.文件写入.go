package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//根路径，可以直接运行生成文件，不需要在终端再跑一次
	//OpenFile("文件路径","文件权限","文件权限位")
	// 常见权限组合
	//0644   用户可读写，组和其他人只读
	//0755   用户可读写执行，组和其他人可读执行
	//0600   仅用户可读写
	//0777   所有人可读写执行（危险！）
	//一次性写
	file, err := os.OpenFile("go1/25.文件操作/w.txt", os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//file.Write([]byte("你好!"))
	//byteData, err := io.ReadAll(file)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("写入成功", string(byteData))

	/**
	os.WriteFile是最简单的文件写入方法：
	一次调用完成所有操作
	自动创建文件（如果需要）
	会覆盖已有文件
	适合小到中等大小的文件
	*/
	os.WriteFile("go1/25.文件操作/w1.txt", []byte("你好!"), 0666)
	if err != nil {
		panic(err)
		return

	}

	fmt.Println("写入成功")

	//复制文件
	//读取的文件
	rFile, err := os.Open("C:\\Users\\Admin\\Desktop\\杂物\\th.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rFile.Close()

	//写入的文件
	wFile, err := os.OpenFile("go1/25.文件操作/sheep.jpg", os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("复制成功")
	defer wFile.Close()
	io.Copy(wFile, rFile)

	//目录结构
	dir, err := os.ReadDir("go1/25.文件操作")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range dir {
		info, _ := file.Info()
		fmt.Println(file.Name(), file.IsDir(), info.Size())

	}
}
