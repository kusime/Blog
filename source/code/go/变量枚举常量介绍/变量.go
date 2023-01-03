package main
import "fmt"
func main() {
	// !  var name type
	var v1 int = 44
	var v2 string = "ming"
	// var v33,v44 int = 1,2
	fmt.Println(v1,v2)
	var (
		v3,v4=3,4
		str1,str2="str1","str2"
	)
	fmt.Println(v3,v4,str1,str2)
	noDeclareTypeVar := "不用声明类型，自动判断\nvar noDeclareTypeVar string"
	fmt.Println(noDeclareTypeVar)
}
//
// ──────────────────────────────────────────────────── I ──────────
//   :::::: O U T P U T : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────
//



/* 
44 ming
3 4 str1 str2
不用声明类型，自动判断
var noDeclareTypeVar string
*/
