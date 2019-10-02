package main

/**
Go语言的错误处理思想及设计包含以下特征：

一个可能造成错误的函数，需要返回值中返回一个错误接口（error）。如果调用是成功的，错误接口将返回 nil，否则返回错误。
在函数调用后需要检查错误，如果发生错误，进行必要的错误处理。

Go语言没有类似 Java 或 .NET 中的异常处理机制，虽然可以使用 defer、panic、recover 模拟，但官方并不主张这样做.Go语言的设计者认为其
他语言的异常机制已被过度使用，上层逻辑需要为函数发生的异常付出太多的资源。同时，如果函数使用者觉得错误处理很麻烦而忽略错误，那么程
序将在不可预知的时刻崩溃。

Go语言希望开发者将错误处理视为正常开发必须实现的环节，正确地处理每一个可能发生错误的函数。同时，Go语言使用返回值返回错误的机制，也
能大幅降低编译器、运行时处理错误的复杂度，让开发者真正地掌握错误的处理。
*/
func main() {

}