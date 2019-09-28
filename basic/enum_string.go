package main

import "fmt"

// 将 int 声明为 ChipType 芯片类型。
type ChipType int

// 将 const 里定义的一句常量值设为 ChipType 类型，且从 0 开始，每行值加 1。
const (
	None ChipType = iota
	CPU           // 中央处理器
	GPU           // 图形处理器
)

// 定义 ChipType 类型的方法 String()，返回字符串。
func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

// 使用 String() 方法的 ChipType 在使用上和普通的常量没有区别。当这个类型需要显示为字符串时，Go 语言会自动寻找 String() 方法并进行调用。
func main() {
	// 输出CPU的值并以整型格式显示
	fmt.Printf("%s %d", CPU, CPU)
}
