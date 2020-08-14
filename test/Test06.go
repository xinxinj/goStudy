package main

import (
	"add" //导入 add 包
	"fmt"
)

func main() {
	c := add.Add(1, 2) //调用 add 包中的 add 函数
	fmt.Println(c)
}
