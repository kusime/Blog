package main

import "fmt"

func yep() int {
	return 1
}

func main() {
	a := 1
	if a == 1 {
		fmt.Printf("a == 1 \n")
	}
	// ! 这个条件语句格式和c 语言的很像
	// ! 但是这里没有用括号

	if a == 2 {
		fmt.Printf("a == 1 \n")
	} else { // ! 对于这个else 代码块就只可以紧接着这个if代码块
		fmt.Printf("a == 2\n")
	}
	/*
		else
		{
			fmt.Printf("a==2\n")
			// !这个是不允许的
			// ! expected statement, found 'else'syntax
		}
	*/
	if a == 2 {
		fmt.Printf("a == 2\n")
	} else if a == 1 {
		fmt.Printf("a == 1 \n")
	}
	// ! 和else 一样 else if 必须和紧跟着那个if 代码快

	if a=yep();a==1{
		fmt.Printf("a == 1 \n")
	}
	// ! 前面这个东西就是初始化子语句
	// ! 可以直接赋值也可以直接调用函数
}
