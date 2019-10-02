package main

import "fmt"

/**
提示：在使用递归时，需要设置退出条件，否则递归将陷入无限循环中。

许多问题都可以使用优雅的递归来解决，比如说著名的快速排序算法。

在使用递归函数时经常会遇到的一个重要问题就是栈溢出：一般出现在大量的递归调用导致的程序栈内存分配耗尽。这个问题可以通过一个名为懒
惰求值的技术解决，在Go语言中，我们可以使用管道（channel）和 goroutine 来实现。

构成递归需具备的条件：

	子问题须与原始问题为同样的事，且更为简单。
	不能无限制地调用本身，须有个出口，化简为非递归状况处理。

Go语言中也可以使用相互调用的递归函数：多个函数之间相互调用形成闭环。因为Go语言编译器的特殊性，这些函数的声明顺序可以是任意的。下面
这个简单的例子展示了函数 odd 和 even 之间的相互调用。
*/
func main() {
	fmt.Printf("%d is even: is %t\n", 16, even(16)) // 16 is even: is true
	fmt.Printf("%d is odd: is %t\n", 17, odd(17))
	// 17 is odd: is true
	fmt.Printf("%d is odd: is %t\n", 18, even(18))
	// 18 is odd: is false
}

// 偶数
func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr) - 1)
}

// 奇数
func odd(nr int) bool { // 15
	if nr == 0 {
		return false
	}
	return even(RevSign(nr) - 1)
}

// 返回一个正数
func RevSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}
