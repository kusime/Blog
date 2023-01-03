package main

import "fmt"

// Jiaoer 这里实现了个都能实现jiao这个行为的接口
type Jiaoer interface {
	叫()
}

type people struct{}

func (r people) 叫() {
	fmt.Printf("人会叫\n")
}

type gou struct{}

func (r gou) 叫() {
	fmt.Printf("狗会叫\n")
}

type mao struct{}

func (r mao) 叫() {
	fmt.Printf("猫会叫\n")
}

func main() {
	人类 := new(people)
	猫 := new(mao)
	狗 := new(gou)

	var 会叫的 Jiaoer

	会叫的 = 人类
	会叫的.叫()

	会叫的 = 猫
	会叫的.叫()

	会叫的 = 狗
	会叫的.叫()

}

/* 
❯ go run "c:\Users\Kusime\Desktop\blog\content\posts\code\go\接口\基本实例.go"
人会叫
猫会叫
狗会叫
*/
