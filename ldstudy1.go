package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sort"
	"strconv"
)

type RandNum struct {
	num [10]int
}

func main() {
	var num RandNum
	fmt.Print("随机数：")
	fmt.Println(num.RandomValues())
	fmt.Print("最大值：")
	fmt.Println(num.Max())
	fmt.Print("最小值：")
	fmt.Println(num.Min())
	fmt.Print("平均数：")
	fmt.Println(num.Average())
	if num.CheckNum() {
		fmt.Println("There's 55 in the array")
	} else {
		fmt.Println("There is no 55 in the array")
	}
}

func (r *RandNum) RandomValues () interface{}{
	max := big.NewInt(100)
	for i := 0; i < 10; i++ {
		result, _ := rand.Int(rand.Reader, max)
		re := result.String()
		r.num[i], _ = strconv.Atoi(re)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(r.num[:])))
	return r.num
}

func (r *RandNum) Max() int {
	return r.num[0]
}

func (r *RandNum) Min() int {
	return r.num[9]
}

func (r *RandNum) Average() int {
	var n int
	for _, v := range r.num{
		n += v
	}
	return n/len(r.num)
}

func (r RandNum) CheckNum() bool  {
	n := 0
	for _, v := range r.num{
		if v == 55 {
			n = 1
			break
		}
	}
	if n == 0 {
		return false
	}
	return  true
}