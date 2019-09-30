package main

import (
	"fmt"
	"sort"
	"strings"
)

var original = []string{
	"Nonmetals",
	"    Hydrogen",
	"    Carbon",
	"    Nitrogen",
	"    Oxygen",
	"Inner Transitionals",
	"    Lanthanides",
	"        Europium",
	"        Cerium",
	"    Actinides",
	"        Uranium",
	"        Plutonium",
	"        Curium",
	"Alkali Metals",
	"    Lithium",
	"    Sodium",
	"    Potassium",
}

/**
本节要讲解的例子就是如何对字符串进行排序。这个函数之所以特别（以及标准库的 sort.Strings() 函数为何无法满足此需求）是因为这
个字符串是按照等级来排序的，也就是它们内部缩进的级别，完整代码如下所示。
*/
func main() {
	fmt.Println("|     Original      |       Sorted      |")
	fmt.Println("|-------------------|-------------------|")
	sorted := SortedIndentedStrings(original) // 最初是 []string
	for i := range original {                 // 在全局变量中设置
		fmt.Printf("|%-19s|%-19s|\n", original[i], sorted[i])
	}
}

/**
其中，函数 SortedIndentedStrings() 和它的辅助函数与类型使用到了递归函数、函数引用以及指向切片的指针等。尽管我们很容易看出程
序做了哪些事情，但是要真正实现这个需求还是要有一些思考的。

导出（公有）的 SortedIndentedStrings() 函数就做了这个工作，虽然我们已经对它进行了重构，让它把所有东西都传递给辅助函数。函数
populateEntries() 传入一个 []string 并返回一个对应的 Entries([]Entry 类型)。
*/
func SortedIndentedStrings(slice []string) []string {
	entries := populateEntries(slice)
	return sortedEntries(entries)
}

// 函数 sortedEntries() 需要传入一个 Entries，然后返回一个排过序的 []string(根据缩进的级别进行排序)
func populateEntries(slice []string) Entries {
	indent, indentSize := computeIndent(slice)
	entries := make(Entries, 0)
	for _, item := range slice {
		i, level := 0, 0
		for strings.HasPrefix(item[i:], indent) {
			i += indentSize
			level++
		}
		key := strings.ToLower(strings.TrimSpace(item))
		addEntry(level, key, item, &entries)
	}
	return entries
}

func computeIndent(slice []string) (string, int) {
	for _, item := range slice {
		if len(item) > 0 && (item[0] == ' ' || item[0] == '\t') {
			// 这里返回 ' ' 或者 '\t' 的字符编码的数字
			whitespace := rune(item[0])
			// i 下标从 0 开始
			for i, char := range item[1:] {
				// 匹配最后一个下标不是 '' 或者 '\t' 的字符
				if char != whitespace {
					// i 下标从 0 开始，打印空白符个数 +
					i++
					return strings.Repeat(string(whitespace), i), i
				}
			}
		}
	}
	return "", 0
}

// 这个递归算法牛逼了
func addEntry(level int, key, value string, entries *Entries) {
	if level == 0 {
		*entries = append(*entries, Entry{key, value, make(Entries, 0)})
	} else {
		addEntry(level-1, key, value, &((*entries)[entries.Len()-1].children))
	}
}

func sortedEntries(entries Entries) []string {
	var indentedSlice []string
	sort.Sort(entries)
	for _, entry := range entries {
		populateIndentedStrings(entry, &indentedSlice)
	}
	return indentedSlice
}

func populateIndentedStrings(entry Entry, indentedSlice *[]string) {
	*indentedSlice = append(*indentedSlice, entry.value)
	sort.Sort(entry.children)
	for _, child := range entry.children {
		populateIndentedStrings(child, indentedSlice)
	}
}

/**
在我们给岀的参考答案里，最关键的地方就是自定义的 Entry 和 Entries 类型。对于在原切片里的每一个字符串，我们为它创建一个 Entry 的“键/值”结构，
键字段是用来排序的，值字段保存原字符串，而 children 字段则是该字符串的孩子 Entry 切片（children 可能为空，如果不为空，它包含的 Entry 自身也
还可能包含子 Entry，以此类推）。
*/
type Entry struct {
	key      string
	value    string
	children Entries
}

// sort.Interface 接口定义了 3 个方法 Len()、Less()和 Swap()。它们的函数签名和 Entries 中的同名方法是一样的。
// 这就意味着我们可以使用标准库里的 sort.Sort() 函数来对一个 Entries 进行排序
type Entries []Entry

// 返回切换的长度
func (entries Entries) Len() int { return len(entries) }

// 判断当前切片的索引
func (entries Entries) Less(i, j int) bool {
	return entries[i].key < entries[j].key
}

// 交换元素
func (entries Entries) Swap(i, j int) {
	entries[i], entries[j] = entries[j], entries[i]
}
