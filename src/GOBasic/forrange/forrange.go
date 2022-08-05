package forrange

import (
	"fmt"
)

func Init() {
	sum := 1
	// 标准for 循环
	for i := 0; i < 10; i++ {
		sum += i
	}
	// 替代while
	for sum < 100 {
		sum += sum
	}

	fmt.Println(sum)
	numbers := []int{1, 2, 3, 4, 5, 6}

	// 相当于 enumerate(numbers)
	for i, x := range numbers {
		fmt.Printf("%d: %d\n", i, x)
	}
	
	// 相当于 while true
	// for  {
	// 	fmt.Println("line")
	// }
}