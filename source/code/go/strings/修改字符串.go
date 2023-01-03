package main

import (
	"fmt"
)

func main() {
	原来的字符 := "hello golang"
	新的字符 := []byte(原来的字符)
	新的字符[5] = '1'
	//!新的字符[5]="1"
	//! cannot use "1" (untyped string constant) as byte value in assignmen
	// ! 这里双引号是字符串，单引号为字符
	fmt.Println(新的字符) // ! 这里现在孩纸字节码的输出
	fmt.Printf("%s\n", 新的字符)

	修改多字节 := "名云"
	修改后的 := []rune(修改多字节)
	修改后的[0] = '铭'
	fmt.Println(string(修改后的))
	// !好像不同的东西，打印的方式不一样

}
