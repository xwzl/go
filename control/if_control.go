package main

import (
	"fmt"
	"math"
)

// if 语句除了没有了 `( )` 之外（甚至强制不能使用它们），看起来跟 C 或者 Java 中的一样，而 `{ }` 是必须的。
func main() {
	i := math.Round(10)
	// go 中 for 与 if 关键字判断条件中不能使用 `()`划分边界，除此之外和 Java 没有什么区别
	if i < 5.0 {
		fmt.Printf("%v 小于 5.0 \n", i)
	}
	c := 2i
	fmt.Println(c * c)

	fmt.Println(sqrt(2), sqrt(-4))
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

/**
if 的便捷语句
跟 for 一样，`if` 语句可以在条件之前执行一个简单的语句。
由这个语句定义的变量的作用域仅在 if 范围之内。
（在最后的 return 语句处使用 v 看看。）
*/
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
