package main
import "fmt"

func toBeOne(a * int )  {
	* a = 1
	// ! 这里实现了直接使用指针来修改指针指向的变量

	// ! 取指针的值的运算符和C语言一样 ， 对 * 号进行了重载
}

func main() {
	var 这个是变量 int = 666
	var 这个是变量的指针 * int = & 这个是变量
	toBeOne(这个是变量的指针)
	fmt.Printf("after One > %d\npointer is > %p \n",这个是变量,这个是变量的指针)

}

/* 
after One > 1
pointer is > 0xc0000b8010 
*/