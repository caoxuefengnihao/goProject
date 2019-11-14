package project01

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"strconv"
)

var (
	cd int = 10
)

/*
_本身就是一个特殊的标识符，被称为空白标识符，它可以像其他标识符那样用于变量的声明或赋值（任何类型都可以赋值给它），
但任何赋给这个标识符的值都将被抛弃，
因此这些值不能在后续的代码中使用，也不可以使用_作为变量对其它变量进行赋值或运算。
go语言的数据类型
bool
数值类型 int(8 16 32 64) float32 float64
字符串
派生类型  指针类型（Pointer） 数组类型 结构化类型（struct） Channel类型 函数类型 切片类型  接口类型 Map类型

Go语言的字符有以下两种：
一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
另一种是 rune 类型，代表一个 UTF-8 字符，
当需要处理中文、日文或者其他复合字符时，
则需要用到 rune 类型。rune 类型等价于 int32 类型。



go语言变量 关键字 var 可以一次声明多个变量
第一种 指定变量类型 如果没有初始化 则变量默认为0
第二种 根据值自行进行判断变量类型
第三种 省略var 注意 ：= 左侧没有声明新的变量 就餐生编译错误 不能提供数据类型 只能在函数中定义
var（变量） 声明的是全局变量上升到包的变量中 而且全局变量是可以只声明不使用的 批量定义
//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"
go常量定义 关键字 const identifyer [type] = value

go语言中不允许将数字类型转换成bool类型
go语言不存在隐式类型转换 因此所有的类型转换都必须显示的声明
valueOfTypeB= typeB（valueOfTypeA）
只有相同底层类型的变量之间可以进行相互转换（如将 int16 类型转换成 int32 类型），不同底层类型的变量相互转换时会引发编译错误

go语言的预算符  & 返回变量的实际地址 * 指针变量 *a 是一个指针变量

go语言的循环语句  只有for循环 循环控制语句  break continue goto（将控制转移到被标记的语句）

go语言数组 声明数组 需要指定元素类型及土元素个数  var balance [10] float32

go语言指针
类似于变量和常量 在使用指针的前你需要声明指针 格式如下
var var_name *var-type
var a *int 指向整型
指针的使用流程
	定义一个指针变量
	为指针变量赋值
	访问指针变量中的指向地址的值
在指针类型前面加上*号 来获取指针所指向的内容
每一个变量在运行时都会有一个指针地址 用&可以获取内存地址
go 空指针
当一个指针被定义后没有分配到任何变量时 它的值为nil
nil指针也称空指针

除了字符串 文字符号 和字节之间的转换 我们常常也需要在数值和字符串之间进行相互转化
go 语言的strconv 包提供了字符串与数值类型之间的相互转化功能
strconv 包提供了很多的函数 大概分为几类
字符串转int Atoi()
int转字符串 Itoa()
parse类函数将字符串转化为指定的数值类型 ParseBool ParseFloat ParseInt ParseUnit
因为string类型转数字类型可能会失败 所以这些函数都有第二个返回值表示转化是否成功

Format类函数将指定的数值类型转化为string 类型 FormatBool

Append类函数用于将指定的数值类型转化成字符串后追加到一个切片中

当有些类型无法转换时 将会报错并返回的错误信息



*/

func main() {
	var a string = "go"
	fmt.Println(a)
	var b, c int = 1, 2
	var ip *string //声明指针变量
	ip = &a
	fmt.Println(ip) //指针变量的存储地址
	fmt.Println(*ip)
	fmt.Println(b + c)
	fmt.Println(cd)
	fmt.Println(&a)

	// 图片大小
	const size = 300
	// 根据给定大小创建灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))
	// 遍历每个像素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			// 填充为白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}
	// 从0到最大像素生成x坐标
	for x := 0; x < size; x++ {
		// 让sin的值的范围在0~2Pi之间
		s := float64(x) * 2 * math.Pi / size
		// sin的幅度为一半的像素。向下偏移一半像素并翻转6
		y := size/2 - math.Sin(s)*size/2
		// 用黑色绘制sin轨迹
		pic.SetGray(x, int(y), color.Gray{0})
	}
	// 创建文件
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	// 使用png格式将数据写入文件
	png.Encode(file, pic) //将image信息写入文件中
	// 关闭文件
	file.Close()

	i, _ := strconv.Atoi("3")
	fmt.Print(i)
	d := 5.0
	e := int(d)
	fmt.Println(e)

}
