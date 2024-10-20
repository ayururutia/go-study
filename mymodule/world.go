package mymodule

import "fmt"

// 模块初始化函数import 包时被调用
func init() {
	fmt.Println("world module init function")
}

func World() string {
	return "world"
}
