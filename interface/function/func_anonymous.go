package main

import (
	"flag"
	"fmt"
)

var skillParam = flag.String("skill", "soldier run", "skill to perform")

/**
Go语言中支持匿名函数，即函数可以像普通变量一样被传递或使用，这与C语言的回调函数比较类似。不同的是，Go语言支持随时在代码里定义匿名函数。

匿名函数是指不需要定义函数名的一种函数实现方式，由一个不带函数名的函数声明和函数体组成，下面来具体介绍一下匿名函数的定义及使用。
定义一个匿名函数

匿名函数的定义格式如下：

	func(参数列表)(返回参数列表){
		函数体
	}

匿名函数的定义就是没有名字的普通函数定义。

1) 在定义时调用匿名函数

匿名函数可以在声明后调用，例如：

	func(data int) {
		fmt.Println("hello", data)
	}(100)

注意第3行“}”后的(100)，表示对匿名函数进行调用，传递参数为 100。

2) 将匿名函数赋值给变量

匿名函数体可以被赋值，例如：

	// 将匿名函数体保存到f()中
	f := func(data int) {
		fmt.Println("hello", data)
	}
	// 使用f()调用
	f(100)

匿名函数的用途非常广泛，匿名函数本身是一种值，可以方便地保存在各种容器中实现回调函数和操作封装。
*/
func main() {
	visitFunction([]int{1, 2, 3, 4, 5, 6}, func(a int) {
		fmt.Printf("匿名函数调用: %v \n", a)
	})

	flag.Parse()
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}
	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
}

func visitFunction(list []int, f func(int)) {
	for _, value := range list {
		f(value)
	}
}
