package main

import "fmt"

func main() {

	fmt.Println("测试1")

	var a bool
	a = true
	var b = (2 == 3) //b也会被推导为bool类型

	fmt.Println(a, b)

	//错误示范
	//var b bool
	//b = 1 //编译错误
	//b = bool(1) //编译错误
}
