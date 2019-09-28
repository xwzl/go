package main

import "fmt"

/**
在对普通变量使用&操作符取地址获得这个变量的指针后，可以对指针使用*操作，也就是指针取值
取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。

变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
	对变量进行取地址（&）操作，可以获得这个变量的指针变量。
	指针变量的值是指针地址。
	对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
*/
func main() {

	// 准备一个字符串并赋值
	var house = "Malibu Point 10880, 90265"

	// 对字符串取地址，将指针保存到 ptr 中, ptr类型为*string,指针地址对应指定的具体类型加 *
	ptr := &house
	// 打印 ptr 变量的类型，类型为 *string
	fmt.Printf("ptr type: %T\n", ptr)
	// 打印 ptr 的指针地址，每次运行都会发生变化
	fmt.Printf("address: %p\n", ptr)

	// 对 ptr 指针变量进行取值操作，value 变量类型为 string
	value := *ptr
	// 取值后的类型
	fmt.Printf("value type: %T\n", value)
	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)
}
