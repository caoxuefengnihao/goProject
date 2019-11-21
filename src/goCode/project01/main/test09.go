package main

import "fmt"

func main() {
	f := &file{}
	var writer DataWriter
	writer = f
	writer.WriteData("data")

}

// 定义一个数据写入器
type DataWriter interface {
	WriteData(data interface{}) error
}

//定义文件结构 用于实心DataWrite
type file struct {
}

func (f *file) WriteData(data interface{}) error {
	fmt.Println("WriterData", data)
	return nil
}
