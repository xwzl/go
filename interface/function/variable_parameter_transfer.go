package main

import "fmt"

/**
在多个可变参数函数中传递参数

可变参数变量是一个包含所有参数的切片，如果要在多个可变参数中传递参数，可以在传递时在可变参数变量中默认添加...，将切片中的
元素进行传递，而不是传递可变参数变量本身。

下面的例子模拟 print() 函数及实际调用的 rawPrint() 函数，两个函数都拥有可变参数，需要将参数从 print 传递到 rawPrint 中。
*/
func main() {
	print(1, 2, 3)
}

// 实际打印的函数
func rawPrint(rawList ...interface{}) {
	// 遍历可变参数切片
	for _, a := range rawList {
		// 打印参数
		fmt.Println(a)
	}
}

// 打印函数封装
// 此时，slist（类型为 []interface{}）将被作为一个整体传入 rawPrint()，rawPrint() 函数中遍历的变量也就是 slist 的切片值。
// 可变参数使用...进行传递与切片间使用 append 连接是同一个特性。
func print(slist ...interface{}) {
	// 将slist可变参数切片完整传递给下一个函数 ... 表示可变参数 ,不加传递的是切片
	rawPrint(slist...)
}
