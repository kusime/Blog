package main

import "fmt"

func main() {
	ret1 := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Printf("%d\n", ret1)
	// ! 这里直接调用匿名函数并且返回了数值

	funCall := func(x, y int) int {
		// ! 这里直接是直接把函数赋值给一个变量
		return x + y
	}

	fmt.Printf("funCall OutPut >> %d\n", funCall(1, 2))
	// ! fmt.Printf("funCall  >> %p\n", 4820960(1,2))
	// ! cannot call non-function 4820960 (type int)
	// ! 但是好像又不能直接调用对应的十六进制地址
	// ! 这里本质上就是直接赋值了对应匿名函数的函数指针，然后调用

}
