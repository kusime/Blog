package main
import "fmt"
func main() {
  for i := 0; i < 10; i++ {
    if i==5 {
        fmt.Println(i)
        break
    }
}
}
/*
上面这个就会直接到五的时候就退出这个循环了
但是这个对于嵌套循环就有点不一样了

5
*/