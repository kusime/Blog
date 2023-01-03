package main

import "fmt"

// AnyType 这里定义接受任意类型的数据,然后使用接口的语法实现不同类型数据的判断
func AnyType(items ...interface{}) {
	for _, v := range items {
		switch v.(type){
		case bool:
			fmt.Printf("bool这里\n")
		case string:
			fmt.Printf("string这里\n")
		case int:
			fmt.Printf("int这里\n")
		
		}
		
	}
}

func main() {
	AnyType(true,"1",1)
}
/* 
❯ go run "c:\Users\Kusime\Desktop\blog\content\posts\code\go\接口\空接口接受多类型数据实现判断.go"
bool这里
string这里
int这里
*/