package main

import "fmt"

// 练习切片 数组 方法调用、 输入 输出 字符串 数值等

// 菲波那切数列
func fbs(n int) []uint64 {

	fbsclice := make([]uint64, n)

	fbsclice[0] = 1
	fbsclice[1] = 1

	for i := 2; i < n; i++ {
		fbsclice[i] = fbsclice[i-1] + fbsclice[i-2]
	}

	return fbsclice
}

// 冒泡排序
func bubbling(arr *[5]uint) [5]uint {

	var temp uint

	for i := 0; i < len(*arr); i++ {
		for j := i; j < len(*arr)-1-i; j++ {

			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}

	return *arr
}

func main() {

	// 斐波那契数列
	// fbslice := fbs(10)
	// fmt.Println(fbslice)

	// 冒泡排期
	// arr := [5]uint{3, 2, 9, 19, 16}
	// nl := bubbling(&arr)

	// fmt.Println(nl)

	// 顺序查找
	var hum [4]string
	hum = [4]string{"白色", "黑色", "蓝色", "紫色"}

	var humname string

	for {
		fmt.Println("请输入您要查找的字符串")

		fmt.Scanln(&humname)

		for i := 0; i < len(hum); i++ {

			if hum[i] == humname {
				fmt.Println("存在")
				break
			} else if i == (len(hum) - 1) {
				fmt.Println("不存在")
			}
		}
	}

}
