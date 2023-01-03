package main

import "fmt"

// Animal 对于公有结构体不给出事会出警告
type Animal struct {
	age int
}

// People 这个是人类，里面继承了匿名结构体 Animal
type People struct {
	Animal
	name string
}

// GetAge 这个获取动物的年龄
func (r *Animal) GetAge() int {
	return r.age
	/* 
	18
	*/
}

// GetAge 这个接受人类，获取年龄，但是因为getage方法以及被重名了，所以这里优先使用外围方法而不是内部匿名构体的方法
func (r *People) GetAge() int {
	return r.age*10
	/* 
	180
	*/
}


// GetName 这个获取人类的名字
func (r *People) GetName() string {
	return r.name
}

func main() {
	tom := People{Animal{18},"tom"}
	// ! 这里创建一个人类实例 tom
	// ! 因为人类结构体里面包含匿名的动物结构体
	// ! 所以人类可以访问到动物的获取年龄的方法
	// ! 这个相当于继承了
	fmt.Println(tom.GetAge())
}
