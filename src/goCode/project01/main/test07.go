package main

/*
在本例中，使用一个数值表示时间中的“秒”值，
然后使用 resolveTime() 函数将传入的秒数转换为天、小时和分钟等时间单位。




go 语言里面有三种类型的函数
普通的带有名字的函数
匿名函数或者lambda函数
方法
普通函数的声明
func 函数名(形式参数列表)(返回值参数列表){
	函数体
}
go语言中传入与返回参数在调用和返回的时候都是用的是值传递 这里需要注意的是指针 切片 和 map等
引用类型对象在参数传递中不会发生复制 而是将指针进行复制 类似于创建一次引用
看下面的例子 详细了解go语言的参数传递

在go语言中 函数也是一种类型 可以和其他类型一样保存到变量中 下面的代码定义了一个函数变量f
并将一个函数名为fire（）的函数赋给函数变量f 这样调用函数变量f时 实际调用的就是fire()方法

go语言支持匿名函数 即在需要使用函数是在定义函数 匿名函数没有函数名 只有函数体 函数可以作为一种类型被赋值给
函数类型的变量 匿名函数也往往以变量方式传递 go语言支持随时在代码里定义匿名函数
定义一个匿名函数
func(参数列表)(返回参数列表){
    函数体
}
可以在声明后调用
func(data int) {
    fmt.Println("hello", data)
}(100)
将匿名函数赋值给变量
// 将匿名函数体保存到f()中
f := func(data int) {
    fmt.Println("hello", data)
}
// 使用f()调用
f(100)

匿名函数的用途非常广泛 它本身就是一种值 可以方便地保存在各种容器中 实现回调函数和操作封装 与scala差不多
函数可以作为其它函数的参数进行传递，然后在其它函数内调用执行，一般称之为回调


go语言函数类型实现接口--把函数作为接口来调用
函数和其他类型的函数一样 也可以实现接口


go语言闭包
函数 + 引用环境
*/
import "fmt"

//定义常量
const (
	// 定义每分钟的秒数
	SecondsPerMinute = 60
	// 定义每小时的秒数
	SecondsPerHour = SecondsPerMinute * 60
	// 定义每天的秒数
	SecondsPerDay = SecondsPerHour * 24
)

func main() {

	//day, hour, minute := resolveTime(1000)
	//fmt.Printf("%d,%d,%d", day, hour, minute)
	//初始化一个结构体
	in := Data{
		complax:  []int{1, 2, 3},
		instance: InnerData{a: 5},
		ptr:      &InnerData{a: 1},
	}
	fmt.Println(in)
	fmt.Printf("input ptr%p\n", &in)
	out := passByValue(in)
	fmt.Println(out)
	fmt.Printf("out ptr%p\n", &out)
	/*
		经过代码示例
		发现 无论是将结构体传入到方法中  还是方法返回一个结构体 都是值传递类型
		结构体的地址发生改变 意味着所有的结构体都是一块新的内存地址 无论是将Data结构传入函数内容
		还是通过函数返回值传回Data都会发生复制行为
		所有的结构体的成员值没有发生改变 原样传递 意味着所有的参数
		都是值传递
		结构体的ptr成员在传递过程中保存一致 表示指针在函数传递中传递的只是指针 不会复制指针指向的部分
	*/

	var f func()
	f = fire
	f()

	f1 := func(data int) {
		fmt.Println(data)
	}
	visit([]int{1, 2, 3}, f1)
}

func resolveTime(seconds int) (day, hour, minute int) {
	day = seconds / SecondsPerDay
	hour = seconds / SecondsPerHour
	minute = seconds / SecondsPerMinute

	return day, hour, minute
}

//用于测试值传递效果的结构体
type Data struct {
	complax  []int      //测试切片在参数传递的效果
	instance InnerData  //实例分配的innerData
	ptr      *InnerData //将ptr声明为InnerData
}
type InnerData struct {
	a int
}

//传递值测试函数
func passByValue(infunc Data) Data {
	fmt.Println("inFunc value:", infunc)
	fmt.Printf("inFunc ptr%p\n", &infunc)
	return infunc
}
func fire() {
	fmt.Printf("fire")
}

//匿名函数
func visit(list []int, f func(int2 int)) {
	for _, v := range list {
		f(v)
	}
}

//调用器接口
type Invoker interface {
	//需要实现一个Call方法
	Call(interface{})
	//这个接口需要实现 Call() 方法，调用时会传入一个 interface{} 类型的变量，这种类型的变量表示任意类型的值。
}
