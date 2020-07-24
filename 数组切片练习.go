package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 倒序排序 -- 冒泡
func sortBubbling(arr *[10]int) ([10]int, int, int) {

	var temp, max, min int

	for i := 0; i < len(arr); i++ {
		for j := 0; j < (len(arr) - 1 - i); j++ {

			if (*arr)[j] < (*arr)[j+1] {
				temp = (*arr)[j+1]
				(*arr)[j+1] = (*arr)[j]
				(*arr)[j] = temp

				max = j + 1
			}
		}
	}

	return *arr, max, min

}

// 练习题1
func test1() (arr [10]int, arrSort [10]int, max int, min int, isFind bool) {

	// 将时间戳设置成种子数 每次随机数都重新生成
	rand.Seed(time.Now().UnixNano())

	// 生成随机数
	var v int
	for i := 0; i < len(arr); i++ {
		v = rand.Intn(100) + 1
		arr[i] = v

		if v == 55 {
			isFind = true
		}
	}

	// 数组倒序排序
	arrSort, max, min = sortBubbling(&arr)

	return

}

// 练习题2
func insert(arr []int, v int) []int {
	var k int

	for i := 0; i < len(arr); i++ {
		if arr[i] > v {
			k = i + 1
		}
	}

	arr = append(arr, 0)
	copy(arr[(k+1):], arr[k:])
	arr[k] = v

	return arr
}

func main() {

	// 练习题1：随机生成10个整数(1-100)保存到数组，倒序打印、求平均值、最大值及最小值下标，查找里面是否包含55
	arr, arrSort, max, min, isFind := test1()

	fmt.Println(arr, arrSort, max, min, isFind)

	// 练习题2：已知已经有个降序数组，现在需要插入一个元素，然后打印该数组 顺序依然是降序

	var v int
	fmt.Scanln(&v)

	res := insert(arr[:], v)

	fmt.Println(res)

}
