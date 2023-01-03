package main

import "fmt"

// Jiaoer 这里实现了个都能实现jiao这个行为的接口
type 会叫的 interface {
	叫()
}
type 会吃的 interface {
	吃()
}

type 又会叫又会吃的 interface {
	会叫的
	会吃的
	// ! 这里就是嵌套了上述两个实现了吃和叫的接口
	// ! 这里就可以实现对两个方法的嵌套
}

type 又会叫又会吃的2 interface {
	吃()
	叫()
	// ! 这里直接在新接口里面对两个方法进行实现也可以的
	// ! 但是这个方法灭有原来代码的简洁
}
type 又会叫又会吃的3 interface {
	会吃的
	叫()
	// ! 这里就是对接口混合，上面的书写方式达到的效果都是一样的

}

type 人 struct{}

func (r 人) 叫() {
	fmt.Printf("人会叫\n")
}
func (r 人) 吃() {
	fmt.Printf("人会吃\n")
}

type 狗 struct{}

func (r 狗) 叫() {
	fmt.Printf("狗会叫\n")
}
func (r 狗) 吃() {
	fmt.Printf("狗会吃\n")
}

type 猫 struct{}

func (r 猫) 叫() {
	fmt.Printf("猫会叫\n")
}
func (r 猫) 吃() {
	fmt.Printf("猫会吃\n")
}

func main() {
	yiming := new(人)
	/* 	var 某个生物 又会叫又会吃的
	   	某个生物=yiming */
	某个生物 := 又会叫又会吃的(yiming) //上面的申明方式也可以的
	某个生物.叫()
	某个生物.吃()
}

/*
人会叫
人会吃
*/
