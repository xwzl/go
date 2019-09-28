package main

import "fmt"

// 取出 x 和 y 的地址作为参数传给 swap() 函数进行调用。
func swap(a, b *int) {
	//fmt.Printf("value type %T\n",a)
	//fmt.Printf("value type %T\n",b)
	// 取a指针的值, 赋给临时变量t
	t := *a
	// 取b指针的值, 赋给a指针指向的变量
	*a = *b
	// 将a指针的值赋给b指针指向的变量
	*b = t
}

/**
操作符作为右值时，意义是取指针的值；作为左值时，也就是放在赋值操作符的左边时，表示 a 指向的变量。
其实归纳起来，*操作符的根本意义就是操作指针指向的变量。当操作在右值时，就是取指向变量的值；当操作
在左值时，就是将值设置给指向的变量。
*/
func main() {
	// 准备两个变量, 赋值1和2
	x, y := 1, 2
	// 交换变量值
	swap(&x, &y)
	// 输出变量值
	fmt.Println(x, y)
	// 交换地址值
	swapAddress(&x, &y)
	fmt.Println(x, y)
}

func swapAddress(x, y *int) {
	// 交换 x,y 变量的地址值，只是交换入参参数的地址值
	x, y = y, x
	fmt.Println(*x, *y)
}
