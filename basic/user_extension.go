// 这个示例程序展示当内部类型和外部类型要
// 实现同一个接口时的做法
package main

import (
	"fmt"
)

// Notifier 是一个定义了
// 通知类行为的接口
type Notifier interface {
	notify()
}

// users 在程序里定义一个用户类型
type users struct {
	name  string
	email string
}

// 通过users 类型值的指针
// 调用的方法
func (u *users) notify() {
	fmt.Printf("Sending users email to %s<%s>\n",
		u.name,
		u.email)
}

// admins 代表一个拥有权限的管理员用户
type admins struct {
	users
	level string
}

// 通过 admins 类型值的指针
// 调用的方法
func (a *admins) notify() {
	fmt.Printf("Sending admins email to %s<%s>\n",
		a.name,
		a.email)
}

/**
如果外部类型并不需要使用内部类型的实现，而想使用自己的一套实现，该怎么办？让我们看另一个示例程序是如何解决这个问题的，代码如下所示。
外部重写相同的方法会阻止方法的提升，语法糖效果

上面代码所示的示例程序的大部分和之前的程序相同，只有一些小变化，这个示例程序为 admin 类型增加了 notifier 接口的实现。当 admin 类型
的实现被调用时，会显示 "Sending admin email"。作为对比，user 类型的实现被调用时，会显示 "Sending user email"。

main 函数里也有一些变化，在代码的第 46 行，我们再次创建了外部类型的变量 ad。在第 57 行，将 ad 变量的地址传给 sendNotification 函数，
这个指针实现了接口所需要的方法集。

在第 60 行，代码直接访问 user 内部类型，并调用 notify 方法。最后，在第 63 行，使用外部类型变量 ad 来调用 notify 方法。这个示例程序
的输出结果如下所示。
Sending admin email to john smith<john@yahoo.com>
Sending user email to john smith<john@yahoo.com>
Sending admin email to john smith<john@yahoo.com>

这次我们看到了 admin 类型是如何实现 notifier 接口的，以及如何由 sendNotification 函数以及直接使用外部类型的变量 ad 来执行 admin 类型
实现的方法。这表明，如果外部类型实现了 notify 方法，内部类型的实现就不会被提升。

不过内部类型的值一直存在，因此还可以通过直接访问内部类型的值，来调用没有被提升的内部类型实现的方法。
*/
func main() {
	// 创建一个 admins 用户
	ad := admins{
		users: users{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}
	// 给admins 用户发送一个通知
	// 接口的嵌入的内部类型实现并没有提升到
	// 外部类型
	sendNotifications(&ad)
	// 我们可以直接访问内部类型的方法
	ad.users.notify()
	// 内部类型的方法没有被提升
	ad.notify()
}

// sendNotification 接受一个实现了 Notifier 接口的值
// 并发送通知
func sendNotifications(n Notifier) {
	n.notify()
}
