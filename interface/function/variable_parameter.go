package main

import "fmt"

/**
在C语言时代大家一般都用过 printf() 函数，从那个时候开始其实已经在感受可变参数的魅力和价值。如同C语言中的 printf() 函数，
Go语言标准库中的 fmt.Println() 等函数的实现也严重依赖于语言的可变参数功能。

本节我们将介绍可变参数的用法。合适地使用可变参数，可以让代码简单易用，尤其是输入输出类函数，比如日志函数等。

可变参数类型

可变参数是指函数传入的参数个数为不定数量。为了做到这点，首先需要将函数定义为接受可变参数类型：

	func myfunc(args ...int) {
		for _, arg := range args {
			fmt.Println(arg)
		}
	}

上面这段代码的意思是，函数 myfunc() 接受不定数量的参数，这些参数的类型全部是 int，所以它可以用如下方式调用：

	myfunc(2, 3, 4)
	myfunc(1, 3, 7, 13)

形如...type格式的类型只能作为函数的参数类型存在，并且必须是最后一个参数。它是一个语法糖（syntactic sugar），即这种语法对语
言的功能并没有影响，但是更方便程序员使用。通常来说，使用语法糖能够增加程序的可读性，从而减少程序出错的机会。

从内部实现机理上来说，类型...type本质上是一个数组切片，也就是[]type，这也是为什么上面的参数 args 可以用 for 循环来获得每个
传入的参数。

假如没有...type这样的语法糖，开发者将不得不这么写：

	func myfunc2(args []int) {
		for _, arg := range args {
			fmt.Println(arg)
		}
	}

从函数的实现角度来看，这没有任何影响，该怎么写就怎么写。但从调用方来说，情形则完全不同：

	myfunc2([]int{1, 3, 7, 13})

会发现，我们不得不加上[]int{}来构造一个数组切片实例。但是有了...type这个语法糖，我们就不用自己来处理了。

可变参数的传递

假设有另一个变参函数叫做 myfunc3(args ...int)，下面的例子演示了如何向其传递变参：

	func myfunc(args ...int) {
		// 按原样传递
		myfunc3(args...)
		// 传递片段，实际上任意的int slice都可以传进去
		myfunc3(args[1:]...)
	}

任意类型的可变参数

之前的例子中将可变参数类型约束为 int，如果你希望传任意类型，可以指定类型为 interface{}。下面是Go语言标准库中 fmt.Printf() 的函数原型：

	func Printf(format string, args ...interface{}) {
		// ...
	}
*/
func main() {
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	MyPrintf(v1, v2, v3, v4)
}

// 获取参数类型
func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}
