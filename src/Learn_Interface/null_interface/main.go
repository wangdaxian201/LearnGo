package main

import "fmt"

type null interface {
}

type person struct {
	name string
	age  int
}

func say(p person) {
	fmt.Println(p.name)
}

func main() {
	// 空接口可以接收任意类型
	var nu null

	nu = 2
	nu = "dddd"
	nu = 3.44
	a := person{"xiaowang", 18}
	say(a)
	fmt.Printf("%v", nu)
}
