package main

import "fmt"

/**
链式调用的本质是......
*/
func main() {

	var list = []string{"Java", "Python", "Go", "Go 中链式调用"}

	chain := []func(string) string{
		test1,
		test2,
		test3,
	}

	chainProcess(list, chain)

}

func chainProcess(list []string, chain []func(string) string) {
	for _, value := range list {
		for _, prop := range chain {
			prop(value)
		}
	}
}
func test1(value1 string) string {
	fmt.Println(value1 + "1")
	return value1
}

func test2(value1 string) string {
	fmt.Println(value1 + "2")
	return value1
}

func test3(value1 string) string {
	fmt.Println(value1 + "3")
	return value1
}
