package main

import "fmt"

/**
从数组或切片生成新的切片

切片默认指向一段连续内存区域，可以是数组，也可以是切片本身。

从连续内存区域生成切片是常见的操作，格式如下：

slice [开始位置:结束位置]

	slice 表示目标切片对象。
	开始位置对应目标切片对象的索引。
	结束位置对应目标切片的结束索引。

从数组生成切片，代码如下：
	var a  = [3]int{1, 2, 3}
	fmt.Println(a, a[1:2])l

从数组或切片生成新的切片拥有如下特性：

	取出的元素数量为：结束位置-开始位置。
	取出元素不包含结束位置对应的索引，切片最后一个元素使用 slice[len(slice)] 获取。
	当缺省开始位置时，表示从连续区域开头到结束位置。
	当缺省结束位置时，表示从开始位置到整个连续区域末尾。
	两者同时缺省时，与切片本身等效。
	两者同时为0时，等效于空切片，一般用于切片复位。

根据索引位置取切片 slice 元素值时，取值范围是（0～len(slice)-1），超界会报运行时错误。生成切片时，结束位置可以填写 len(slice) 但不会报错。
*/
func main() {
	var a = [3]int{1, 2, 3}
	// 切片，数量为结束索引减开始索引，值包含结束索引
	fmt.Println(a, a[1:2])

	b := a[:]
	fmt.Println(b)
}
