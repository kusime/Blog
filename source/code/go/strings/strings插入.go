package main

import (
	"fmt"
	"strings"
)
func main() {
	等待转化的字符串:="以空格 为切片 切割对象"
	a1:=strings.Fields(等待转化的字符串)
	// ! Fields 把对应的字符串转变为切片对象
	fmt.Println(a1)
	for index,pattern :=range a1{
		// ! 这个语法用于赋值 对应切片的的index，对应子串
		// ! range str 这个方法返回多个数值
		// ! 这里其实就是对切片的遍历
		fmt.Println(index,pattern)

		fmt.Printf("%s\n",pattern)
		// ! 上面的语句可以看出来其是字符串对象，
		// ! 然后这个切片元素就可以理解为Python的
		// ! ["str1","str2","str3"]
	}
	newstring:=strings.Join(a1,"-")
	// ! 方法接受的是切片，然后后面的是rune类型
	// ! 返回的值就是字符串
	fmt.Println(newstring)
}
// ─── OUTPUT ─────────────────────────────────────────────────────────────────────

/* 
[以空格 为切片 切割对象]
0 以空格
1 为切片
2 切割对象
*/