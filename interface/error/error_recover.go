package main

import (
	"fmt"
	"runtime"
)

/**
Recover 是一个Go语言的内建函数，可以让进入宕机流程中的 goroutine 恢复过来。recover 仅在延迟函数 defer 中有效。在正常的执行过程中，调用 recover 会返回 nil 并
且没有其他任何效果。如果当前的 goroutine 陷入恐慌，调用 recover 可以捕获到 panic 的输入值，并且恢复正常的执行。

recover 函数用于终止错误处理流程。一般情况下，recover 应该在一个使用 defer 关键字的函数中执行以有效截取错误处理流程。如果没有在发生异常的 goroutine 中明确调用
恢复过程（使用 recover 关键字），会导致该 goroutine 所属的进程打印异常信息后直接退出。

通常来说，不应该对进入 panic 宕机的程序做任何处理，但有时，也许我们可以从宕机中恢复，至少我们可以在程序崩溃前，做一些操作。举个例子，当 web 服务器遇到不可预料的
严重问题时，在崩溃前应该将所有的连接关闭；如果不做任何处理，会使得客户端一直处于等待状态。如果 web 服务器还在开发阶段，服务器甚至可以将异常信息反馈到客户端，帮助
调试。

如果在 deferred 函数中调用了内置函数 recover，并且定义该 defer 语句的函数发生了 panic 异常，recover 会使程序从 panic 中恢复，并返回 panic value。导致 panic 异
常的函数不会继续运行，但能正常返回。在未发生 panic 时调用 recover，recover 会返回 nil。

提示

在其他语言里，宕机往往以异常的形式存在。底层抛出异常，上层逻辑通过 try/catch 机制捕获异常，没有被捕获的严重异常会导致宕机，捕获的异常可以被忽略，让代码继续运行。

Go语言没有异常系统，其使用 panic 触发宕机类似于其他语言的抛出异常，那么 recover 的宕机恢复机制就对应 try/catch 机制。

让程序在崩溃时继续执行

下面的代码实现了 ProtectRun() 函数，该函数传入一个匿名函数或闭包后的执行函数，当传入函数以任何形式发生 panic 崩溃后，可以将崩溃发生的错误打印出来，同时允许后面的
代码继续运行，不会造成整个进程的崩溃。

panic和recover的关系

panic 和 defer 的组合有如下特性：

	有 panic 没 recover，程序宕机。
	有 panic 也有 recover 捕获，程序不会宕机。执行完对应的 defer 后，从宕机点退出当前函数后继续执行。

提示

虽然 panic/recover 能模拟其他语言的异常机制，但并不建议代表编写普通函数也经常性使用这种特性。

在 panic 触发的 defer 函数内，可以继续调用 panic，进一步将错误外抛直到程序整体崩溃。

如果想在捕获错误时设置当前函数的返回值，可以对返回值使用命名返回值方式直接进行设置。
*/
func main() {
	fmt.Println("运行前")
	// 允许一段手动触发的错误
	ProtectRun(func() {
		fmt.Println("手动宕机前")
		// 使用panic传递上下文
		panic(&panicContext{
			"手动触发panic",
		})
		fmt.Println("手动宕机后")
	})
	// 故意造成空指针访问错误
	ProtectRun(func() {
		fmt.Println("赋值宕机前")
		var a1 *int
		*a1 = 1
		fmt.Println("赋值宕机后")
	})
	fmt.Println("运行后")
}

// 崩溃时需要传递的上下文信息
type panicContext struct {
	function string // 所在函数
}

// 保护方式允许一个函数
func ProtectRun(entry func()) {
	// 延迟处理的函数
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error: // 运行时错误
			fmt.Println("runtime error:", err)
		default: // 非运行时错误
			fmt.Println("error:", err)
		}
	}()
	entry()
}
