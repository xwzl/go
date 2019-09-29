package main

import "fmt"

/**
从指定范围中生成切片

切片和数组密不可分。如果将数组理解为一栋办公楼，那么切片就是把不同的连续楼层出租给使用者。出租的过程需要选择开始楼层和结束楼层，这个过程就会生成切片。
*/
func main() {
	//sliceSplits()
	//announceSlice()
}

/**
除了可以从原有的数组或者切片中生成切片，你也可以声明一个新的切片。每一种类型都可以拥有其切片类型，表示多个类型元素的连续集合。因此切片类型也可以被声明。
切片类型声明格式如下：

	var name []T

	name 表示切片类型的变量名。
	T 表示切片类型对应的元素类型。
*/
func announceSlice() {
	// 声明字符串切片
	var strList []string
	// 声明整型切片
	var numList []int
	// 声明一个空切片,本来会在{}中填充切片的初始化元素，这里没有填充，所以切片是空的。但此时 numListEmpty 已经被分配了内存，只是还没有元素
	var numListEmpty = []int{}
	// 输出3个切片
	fmt.Println(strList, numList, numListEmpty)
	// 输出3个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty))
	// 切片判定空的结果
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	// numListEmpty 已经被分配到了内存，但没有元素，因此和 nil 比较时是 false
	// 切片是动态结构，只能与nil判定相等，不能互相判等时。
	// 声明新的切片后，可以使用 append() 函数来添加元素。
	fmt.Println(numListEmpty == nil)
}

// 切片分组示例
func sliceSplits() {
	var highRiseBuilding [30]int
	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}
	// 区间
	fmt.Println(highRiseBuilding[10:15])
	// 中间到尾部的所有元素
	fmt.Println(highRiseBuilding[20:])
	// 开头到中间的所有元素
	fmt.Println(highRiseBuilding[:2])
	// 重置切片
	a := []int{1, 2, 3}
	fmt.Println(a[0:0])
}
