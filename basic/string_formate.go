package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	//lenSample()
	//stringFormat()
	//substr()
	//asciiRange()
	//unicodeRange()
	//multiStr()
	//stringBuilder()
	// 需要处理的字符串
	message := "Away from keyboard. https://golang.org/"
	// 编码消息
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(message))
	// 输出编码完成的消息
	fmt.Println(encodedMessage)
	// 解码消息
	data, err := base64.StdEncoding.DecodeString(encodedMessage)
	// 出错处理
	if err != nil {
		fmt.Println(err)
	} else {
		// 打印解码完成的数据
		fmt.Println(string(data))
	}
	//fmt.Println(utf8.RuneCountInString("中文字符长度获取用这个方法"))
}

func stringBuilder() {
	hammer := "吃我一锤"
	sickle := "死吧"
	// 声明字节缓冲
	var stringBuilder bytes.Buffer
	// 把字符串写入缓冲
	stringBuilder.WriteString(hammer)
	stringBuilder.WriteString(sickle)
	// 将缓冲以字符串形式输出
	fmt.Println(stringBuilder.String())
}

func unicodeRange() {
	theme := "中文字符串遍历"
	for _, value := range theme {
		fmt.Printf("Unicode: %c  %d\n", value, value)
	}
}

/**
中文字符长度被分割成 ascii 码长度个数，这种模式下取到的汉字“惨不忍睹”。由于没有使用 Unicode，汉字被显示为乱码。
*/
func asciiRange() {
	theme := "狙击 start"
	for i := 0; i < len(theme); i++ {
		fmt.Printf("ascii: %c  %d\n", theme[i], theme[i])
	}
}

/**
多行字符串
*/
func multiStr() {
	multiStr := `
				we
				go
		`
	for i := 0; i < len(multiStr); i++ {
		fmt.Printf("asciiL %c %d \n", multiStr[i], multiStr[i])
	}
}

/**
len() 函数的返回值的类型为 int，表示字符串的 ASCII 字符个数或字节长度。
ASCII 字符串长度使用 len() 函数。
Unicode 字符串长度使用 utf8.RuneCountInString() 函数
*/
func lenSample() {
	// 然而根据习惯，“火影忍者”的字符个数应该是 4
	// 这里的差异是由于 Go 语言的字符串都以 UTF-8 格式保存，每个中文占用 3 个字节，因此使用 len() 获得两个中文文字对应的 6 个字节。
	title := "火影忍者"
	fmt.Println(len(title))
	title = "happy new year"
	fmt.Println(len(title))
	// 如果希望按习惯上的字符个数来计算，就需要使用 Go 语言中 UTF-8 包提供的 RuneCountInString() 函数，统计 Uncode 字符数量。
	// 下面的代码展示如何计算UTF-8的字符个数。
	fmt.Println(utf8.RuneCountInString("忍者"))
	fmt.Println(utf8.RuneCountInString("龙忍出鞘,fight!"))
}

/**
截取字符串
*/
func substr() {
	tracer := "死神来了, 12死神bye bye"
	comma := strings.Index(tracer, ", ")
	pos := strings.Index(tracer[comma:], "死神")
	fmt.Println(comma, pos, tracer[comma+pos:])
}

func stringFormat() {
	const f = "%T数据类型,%v替换变量\n"
	var x int = 1
	fmt.Printf(f, x, x)
	var str string = "我们"
	fmt.Printf(f, str, str)
}
