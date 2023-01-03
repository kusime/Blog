package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		for y := 0; y < 10; y++ {
			fmt.Printf("%d break detect\n", y)
			if y == 5 {

				break
				// ! 所以这个break不会导致嵌套循环的结束
				// ! 这个只会结束内层循环，如果需要结束两个循环的话
				// ! 需要使用那个标签
			}
		}
		fmt.Printf("%d still to run \n", i)
	}
}
/* 
0 break detect
1 break detect
2 break detect
3 break detect
4 break detect
5 break detect
0 still to run 
0 break detect
1 break detect
2 break detect
3 break detect
4 break detect
5 break detect
1 still to run 
0 break detect
1 break detect
2 break detect
3 break detect
4 break detect
5 break detect
2 still to run 

*/