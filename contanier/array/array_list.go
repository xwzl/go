package main

import "fmt"

/**
数组
类型 [n]T 是一个有 n 个类型为 T 的值的数组。

表达式

var a [10]int
定义变量 a 是一个有十个整数的数组。

数组的长度是其类型的一部分，因此数组不能改变大小。 这看起来是一个制约，但是请不要担心； Go 提供了更加便利的方式来使用数组。
*/
func main() {
	//arraySample()
	//sliceSample()
	//sliceSplit()
	//sliceMake()
	//sliceNil()
	//addElement()
	//rangeSample()
	kanBuDong()
}

/**
可以通过赋值给 _ 来忽略序号和值,如果只需要索引值，去掉“, value”的部分即可。
*/
func kanBuDong() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func rangeSample() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

/**
向 slice 添加元素是一种常见的操作，因此 Go 提供了一个内建函数 `append`。 内建函数的文档对 append 有详细介绍。

func append(s []T, vs ...T) []T
append 的第一个参数 s 是一个类型为 T 的数组，其余类型为 T 的值将会添加到 slice。

append 的结果是一个包含原 slice 所有元素加上新添加的元素的 slice。

如果 s 的底层数组太小，而不能容纳所有值时，会分配一个更大的数组。 返回的 slice 会指向这个新分配的数组。
*/
func addElement() {
	var a []int
	printSlice("a", a)
	// append works on nil slices.
	a = append(a, 0)
	printSlice("a", a)
	// the slice grows as needed.
	a = append(a, 1)
	printSlice("a", a)
	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printSlice("a", a)
}

/**
slice 的零值是 `nil`,一个 nil 的 slice 的长度和容量是 0。
*/
func sliceNil() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}

/**
slice 由函数 make 创建。这会分配一个零长度的数组并且返回一个 slice 指向这个数组：
a := make([]int, 5)  // len(a)=5

为了指定容量，可传递第三个参数到 `make`：
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
*/
func sliceMake() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

/**
slice 可以重新切片，创建一个新的 slice 值指向相同的数组。

表达式

s[lo:hi]
表示从 lo 到 hi-1 的 slice 元素，含两端。因此

s[lo:lo]
是空的，而

s[lo:lo+1]
有一个元素。
*/
func sliceSplit() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])
	// 省略下标代表从 0 开始
	fmt.Println("p[:3] ==", p[:3])
	// 省略上标代表到 len(s) 结束
	fmt.Println("p[4:] ==", p[4:])
}

/**
一个 slice 会指向一个序列的值，并且包含了长度信息,[]T 是一个元素类型为 T 的 slice。
*/
func sliceSample() {
	p := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("p ==", p)
	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}

func arraySample() {
	var str [2]string
	str[0] = "value1"
	str[1] = "value2"
	fmt.Println(str[0], str[1])
	fmt.Println(str)
}
