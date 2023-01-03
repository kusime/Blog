package main

import (
	"fmt"
	"strings"
)

func main() {

	a1 := strings.Trim("!! Ming-Cloud !!", "!")
	// ! 这里修剪字符串前后字串，匹配到！然后一直裁剪直到遇到空格就停下来了
	a2 := strings.Trim("!! Ming-Cloud !!", "! ")
	// ! 第二个参数会被转化为rune对象然后去修剪匹配到的在子串中的字符代码
	// ! 可以理解为删除掉传入的字符串头尾的第二个字符串中的字符 这里为 空格和！
	a3 := strings.Trim("!! Ming-Cloud !!", " !")
	fmt.Printf("%q\n%q\n%q\n",a1, a2, a3)
	// ! 被裁剪后的字符串通过%q的格式化输出
}
