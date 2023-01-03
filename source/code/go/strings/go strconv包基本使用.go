package main

import (
	"fmt"
	"strconv"
)

func main() {
	// ! 变量声明

	// ! 基本函数介绍
	fmt.Printf("机器位数 》》%d\n", strconv.IntSize)
	// ! strconv.IntSize 这个部分其实就是一个常量可以用来查看int默认的位数
	fmt.Printf("%t,%s\n", strconv.Itoa(666), strconv.Itoa(666))

	// ! fmt.Printf("%s",strconv.Atoi("Ming"))
	// ! 1.Ming不是有效的数字
	// ! 2.因为strconv.Atoi会返回两个数值，所以不应该直接去对其进行格式化
	after, err := strconv.Atoi("666")
	// ! fmt.Printf(after)
	// ! 因为printf这个函数需要格式化字符串和其变量，所以不是直接传入对应的数字的
	fmt.Printf("%T %d,%s", after, after, err)

	
}

// ─── OUTPUT ─────────────────────────────────────────────────────────────────────

/*
机器位数 》》64
%!t(string=666),666
// ! 其实这个算是报错了，这个的意思就是，嘿,这个参数不要用%t来格式化，这个参数是string，然后其值是 xxx
int 666,%!s(<nil>)%
*/
