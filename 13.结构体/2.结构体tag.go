package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age,omitempty"` //忽略空值,不希望出现零值问题
	Password string `json:"-"`             //敏感信息,转json的时候忽略该字段
}

func main() {
	user := User{
		Name:     "dkw",
		Age:      21,
		Password: "123456",
	}
	byteData, _ := json.Marshal(user)
	fmt.Println(string(byteData))

	//fmt.Println(user)

}
