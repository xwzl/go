package main

import "fmt"

/*
结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。每个值称为结构体的成员。Go语言的结构体（struct）和其他语言的类（class）有同等的地位，
但Go语言放弃了包括继承在内的大量面向对象特性，只保留了组合（composition）这个最基础的特性。

组合甚至不能算面向对象特性，因为在C语言这样的过程式编程语言中，也有结构体，也有组合。组合只是形成复合类型的基础。

使用关键字 type 可以将各种基本类型定义为自定义类型，基本类型包括整型、字符串、布尔等。结构体是一种复合的基本类型，通过 type 定义为自定义类型后，使结构体更便于使用。

结构体的定义格式如下：

	type 类型名 struct {
		字段1 字段1类型
		字段2 字段2类型
		…
	}

对各个部分的说明：

	类型名：标识自定义结构体的名称，在同一个包内不能重复。
	struct{}：表示结构体类型，type 类型名 struct{} 可以理解为将 struct{} 结构体定义为类型名的类型。
	字段1、字段2……：表示结构体字段名。结构体中的字段名必须唯一。
	字段1类型、字段2类型……：表示结构体字段的类型。

结构体的定义只是一种内存布局的描述，只有当结构体实例化时，才会真正地分配内存。因此必须在定义结构体并实例化后才能使用结构体的字段。

实例化就是根据结构体定义的格式创建一份与格式一致的内存区域，结构体实例与实例间的内存是完全独立的。

Go语言可以通过多种方式实例化结构体，根据实际需要可以选用不同的写法。

基本的实例化形式

结构体本身是一种类型，可以像整型、字符串等类型一样，以 var 的方式声明结构体即可完成实例化。

基本实例化格式如下：

	var ins T
*/
func main() {
	var point = Point{
		X: 1,
		Y: 2,
	}
	fmt.Println(point)
	fmt.Println(point.X)

	news := new(Point)
	news.X = 1

	fmt.Print(news.X)
	fmt.Printf("%p %T %T %T \n", news, news, point, &point)

	// new 出来的对象是指针对象
	tank := new(Player)
	tank.Name = "Canon"
	tank.HealthPoint = 300

	// 在 C/C++ 语言中，使用 new 实例化类型后，访问其成员变量时必须使用->操作符。
	// 在Go语言中，访问结构体指针的成员变量时可以继续使用.。这是因为Go语言为了方便开发者访问结构体指针的成员变量，使用了语法糖（Syntactic sugar）技术，将 ins.Name 形式转换为 (*ins).Name。
	fmt.Println((*tank).Name)

	var version int = 1
	cmd := &Command{}
	cmd.Name = "version"
	cmd.Var = &version
	cmd.Comment = "show version"
	println(cmd.Name)

	// go 语言传说中的构造方法

}

type Point struct {
	X int
	Y int
}

/**
创建指针类型的结构体

Go语言中，还可以使用 new 关键字对类型（包括结构体、整型、浮点数、字符串等）进行实例化，结构体在实例化后会形成指针类型的结构体。

使用 new 的格式如下：

	ins := new(T)
其中：

	T 为类型，可以是结构体、整型、字符串等。
	ins：T 类型被实例化后保存到 ins 变量中，ins 的类型为 *T，属于指针。

Go语言让我们可以像访问普通结构体一样使用.访问结构体指针的成员。
*/
type Player struct {
	Name        string
	HealthPoint int
	MagicPoint  int
}

/**
取结构体的地址实例化

在Go语言中，对结构体进行&取地址操作时，视为对该类型进行一次 new 的实例化操作。取地址格式如下：

	ins := &T{}

其中：

	T 表示结构体类型。
	ins 为结构体的实例，类型为 *T，是指针类型。

下面使用结构体定义一个命令行指令（Command），指令中包含名称、变量关联和注释等。对 Command 进行指针地址的实例化，并完成赋值过程
*/
type Command struct {
	Name    string // 指令名称
	Var     *int   // 指令绑定的变量
	Comment string // 指令的注释
}

/**
这就是传说中的构造函数
*/
func newCommand(name string, vars *int, comment string) *Command {
	return &Command{
		Name:    name,
		Var:     vars,
		Comment: comment,
	}
}
