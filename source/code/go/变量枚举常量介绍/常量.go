package main
import (
	"fmt"
)
func main()  {
	const const1 int = 555 // 可以明显的声明变量的类型
	const declare = "null" // 也可以自动让其自动判断
	const calculate = 6/3 // 常量可以是计算出来的
	// ! const var_dona_do_this = 11  GO推荐使用驼峰命名法
	const getValueInFunction = len(declare)
	// ! const 不要从用户函数获取数值= 用户函数()
	const 中文="可以作为变量名字"
	// ！ 常量声明可以不应用
	fmt.Println(const1 ,declare,"\n",中文)
	// ! 不推荐fmt.Println以\n结尾
	const ( 
		c1,c2,c3 = 1,2,3
		c4,c5,c6 = 4,5,6
	)

}