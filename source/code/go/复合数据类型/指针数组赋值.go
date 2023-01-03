package main

import "fmt"

func main() {
	var arrP [3]*string
	// ! 上面声明的是，包含三个字符串指针的数组
	arr := [3]*string{new(string), new(string), new(string)}
	// ! 上面声明的是包含三个初始化后的字符串指针的数组
	arrP = arr
	fmt.Println(arrP, "\n", arr)
	/*
		[0xc00009e220 0xc00009e230 0xc00009e240]
	 [0xc00009e220 0xc00009e230 0xc00009e240]
	*/
	// ! 可以看见对应的指向的为同一个字符串
	// ! 因为指向的都是同一个东西了 
}
