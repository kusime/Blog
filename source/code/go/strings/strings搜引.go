package main

import "fmt"
import "strings"

func main() {
	测试字符串 := "hi end hi"
	a1 := strings.Index(测试字符串, "hi")
	a2 := strings.LastIndex(测试字符串, "hi")
	a3 := strings.Index(测试字符串, "none")
	fmt.Println(a1, a2, a3)
	// ! 0 7 -1
	// ! 上面是 ascii 编码
	非ascii测试字符串 := "你好世界"
	a4 := strings.IndexRune(非ascii测试字符串, '世')
	fmt.Println(a4)
	// ! 这个可以正常的定位到对应字符串返回的字符串index
	// ! 但是返回的还是字节单位，而不是字符串的index
}
