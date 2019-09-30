package main

import "fmt"

// 直接调过外层循环
func main() {

OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				continue OuterLoop
			}
			fmt.Println(i, j, 1)
		}
	}

}
