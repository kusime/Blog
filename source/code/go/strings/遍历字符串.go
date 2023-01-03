package main
import "fmt"
func main() {
	被遍历的字符串:= "中文字符"
	for unkonw , visual := range 被遍历的字符串 {
		fmt.Println(unkonw,visual)
		// ! 使用range方法可以方面的遍历出多字节字符
		// ! range str 返回两个数值
		// ! 第一个返回对于字符的 index 
	    // ! 第二个返回对应的字节码
		
	}

	rune方法遍历 := []rune(被遍历的字符串)
	for i := 0; i < len(rune方法遍历); i++ {	
		fmt.Println(rune方法遍历[i])
		fmt.Printf("%c\n",rune方法遍历[i])
	}

}

// ─── OUTPUT ─────────────────────────────────────────────────────────────────────

/* 
0 20013
3 25991
6 23383
9 31526
20013
25991
23383
31526

*/