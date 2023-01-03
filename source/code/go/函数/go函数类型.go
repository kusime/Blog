package main

import (
	"fmt"
)
func bigger(input1 int , input2 int) bool {
	return (input1 > input2)
}

func smaller(input1 int , input2 int) bool {
	return (input1 < input2)
}

func equal(input1 int , input2 int) bool {
	return (input1 == input2)
}


type (
	cmp func(int , int ) bool
// ! 这里其实就是类型别名的语法扩展
// ! 就是之前解释的自定义类型别名的意思

)


// ! 这里只有两个元素，第一个就是函数的输入参数类型，以及输出返回值的类型
// ! 但是这里注意，这里可以参考C语言的原型函数？
// ! 也就是说这个就只是一个函数模板
// ! 作用就是作为参数传入的时候指明传入函数的传入传出参数的类型

 

func compare(input1 int,input2 int,function cmp) bool {
	return (function(input1,input2))
}
func main() {
	test1:=1
	test2:=2
	fmt.Printf("test1 > test2 == %t\n",compare(test1,test2,bigger))
	fmt.Printf("test1 < test2 == %t\n",compare(test1,test2,smaller))
	fmt.Printf("test1 = test2 == %t\n",compare(test1,test2,equal))
	fmt.Printf("function type compare ")
}