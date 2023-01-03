package main
import "fmt"

func adding(a int ) int {
	defer un(trace("adding"))
	fmt.Printf("这里的是adding的代码\n")
	return a  + 1
	// ! 最后输出只是在15行的代码决定的
	// ! 不代表在这个函数中return是在这函数中最后运行的
}


func main() {
	aa := 1
	bb := adding(aa)
	fmt.Printf("%d\n",bb)
}


func trace(s string) string {
	fmt.Printf("function [ %s  ] start \n",s)
	return s
}

func un (s string) {
	fmt.Printf("function [ %s ] end \n" , s)
}


// ! 这里就是为了说明defer可以作为debugger 的一个意思
// ! 因为defer在return之后，defer运行的可以获取到

// ! 在这个函数中的函数变量的数值
