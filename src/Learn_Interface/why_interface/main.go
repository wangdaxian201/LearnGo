package main

import "fmt"

type dog struct{}
type cat struct{}
type person struct {
	name string
	age  int
}

// sayer 接口
type sayer interface {
	say()
}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

func (p person) say() {
	fmt.Println("嘿嘿嘿")
}

func main() {
	var s sayer
	s = dog{}
	s.say()

	s = cat{}
	s.say()

	s = person{"dd", 18}
	s.say()

}
