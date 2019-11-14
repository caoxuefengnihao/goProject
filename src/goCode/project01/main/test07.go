package main

/*
在本例中，使用一个数值表示时间中的“秒”值，
然后使用 resolveTime() 函数将传入的秒数转换为天、小时和分钟等时间单位。

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

	day, hour, minute := resolveTime(1000)
	fmt.Printf("%d,%d,%d", day, hour, minute)

}

func resolveTime(seconds int) (day, hour, minute int) {
	day = seconds / SecondsPerDay
	hour = seconds / SecondsPerHour
	minute = seconds / SecondsPerMinute

	return day, hour, minute
}
