package main

import "fmt"

/*
go 语言不是一种传统的面向对象编程语言 它里面没有类和集成的概念
但是go语言里面有非常灵活的接口概念 通过他可以实现很多面向对象的特性
很多面向对象的语言都有相识的接口 但是go语言中的接口类型的独特之处在于它是
他是隐式实现的  也就是说我们没有必要对于给定的具体类型定义所有瞒住的接口类型 简单地拥有
一些必要的方法就足够了
这种设计可以让你创建一个新的接口类型满足已经存在的具体类型却不会改变这些类型的定义
当我们使用的类型来自于不受我们控制的包时 这种设计尤其有用
接口类型是对其他类型行为的抽象和概括 因为接口类型不会和特定的实现细节绑定在一起 通过这种抽象的方式我们可以让我们的’
函数更加灵活和更具有适应能力
接口是双方约定的一种合作协议 接口实现值者不需要关心接口会被怎样使用 调用者也不需要关心接口的实现细节
接口是一种类型也是一种抽象结构 不会暴露出所含的数据的格式 类型及结构
接口的声明格式
type 接口类型名字 interface{
	方法1（参数列表1） 返回值列表1
	方法2（参数列表2） 返回值列表2
	......
}

如果一个任意类型 T 的方法集为一个接口类型的方法集的超集，则我们说类型 T 实现了此接口类型。T 可以是一个非接口类型，也可以是一个接口类型。

实现关系在 Go语言中是隐式的。两个类型之间的实现关系不需要在代码中显式地表示出来。Go语言中没有类似于 implements 的关键字。 Go编译器将自动在需要的时候检查两个类型之间的实现关系。

接口定义后，需要实现接口，调用方才能正确编译通过并使用接口。接口的实现需要遵循两条规则才能让接口可用。
接口被实现的条件一：接口的方法与实现接口的类型方法格式一致
接口被实现的条件二：接口中所有方法均被实现

一个类型可以实现多个接口
一个接口可以被多个类型实现
四种接口相关的类型转换情形
1:将一个非接口类型值转换为一个接口类型 在这样的转换中 此非接口值得类型必须实现了此接口类型
2:将一个接口值转换为另一个接口类型（前者接口值得类型实现了后者目标接口类型）
3:将一个接口值转换为另一个非接口类型（此非接口类型必须实现了此接口类型）
4:将一个接口值转换为另一个接口类型（前者接口值得类型可以实现了也可以未实现后者目标接口类型）

一个断言类型表达式的语法为i.(T)，其中i是一个接口值，T为一个类型字面表示 类型T可以未任意一个非接口类型 或者任意一个接口类型

一个失败的断言的估值结果为断言类型的0值 看下面示例
如果一个断言失败并且他的可选的第二个结果为呈现则此断言将造成一个恐慌
go语言中使用接口断言将接口转换成另外一个接口 也可以将接口转换成另外的类型 接口的转换在开发中非常常见
使用也非常频繁

将接口转换为其他接口此时可以在两个接口间转换
例如 鸟和猪具有不同的特性 鸟可以飞 猪不能飞 但是两种动物都可以行走

空接口是接口类型的特殊形式 空接口没有任何的方法 因此任何类型都无须实现空接口 从实现的角度看 任何值都满足这个接口的需求
因此空接口类型可以保存任何值 也可以从空接口中取出原值
空接口的内部实现保存了对象的类型和指针 使用空接口保存一个数据的过程会比直接用数据对应类型的变量
保存稍慢 因此在开发中 应在需要的地方使用空接口 而不是在所有地方使用空接口

go语言类型分支 type-switch 流程控制的语法可以被看作是类型断言的增强版
格式
switch 接口变量.(type) {
    case 类型1:
        // 变量是类型1时的处理
    case 类型2:
        // 变量是类型2时的处理
    …
    default:
        // 变量不是所有case中列举的类型时的处理
}



接口与动态类型 因为go语言不像其他语言一样 实现接口要显示的声明 更加简便的多态实现
接受一个或多个接口类型作为参数的函数  可以被实现了该接口的类型实例调用 实现了某个接口的类型可以被传给任何一次接口作为参数的流函数


go语言排序（借助sort.interface接口）

使用空接口实现可以保存任意值得字典

在go语言中 接口与接口间也可以通过嵌套创造出新的接口

一个接口可以包含一个或多个其他的接口 这相当于直接将这些内嵌的方法列举在外层接口中一样
只要接口的所在方法被实现 则这个接口的所有嵌套接口的方法均可以被调用 代码示例如下


*/
func main() {
	s := &Socket{}
	var w Writer
	w = s
	w.Write([]byte{1, 2, 3})

	var x interface{} = 123
	n, ok := x.(int)
	fmt.Println(n, ok)
	a, ok := x.(float64)
	fmt.Println(a, ok) // 0 false
	//a = x.(float64)    // 将产生一个恐慌

	//将值保存到空接口中
	var any interface{}
	any = 1
	fmt.Println(any)
	//从空接口获取值
	var ao int = 1
	var ii interface{} = ao
	//var b int = ii 报错 cannot use i (type interface {}) as type int in assignment: need type assertion
	/*
		空接口类型类似于java中的object 断言就像强制类型转换
	*/
	var b int = ii.(int)
	fmt.Println(b)
	//空接口的值比较 空接口在保存不同的值后 可以喝其他变量值一样使用 == 进行比较操作 空接口的比较有一下几种特性
	//1 类型不同的空接口间的比较结果不相同 2 不能比较空接口中的动态值

	//go的动态接口
	bi := &brid{}
	DuckDance(bi)

}

//类似于java中的ToString
type Stringer interface {
	String() string
}

type Socket struct {
}

func (s *Socket) Write(p []byte) (n int, err error) {
	fmt.Println("Writer", p)
	return 0, nil
}
func (s *Socket) Close() error {
	return nil
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}

//定义字典结构 键值对都是空接口类型
type Dictionary struct {
	data map[interface{}]interface{}
}

//根据键获取值
func (d *Dictionary) Get(key interface{}) interface{} {
	return d.data[key]
}

//设置键值对
func (d *Dictionary) Set(key interface{}, value interface{}) {
	d.data[key] = value
}

type IDuck interface {
	Quack()
	Walk()
}

func DuckDance(duck IDuck) {
	for i := 1; i <= 3; i++ {
		duck.Quack()
		duck.Walk()
	}
}

type brid struct {
}

func (b *brid) Quack() {
	fmt.Println("I am quacking!")
}
func (b brid) Walk() {
	fmt.Println("I am walking!")
}
