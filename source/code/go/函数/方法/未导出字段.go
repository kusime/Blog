package main
import "fmt"

type person struct{
  name string
  age int
}

func (r *person) Setter(input string,age int)  {
  r.name = input
  r.age = age
}

func (r *person) Getter()  (outputName string, outputAge int)  {
  return r.name,r.age
}

func main()  {
  tom := new(person)
  tom.Setter("tom",18)
  fmt.Println(tom.Getter())
}

/* 
tom 18
*/