package main
import "fmt"

var a = 3

func Cur() func (bb,aa int) int {
	// ! 这里说明 Cur这函数不接收数值
	// ! 但是其返回值是一个匿名函数
	// ! 这个匿名函数接受两个数值
	return func (bb,aa int) int{
		return bb + aa
	}
}

func CCur(a int) func (bb,aa int) int {
	const ConstInt int = 1000
	return func (bb,aa int ) int {
		return a + bb + aa + ConstInt
		// ! 因为这个匿名函数在 这个函数的声明之下，所以，这里可以继承到
		// ! CCur函数声明时候的所有变量包括,在这个函数中声明的一个常数，或者是由CCur直接获取的形式参量
	}
}




func main() {
	fmt.Printf("sum is >> %d\n",Cur()(2,3))// ! 到这里为止，匿名函数可以享受到这个为止的所有变量
	// ! 这里调用Cur函数，这个函数返回那个接受两个数值的匿名函数
	// ! 然后后面两个数值被传入给那个匿名函数

	fmt.Printf("sum is >> %d\n",CCur(1)(2,a))
	// ! 可以看见输出为 1006
	// ! 所以 不管是 CCur直接传入的参数也好，还是返回匿名函数后的数值也好，还是在调用CCur函数时候的函数变量也好
	// ! 

}

/* 
sum is >> 5
sum is >> 1006
*/