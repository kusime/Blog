package main
import "fmt"
type any interface{}

func main() {
	test := any(1)
	// ! 空接口类型就是一种普适性的 包装
	fmt.Printf("对应类型 >> %T",test)
	var anyArray  [4]interface{}
	anyArray[0]=1
	anyArray[1]="one"
	anyArray[2]=true
	anyArray[3] = []int{1,2,3,4,5}

	for _, v := range anyArray {
		fmt.Printf("对应类型 >> %T\n",v)
	}
}