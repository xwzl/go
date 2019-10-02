package main

import "fmt"

/**
示例：在解析中使用自定义错误

使用 errors.New 定义的错误字符串的错误类型是无法提供丰富的错误信息的。那么，如果需要携带错误信息返回，就需要借助自定义结构体实现错误接口。

下面代码将实现一个解析错误（ParseError），这种错误包含两个内容：文件名和行号。解析错误的结构还实现了 error 接口的 Error() 方法，返回错误
描述时，就需要将文件名和行号返回。
*/
func main() {
	var e error
	// 创建一个错误实例，包含文件名和行号
	e = newParseError("main.go", 1)
	// 通过error接口查看错误描述
	fmt.Println(e.Error())
	// 根据错误接口具体的类型，获取详细错误信息
	switch detail := e.(type) {
	case *ParseError: // 这是一个解析错误
		fmt.Printf("Filename: %s Line: %d\n", detail.Filename, detail.Line)
	default: // 其他类型的错误
		fmt.Println("other error")
	}
}

// 声明一个解析错误
type ParseError struct {
	Filename string // 文件名
	Line     int    // 行号
}

// 实现error接口，返回错误描述
func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", e.Filename, e.Line)
}

// 创建一些解析错误
func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}
