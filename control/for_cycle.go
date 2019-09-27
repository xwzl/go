package main

import "fmt"

/**
Go 只有一种循环结构——`for` 循环。
基本的 for 循环除了没有了 `( )` 之外（甚至强制不能使用它们），看起来跟 C 或者 Java 中做的一样，而 `{ }` 是必须的。
*/
func main() {
	var index int = 0

	for i := 0; i < 100; i++ {
		index += i
	}
	fmt.Println(index)

	// 跟 C 或者 Java 中一样，可以让前置、后置语句为空。
	sum := 1

	// 类似与  Java 中 while 循环
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// This is a the death of cycle
	// 如果省略了循环条件，循环就不会结束，因此可以用更简洁地形式表达死循环。
	//for {
	//
	//}
}
