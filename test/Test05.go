package main

import "fmt"

func main() {

	fmt.Println("流程控制包含（if,switch,for,goto,select）,其中select用于监听channel（通道）")

	//if   其中 optionalStatement 是可选的表达式，真正决定分支走向的是 booleanExpression1 的值

	//if optionalStatement1; booleanExpression1 {
	//	block1
	//} else if optionalStatement2; booleanExpression2 {
	//	block2
	//} else {
	//	block3
	//}

	//for 语句  for 语句可以遍历数组，切片，映射等类型，也可以用于无限循环。
	//for { // 无限循环
	//	block
	//}
	//
	//for booleanExpression { // while循环，在Go语言中没有while关键字
	//
	//}
	//
	//for index, char := range aString { // 迭代字符串
	//
	//}
	//
	//for item := range aChannel { // 迭代通道
	//
	//}

	//goto跳转语句
	myfunc()

	//switch  switch 分支既可用于常用的分支就像 C 语言中的 switch 一样，也可以用于类型开关，所谓类型开关就是用于判断变量属于什么类型
	//switch optionalStatement; optionalExpression {
	//case expression1: block1
	//	...
	//case expressionN: blockN
	//default: blockD
	//}

	//switch {        // 没有表达式，默认为True值，匹配分支中值为True的分支
	//case value < minimum:
	//	return minimum
	//case value > maximum:
	//	return maximum
	//default:
	//	return value
	//}
	//switch练习
	switchfunc(5, -17.98, "AIDEN", nil, true, complex(1, 1))

}

//创建一个函数，该函数可以接受任意多的任意类型的参数
func switchfunc(items ...interface{}) {

	for i, x := range items {
		switch x := x.(type) { // 创建了影子变量
		case bool:
			fmt.Printf("param #%d is a bool, value: %t\n", i, x)
		case float64:
			fmt.Printf("param #%d is a float64, value: %f\n", i, x)
		case int, int8, int16, int32, int64:
			fmt.Printf("param #%d is a int, value: %d\n", i, x)
		case uint, uint8, uint16, uint32, uint64:
			fmt.Printf("param #%d is a uint, value: %d\n", i, x)
		case nil:
			fmt.Printf("param #%d is a nil\n", i)
		case string:
			fmt.Printf("param #%d is a string, value: %s\n", i, x)
		default:
			fmt.Printf("param #%d's type is unknow\n", i)
		}
	}

}

//跳转语句 使用 goto 关键字实现跳转。goto 语句的语义非常简单，就是跳转到本函数内的某个标签
func myfunc() {
	i := 0
THIS: //定义一个THIS标签
	fmt.Println(i)
	i++
	if i < 1 {
		goto THIS //跳转到THIS标签
	}
}
