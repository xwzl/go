package main

import (
	"fmt"
	"testing"
)

// 函数作为变量传递
func TestFuncVar(test *testing.T) {
	f := varHello
	i := f(1)
	fmt.Println(i)

	hello(12, func(i int) int {
		return i + 1
	})
}

func varHello(a int) int {
	return a
}

// 这不是 lambda 表达式的写法吗？Java 中传递的接口，这里直接传递一定公民函数
func hello(lambda int, hello func(int) int) {
	fmt.Println(hello(lambda))
}
