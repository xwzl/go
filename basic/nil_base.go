package main

/**
nil 是—种数据结构么？没错。不仅仅是Go语言中，每门语言中 nil 都是非常重要的，它代表的是空值的语义。

在不同语言中，表示空这个概念都有细微不同。比如在 scheme 语言（一种 lisp 方言）中，nil 是 true 的，而在 ruby 语言中，
一切都是对象，连 nil 也是—个对象！在C语言中 NULL 跟 0 是等价的。

按照Go语言规范，任何类型在未初始化时都对应一个零值：布尔类型是 false，整型是 0，字符串是 ""，而指针，函数，interface，
slice，channel 和 map 的零值都是 nil。
*/
type Error struct {
	errCode uint8
}

func (e *Error) Error() string {
	switch e.errCode {
	case 1:
		return "file not found"
	case 2:
		return "time out"
	case 3:
		return "permission denied"
	default:
		return "unknown error"
	}
}

// http://c.biancheng.net/view/4776.html
func main() {
	var e *Error
	checkError(e)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
