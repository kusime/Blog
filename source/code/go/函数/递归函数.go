package main
import "fmt"

func FFCall(num uint64) (result uint64){
	if (num > 0 ){
		return num*FFCall(num-1)
	}
	return 1
}

func main() {
	var a uint64 = 15
	fmt.Printf("%d\n",FFCall(a))
}
// ! 这里注意一下最大数值的大小，可能会超出类型的最大数值

/* 
1307674368000
*/