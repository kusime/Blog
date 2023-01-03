package main
import "fmt"
func main() {
	giao := [1000]int{}
	// ! 上面创建了一个很大的数组，如果直接传递给函数
	// ! 函数会创建对应的副本（形参）
	// ! 所以比较好的解决方法就是传递数组的指针
	getInit(&giao)
	fmt.Printf("giao[19] value is >> %d\ngiao type is >> %T\n",giao[19],giao)
	fmt.Printf("giao[20] value is >> %d\n&giao type is >> %T\n",giao[20],&giao)
	
}

func getInit(getPointer *[1000]int) () { // ! 这里的申明方式是直接创建一个形参
	getPointer[19] = 19
	var tmPointer * [1000]int
	tmPointer = getPointer
	// ! 有必要区分一下数组的指针和数组本省
	tmPointer[20]= 20
	fmt.Printf(" tmPointer[20] value is >> %d\ntmPointer type is >> %T\n",tmPointer[19],tmPointer)
	
	
}

/* 
tmPointer[20] value is >> 19
tmPointer type is >> *[1000]int
giao[19] value is >> 19
giao type is >> [1000]int
*/