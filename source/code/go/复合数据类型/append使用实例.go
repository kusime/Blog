package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	testCapAndLen(slice)

}

func testNewArrayBeyond(slice []int, newslice []int) {

	fmt.Printf("\n------------new test---------------------\n")
	fmt.Println(slice, "\n", newslice)
	fmt.Printf("\n------------new test---------------------\n")
}

func testCapAndLen(slice []int) { // ! 这里传递的是切片的副本
	// ! 这里先初始化一个切片，这里的长度和容量都是5 所以之后添加元素的时候
	// ! 切片的容量应该会翻倍的
	fmt.Printf("slice len is %d\ncap is %d\n", len(slice), cap(slice))
	newslice := append(slice, 333)
	fmt.Printf("newslice len is %d\ncap is %d\n", len(newslice), cap(newslice))
	// ! 这里区分一下长度和容量
	// ! 可以看见长度只是在原来的长度下添加了1
	// ! 但是可以看见对应的容量翻倍了
	// !新创建的slice具有一个全新的数组
	newslice[0] = 9999
	testNewArrayBeyond(slice, newslice)
	// ! 从下面的 	 [9999 2 3 4 5 333]
	// ! 可以看出来这里是两个不同的底层数组了

	slice2 := make([]int, 1, 10)
	// ! 这里创建了一个int形式的切片，长度为 1 容量为 10
	slice2[0] = 1
	// ! 这里的切片的长度和容量是不一样的，所以append函数不会创建一个新的底层数组
	// ! 新旧切片的修该会相互影响的
	newslice2 := append(slice2, 2)
	newslice2[0] = 444

	testNewArrayBeyond(slice2, newslice2)
	// ! 结合下面的修改来看，源切片的数值，有因为 	newslice2[0]=444 的修改而修改
	// ! 所以可以确定，两个是公用一个底层数组的

	/*
	slice len is 5
	cap is 5
	newslice len is 6
	cap is 10

	------------new test---------------------
	[1 2 3 4 5]
	 [9999 2 3 4 5 333]

	------------new test---------------------

	------------new test---------------------
	[444]
	 [444 2]

	------------new test---------------------
	*/
}
