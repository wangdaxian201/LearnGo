package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name  string `json:"name" ini:"s_name"`
	Age   int    `json:"age" ini:"s_age"`
	Email string `json:"email" ini:"email"`
}

func (p person) Run() {
	fmt.Printf("%v 正在跑", p.Name)
}

func (p person) Sleep() {
	fmt.Printf("%v 正在睡觉", p.Name)
}

func PrintMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(v.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodObj := v.Method(i)
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("%s", methodObj.Type())

		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args []reflect.Value
		v.Method(i).Call(args)
	}
}

func main() {
	// 结构体反射
	stu1 := person{
		"小王",
		19,
		"123@111",
	}
	//
	//t := reflect.TypeOf(stu1)
	//fmt.Printf("name: %v kind: %v", t.Name(), t.Kind())
	//
	//fmt.Println("name:---", t.NumField())
	//// NumField()  返回结构体成员字段数量。
	//for i := 0; i < t.NumField(); i++ {
	//	// Field 根据索引，返回索引对应的结构体字段的信息
	//	fileObj := t.Field(i)
	//	fmt.Printf("name:%v \n type: %v \n tag: %v \n", fileObj.Name, fileObj.Type, fileObj.Tag)
	//	// obj.Tag.Get()  // 用于获取tag中的字段信息
	//	fmt.Printf("\n json:%v \n", fileObj.Tag.Get("json"))
	//	fmt.Printf("ini:%v \n", fileObj.Tag.Get("ini"))
	//}

	PrintMethod(stu1)

}
