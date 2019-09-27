package main

import (
	"fmt"
	"math"
)

func main() {
	// 类似平方的效果
	fmt.Println(math.Pow(2, 10))
	fmt.Println(
		plows(3, 2, 10),
		plows(3, 3, 20),
	)
}

/**
平方在比较值
*/
func plows(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g > %g \n", v, lim)
	}
	return lim
}
