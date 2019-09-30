package main

import "sync"

/**
正确的写法是使用函数参数做一次数据复制，而不是闭包。

可以看到新程序的运行结果符合预期。这个不能说是缺陷，而是Go语言设计者为了性能而选择的一种设计方案，因为大多情况下 for
循环块里的代码是在同一个 goroutine 里运行的，为了避免空间的浪费和 GC 的压力，复用了 range 迭代临时变量。语言使用者
明白这个规约，在 for 循环下调用并发时要复制造代变量后再使用，不要直接引用 for 迭代变量。
 */
func main() {
	//rangeError()
	wg := sync.WaitGroup{}
	si := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range si {
		wg.Add(1)
		//这里有一个实参到形参的值拷贝
		go func(a int) {
			println(a)
			wg.Done()
		}(i)
	}
	wg.Wait ()
}

/**
程序结果并没有如我们预期一样遍历切片白，而是全部打印 9，有两点原因导致这个问题：

	for range 下的迭代变量 i 的值是共用的。
	main 函数所在的 goroutine 和后续启动的 goroutines 存在竞争关系。

使用 go run -race 来看一下数据竞争情况：

可以看到 Goroutine 13 和 main goroutine 存在数据竞争，更进一步证实了 range 共享临时变量。range 在迭代写的过程中，
多个 goroutine 并发地去读。
*/
func rangeError() {
	wg := sync.WaitGroup{}
	si := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range si {
		wg.Add(i)
		go func() {
			println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}
