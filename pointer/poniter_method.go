package main

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

/**
指针直接引用方法
*/
func (v *Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/**
接收者为指针的方法
方法可以与命名类型或命名类型的指针关联。

刚刚看到的两个 Abs 方法。一个是在 *Vertex 指针类型上，而另一个在 MyFloat 值类型上。 有两个原因需要使用指针接收者。首先避免在每个方法调用中拷贝值（如果值类型是大的结构体的话会更有效率）。其次，方法可以修改接收者指向的值。

尝试修改 Abs 的定义，同时 Scale 方法使用 Vertex 代替 *Vertex 作为接收者。

当 v 是 Vertex 的时候 Scale 方法没有任何作用。`Scale` 修改 `v`。当 v 是一个值（非指针），方法看到的是 Vertex 的副本，并且无法修改原始值。

Abs 的工作方式是一样的。只不过，仅仅读取 `v`。所以读取的是原始值（通过指针）还是那个值的副本并没有关系。
*/
func (v *Vertex2) Ab(x float64) {
	v.X = v.X * x
	v.Y = v.Y * x
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	// 指针对象调用方法
	v := &Vertex2{3, 4}
	v.Ab(5)
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
