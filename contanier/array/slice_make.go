package main

import "fmt"

/**
使用 make() 函数构造切片

如果需要动态地创建一个切片，可以使用 make() 内建函数，格式如下：

	make( []T, size, cap )

	T：切片的元素类型。
	size：就是为这个类型分配多少个元素。
	cap：预分配的元素数量，这个值设定后不影响 size，只是能提前分配空间，降低多次分配空间造成的性能问题。

容量不会影响当前的元素个数，因此 a 和 b 取 len 都是 2。

温馨提示

使用 make() 函数生成的切片一定发生了内存分配操作。但给定开始与结束位置（包括切片复位）的切片只是将新的切片结构指向已经分配好的内存区域，
设定开始与结束位置，不会发生内存分配操作。

切片不一定必须经过 make() 函数才能使用。生成切片、声明后使用 append() 函数均可以正常使用切片。
*/
func main() {
	a := make([]int, 2, 2)
	b := make([]int, 2, 5)

	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))
	fmt.Println(a, b)

	a = append(a, 1)
	fmt.Println(a)
	fmt.Println(len(a), cap(a))
}
