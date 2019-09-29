package main

import "fmt"

/**
当迭代切片时，关键字 range 会返回两个值。第一个值是当前迭代到的索引位置，第二个值是该位置对应元素值的一份副本

需要强调的是，range 创建了每个元素的副本，而不是直接返回对该元素的引用，如下所示。如果使用该值变量的地址作为指向每个元素的指针，
就会造成错误。让我们看看是为什么。
*/
func main() {

	// 创建一个整型切片
	// 其长度和容量都是4 个元素
	slice := []int{10, 20, 30, 40}
	// 迭代每个元素，并显示值和地址
	for index, value := range slice {
		// 因为迭代返回的变量是一个迭代过程中根据切片依次赋值的新变量，所以 value 的地址总是相同的。要想获取每个元素的地址，可以使用切片变量和索引值。
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &slice[index])
	}
}
