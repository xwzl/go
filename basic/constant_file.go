package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

/**
常量的定义与变量类似，只不过使用 const 关键字。
常量可以是字符、字符串、布尔或数字类型的值。
常量不能使用 := 语法定义。
*/
func main() {
	// 常量首字母大写,常量和 Java 一样不可改变
	const Fix = "这是一个世界"
	// 编译不通过
	//const Six := 6
	fmt.Println(Fix)

	// 01 -> 1 => 10 -> 2 二进制
	fmt.Println(1 << 1)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

/**
数值常量

数值常量是高精度的 _值_。
一个未指定类型的常量由上下文来决定其类型。
也尝试一下输出 needInt(Big) 吧。
*/
