package main

import "fmt"

type X2 struct {
	a int
}

type Y2 struct {
	X2
	b int
}

type Z2 struct {
	Y2
	c int
}

func (x X2) Print() {
	fmt.Printf("In X, a=%d\n", x.a)
}

func (x X2) XPrint() {
	fmt.Printf("In X, a=%d\n", x.a)
}

func (y Y2) Print() {
	fmt.Printf("In Y2, b=%d\n", y.b)
}

func (z Z2) Print() {
	fmt.Printf("In Z2, c=%d\n", z.c)
	//显式的完全路径调用内嵌字段的方法
	z.Y2.Print()
	z.Y2.X2.Print()
}

/**

struct 类型方法调用也使用点操作符，不同嵌套层次的字段可以有相同的方法，外层变量调用内嵌字段的方法时也可以像嵌套字段的访问一样使用简化模式。
如果外层字段和内层字段有相同的方法，则使用简化模式访问外层的方法会覆盖内层的方法。

即在简写模式下，Go 编译器优先从外向内逐层查找方法，同名方法中外层的方法能够覆盖内层的方法。这个特性有点类似于面向对象编程中，子类覆盖父类的
同名方法

最外层的方法会覆盖最内层的方法

不推荐在多层的 struct 类型中内嵌多个同名的字段；但是并不反对 struct 定义和内嵌字段同名方法的用法，因为这提供了一种编程技术，
使得 struct 能够重写内嵌字段的方法，提供面向对象编程中子类覆盖父类的同名方法的功能。
*/
func main() {
	x2 := X2{
		a: 1,
	}
	y2 := Y2{
		X2: x2,
		b:  2,
	}
	z2 := Z2{
		Y2: y2,
		c:  3,
	}

	//从外向内查找，首先找到的是 Z 的 Print() 方法
	z2.Print()

	//从外向内查找，最后找到的是 x 的 XPrint()方法
	z2.XPrint()
}
