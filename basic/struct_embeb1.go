package main

type X1 struct {
	a int
}
type Y1 struct {
	X1
	a int
}
type Z1 struct {
	Y1
	a int
}

// 在 struct 的多层嵌套中，不同嵌套层次可以有相同的字段，此时最好使用完全路径进行访问和初始化。
// 在实际数据结构的定义中应该尽量避开相同的字段，以免在使用中出现歧义。
func main() {
	x := X1{a: 1}
	y := Y1{
		X1: x,
		a:  2,
	}
	z := Z1{
		Y1: y,
		a:  3,
	}
	//此时的z.a, z.Y1.a, z.Y1.X1.a 代表不同的字段
	println(z.a, z.Y1.a, z.Y1.X1.a) // 3 2 1
	z = Z1{}
	z.a = 4
	z.Y1.a = 5
	z.Y1.X1.a = 6
	//此时的z.a, z.Y1.a, z.Y1.X1.a 代表不同的字段
	println(z.a, z.Y1.a, z.Y1.X1.a) // 4 5 6
}
