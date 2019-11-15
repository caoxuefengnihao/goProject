package main

/**
go 语言数组
数组是一个有固定长度的特定类型元素组成的序列
学习内容 数组 切片 映射 列表的增加 修改 删除
切片可以增加和收缩 的动态序列
数组的声明语法如下
var 数组变量名[元素数量]Type

在数组的定义中 如果在数组长度的位置出现 ... 省略号 则表示数组的长度是
根据初始化的个数来计算
比较两个数组是否相等 如果两个数组类型相同（包括类型和长度） 可以直接通过比较运算符 == ！=

从数组或切片生成新的切片拥有以下特性
取出的元素数量为 结束为止-开始位置
取出的元素不包括结束位置对应的索引 切片最后一个元素使用slice[len(slice)] 获取
当缺省开始位置是 表示从连续区域开头到 结束位置
当缺省结束位置时 表示冲开始位置到整个连续区域末尾
两者同时缺省时 与切片本身等效
两者同时为0时 等效于空切片 一般用于切片复位
声明新的切片的格式如下
var name []type 之后可以使用append（）函数想切片中添加元素
如果需要动态的创建一个切片 可以使用make（）方法
两者区别
使用make方法生成的切片一定发生了内存的分配操作 但给定开始于结束标志的
的切片只是将新的切片结构指向已经分配好的内存区域不会发生内存分配操作
Go语言中删除切片元素的本质是，以被删除元素为分界点，将前后两个部分的内存重新连接起来。

go语言map
声明方式
var mapname map[keyType]valueType
未初始化的 map 的值是 nil
map 的删除和清空 delete（map，键）方法
在go语言中并没有为map提供任何清空所有元素的函数方法 清空map的唯一办法就是
重新make一个新的map
sync.map
go语言中的map在并发情况下 只读是线程安全的 同时读写是线程不安全的
特性  无需初始化 直接声明即可
sycn.Map不能使用map的方式进行取值和设置等操作 而是使用sync.Map 的方法进行调用
Store表示存储 load 表示获取 Delete 表示删除
使用Range配合一个回调函数进行遍历操作 通过回调函数返回内部遍历出来的值
Range 参数中回调函数的返回值在需要继续迭代是 返回true 终止遍历时 返回 false
在sync.Map 中 没有获取数量的方法  只能执行计算数量

go语言 list
在Go语言中，列表使用 container/list 包来实现，内部的实现原理是双链表，列表能够高效地进行任意位置的元素插入和删除操作。
初始化链表  有两种方法 分别是使用New（）  和 var 关键字声明 两种方法的初始化效果都是一致的
链表与切片 map 不同的是 列表没有具体元素类型的限制 因此列表的元素可以是任意类型
的元素， 这即带来了便利 也引来了一些问题
双向链表支持从队列的前方或后方插入元素 分别对应的方法是PushFront PushBack



go语言中make 与 new 关键字的区别以及实现原理
make关键字的主要作用是创建切片 哈希表 和 channel 等内置的数据结构 而new的主要作用是为
类型申请一片内存空间 并返回指向这片内存的指针


*/
import (
	"container/list"
	"fmt"
	"sync"
)

func main() {
	fmt.Println("a")
	fmt.Println("dd")

	for y := 1; y <= 9; y++ {
		for x := 1; x <= y; x++ {
			//在这里 我们学一下字符串的格式化输出
			//格式化样式 格式化动词以%开头
			//参数列表 多个参数以逗号分隔，个数必须与格式化样式中的个数一一对应 否则运行时会报错
			fmt.Printf("%d*%d = %d ", x, y, x*y)
		}
		fmt.Println()
	}

	var a []int
	//打印索引和元素

	for i := 0; i < 10; i++ {
		a = append(a, i+10)
	}
	for _, e := range a {
		fmt.Println(e)
	}

	var highRiseBuilding []int
	for i := 0; i < 31; i++ {
		highRiseBuilding = append(highRiseBuilding, i+10)
	}
	fmt.Println(highRiseBuilding[10:15])
	//重置切片 清空拥有的元素
	fmt.Println(highRiseBuilding[0:0])
	//声明新的切片
	mapname := make(map[string]int)
	mapname["1"] = 1
	mapname["2"] = 2
	for k, v := range mapname {
		fmt.Printf("%d,%d\n", k, v)
	}
	//删除map中
	delete(mapname, "1")
	for k, v := range mapname {
		fmt.Printf("%d,%d\n", k, v)
	}
	//清空map中的所有元素  无

	//异步map
	var syncmap sync.Map
	//将键值对保存到sync.Map
	syncmap.Store("green", 97)

	var listname list.List
	listname1 := list.New()
	listname.PushBack("vvv")
	listname1.PushBack("gagggg")
	fmt.Println(listname1)
	//遍历双向链表 需要配合Front（）函数获取头元素 遍历时 只要元素不为空就可以继续进行
	//每一次遍历都会调用元素的Next()函数
	for i := listname1.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	for i := listname.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
