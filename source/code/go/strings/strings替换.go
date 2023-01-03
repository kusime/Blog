package main

import "fmt"
import "strings"

func main() {
	被替换字符串 := "你好世界，世界你好"
	oldone := "世界"
	newone := "地球"
	index := 2 
	// ! 这个指明直到替换到index个匹配到的字符串
	// ! -1位匹配到所有，大于现有字符的匹配子串被视为匹配所有
	fmt.Println(strings.Replace(被替换字符串, oldone, newone, index))
}
