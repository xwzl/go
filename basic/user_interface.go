// 这个示例程序展示如何将嵌入类型应用于接口
package main

import (
	"fmt"
)

// notifier 是一个定义了
// 通知类行为的接口
type notifier interface {
	notify()
}

// User 在程序里定义一个用户类型
type User struct {
	name  string
	email string
}

// 通过 User 类型值的指针
// 调用的方法
func (u *User) notify() {
	fmt.Printf("Sending User email to %s<%s>\n",
		u.name,
		u.email)
}

// Admin 代表一个拥有权限的管理员用户
type Admin struct {
	User
	level string
}

/**
上面代码所示的示例程序的大部分和之前的程序相同，只有一些小变化，在代码的第 8 行，声明了一个 notifier 接口。之后在第 53 行，
有一个 sendNotification 函数，接受 notifier 类型的接口的值。从代码可以知道，user 类型之前声明了名为 notify 的方法，该方法
使用指针接收者实现了 notifier 接口。

之后，让我们看一下 main 函数的改动，这里才是事情变得有趣的地方。在代码的第 37 行，我们创建了一个名为 ad 的变量，其类型是外部
类型 admin。这个类型内部嵌入了 user 类型。

之后在第 48 行，我们将这个外部类型变量的地址传给 sendNotification 函数。编译器认为这个指针实现了 notifier 接口，并接受了这个
值的传递。不过如果看一下整个示例程序，就会发现 admin 类型并没有实现这个接口。

由于内部类型的提升，内部类型实现的接口会自动提升到外部类型。这意味着由于内部类型的实现，外部类型也同样实现了这个接口。运行这个
示例程序，会得到如下所示的输出。
*/
func main() {
	// 创建一个 Admin 用户
	ad := Admin{
		User: User{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// 给 Admin 用户发送一个通知
	// 用于实现接口的内部类型的方法，被提升到 外部类型
	// 这个必须使用指针吗？
	sendNotification(&ad)
}

// sendNotification 接受一个实现了notifier 接口的值
// 并发送通知
func sendNotification(n notifier) {
	n.notify()
}
