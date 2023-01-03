package main

import "fmt"

func main() {
	解释型 := "这个会解释转义字符，用双引号声明\n"
	非解释型 := `这个使用反引号来声明，
	还可以跨行声明，
	里面的 \n 转移字符不会被解释`
	fmt.Println(解释型)
	fmt.Println(非解释型)
	可以遍历 := "yes你好"
	for i := 0; i < len(可以遍历); i++ {
		fmt.Println(可以遍历[i])
		// ! 这样输出的是字节码
		// ! fmt.Println("%s",可以遍历[i])
		// ! 上面的同构%s来把字符代码转化为字符
	}
	裁剪对象byte := "ming-cloud"
	fmt.Println(裁剪对象byte[0:4])
	裁剪对象multiply := "你好"
	fmt.Println(裁剪对象multiply[0:1])
	// !这只可以裁剪出第一个字节
	// ! 但是一个汉字有三个字节
	fmt.Println(裁剪对象multiply[0:3])

}
