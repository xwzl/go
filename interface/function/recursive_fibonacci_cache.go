package main

import (
	"fmt"
	"time"
)

const LIM = 100

var fibs [LIM]uint64

func main() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacciPlus(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func fibonacciPlus(n int) (res uint64) {
	// 记忆化：检查数组中是否已知斐波那契（n）
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacciPlus(n-1) + fibonacciPlus(n-2)
	}
	fibs[n] = res
	return
}
