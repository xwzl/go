package main

import (
	"fmt"
	"sort"
)

type statistics struct {
	numbers []float64
	mean    float64
	mdian   float64
}

/**
getStats 函数的作用就是对传入的 []float64 切片（这些数据都在 processRequest() 里得到）进行统计，然后将相应的结果保存到 stats 结果变量中。

其中计算中位数使用了 sort 包里的 Float64s() 函数对原数组进行升序排列（原地排序），也就是说 getStats() 函数修改了它的参数，这种情况在传切片、
引用或者函数指针到函数时是很常见的。

如果需要保留原始切片，可以使用Go语言内置的 copy() 函数将它赋值到一个临时变量，使用临时变量来工作。

结构体中的 mean（通常也叫平均数）是对一连串的数进行求和然后除以总个数得到的结果。这里我们使用一个辅助函数 sum() 求和，使用内置的 len() 取得切片
的大小（总个数）并将其强制转换成 float64 类型的变量（因为 sum() 函数返回一个 float64 的值)。

这样我们也就确保了这是一个浮点除法运算，避免了使用整数类型可能带来的精度损失问题。median 是用来保存中位数的，我们使用 median() 函数来单独计算它。

我们没有检查除数为 0 的情况，因为在我们的程序逻辑里，getStats() 函数只有在至少有 1 个数据的时候才会被调用，否则程序会退出并产生一个运行时异常 (runtime panic)。

对于一个关键性应用当发生一个异常时程序是不应该被结束的，我们可以使用 recover() 来捕获这个异常，将程序恢复到一个正常的状态，让程序继续运行。
*/
func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.mdian = median(numbers)
	return stats
}

func sum(numbers []float64) (result float64) {
	for _, value := range numbers {
		result += value
	}
	return result
}

func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}

func main() {
	stats := getStats([]float64{10, 30, 20, 40})
	fmt.Println(stats)
}
