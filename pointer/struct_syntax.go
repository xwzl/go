package main

import "fmt"

var (
	v1 = VertexPlus{1, 2}  // 类型为 Vertex
	v2 = VertexPlus{X: 1}  // Y:0 被省略
	v3 = VertexPlus{}      // X:0 和 Y:0
	p  = &VertexPlus{1, 2} // 类型为 *Vertex
)

/**
结构体文法

结构体文法表示通过结构体字段的值作为列表来新分配一个结构体。
使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）
特殊的前缀 & 返回一个指向结构体的指针。
*/
func main() {
	fmt.Println(v1, p, v2, v3)
}

type VertexPlus struct {
	X, Y int
}
