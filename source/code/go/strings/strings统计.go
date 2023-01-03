package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	字符串 := "Google is Cool呀呀"
	字段 := "o"
	fmt.Println(strings.Count(字符串, 字段))
	// ! 这个统计个数是通过rune类型来进行运算的
	// ! 也就是说输入的数值要是字符串类型
	// ! 这也就兼容了统计的对象为多个字符的情况
	fmt.Println(len([]rune(字符串)))
	// ! 上面的转化为rune类型对象，这样统计的个数就是字符串的长度
	// ! 而不是字节长度，assic字符串长度字节长度一致都是1个字节

	fmt.Println(utf8.RuneCountInString(字符串))
	// ! 忘记引用的包，vscode会自动引用
	// ! 这个是用utf8的包中的方法来统计个数
}
