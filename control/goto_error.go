package main

// 如果使用 goto 语句来实现同样的逻辑：
func main() {
	//	err := firstCheckError()
	//	if err != nil {
	//		goto onExit
	//	}
	//	err = secondCheckError()
	//	if err != nil {
	//		goto onExit
	//	}
	//	fmt.Println("done")
	//	return
	//onExit:
	//	fmt.Println(err)
	//	exitProcess()
}

// 多处错误处理存在代码重复时是非常棘手的，goto 实现则简略得多
func errors() {
	//err := firstCheckError()
	//if err != nil {
	//fmt.Println(err)
	//exitProcess()
	//return
	//}
	//err = secondCheckError()
	//if err != nil {
	//fmt.Println(err)
	//exitProcess()
	//return
	//}
	//fmt.Println("done")
}
