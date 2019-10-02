package main

import (
	"bytes"
	"fmt"
)

/**
获得可变参数类型——获得每一个参数的类型

当可变参数为 interface{} 类型时，可以传入任何类型的值。此时，如果需要获得变量的类型，可以通过 switch 类型分支获得变量的类型。
下面的代码演示将一系列不同类型的值传入 printTypeValue() 函数，该函数将分别为不同的参数打印它们的值和类型的详细描述。
*/
func main() {
	// 将不同类型的变量通过printTypeValue()打印出来
	fmt.Println(printTypeValue(100, "str", true))
}

func printTypeValue(slist ...interface{}) string {
	// 字节缓冲作为快速字符串连接
	var b bytes.Buffer
	// 遍历参数
	for _, s := range slist {
		// 将interface{}类型格式化为字符串
		str := fmt.Sprintf("%v", s)
		// 类型的字符串描述
		var typeString string
		// 对s进行类型断言
		switch s.(type) {
		case bool: // 当s为布尔类型时
			typeString = "bool"
		case string: // 当s为字符串类型时
			typeString = "string"
		case int: // 当s为整型类型时
			typeString = "int"
		}
		// 写值字符串前缀
		b.WriteString("value: ")
		// 写入值
		b.WriteString(str)
		// 写类型前缀
		b.WriteString(" type: ")
		// 写类型字符串
		b.WriteString(typeString)
		// 写入换行符
		b.WriteString("\n")
	}
	return b.String()
}
