package main

import "fmt"

type UserInfo struct {
	Name string `json:"name"`
}

// SetName 用指针接收，才可以修改name
func (u *UserInfo) SetName(name string) {
	u.Name = name
}

func main() {
	user := UserInfo{
		Name: "dkw",
	}
	user.SetName("dkw1")
	fmt.Println(user.Name)

}
