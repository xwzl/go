package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
switch
一个结构体（`struct`）就是一个字段的集合。
除非以 fallthrough 语句结束，否则分支会自动终止。
*/
func main() {
	//acquireOs()
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

/**
获取操作系统
*/
func acquireOs() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Printf("OS X.")
	case "linux":
		fmt.Printf("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}
