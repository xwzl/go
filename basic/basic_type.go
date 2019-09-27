package main

import (
	"fmt"
	"math/cmplx"
)

var (
	// 布尔值
	flag   bool       = true
	mixInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 128i)
)

/**
Go 的基本类型有Basic types

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名
rune // int32 的别名

// 代表一个Unicode码

float32 float64

complex64 complex128
这个例子演示了具有不同类型的变量。 同时与导入语句一样，变量的定义“打包”在一个语法块中*/
func main() {
	basic()
}

func basic() {
	// %T 表示数据的类型吗？ %v 表示值吗？
	const f = "%T(%v)\n"
	fmt.Printf(f, flag, flag)
	fmt.Printf(f, mixInt, mixInt)
	fmt.Printf(f, z, z)
}
