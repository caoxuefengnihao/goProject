package main

/**
range 关键字
*/
import "fmt"

func main() {
	slice := []int{10, 20, 30, 40}
	for k, v := range slice {
		fmt.Printf("%d %d\n", k, v)
	}
	/*
		当迭代切片时 关键字range 会返回两个值 第一个值是当前迭代到的索引位置
		第二个值为对应元素


	*/
}
