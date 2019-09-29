package main

import "fmt"

// 清空 map 中的所有元素
// 有意思的是，Go语言中并没有为 map 提供任何清空所有元素的函数、方法。清空 map 的唯一办法就是重新 make 一个新的
// map。不用担心垃圾回收的效率，Go语言中的并行垃圾回收效率比写一个清空函数高效多了
func main() {

	scene := make(map[string]int)
	scene["1"] = 1
	scene["2"] = 2

	delete(scene, "1")

	fmt.Println(scene)
}
