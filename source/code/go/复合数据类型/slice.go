package main
import "fmt"
func main() {
	slice := []int{1,2,3,4,5}
	Test(slice)// ! [999 2 3 4 5]
	fmt.Println(slice)
	newSlice := slice[1:3]
	newSlice2 := slice[0:3:4]
	newSlice[0]=999
	newSlice2[0]=555
	fmt.Println("newSlice >>",newSlice,"\nnewSlice2 >>",newSlice2)
	fmt.Println("slice >> ",slice)
}

func Test(arr []int)  {
	arr[0]=999	
}