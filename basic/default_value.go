package main

import "fmt"

/**
变量在定义时没有明确的初始化时会赋值为_零值_。
零值是：
数值类型为 `0`，
布尔类型为 `false`，
字符串为 `""`（空字符串）。
*/
func main() {
	defaultValue()
}

func defaultValue() {
	var i int
	var f float64
	var b bool
	var s string
	// %v 不能替换字符串，字符串替换使用 %q
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
