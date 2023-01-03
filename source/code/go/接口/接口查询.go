package main

import "fmt"

type 某个东西 struct {
	name string
	age  int
}
type 大接口 interface {
	名字是什么()
	获取全部信息()
}

type 小接口 interface {
	名字是什么()
}

func (r 某个东西) 名字是什么() {
	fmt.Printf("name == %s\n", r.name)
}

func (r 某个东西) 获取全部信息() {
	fmt.Printf("name == %s\nage == %d", r.name, r.age)
}

func main() {
	var 小接口实例 小接口
	var 大接口实例 大接口

	// 这里主要的判断还是一个接口时间的集合关系
	大接口实例 = 某个东西{"ming",1}
	小接口实例 = 大接口实例 // 这里就是把一个提供了足够方法的接口传递给了一个较为有限的接口之中
	小接口实例.名字是什么()

	// 然后判断是否 某个接口实例是否完全的实现另一个接口
	// 如果判断成功就说明,小接口这个接口里面的所有方法,对于大接口实例(某个实例)来说,小接口里面的方法,这个实例都可以使用
	// 这个判断就是i接口查询的价值所在
	if value , ok := 大接口实例.(小接口);ok {
		fmt.Printf("%T",value)
	}
}
