package main

import "fmt"

func main() {
	var array [5]int // ! 完整的声明方式
	/* 	array2 := [5]int{} // ! 这个是var 语法声明的简略方式
	   	array3 := [5]int{1,2,3,4,5}// ! 这个是声明加初始化
	   	array4 := [...]int{1,2,3}// ! 这个是根据初始化的数量来动态的决定数组的产股
	   	array5 := [5]int{1:111,2:222}  // ! 这个是声明加初始化指定位置的数据
	*/
	array[1] = 1
	array[2] = 2
	array[3] = 3
	array[4] = 4
	/*
		var arrayP * int = &array
		# 这样会报错，也就是这里和C语言里面有些差距
		在C语言里面，数组名字本身就是对应数组的数据类型的指针
		数组为 int a[19]
		呢么对应的a 就是 int* 类型
		但是结合前文来看
		" [20]int , [100]int 这里两个是不同的类型"
		居然是不同类型的话，那么也会有不同类型的指针，所以只要在对应的数据类型前面加 * 就是代表对应数据类型的指针
		所以正确的指明指针的方式应该是
	*/

	/*
		var arrayP *[5]int = &array
		   output >> 0
		   output >> 1
		   output >> 2
		   output >> 3
		   output >> 4
	*/

	/*
		arrayP := &array // ! 这里就是直接的让go来自动判断类型
		output >> 0
		output >> 1
		output >> 2
		output >> 3
		output >> 4
	*/
	arrayP := &array
	for i := 0; i < len(array); i++ {
		// ! 这里验证和C语言是不是差不多的东西
		fmt.Printf("1 output arrayP[i] >> %d\n", arrayP[i])
		fmt.Printf("2 output array[i] >> %d\n", array[i])
	}
	fmt.Printf("arrayP Type is >> %T\n", arrayP)
	fmt.Printf("array Type is >> %T\n", array)
	/*
		output >> 0
	output >> 1
	output >> 2
	output >> 3
	output >> 4
	Type is >> *[5]int
	// ! k可以看见，这里直接可以输出对应的数据类型数值一样，和我们的推测一样
	*/
}
