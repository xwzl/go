package main

import "fmt"

func main() {
	const f = "%T数据类型,%v替换变量\n"
	var x int = 1
	fmt.Printf(f, x, x)
	var str string = "我们"
	fmt.Printf(f, str, str)
}
