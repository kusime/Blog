package main
import "fmt"

type date struct {
  mouth int
  day int
}

func (r *date) PTRNow()  {// ! 上面是基于date类型指针的方法。引用的是date实例的结构体指针
  fmt.Println(r,r.mouth , r.day)
}

func (r date) Now()  {// ! 上面是基于date类型指针的方法。引用的是date实例的结构体指针
	fmt.Println(r,r.mouth , r.day)
  }

func (r *date) PTRChangeNow()  {
	r.mouth = 1
	r.day = 1
  }


func (r date) ChangeNow()  {// ! 上面的是date类型，引用的是date实例，这样会复制整个副本，所以下面修改不会影响到传入的date实例
  r.mouth = 1
  r.day = 1
}

func main()  {
  var test date
  // 上面定义了一个值类型
  fmt.Printf("test 数值是 > %v\n",test)
  test.ChangeNow() // 这里test被传入到 r ,然后r只是一个形式参数副本所以不会改变
  fmt.Printf("test 类型是 > %T\n第一次修改后的数值是 %v\n",test,test)

  test.PTRChangeNow()// 按道理来说,test的指针是指针,值是值,两个是不同的东西,但是golang有语法糖,具体点说就是会传入的时候自动取其指针
  // 上面个的语句本质上等价于 (&test).PTRChangeNow
  // 所以说,test的地址副本被复制给 PRTChangeNow 这个形式参量中,然后取得地址后,自然也就是可以修改了
  fmt.Printf("第二次修改后的数值是 %v",test)
  

}


/* 
{0 0} 0 0
&{0 0} 0 0
{0 0} 0 0
&{0 0} 0 0
{1 1} 1 1
&{1 1} 1 1


结合第一列来看，可以分析出，具体是指针还是值类型
是由方法的接受者类型决定的

某个数据类型的指针
和某个数据类型，
她们的方法是通用的，也就是说，某数据类型的指针类型和某数据类型，
她们的方法是不可以重名的

posts/code/go/函数/方法/指针对象.go:19:6: method redeclared: date.ChangeNow
        method(*date) func()
        method(date) func()

*/