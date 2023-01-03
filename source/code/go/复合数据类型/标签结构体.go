package main

import "fmt"
import "reflect"

type TagType struct {
	filed1 bool   "is right"
	filed2 string "name here"
	filed3 int    "age here"
}

func main() {
	tt := TagType{true, "ming-cloud", 17}
	ttype := reflect.TypeOf(tt)
	fmt.Println(ttype.Field(0).Tag)
	// ! 这里其实获取的是 ttype.Field(0) 的某个部分
	// ! fmt.Println(ttype.Field(0)） 的返回是 reflect 的东西，用来反查结构体的东西的东西 emm
	// ! {filed1 main bool is right 0 [0] false}

}
/* 
is right
*/