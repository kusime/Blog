package main
import "fmt"

// Camera 这个是一个单独的结构体，可以理解为一个单独的类
type Camera struct{}
// TakeShot 这个是上述结构体的方法
func (r Camera) TakeShot()  {
	fmt.Printf("茄子！！")
}


// Phone 这个是一个单独的结构体，可以理解为一个单独的类
type Phone struct{}
// Call 这个是上述结构体的方法
func (r Phone) Call()  {
	fmt.Printf("电话来了！！")
}

// CameraPhone 这里对上述两个结构体进行多重继承
type CameraPhone struct{
	Camera
	Phone
}

func main() {
	tomSPhone:=	CameraPhone{}
	// ! 上面声明多重继承的结构体实例
	tomSPhone.Call()
	tomSPhone.TakeShot()
	// ! 上面两个分别调用来自于不同类的方法
/* 
电话来了！！茄子！！
*/

}