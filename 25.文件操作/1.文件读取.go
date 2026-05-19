package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//go1/25.文件操作/hello.txt 根路径

	file, err := os.Open("D:\\Go项目\\go1\\25.文件操作\\hello.txt") //绝对路径
	if err != nil {
		panic(err)

	}

	//关闭文件
	defer file.Close()

	//按字节读
	//for {
	//	//windows有2个换行符\r\n,所以长度是12+2=14
	//	var byteData = make([]byte, 14)
	//	n, err := file.Read(byteData)
	//	//读完了
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Println(string(byteData), n)
	//}

	//按行读
	//buf := bufio.NewReader(file)
	//
	//for {
	//	line, _, err := buf.ReadLine()
	//
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Println(string(line), err)
	//}

	//指定分割符
	buf := bufio.NewScanner(file)
	//按照单词ScanWordds
	buf.Split(bufio.ScanWords)
	var index int
	for buf.Scan() {
		index++
		fmt.Println(index, buf.Text())

	}

}
