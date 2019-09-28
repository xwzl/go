package main

import (
	"fmt"
	"math"
)

/**
接口类型是由一组方法定义的集合,接口类型的值可以存放实现这些方法的任何值
*/
type user interface {
	printInfo() float64
}

type Name float64

func (name Name) printInfo() float64 {
	if name < 0 {
		return float64(-name)
	}
	return float64(name)
}

type Fix struct {
	x, y float64
}

func (fix *Fix) printInfo() float64 {
	return fix.x + fix.y
}

func main() {

	var users user
	name := Name(-math.Sqrt2)
	fix := Fix{12.0, 13.0}

	users = name
	fmt.Println(users.printInfo())

	// 指针执行接口的实现
	users = &fix
	fmt.Println(users.printInfo())
}
