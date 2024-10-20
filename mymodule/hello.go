package mymodule

import "fmt"

// 模块初始化函数 import 包时被调用
func init() {
	fmt.Println("hello module init function")
}

func Hello() string {
	return "hello"
}
