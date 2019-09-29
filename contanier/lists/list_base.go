package main

import (
	"container/list"
	"fmt"
)

/**
在Go语言中，将列表使用 container/list 包来实现，内部的实现原理是双链表。列表能够高效地进行任意位置的元素插入和删除操作。

初始化列表

list 的初始化有两种方法：New 和声明。两种方法的初始化效果都是一致的。

1) 通过 container/list 包的 New 方法初始化 list

	变量名 := list.New()

2) 通过声明初始化list

	var 变量名 list.List

列表与切片和 map 不同的是，列表并没有具体元素类型的限制。因此，列表的元素可以是任意类型。这既带来便利，也会引来一些问题。给一个
列表放入了非期望类型的值，在取出值后，将 interface{} 转换为期望类型时将会发生宕机。

在列表中插入元素

双链表支持从队列前方或后方插入元素，分别对应的方法是 PushFront 和 PushBack。

提示

这两个方法都会返回一个 *list.Element 结构。如果在以后的使用中需要删除插入的元素，则只能通过 *list.Element 配合 Remove() 方法进行
删除，这种方法可以让删除更加效率化，也是双链表特性之一。

列表插入元素的方法如下表所示。
													方  法	功  能

	InsertAfter(v interface {}, mark * Element) * Element	在 mark 点之后插入元素，mark 点由其他插入函数提供
	InsertBefore(v interface {}, mark * Element) *Element	在 mark 点之前插入元素，mark 点由其他插入函数提供
	PushBackList(other *List)								添加 other 列表元素到尾部
	PushFrontList(other *List)								添加 other 列表元素到头部

从列表中删除元素

列表的插入函数的返回值会提供一个 *list.Element 结构，这个结构记录着列表元素的值及和其他节点之间的关系等信息。从列表中删除元素时，需要
用到这个结构进行快速删除

						  列表元素操作的过程
			操作内容						列表元素
		l.PushBack("canon")					canon
		l.PushFront(67)	67, 				canon
		element := l.PushBack("fist")		67, canon, fist
		l.InsertAfter("high", element)		67, canon, fist, high
		l.InsertBefore("noon", element)		67, canon, noon, fist, high
		l.Remove(element)					67, canon, noon, high
*/
func main() {
	//var list1 list.List // or var list1 =list.New()
	//list1.PushBack("1111")
	//list1.PushFront("2222")

	l := list.New()
	// 尾部添加
	l.PushBack("canon")
	// 头部添加
	l.PushFront(67)
	// 尾部添加后保存元素句柄
	element := l.PushBack("fist")
	// 在fist之后添加high
	l.InsertAfter("high", element)
	// 在fist之前添加noon
	l.InsertBefore("noon", element)
	// 使用
	fmt.Println(l)
	// 便利元素
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
