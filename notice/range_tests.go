package main

import "fmt"

func main() {
	slice := []int{0}
	for value := range []int{1, 2, 3, 4, 5, 6} {
		slice = append(slice, value)
	}
	fmt.Println(slice)

}
