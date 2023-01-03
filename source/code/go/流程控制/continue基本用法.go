package main

import "fmt"

func main() {
	/* LOOP1:
	   	i := 0
	   这里会报错，好像是标签后面只允许代码块
	   类似于 for ， switch ， while等
	   但是那个就是个赋值语句
	*/
LOOP1:
	for i := 0; i < 5; i++ {
		switch {
		// ! 因为下面直接指明了case比较的对象所以这里没有规定switch的对象
		case i == 1:
			fmt.Printf("Loop 1 continue\n")
			continue LOOP1 // ! 回到LOOP1的时候不会继续运行对应的初始化语句

		}
		fmt.Printf("Loop 1 continue --- i is>>%d\n", i)
	}
}
/* 
Loop 1 continue --- i is>>0
Loop 1 continue // ! 这里就直接回到外循环的开头了就直接忽略了 fmt.Printf("Loop 1 continue --- i is>>%d\n", i) 这个语句
Loop 1 continue --- i is>>2
Loop 1 continue --- i is>>3
Loop 1 continue --- i is>>4

*/