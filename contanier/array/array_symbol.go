package main

import "fmt"

type Currency int

// 我们将会发现，数组、slice、map 和结构体字面值的写法都很相似。上面的形式是直接提供顺序初始化值序列，但是也可以指定
// 一个索引和对应值列表的方式初始化，就像下面这样：
const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)

func main() {
	// 数组 key:value 键值对赋值
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB]) // "3 ￥"

	// 在这种形式的数组字面值形式中，初始化索引的顺序是无关紧要的，而且没用到的索引可以省略，和前面提到的规则一样，未指定初
	// 始值的元素将用零值初始化。例如，
	r := [...]int{99: -1}
	fmt.Println(r[99])

	rtx := [...]string{0: "泰坦", 1: "RTX 2080 Ti", 2: "RTX 2080", 3: "RTX 2070 Ti", 4: "RTX 2070 Super"}
	for key, value := range rtx {
		fmt.Printf("Rtx 系列：第 %d 名显卡 %s\n", key+1, value)
	}
}
