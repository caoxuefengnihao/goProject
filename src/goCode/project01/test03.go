package project01

/*
栈内存 与 堆内存
变量逃逸分析 -- 自动决定变量的分配方式 提高运行效率
通过编译器分析代码的特征与代码的生命周期 决定应该使用堆还是栈来进行内存分配

*/
import "fmt"

func dummy(b int) int {
	var c int
	c = b
	return c
}

// 空函数 什么也不做
func void() {

}
func main() {
	var a int
	void()
	fmt.Println(a, dummy(0))
}
