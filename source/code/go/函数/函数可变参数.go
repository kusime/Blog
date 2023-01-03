package main

import "fmt"

func getPrint(sliceInput ...int) {
	fmt.Printf("sliceInput pass to next Function\n")
	getPrintCall(sliceInput...)

}

// ! 边长参数支持二次传递
// ! 本质上来说的这个参数就是输出切片也即是
// ! []type类型 这也就是为什么可以用range来迭代的原因
func getPrintCall(sliceInput ...int) {
	for _, v := range sliceInput {
		fmt.Printf("%d\n", v)
	}
}

func main() {
	testSlice := []int{1, 2, 3, 4, 5}
	getPrint(testSlice...)
	fmt.Printf("-------\n")
	getPrint(6, 6, 6, 6, 6)
	// ! 这里多个参数的传入和slice的传入是一样的
	fmt.Printf("-------\n")
	getPrint([]int{5, 5, 5, 5, 5}...)
	// ! 这里是直接传入数组类了


	// ! -------------------------------------
	// ! 因为前面在函数说明了 可变参数的类型
	// ! 如果希望传入任意类型可以尝试使用接口
	// ! fmt.Printf()的实现就是这样的
	/*
		func Printf(format string , args ...interface{}){
			// !
		}
	*/
}

// ─── OUTPUT ─────────────────────────────────────────────────────────────────────

/* 

sliceInput pass to next Function
1
2
3
4
5
-------
sliceInput pass to next Function
6
6
6
6
6
-------
sliceInput pass to next Function
5
5
5
5
5


*/