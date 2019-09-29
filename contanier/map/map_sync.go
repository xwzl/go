package main

import (
	"fmt"
	"sync"
)

/**
需要并发读写时，一般的做法是加锁，但这样性能并不高。Go语言在 1.9 版本中提供了一种效率较高的并发安全的 sync.Map。sync.Map 和map
不同，不是以语言原生形态提供，而是在 sync 包下的特殊结构。

sync.Map有以下特性：

	无须初始化，直接声明即可。
	sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用。Store 表示存储，Load 表示获取，Delete
    表示删除。
	使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值。Range 参数中的回调函数的返回值功能是：需要继续迭
    代遍历时，返回 true；终止迭代遍历时，返回 false。
*/
func main() {

	var scene sync.Map

	scene.Store("key", "values")
	scene.Store("table", "tables")
	scene.Store("Java", "java")

	// 获取
	fmt.Println(scene.Load("Java"))

	// 删除
	scene.Delete("Java")

	scene.Range(func(key, value interface{}) bool {
		fmt.Println("iterate:", key, value)
		return true
	})

}
