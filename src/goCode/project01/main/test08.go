package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

/*
go语言的结构体

使用关键字 type 可以将各种基本类型定义为自定义类型，基本类型包括整型、字符串、布尔等。结构体是一种复合的基本类型，通过 type 定义为自定义类型后，使结构体更便于使用。
type 类型名 struct {
    字段1 字段1类型
    字段2 字段2类型
    …
}
结构体的定义只是一种内存布局的描述 只有当结构体实例化时 才会真正的分配内存 因此必须在定义结构体并实例化后才能使用结构体的字段

基本的实例化形式
结构体本身是一种类型 可以像整型 字符串类型一样 以var的方式声明结构体即可完成实例化
格式：var ins T
创建指针类型的结构体  go语言中还可以使用new 关键字对类型进行实例化 结构体在实例化后会形成指针类型的结构体
格式：ins := new(T)
取结构体的地址实例化 go语言中 对结构体进行& 取地址操作是 视为对该类型进行一次new的实例化操作 取地址格式如下
格式：ins ：=&T{}

其中取地址实例化是最广泛的一种结构体实例化方式 可以使用函数封装上面的初始化过程

结构体在实例化时可以直接对成员变量进行初始化初始化有两种形式 分别是以字段
键值对形式 和 多个值得列表形式 键值对形式的初始化适合选择性填充字段较多的结构体 多格式的
列表形式 适合填充字段较少的结构体

键值对形式初始化结构体
键 ---> 结构体中的一个字段 值 ------> 对应的初始化的值
书写格式
ins := 结构体类名{
	字段一:字段一的值,
	........
}

使用多个值得列表初始化结构体
书写格式
ins := 结构体类名{
	字段1的值，
	字段2的值，
	......
}
需要注意
必须初始化结构体的所有字段
每一个初始值的填充顺序必须与字段在结构体的声明顺序一致
与键值对的形式不能够混用

go语言的类型或结构体没有构造函数的功能
但是我们可以使用结构体初始化的过程来模拟实现构造函数


结构体的方法  ---> 接收器
接收器类型可以是任何类型 但是接收器不能是一个接口类型 也不能是指针类型
类型T上的所有方法的集合 叫做类型T的方法集
因为方法是函数，所以同样的，不允许方法重载，
即对于一个类型只能有一个给定名称的方法，但是如果基于接收器类型，
是有重载的：具有同样名字的方法可以在 2 个或多个不同的接收器类型上存在
，比如在同一个包里这么做是允许的
只是Go语言建立的“接收器”强调方法的作用对象是接收器，也就是类实例，而函数没有作用对象

go语言可以为任何类型添加方法 给一种类型添加方法就好像给结构体添加方法一样 因为结构体也是一种类型

为基本类型添加方法
在go语言中 使用type关键字可以定义出新的自定义类型 之后就可以为自定义类型添加各种方法
为HTTP添加方法
go语言提供的HTTP包里也大量使用了类型方法 go语言使用HTTP包进行HTTP的请求
使用HTTP包的NewRequest方法 可以创建一个HTTP请求 填充请求中的HTTP头 在调用 HTTP.client的Do方法
将传入的HTTp请求发送出去
time包中的类型方法
go语言提供的time包主要用于时间的获取和计算等，在这个包中，也使用了类型方法

go语言可以将类型的方法与普通的方法视为一个概念 从而简化方法和函数混为合作为回调类型时的复杂性
调用者无需关心谁来支持调用 系统会自动处理是否调用普通函数或类型的方法



*/
func main() {
	var p Point
	p.X = 0
	p.Y = 1
	fmt.Println(p.Y)
	p1 := new(Point)
	p1.X = 3
	p1.Y = 2
	play1 := new(Player)
	play1.Name = "Canon"
	play1.HealthPoint = 300

	var version int = 1
	cmd := &Command{}
	cmd.Name = "version"
	cmd.Comment = "show version"
	cmd.Var = &version

	cat := NewCatName("zhagnsan")
	fmt.Println(cat.Name)

	var b bag
	b.item = append(b.item, 100)
	Insert02(&b, 200)
	fmt.Println(b)

	b2 := Insert(b, 300)
	fmt.Println(b2)
	fmt.Println(&b, &b2)

	var name map[int]string
	name = map[int]string{1: "2"}
	fmt.Println(name)

	var i MyInt
	fmt.Println(i.IsZero())
	i = 1
	fmt.Println(i.IsZero())
	fmt.Println(i.Add(5))

	client := &http.Client{}
	//创建一个HTTP请求
	re, err := http.NewRequest("POST", "http://www.163.com", strings.NewReader("key=value"))
	if err != nil {

		fmt.Println(err)
		os.Exit(1)
		return
	}
	re.Header.Add("User-Agent", "myClient")
	resq, err := client.Do(re)
	data, _ := ioutil.ReadAll(resq.Body)
	fmt.Println(string(data))
	defer resq.Body.Close()

}

type Point struct {
	X int
	Y int
}

type Player struct {
	Name        string
	HealthPoint int
	MagicPoint  int
}

type Command struct {
	Name    string
	Var     *int
	Comment string
}
type Cat struct {
	Color string
	Name  string
}

func NewCatName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}
func NewCatColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}

type bag struct {
	item []int
}

//目的是往背包里放元素的一个方法  类似于java类中成员方法
/*
但是这种写法在实际的适用当中 并不是每个人都习惯把操作对象放到首位 一定程度上让代码失去了
一些规范性和描述性 同时 Insert函数与bag没有任何归属概念
改革  就是出现了接收器  看下面代码
每个方法只能有一个接收器
注意
方法接收者为对象的指针与值有什么区别呢？
如果方法接收者为对象的指针，则会修改原对象，
如果方法接收者为对象的值，那么在方法中被操作的是原对象的副本，
不会影响原对象。示例如下：

在计算机中，小对象由于值复制时的速度较快，
所以适合使用非指针接收器，大对象因为复制性能较低，
适合使用指针接收器，在接收器和参数间传递时不进行复制，
只是传递指针
*/
func Insert(b bag, item int) bag {
	return bag{
		item: append(b.item, item),
	}
}
func Insert02(b *bag, item int) {
	b.item = append(b.item, item)
}
func (b *bag) Insert01(item int) {
	b.item = append(b.item, item)
}

type MyInt int

func (m MyInt) IsZero() bool {
	return m == 0
}
func (m MyInt) Add(order int) MyInt {
	return MyInt(order + int(m))
}
