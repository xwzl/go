package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
string 和 slice

string 的空值是 ""，它是不能跟 nil 比较的。即使是空的 string，它的大小也是两个机器字长的。slice 也类似，它的空值并不是一个空指针，
而是结构体中的指针域为空，空的 slice 的大小也是三个机器字长的。

channel 和 map

channel 跟 string 或 slice 有些不同，它在栈上只是一个指针，实际的数据都是由指针所指向的堆上面。

跟 channel 相关的操作有：初始化/读/写/关闭。channel 未初始化值就是 nil，未初始化的 channel 是不能使用的。下面是—些操作规则：

	读或者写一个 nil 的 channel 的操作会永远阻塞。
	读一个关闭的 channel 会立刻返回一个 channel 元索类型的零值。
	写一个关闭的 channel 会导致 panic。

map 也是指针，实际数据在堆中，未初始化的值是 nil。
*/
func main() {
	//var i []int
	//if(i == nil){
	//	fmt.Println("111")
	//}
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	//注意：忽略 input.Err() 中可能的错误
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
