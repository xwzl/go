package main

import "fmt"

// 循环嵌套循环时，可以在 break 后指定标签。用标签决定哪个循环被终止。
func main() {

OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				break OuterLoop
			case 3:
				fmt.Println(i, j)
				break OuterLoop
			}
		}
	}
	fmt.Println("gg")
}
