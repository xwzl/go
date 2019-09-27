package main

import (
	"fmt"
	"math"
)

// 定义全局变量
var p, q int = 1, 2

// 主程序
func main() {
	// hello world
	fmt.Println("Hello World")

	// 生成随机数
	fmt.Println("This is a random ", math.Round(10))

	//  格式化字符串输出，换行暂时未知
	fmt.Printf("Now you have %g problems", math.Nextafter(2, 3))

	fmt.Println(math.Pi)

	// 函数
	fmt.Println(add(11, 12))

	// 参数缩写
	fmt.Println(addAbridge(11, 11))

	// 多值返回
	a, b := multivaluedReturn("value1", 1)
	fmt.Println(a, b)

	// 返回值可以命名，就像入参一样
	x, y := returnValueName(21)
	fmt.Println(x, y)

	// var 类型
	fmt.Println(p, q)

	// var 类型参数测试
	// 变量定义可以包含初始值，每个变量对应一个。
	// 如果初始化是使用表达式，则可以省略类型；变量从初始值中获得类型。
	var x1, x2, x3 = true, false, "test"
	fmt.Println(x1, x2, x3)

	// 短声明变量
	// 在函数中,`：=`间接赋值语句在明确的地方，可以用于替代 var 定义。
	// 函数外的每个语句都必须以关键字开始（`var`、`func`、等等），`:=` 结构不能使用在函数外。
	var temp1, temp2 int = 1, 2
	temp3 := 3
	fmt.Println(temp1, temp2, temp3, ":= 不能不使用在函数外")
}

/**
Go 的返回值可以被命名，并且像变量那样使用。
返回值的名称应当具有一定的意义，可以作为文档使用。
没有参数的 return 语句返回结果的当前值。也就是`直接`返回。
直接返回语句仅应当用在像下面这样的短函数中。在长的函数中它们会影响代码的可读性。
*/
func returnValueName(sum int) (x, y int) {
	x = sum / 3
	y = sum - 3
	//return x,y
	// 可以没有返回值
	return
}

/**
多值返回，函数可以返回任意数量的返回值
*/
func multivaluedReturn(str1 string, str2 int) (string, int) {
	return str1, str2
}

/**
当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。
*/
func addAbridge(x, y int) int {
	return x + y
}

/**
函数调用，类似 ts 写法 返回值写在最后
*/
func add(x int, y int) int {
	return x + y
}

/**
导出名
在导入了一个包之后，就可以用其导出的名称来调用它。
在 Go 中，首字母大写的名称是被导出的。
Foo 和 FOO 都是被导出的名称。名称 foo 是不会被导出的。
执行代码。然后将 math.pi 改名为 math.Pi 再试着执行一下。
*/
func Export() {
	fmt.Println("这是会被导出")
}
