package main

type X3 struct {
	a int
}

type Y3 struct {
	X3
}

type Z3 struct {
	*X3
}

func (x X3) Get() int {
	return x.a
}
func (x *X3) Set(i int) {
	x.a = i
}

/**
组合的方法集

组合结构的方法集有如下规则：

	若类型 S 包含匿名字段 T，则 S 的方法集包含 T 的方法集。
	若类型 S 包含匿名字段 *T，则 S 的方法集包含 T 和 *T 方法集。
	不管类型 S 中嵌入的匿名字段是 T 还是*T，*S 方法集总是包含 T 和 *T 方法集。

下面举个例子来验证这个规则的正确性，前面讲到方法集时提到 Go 编译器会对方法调用进行自动转换，为了阻止自动转换，本示例使用方法表达式
的调用方式，这样能更清楚地理解这个方法集的规约

到目前为止还没有发现方法集有多大的用途，而且通过实践发现，Go 编译器会进行自动转换，看起来不需要太关注方法集，这种认识是错误的。编译器
的自动转换仅适用于直接通过类型实例调用方法时才有效，类型实例传递给接口时，编译器不会进行自动转换，而是会进行严格的方法集校验。

Go 函数的调用实参都是值拷贝，方法调用参数传递也是一样的机制，具体类型变量传递给接口时也是值拷贝，如果传递给接口变量的是值类型，但调用方
法的接收者是指针类型，则程序运行时虽然能够将接收者转换为指针，但这个指针是副本的指针，并不是我们期望的原变量的指针。

所以语言设计者为了杜绝这种非期望的行为，在编译时做了严格的方法集合的检查，不允许产生这种调用；如果传递给接口的变量是指针类型，则接口调用
的是值类型的方法，程序运行时能够自动转换为值类型，这种转换不会带来副作用，符合调用者的预期，所以这种转换是允许的，而且这种情况符合方法集
的规约。
*/
func main() {
	x := X3{a: 1}
	y := Y3{
		X3: x,
	}
	println(y.Get()) // 1
	//此处编译器做了自动转换
	y.Set(2)
	println(y.Get()) // 2
	//为了不让编译器做自动转换，使用方法表达式调用方式
	//Y3 内嵌字段 X3，所以 type y 的方法集是 Get, type *Y3 的方法集是 Set Get
	(*Y3).Set(&y, 3)
	//type y 的方法集合并没有 Set 方法，所以下一句编译不能通过
	//Y3.Set(y, 3)
	println(y.Get()) // 3
	z := Z3{
		X3: &x,
	}
	//按照嵌套字段的方法集的规则
	//Z3 内嵌字段＊X3 ，所以 type Z3 和 type *Z3 方法集都包含类型 X3 定义的方法 Get 和 Set
	//为了不让编译器做自动转换，仍然使用方法表达式调用方式
	Z3.Set(z, 4)
	println(z.Get()) // 4
	(*Z3).Set(&z, 5)
	println(z.Get()) // 5
}
