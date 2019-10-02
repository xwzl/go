package main

import "fmt"

// 当一个函数在其函数体内调用自身时，则称之为递归。最经典的例子便是计算斐波那契数列，即每个数均为前两个数之和。
func main() {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
