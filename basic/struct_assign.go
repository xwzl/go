package main

import "fmt"

type x struct {
	username string
	password string
}

// 结构体赋值
func main() {
	x := x{
		username: "张三",
		password: "李四",
	}
	fmt.Printf("Username: %s,Password: %s", x.username, x.password)
}
