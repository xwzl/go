package main

import "fmt"

/**
for range 结构是Go语言特有的一种的迭代结构，在许多情况下都非常有用。for range 可以迭代任何一个集合（包括数组和 map）。语法上很类似其
它语言中 foreach 语句，但依旧可以获得每次迭代所对应的索引。一般形式为：

	for ix, val := range coll { } 。

需要要注意的是， val 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值（如果 val 为
指针，则会产生指针的拷贝，依旧可以修改集合中的原值）。一个字符串是 Unicode 编码的字符（或称之为 rune ）集合，因此也可以用它迭代字符串：

	for pos, char := range str {
		...
	}

每个 rune 字符和索引在 for-range 循环中是一一对应的。它能够自动根据 UTF-8 规则识别 Unicode 编码的字符。

Go语言中可以使用 for range 遍历数组、切片、字符串、map 及通道（channel）。通过 for range 遍历的返回值有一定的规律：

	数组、切片、字符串返回索引和值。
	map 返回键和值。
	通道（channel）只返回通道内的值。

我们总结一下 for 的功能：

	Go语言的 for 包含初始化语句、条件表达式、结束语句，这 3 个部分均可缺省。
	for range 支持对数组、切片、字符串、map、通道进行遍历操作。
	在需要时，可以使用匿名变量对 for range 的变量进行选取。
*/
func main() {

	//rangArray()
	//rangeChar()
	//rangeMap()
	//rangeChannel()
	//rangeValue()
}

/**
在遍历中选择希望获得的变量

在使用 for range 循环遍历某个对象时，一般不会同时需要 key 或者 value，这个时候可以采用一些技巧，让代码变得更简单。
下面将前面的例子修改一下，参考下面的代码示例：

在例子中将 key 变成了下划线，那么这里的下划线就是匿名变量。什么是匿名变量？

	可以理解为一种占位符。
	本身这种变量不会进行空间分配，也不会占用一个变量的名字。
	在 for range 可以对 key 使用匿名变量，也可以对 value 使用匿名变量。
*/
func rangeValue() {
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for _, value := range m {
		fmt.Println(value)
	}
}

/**
遍历通道（channel）——接收通道数据

for range 可以遍历通道（channel），但是通道在遍历时，只输出一个值，即管道内的类型对应的数据。


*/
func rangeChannel() {
	// 创建了一个整型类型的通道
	c := make(chan int)
	// 启动了一个 goroutine，其逻辑的实现体现在第 5～8 行，实现功能是往通道中推送数据 1、2、3，然后结束并关闭通道
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	// 使用 for range 对通道 c 进行遍历，其实就是不断地从通道中取数据，直到通道被关闭
	for v := range c {
		fmt.Println(v)
	}
}

/**
遍历map——获得map的键和值

对于 map 类型来说，for range 遍历时，key 和 value 分别代表 map 的索引键 key 和索引对应的值，一般被称为 map 的键值对，
因为它们总是一对一对的出现。下面的代码演示了如何遍历 map。

注意

对 map 遍历时，遍历输出的键值是无序的，如果需要有序的键值对输出，需要对结果进行排序。

*/
func rangeMap() {
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
}

/**
遍历字符串——获得字符

Go语言和其他语言类似，可以通过 for range 的组合，对字符串进行遍历，遍历时，key 和 value 分别代表字符串的索引（base0）和字符串中的每一个字符。

代码中的 v 变量，实际类型是 rune，实际上就是 int32，以十六进制打印出来就是字符的编码。
*/
func rangeChar() {
	var str = "hello 你好"
	for key, value := range str {
		fmt.Printf("key:%d value:0x%x %c\n", key, value, value)
	}
}

/**
在遍历代码中，key 和 value 分别代表切片的下标及下标对应的值。下面的代码展示如何遍历切片，数组也是类似的遍历方法：
*/
func rangArray() {
	for key, value := range []int{1, 2, 3, 4, 5} {
		fmt.Printf("key %d,vlaue %d\n", key, value)
	}
}
