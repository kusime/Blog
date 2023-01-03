package main
import "fmt"
import "reflect"

func main() {
	var v1 float64
	v1 = 1 // !声明了是浮点类型的话，不管有没有小数点就是浮点形式
	v2 :=2 // ! 没有小数点，而且没有声明变量的类型，被推断为整型
	v3 :=3.0 // ! 浮点类型
	// v := v1+v2 不同数据类型不可运算 float + int
	v := v1 + v3
	fmt.Println(v1,v2,v3,v)
	fmt.Println("type of v",reflect.TypeOf(v))
}

// ─── OUTPUT ─────────────────────────────────────────────────────────────────────

/* 
1 2 3 4
type of v float64
*/