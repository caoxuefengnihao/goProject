package main

/*
练习三 日志系统
*/
import (
	"errors"
	"fmt"
	"os"
)

func main() {
	logger := NewLogger()
	writer := newfileWriter()
	writer.SetFile("log.log")
	logger.RegisterWriter(writer)
	logger.Log("hello\nhello")
}

/*
首先定义一个日志堆外接口
本例中定义一个日志写入器接口（LogWrite） 要求写入设备必须遵循这个接口协议才能被日志器(Logger)注册日志器里
有一个写入器的注册方法（Logger的RegisterWriter（）方法）
*/
type LogWriter interface {
	write(data interface{}) error
}

/**
日志器 里面包含很多的日志写入器
并且实现有一个注册方法
*/
type Logger struct {
	writerList []LogWriter
}

/*
注册日志写入器的方法
*/
func (log *Logger) RegisterWriter(write LogWriter) {
	log.writerList = append(log.writerList, write)
}

/*
将一个data类型的数据写入日志
*/
func (l *Logger) Log(data interface{}) {
	for _, writer := range l.writerList {
		writer.write(data)
	}
}

/*
定义一个Logger生成器 相当于java中的构造方法
*/
func NewLogger() *Logger {
	return &Logger{}
}

/*
接下来定义几个文件写入器 首先是文件写入器
声明文件写入器
*/
type fileWriter struct {
	file *os.File
}

/*
设置文件写入器的文件名
*/
func (f *fileWriter) SetFile(name string) (err error) {
	//如果文件已经打开关闭前一个文件
	if f.file != nil {
		f.file.Close()
	}
	//创建一个文件并保存句柄
	file, err := os.Create(name)
	f.file = file
	return err
}

//实现 LogWriter 的 Write方法 遵守接口协议

func (fi *fileWriter) write(data interface{}) (err error) {
	// 日志文件可能没有创建成功
	if fi.file == nil {
		// 日志文件没有准备好
		return errors.New("file not created")
	}
	// 将数据序列化为字符串
	str := fmt.Sprintf("%v\n", data)
	// 将数据以字节数组写入文件中
	_, err = fi.file.Write([]byte(str))
	return err
}

// 创建命令行写入器实例
func newfileWriter() *fileWriter {
	return &fileWriter{}
}
