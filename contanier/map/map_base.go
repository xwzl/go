package main

import "fmt"

/**
map 概念

map 是引用类型，可以使用如下声明：

	var map1 map[keytype]valuetype
	var map1 map[string]int

其中：

	keytype 为键类型。
	valuetype 是键对应的值类型。

提示：[keytype] 和 valuetype 之间允许有空格，但是 gofmt（格式化代码工具）会移除空格。

在声明的时候不需要知道 map 的长度，map 是可以动态增长的。未初始化的 map 的值是 nil。

key 可以是任意可以用 == 或者 != 操作符比较的类型，比如 string、int、float。所以数组、切片和结构体不能作为 key，但是指针和接口类型可以。
如果要用结构体作为 key 可以提供 Key() 和 Hash() 方法，这样可以通过结构体的域计算出唯一的数字或者字符串的 key。

value 可以是任意类型的；通过使用空接口类型，我们可以存储任意值，但是使用这种类型作为值时需要先做一次类型断言。

map 传递给函数的代价很小：在 32 位机器上占 4 个字节，64 位机器上占 8 个字节，无论实际上存储了多少数据。通过 key 在 map 中寻找值是很快的，
比线性查找快得多，但是仍然比从数组和切片的索引中直接读取要慢 100 倍；所以如果很在乎性能的话还是建议用切片来解决问题。

map 也可以用函数作为自己的值，这样就可以用来做分支结构：key 用来选择要执行的函数。

如果 key1 是 map1 的 key，那么 map1[key1] 就是对应 key1 的值，就如同数组索引符号一样（数组可以视为一种简单形式的 map，key 是从 0 开始的
整数）。

key1 对应的值可以通过赋值符号来设置为val1 : map1[key1] = val1。

令v := map1[key1]可以将 key1 对应的值赋值为 v；如果 map 中没有 key1 存在，那么 v 将被赋值为 map1 的值类型的空值。

常用的 len(map1) 方法可以获得 map 中的 pair 数目，这个数目是可以伸缩的，因为 map-pairs 在运行时可以动态添加和删除。

map 容量
和数组不同，map 可以根据新增的 key-value 对动态的伸缩，因此它不存在固定长度或者最大限制。但是也可以选择标明 map 的初始容量 capacity，格式如下：
make(map[keytype]valuetype, cap)

例如：
map2 := make(map[string]float, 100)

当 map 增长到容量上限的时候，如果再增加新的 key-value 对，map 的大小会自动加 1。所以出于性能的考虑，对于大的 map 或者会快速扩张的 map，
即使只是大概知道容量，也最好先标明
*/
func main() {

	var mapLit map[string]int
	//var mapCreated map[string]float32
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])       // 1
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"]) // 3.14
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])      // 3
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])       // 0
}
