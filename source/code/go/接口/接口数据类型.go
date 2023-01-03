package main

import "fmt"

// Jiaoer 这里实现了个都能实现jiao这个行为的接口
type 会叫的 interface {
	叫()
}

type 人 struct {
	age int
}

func (r 人) 叫() {
	fmt.Printf("人会叫\n")
}

type 狗 struct {
	name string
}

func (r 狗) 叫() {
	fmt.Printf("狗会叫\n")
}

type 猫 struct {
	age  int
	name string
}

func (r 猫) 叫() {
	fmt.Printf("猫会叫\n")
}

func main() {
	会叫的数组 := []会叫的{会叫的(人{13}), 会叫的(狗{"kusime"}), 会叫的(猫{13, "kusime"})}
	for _, v := range 会叫的数组 {
		v.叫()
	}
}

/*
❯ go run "c:\Users\Kusime\Desktop\blog\content\posts\code\go\接口\接口数据类型.go"
人会叫
狗会叫
猫会叫
*/
