package main

import "fmt"

type Weapon int

const (
	Arrow Weapon = iota // 开始生成枚举值, 默认为0
	Shu
	SniperRifle
	Rifle
	Blower

	// 时间上 iota 的值还是从上一次自增为5
	FlagNone = 1 << iota
	FlagRed
	FlagGreen
	FlagBlue
)

func main() {
	// 输出所有枚举值
	fmt.Println(Arrow, Shu, SniperRifle, Rifle, Blower)
	// 使用枚举类型并赋初值
	var weapon Weapon = Blower
	fmt.Println(weapon)
	fmt.Println(FlagNone, FlagRed, FlagGreen, FlagBlue)
}
