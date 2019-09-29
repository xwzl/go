package main

import "fmt"

/**
不过要注意的是，在容量不足的情况下， append 的操作会导致重新分配内存（扩容），可能导致巨大的内存分配和复制数据代价。即使容量足够，
依然需要用 append 函数的返回值来更新切片本身，因为新切片的长度已经发生了变化。
*/
func main() {
	//first()
	//second()
	//third()
	//four()
	a := []int{1, 2, 3}
	a = add(1, 11, a)
	add(1, 12, a)
	fmt.Print(a)
}

func add(x int, value int, array []int) []int {
	// array[:x] 切片 从 0 开始但是不包含 X
	// 第二个参数为切片时追加 .... 字符
	array = append(array[:x], append([]int{value}, array[x:]...)...)
	return array
}

// 由于 append 函数返回新的切片，也就是它支持链式操作。我们可以将多个 append 操作组合起来，实现在切片中间插入元素：
// 每个添加操作中的第二个 append 调用都会创建一个临时切片，并将 a[i:] 的内容复制到新创建的切片中，然后将临时创建的切片再追加到 a[:i] 。
func four() {
	a := []int{1, 2, 3, 4}
	// 先想 a[1:] 切片前面追加一个元素，最后在拼接到 a[:1]
	a = append(a[:1], append([]int{10}, a[1:]...)...)
	fmt.Println(a)
}

// 除了在切片的尾部追加，我们还可以在切片的开头添加元素：
// 在开头一般都会导致内存的重新分配，而且会导致已有的元素全部复制 1 次。
// 因此，从切片的开头添加元素的性能一般要比从尾部追加元素的性能差很多。
func third() {
	var a = []int{1, 2, 3}
	a = append([]int{0}, a...)          // 在开头添加1个元素
	a = append([]int{-3, -2, -1}, a...) // 在开头添加1个切片
	fmt.Println(a)
}

// 切片在扩容时，容量的扩展规律按容量的 2 倍数扩充，例如 1、2、4、8、16……
func second() {
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}
}

func first() {
	var a []int
	a = append(a, 1)
	// 追加1个元素
	fmt.Println(len(a), cap(a))
	a = append(a, 1, 2, 3)
	// 追加多个元素, 手写解包方式
	fmt.Println(len(a), cap(a))
	// 是切片必须追加 ...
	a = append(a, []int{1, 2, 3}...)
	// 追加一个切片, 切片需要解包
	fmt.Println(len(a), cap(a))
}
