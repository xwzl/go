package main

import (
	"fmt"
	"testing"
)

/**
test 简单用法

每一个测试文件必须导入 testing 包。

功能测试函数必须以 Test 开头，可选的后缀名称必须以大写字母开头

参数 t 提供了汇报测试失败和日志记录的功能。
*/
func TestMainTwo(t *testing.T) {
	fmt.Println("方法名以 Test 开头，参数是 *testing.T 类型")
}
