package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/**
像C语言或其他语言中的 printf 一样，函数 fmt.Printf 从一个表达式列表生成格式化的输出。它的第一个参数是格式化指示字符串，由它指定其他参数如何格式化。
每一个参数的格式是一个转义字符、一个百分号加一个字符。例如：％d 将一个整数格式化为十进制的形式，％s 把参数展开为字符串变量的值。

Printf 函数有超过 10 个这样的转义字符，Go 程序员称为 verb。下表远不完整，但是它说明有很多可以用的功能：

		verb				描述
	%d				十进制整数
	%x, %o, %b		十六进制、八进制、二进制整数
	%f, %g, %e		浮点数：如 3.141593, 3.141592653589793, 3.141593e+00
	%t				布尔型：true 或 false
	%c				字符（Unicode 码点）
	%s				字符串
	%q				带引号字符串（如 "abc"）或者字符（如 'c'）
	%v				内置格式的任何值
	%T				任何值的类型
	%%				百分号本身（无操作数）

程序 dup1 中的格式化字符串还包含一个制表符 \t 和一个换行符 \n。字符串字面量可以包含类似转义序列（escape sequence）来表示不可见字符。Printf 默认不
写换行符。

按照约定，诸如 log.Printf 和 fmt.Errorf 之类的格式化函数以 f 结尾，使用和 fmt.Printf 相同的格式化规则；而那些以 ln 结尾的函数（如 Println）则使用
％v 的方式来格式化参数，并在最后追加换行符。

http://c.biancheng.net/view/4778.html
语言进制转换 http://c.biancheng.net/view/5575.html
*/
func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\r") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
