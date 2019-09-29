package main

import "fmt"

/**
Go语言的内置函数 copy() 可以将一个数组切片复制到另一个数组切片。如果加入的两个数组切片不一样大，就会按其中较小的那个数组切片的元素个数进行复制。
*/
func main() {
	//copyFirst()
	// 设置元素数量为1000
	const elementCount = 1000
	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)
	// 将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	// 引用切片数据
	refData := srcData
	// 预分配足够多的元素切片
	copyData := make([]int, elementCount)
	// 将数据复制到新的切片空间中
	copy(copyData, srcData)
	// 修改原始数据的第一个元素
	srcData[0] = 999
	// 打印引用切片的第一个元素
	fmt.Println(refData[0])
	// 打印复制切片的第一个和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount-1])
	// 复制原始数据从4到6(不包含)
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}
}

// 虽然通过循环复制元素更直接，不过内置的 copy 函数使用起来更加方便。copy 函数的第一个参数是要复制的目标 slice，第二个参数是源 slice，目标和源
// 的位置顺序和 dst = src 赋值语句是一致的。两个 slice 可以共享同一个底层数组，甚至有重叠也没有问题。
// copy 函数将返回成功复制的元素的个数，等于两个 slice 中较小的长度，所以我们不用担心覆盖会超出目标 slice 的范围。
func copyFirst() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1)
	// 只会复制slice1的前3个元素到slice2中
	copy(slice1, slice2)
	// 只会复制slice2的3个元素到slice1的前3个位置
}
