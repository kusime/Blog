package main

import "fmt"

var x interface{}// ! 空接口

func main() {
	a := 1
	switch a {
	case 1:
		fmt.Printf("a==1\n")
		fallthrough // ! 下面就不判断条件了
	case 2:
		fmt.Printf("a==2\n")
		// ! 直接就运行上面这个语句了
	}
	x = 1
	
	switch  i:= x.(type) {
	case int:
		fmt.Printf("this is %T\n",i)
	default: // ! 这个和linux
		fmt.Printf("not match ")
	}
	// ! 上面是类型


}

/*
a==1
a==2
*/
