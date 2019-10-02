package main

import "fmt"

func main() {
	func(a int) {
		fmt.Printf("闭包调用：%v \n", a)
	}(100)

	a := 100

	f := func(a int) {
		fmt.Printf("闭包动态变量调用：%v \n", a)
	}

	f(a)
}
