package main
import "fmt"
func main() {
	var v1 float64 // ! v1最多
	v1=1
	v2 := 1.00000000000000000001 // ! 因为精度过高后面的值就被忽略了
	if v1 == v2 {
		fmt.Println("yes")
	}

}