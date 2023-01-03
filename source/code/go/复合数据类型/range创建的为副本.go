package main
import "fmt"
func main() {
	slice := []int{1,2,3,4,5}
	for index,value:=range slice{
		fmt.Printf("value pointer is %X , slice pointer is %X \n",&value,&slice[index])
		// ! 上面的语法是分别查看对应切片的实际位置，以及遍历到的元素的位置
		// ! 这个代码为了说明range里边的value只是实际元素地址的值的引用
	}
}
// ! 可以看见，这两个的地址是不一样的，所以对range的理解就是

// ! range创建了一个大小为切片的数据元素，然后每个被range遍历的值就会
// ! 复制到这个位置中，而且总之会覆盖写入这位置
/* 
value pointer is C0000160B0 , slice pointer is C000020180 
value pointer is C0000160B0 , slice pointer is C000020188 
value pointer is C0000160B0 , slice pointer is C000020190 
value pointer is C0000160B0 , slice pointer is C000020198 
value pointer is C0000160B0 , slice pointer is C0000201A0 
*/