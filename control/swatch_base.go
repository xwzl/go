package main

import "fmt"

/**
基本写法

Go语言改进了 switch 的语法设计，避免人为造成失误。Go语言的 switch 中的每一个 case 与 case 间是独立的代码块，不需要通过 break 语句
跳出当前 case 代码块以避免执行到下一行。
*/
func main() {

	//switchBase()
	//switchMulti()
	//switchExpress()
	//switchInterface()
}

/**
switch 还可用于获得一个接口变量的动态类型。这种类型 switch 使用类型断言的语法，在括号中使用关键字 type 。如果 switch 在表
达式中声明了一个变量，则变量会在每个子句中具有对应的类型。比较符合控制结构语言习惯的方式是在这些 case 里重用一个名字，实际上
是在每个 case 里声名一个新的变量，其具有相同的名字，但是不同的类型。
*/
func switchInterface() {
	var t interface{}
	t = 1
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T", t) // %T 打印任何类型的t
	case bool:
		fmt.Printf("boolean %t\n", t) // t 的类型是 bool
	case int:
		fmt.Printf("integer %d\n", t) // t 的类型是 int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t 的类型是 *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t 的类型是 *int
	}
}

// 条件表达式
func switchExpress() {
	var r int = 11
	switch {
	case r > 10 && r < 20:
		fmt.Println(r)
	}
}

// 一分支多值,当出现多个 case 要放在一起的时候,不同的 case 表达式使用逗号分隔。
func switchMulti() {
	var x = "hello"
	switch x {
	case "hello", "world":
		fmt.Println("one branch multi value")
	}
}

func switchBase() {
	var x = "hello"
	switch x {
	case "switch":
		fmt.Println("switch")
	case "hello":
		fmt.Println("hello switch")
	default:
		fmt.Println("unknown")
	}
}
