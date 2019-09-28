package main

import (
	"fmt"
	"math"
)

type service interface {
	Abs() float64
}

func main() {
	var a service
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat 实现了 Abser
	fmt.Println(a.Abs())
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 service。
	//a = v

	fmt.Println(a.Abs())
}

type Index int

// 这种方法的使用，必须通过结构体来实现，相当于一个 Java 构造方法
func (i Index) Abs() {

}

type MyFloat float64

// 参数也是结构体吗? Abs 相当于 service 的实现
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// 结构体
type Vertex struct {
	X, Y float64
}

// 结构体当做参数
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
