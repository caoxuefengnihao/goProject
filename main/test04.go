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
map 的删除和清空 delete（map，键）方法
*/
import (
	"fmt"
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
		highRiseBuilding[i] = i + 1
	}
	fmt.Println(highRiseBuilding[10:15])
	//重置切片 清空拥有的元素
	fmt.Println(highRiseBuilding[0:0])
	//声明新的切片

	var mapname map[string]int
	mapname["1"] = 1
	mapname["2"] = 2
	for k, v := range mapname {
		fmt.Printf("%d,%d\n", k, v)
	}
	//删除map中

}
