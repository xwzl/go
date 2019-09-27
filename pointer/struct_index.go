package main

import "fmt"

func main() {

	v := Vertex{9, 21}
	v.X = 100
	fmt.Printf("%T 结构体类型，%v \n", v, v.X)

	fmt.Println(Vertex{1, 2})

	// 结构体字段可以通过结构体指针来访问,通过指针间接的访问是透明的。
	p := &v
	p.X = 10011
	fmt.Println(v)

}

/**
结构体

一个结构体（`struct`）就是一个字段的集合。（而 type 的含义跟其字面意思相符。）
*/
type Vertex struct {
	X int
	Y int
}
