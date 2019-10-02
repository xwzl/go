package main

import "fmt"

// 本节将通过示例来演示一下如何使用闭包将函数作为返回值
func main() {
	addReturn := addReturn()
	// 3 调用的 addReturn 返回的闭包变量
	fmt.Printf("Call Add2 for 3 gives: %v\n", addReturn(3))
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))

	var f = AdderPlus()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))

	go func(a int) {
		fmt.Println(a)
	}(100)

	fmt.Println(1)
}

func addReturn() func(a int) int {
	return func(a int) int {
		return a + 1
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

/**
三次调用函数 f 的过程中函数 Adder() 中变量 delta 的值分别为：1、20 和 300。

我们可以看到，在多次调用中，变量 x 的值是被保留的，即 0 + 1 = 1 ，然后 1 + 20 = 21，最后 21 + 300 = 321：闭包函数保存并
积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量。

这些例子清楚地展示了如何在Go语言中使用闭包。在闭包中使用到的变量可以是在闭包函数体内声明的，也可以是在外部函数声明的：

	var g int
	go func(i int) {
		s := 0
		for j := 0; j < i; j++ { s += j }
		g = s
	}(1000) // Passes argument 1000 to the function literal.

这样闭包函数就能够被应用到整个集合的元素上，并修改它们的值。然后这些变量就可以用于表示或计算全局或平均值。
*/
func AdderPlus() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
