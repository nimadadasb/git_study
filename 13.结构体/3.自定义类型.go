package main

import "fmt"

type Code int

const (
	SuccessCode      Code = 0
	ServiceErrorCode Code = 1001 //服务错误
	NetworkErrorCode Code = 1002 //网络错误

)

func (c Code) GetMsg() string {
	switch c {
	case SuccessCode:
		return "success"
	case ServiceErrorCode:
		return "service error"
	case NetworkErrorCode:
		return "network error"
	}
	return "unknown code"
}

func (c Code) GetAll() (code Code, msg string) {
	fmt.Printf("GetAll return -> code=%d, msg=%s\n", c, c.GetMsg())
	return c, c.GetMsg()
}

//func getCodeMessage(code Code) (msg string) {
//	switch code {
//	case SuccessCode:
//		return "success"
//	case ServiceErrorCode:
//		return "service error"
//	case NetworkErrorCode:
//		return "network error"
//
//	}
//	return "unknown error"
//}

func webServer(name string) (code Code, msg string) {
	if name == "1" {
		return SuccessCode.GetAll()
	}
	//Go 更推荐：提前 return，少用 else
	if name == "2" {
		return NetworkErrorCode.GetAll()
	}
	return ServiceErrorCode.GetAll()

}

func main() {
	webServer("1")

}
