package main
import "fmt"

type 生物 interface{
	我是谁()
}

type 人类 struct{
	name string
}

func (r 人类) 我是谁()  {
	fmt.Printf("我是人类\n")
}

type 猪 struct{}

func (r 猪) 我是谁()  {
	fmt.Printf("我是猪\n")
}



func main() {
	某个生物 := 生物(人类{"yiming"})
	if v,ok := 某个生物.(人类); ok {
		fmt.Printf("看来我是%T,名字是%s\n",v,v.name)
	}
/* 
❯ go run "c:\Users\Kusime\Desktop\blog\content\posts\code\go\接口\接口类型判断.go"
看来我是main.人类,名字是yiming
*/

}