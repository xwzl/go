package main

import "fmt"

type HashMap map[int]string

// 遍历输出元素的顺序与填充顺序无关。不能期望 map 在遍历时返回某种期望顺序的结果。
func main() {
	//var maps HashMap = HashMap{
	//	1: "张三",
	//	2: "李四",
	//	3: "王五",
	//}
	//
	//for key, value := range maps {
	//	fmt.Printf("乱序打印  %d : %s\n", key, value)
	//}
	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	for k, v := range scene {
		fmt.Println(k, v)
	}
}
