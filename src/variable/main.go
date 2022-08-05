package main

import (
	"fmt"
)

/*
1. 函数外的每个语句都必须以关键字开始 （var, const, func...）
2. := 类型推导不可以使用在函数外
3. _ 多用于占位, 忽略
*/
func main() {
	// 标准声明
	//var a int
	//var b string
	//var c bool
	//var d float32

	//fmt.Println(a, b, c, d)

	fmt.Println("=============")
	// 批量声明
	var (
		a int
		b string
		c bool
		d float32
	)

	a = 10
	b = "haha"
	c = false
	d = 2.3
	fmt.Println(a, b, c, d)

}
