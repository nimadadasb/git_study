package main

import (
	"fmt"
	"io"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		var byteData = make([]byte, 1024)
		n, err := conn.Read(byteData)
		if err == io.EOF {
			break
		}
		fmt.Println(string(byteData[:n]))
	}
}
