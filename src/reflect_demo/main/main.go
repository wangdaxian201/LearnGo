package main

import (
	"fmt"
	"reflect"
)

func reflectType(t interface{}) {
	v := reflect.TypeOf(t)
	fmt.Println(v, v.Name(), v.Kind())
	fmt.Printf("type: %v\n", v)
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Elem().Kind()

	switch k {
	case reflect.Float32:
		v.Elem().SetFloat(1.23333)
	case reflect.Int32:
		v.Elem().SetInt(1000)

	}
}

func main() {
	// 反射
	var a float32 = 3.3333
	reflectType(a)

	var b string = "aaaaa"
	reflectType(b)

	var c bool
	c = false
	reflectType(c)

	var d float32 = 1.111
	reflectSetValue(&d)
	fmt.Println("d:", d)

	var g int32 = 10
	reflectSetValue(&g)
	fmt.Println("g:", g)
}
