package main

import "fmt"

// 函数作为变量使用
func main() {

	var f func()

	f = say

	sayPlus(f)

}

func say() {
	fmt.Println("函数作为变量被调用")
}

func sayPlus(say func()) {
	say()
}
