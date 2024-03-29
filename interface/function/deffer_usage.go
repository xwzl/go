package main

import (
	"fmt"
	"os"
	"sync"
)

var (
	// 一个演示用的映射
	valueByKey = make(map[string]int)
	// 保证使用映射时的并发安全的互斥锁
	valueByKeyGuard sync.Mutex
)

// 根据键读取值
//func readValue(key string) int {
//	// 对共享资源加锁
//	valueByKeyGuard.Lock()
//	// 取值
//	v := valueByKey[key]
//	// 对共享资源解锁
//	valueByKeyGuard.Unlock()
//	// 返回值
//	return v
//}

// 上面的代码中第 6~8 行是对前面代码的修改和添加的代码，代码说明如下：
//
// 		第 6 行在互斥量加锁后，使用 defer 语句添加解锁，该语句不会马上执行，而是等 readValue() 返回时才会被执行。
// 		第 8 行，从 map 查询值并返回的过程中，与不使用互斥量的写法一样，对比上面的代码，这种写法更简单。
func readValue(key string) int {
	valueByKeyGuard.Lock()

	// defer后面的语句不会马上调用, 而是延迟到函数结束时调用
	defer valueByKeyGuard.Unlock()
	return valueByKey[key]
}

/**
Go语言中关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数（为什么要在返回之后才执行这些语句？
因为 return 语句同样可以包含一些操作，而不是单纯地返回某个值）。

关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源。典型的例子就是对一个互斥解锁，或者关闭
一个文件。
*/
func main() {
	//multipleDeffer()
}

/**
结果分析如下：

	代码的延迟顺序与最终的执行顺序是反向的。
	延迟调用是在 defer 所在函数结束时进行，函数结束可以是正常返回时，也可以是发生宕机时。
*/
func multipleDeffer() {
	fmt.Println("defer begin")
	// 将defer放入延迟调用栈
	defer fmt.Println(1)
	defer fmt.Println(2)
	// 最后一个放入, 位于栈顶, 最先调用
	defer fmt.Println(3)
	fmt.Println("defer end")
}

/**
使用延迟释放文件句柄,根据文件名查询其大小

文件的操作需要经过打开文件、获取和操作文件资源、关闭资源几个过程，如果在操作完毕后不关闭文件资源，进程将一直无法释放文件资源。
在下面的例子中将实现根据文件名获取文件大小的函数，函数中需要打开文件、获取文件大小和关闭文件等操作。由于每一步系统操作都需要进
行错误处理，而每一步处理都会造成一次可能的退出，因此就需要在退出时释放资源，而我们需要密切关注在函数退出处正确地释放文件资源。
*/
func fileSize(filename string) int64 {
	// 根据文件名打开文件, 返回文件句柄和错误
	f, err := os.Open(filename)
	// 如果打开时发生错误, 返回文件大小为0
	if err != nil {
		return 0
	}
	// 取文件状态信息
	info, err := f.Stat()

	// 如果获取信息时发生错误, 关闭文件并返回文件大小为0
	if err != nil {
		f.Close()
		return 0
	}
	// 取文件大小
	size := info.Size()
	// 关闭文件
	f.Close()

	// 返回文件大小
	return size
}

func fileSizeDeffer(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		return 0
	}
	// 延迟调用Close, 此时Close不会被调用,不能放在第一行后面，如果打开文件出现错误，f 值为 null
	// 在延迟语句触发时，将触发宕机错误。
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		// defer机制触发, 调用Close关闭文件
		return 0
	}
	size := info.Size()
	// defer机制触发, 调用Close关闭文件
	return size
}
