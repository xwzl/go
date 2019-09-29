package main

/**
使用循环语句时，需要注意的有以下几点：

	左花括号{必须与 for 处于同一行。
	Go语言中的 for 循环与C语言一样，都允许在循环条件中定义和初始化变量，唯一的区别是，Go语言不支持以逗号为间隔的多个赋值语句，
必须使用平行赋值的方式来初始化多个变量。
	Go语言的 for 循环同样支持 continue 和 break 来控制循环，但是它提供了一个更高级的 break，可以选择中断哪一个循环，如下例
*/
func main() {
	//for j := 0; j < 5; j++ {
	//	for i := 0; i < 10; i++ {
	//		if i > 5 {
	//			break JLoop
	//		}
	//		fmt.Println(i)
	//	}
	//}
	//JLoop:

	var i int
	for {
		if i > 10 {
			break
		}
		i++
	}

	for i <= 10 {
		i++
	}
}
