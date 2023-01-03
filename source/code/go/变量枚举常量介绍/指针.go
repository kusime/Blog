package main

import "fmt"

func main() {
	测试字符串 := "Ming-Cloud"
	var strPointer *string
	var strPPointer **string
	strPointer = &测试字符串
	strPPointer = &strPointer
	fmt.Printf("%s \n strPointer>>%p \n", 测试字符串, strPointer) // ! ,strPointer+1) 不允许指针操作

	// ! 可以获取字符串的指针地址，也就是第一个字符的地址位置
	fmt.Printf("strPPointer>>%p\n*strPPointer=%p\n", strPPointer, *strPPointer)
	// ! 上面的就是操作套娃指针的意思，操作感觉和C语言是很像的
}

// ─── OUTPUT ─────────────────────────────────────────────────────────────────────

/*
Ming-Cloud 
 strPointer>>0xc00008a220 
strPPointer>>0xc0000a6018
*strPPointer=0xc00008a220
*/
