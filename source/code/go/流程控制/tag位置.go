package main

import "fmt"

func main() {

TAG2:
	fmt.Printf("tag2 select")

TAG3:
	fmt.Printf("tag3 select")

	/* BREAK:
	fmt.Printf("break here")
	break
	// ! break not in for, switch, or select statement
	可以使用break标签来指定对应的break位置
	但是要注意的是，这个break标签需要在允许的语句代码块中出现
	*/

	// ! 作用一，switch

	var a int = 1
	switch a {
	// ! 这里不出意料的会死循环
	case 1:
		goto TAG1
	case 2:
		goto TAG2

	case 3:
		goto TAG3

		/* 	DEFA:
		fmt.Printf("done") 
		// ! goto DEFA jumps into block starting at
		*/
		
	}
TAG1:
	goto DEFA

	/* DEFA:
	fmt.Printf("done")
	// ! goto DEFA jumps over variable declaration at line
	*/
DEFA:
	fmt.Printf("done")
	// ! 这个才对
	// ! 总结就是，虽然tag可以实现随意的位置跳转
	// ! 但是我们最好还是使用其来跳过不想运行的语句
	// ! 但是切记不要随便去逆转程序的运行方向 （从下往上运行）
	// ! 相当有可能出现死循环
	// ! 所以要特别注意TAG摆放的位置

}
