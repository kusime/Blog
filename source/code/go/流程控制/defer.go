package main
import "fmt"
var i int =0
func print(){
	fmt.Println(i)
}

func printi(i int ){
	fmt.Println(i)
}


func main() {
	defer fmt.Printf("delay to run \n")
	// ! 最先进去的话·就是最后输出来的哦
	fmt.Printf("first to run \n")
	
	for i =0; i < 5; i++ {
		print() // ! 这里走正常的程序流，
		// ! 所以第一段直接正常运行
	}

	fmt.Printf("---------")
	// ! 这里也是正常的输出

	 // ! 重置之前的i == 5 
	for i =0; i < 5; i++ {
		defer print()
		// ! 这里调用defer
		// ! 函数执行顺序被压在底下
		// ! 输出五个五
		// ! 因为函数返回的时候没有指定i，而且换个角度来说
		// ! 5之前的数早就不存在了，因为没有被形式参量保存
		// ! 这个函数也只能是采取现有i 的数值，也就是递增后5
	
	}
	
	defer fmt.Printf("----sp first-----\n")
	
	var x int = 0
	for ; x < 5; x++ {
		defer printi(x)
		// ! 这里调用defer
		// ! 但是因为整个代码里面有不止其一次defer调用 （这里的，上面的）
		// ! defer的执行遵循 后进先出 
		// ! 也就是说，上面的defer是倒数第五个输出的
		// ! 4
		// ! 3
		// ! 2
		// ! 1
		// ! 0
	}
 	defer fmt.Printf("-----sp second-------- \n")
}
// ─── OUTPUT ─────────────────────────────────────────────────────────────────────

/* 
first to run 
0
1
2
3
4
--------------sp second-------- 
4
3
2
1
0
----sp first-----
5
5
5
5
5
delay to run 

个人感觉，对于defer 修饰的函数只要从后往前看就好了
*/