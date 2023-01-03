package main

import (
	"fmt"
	"strings"
)
func main() {
	等待分割的字符串:="Ming-Cloud"
	a1:=strings.Split(等待分割的字符串,"-")
	// ! 这里的操作对象为rune对象，也就是多字节字符
	// ! 这里的方法可以理解为Python中的split函数
	// ! go语言返回的叫做切片，但是Python中的叫做数组
	fmt.Println(a1) // [Ming Cloud]
	fmt.Printf("%q",a1)
	// ! 使用%q来格式化对应的切片
	
}

// ─── OUTPUT ─────────────────────────────────────────────────────────────────────

/* 
[Ming Cloud]
["Ming" "Cloud"]
*/