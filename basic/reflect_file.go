package main

import (
	"fmt"
	"reflect"
)

type BaseInfo interface {
	printBaseInfo()
}

type Class struct {
	// 班级名称
	className string

	mainTeacher string
}

func (c Class) init() {
	c.className = "成都七中 高一一班"
	c.mainTeacher = "张佳丽"
}

func (c Class) printBaseInfo() {
	fmt.Println(c.className+"1", c.mainTeacher)
}

// 定义 Class 结构体的别名为 C
type C = Class

type School struct {
	Class
	C
}

func main() {

	var s School

	s.Class.init()
	ta := reflect.TypeOf(s)

	for i := 0; i < ta.NumField(); i++ {
		// school 的成员信息
		field := ta.Field(i)
		fmt.Printf("FieldName: %v, FieldType: %v\n", field.Name, field.Type.Name())
	}

	s.Class.className = "111"
	s.Class.mainTeacher = "222"

	s.Class.printBaseInfo()

}
