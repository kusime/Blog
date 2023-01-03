package main

import "fmt"




func main() {
	var arr [4][2]int // ! 那个快速声明字面量声明的特权
	arr[0][0] = 1
	/* var arr3 [2][2]int = {{1,1},{2,2}}
	   # command-line-arguments
	   posts/code/go/复合数据类型/多维数组赋值.go:
	   6:22: syntax error: unexpected {,
	   	expecting expression
	*/
	arr1 := [2][2]int{{1, 1}, {2, 2}} // ! 只有这个语法可以快速赋值
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("arr1 [%d][%d] >> %d \n", i, j, arr1[i][j])
		}
	}
	/*
	   arr1 [0][0] >> 1
	   arr1 [0][1] >> 1
	   arr1 [1][0] >> 2
	   arr1 [1][1] >> 2
	*/
	var arr3 [2]int
	arr3 = arr1[1]
	// ! 这里是二维数组的赋值
	for i := 0; i < 2; i++ {
		fmt.Printf("%d\n",arr3[i])
	}

}
