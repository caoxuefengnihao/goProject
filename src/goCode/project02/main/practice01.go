package main

import "math"

/*
示例1 二维矢量模拟玩家移动 在游戏中 一般使用二维矢量保存玩家的位置 使用矢量运算
可以计算出玩家移动的位置
本例子中 首先实现二维矢量对象 接着构造玩家对象 最后使用矢量对象和玩家对象共同模拟玩家移动的过程

*/

func main() {

}

type Vec2 struct {
	X, Y float32
}

//加
func (v1 Vec2) Add(v2 Vec2) *Vec2 {
	return &Vec2{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

//减
func (v1 Vec2) Sub(v2 Vec2) *Vec2 {
	return &Vec2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
}

//乘
func (v1 Vec2) Scale(s float32) *Vec2 {
	return &Vec2{
		X: v1.X * s,
		Y: v1.Y * s,
	}
}

//距离
func (v1 Vec2) DistanctTo(v2 Vec2) float32 {
	dx := v1.X - v2.X
	dy := v1.Y - v2.Y
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}
