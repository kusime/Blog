package main

import "fmt"

func main() {
	const (
		enum1 = 1
		enum2 = 2
		enum3 = 3
	)
	// ! 常量用作枚举类型
	// ! iota 重要的知识点
	const (
		e0 = iota
		e1
		e2, ee2, eee2 = iota, iota, iota // !
		brk           = "iota continue adding \n"
		brkContinue
		e5 = iota
	)
	fmt.Println(e0, "\n", e1, "\n", e2, ee2, eee2, "\n", brk, brkContinue, e5)
}
//
// ──────────────────────────────────────────────────── I ──────────
//   :::::: O U T P U T : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────
//


/* out put
0
 1
 2 2 2
 iota continue adding
 iota continue adding
 5
*/
