package main

import "fmt"
import "strings"

func main() {
	// ! strings.HasPrefix(str,"字段")
	// ! strings.HasSuffix(str，”字段“)
	// ! strings.Contains(str，”字段“)
	// ! 上面分别用来判断输出的字段是否存在于对应字符串中
	// ! 分别是前缀。后缀，整个字符串
	// ! 返回值就是true 或者 false
	被比较的字符串 := "ming-cloud"
	a1 := strings.HasPrefix(被比较的字符串, "mi")
	a2 := strings.HasSuffix(被比较的字符串, "cloud")
	a3 := strings.Contains(被比较的字符串, "-")
	fmt.Println(a1,a2,a3)
	// ! 这个字符串的前缀后缀被-给隔断
	// ! 隔断符号可以是空格或者一些标点符号
}

/* 
true true true
*/
