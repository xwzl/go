package main

import "fmt"

/**
Go语言的匿名函数就是一个闭包，闭包是可以包含自由（未绑定到特定对象）变量的代码块，这些变量不在这个代码块内或者任何全局上下文中定义，
而是在定义代码块的环境中定义。要执行的代码块（由于自由变量包含在代码块中，所以这些自由变量以及它们引用的对象没有被释放）为自由变量提
供绑定的计算环境（作用域）。

闭包的价值在于可以作为函数对象或者匿名函数，对于类型系统而言，这意味着不仅要表示数据还要表示代码。支持闭包的多数语言都将函数作为第一
级对象，就是说这些函数可以存储到变量中作为参数传递给其他函数，最重要的是能够被函数动态创建和返回。

Go语言中的闭包同样也会引用到函数外的变量。闭包的实现确保只要闭包还被使用，那么被闭包引用的变量会一直存在。因此，简单的说：

	函数 + 引用环境 = 闭包

同一个函数与不同引用环境组合，可以形成不同的实例，如下图所示。

一个函数类型就像结构体一样，可以被实例化。函数本身不存储任何信息，只有与引用环境结合后形成的闭包才具有“记忆性”。函数是编译期静态的概念，
而闭包是运行期动态的概念。

其它编程语言中的闭包

闭包（Closure）在某些编程语言中也被称为 Lambda 表达式。

闭包对环境中变量的引用过程，也可以被称为“捕获”，在 C++11 标准中，捕获有两种类型：引用和复制，可以改变引用的原值叫做“引用捕获”，捕获的过程
值被复制到闭包中使用叫做“复制捕获”。

在 Lua 语言中，将被捕获的变量起了一个名字叫做 Upvalue，因为捕获过程总是对闭包上方定义过的自由变量进行引用。

闭包在各种语言中的实现也是不尽相同的。在 Lua 语言中，无论闭包还是函数都属于 Prototype 概念，被捕获的变量以 Upvalue 的形式引用到闭包中。

C++ 与 C# 中为闭包创建了一个类，而被捕获的变量在编译时放到类中的成员中，闭包在访问被捕获的变量时，实际上访问的是闭包隐藏类的成员。
*/
func main() {
	//valueModify()
	//add()
	playerGenerator()
}

func playerGenerator() {
	// 创建一个玩家生成器
	generator := playerGen("high noon")
	// 返回玩家的名字和血量
	name, hp := generator()
	// 打印值
	fmt.Println(name, hp)
}

// 创建一个玩家生成器, 输入名称, 输出生成器
func playerGen(name string) func() (string, int) {
	// 血量一直为150
	hp := 150
	// 返回创建的闭包
	return func() (string, int) {
		// 将变量引用到闭包中
		return name, hp
	}
}

// 被捕获到闭包中的变量让闭包本身拥有了记忆效应，闭包中的逻辑可以修改闭包捕获的变量，变量会跟随闭包生命期一直存在，闭包本身就如同变量一样拥有了记忆效应。
// 对比输出的日志发现 accumulator 与 accumulator2 输出的函数地址不同，因此它们是两个不同的闭包实例。
// 每调用一次 accumulator 都会自动对引用的变量进行累加。
func add() {
	// 创建一个累加器, 初始值为1
	accumulator := Accumulate(1)
	// 累加1并打印
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", accumulator)
	// 创建一个累加器, 初始值为1
	accumulator2 := Accumulate(10)
	// 累加1并打印
	fmt.Println(accumulator2())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", accumulator2)
}

// 提供一个值, 每次调用函数会指定对值进行累加
func Accumulate(value int) func() int {
	// 返回一个闭包
	return func() int {
		// 累加
		value++
		// 返回一个累加值
		return value
	}
}

// 闭包对它作用域上部变量的引用可以进行修改，修改引用的变量就会对变量进行实际修改
func valueModify() {
	str := "hello world"
	foo := func() {
		str = "hello closure"
	}
	foo()
	fmt.Println(str)
}
