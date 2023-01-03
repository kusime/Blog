package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Test （） return >> ", Test(), time.Now())
	fmt.Println("TestCC () retrun  >> ", TestCC())
}

func Test() int {
	a := 200
	a_p := &a
	defer func() {
		(*a_p)++
		fmt.Println(a)
		// ! 这里的数值也成功通过指针修改
		// ! 但是就算这样还是没有能够修改这个函数的返回数值
		// ! 之外的东西？？
	}()
	// ! 那么只有可能就是先运行了return函数，在defer函数运行之前就以及把数值传出去了
	// ! 那么要实现对defer对函数数值的修改就只有让 内部的defer能碰到return 返回到的数值
	// ! 那么就是提前声明一个通用的数据
	fmt.Println(">>", a)
	return *(a_p)
}

func TestCC() (tt int) { // ! 这里只可以做默认初始化也就是零
	// ! 这里的返回数值是提前声明了的
	// ! 这里不只是声明数据的返回数值的类型
	// ! tt 确实被实体化成一个具体的数据的！！
	// ! 然后这个数据在下面的代码就都可以使用到这个数据的
	var tt_p *int
	tt_p = &tt
	defer func() {
		*(tt_p)++
		fmt.Printf("poiter is >> %p \nvalue is >> %d \n ", tt_p, *tt_p)
	}() // ! 这里后面的括号就是之类掉用这个函数的意思
	fmt.Printf("after func >> %d\n", tt)
	return tt

	// ! 所以这里意思就是
	// ! tt确实被return 提前返回过去了
	// ! 但是呢，因为 defer 后面的匿名函数可以调用到之前声明的tt的位置
	// ! 所以在defer运行的时候通过这个位置，修改了即将输出的数值
}

/*
>> 200
201
Test （） return >>  200 2021-03-28 13:39:15.887709951 +0800 CST m=+0.000040620
after func >> 0
poiter is >> 0xc000016120
value is >> 1
 TestCC () retrun  >>  1
*/
