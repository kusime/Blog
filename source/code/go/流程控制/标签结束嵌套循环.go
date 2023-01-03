package main
import "fmt"
func main() {
LOOP1:	
	for i := 0; i < 4; i++ {
		fmt.Printf("i == %d \n",i)
		for y := 0; y < 4; y++ {
			if y==1 {
				fmt.Printf("break Loop1 \n")
				// ! 这里也就是直接结束外层循环
				break LOOP1
			}
		}
	}
}