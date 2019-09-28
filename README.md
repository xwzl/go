
http://tour.studygolang.com/list

# 基本命令
    
    // 运行
    go run demo.go
    
    // 编译二进制文件
    go build demo.go

# 哈哈
    
 Go 程序是通过 package 来组织的。

 只有 package 名称为 main 的包可以包含 main 函数。
 
 一个可执行程序有且仅有一个 main 包。
 
 通过 import 关键字来导入其他非 main 包。
 
 可以通过 import 关键字单个导入:

```go
package main
import (
    "fmt"
    "math"
)
func main() {
    fmt.Println(math.Exp2(10))  // 1024
}
```

使用 <PackageName>.<FunctionName> 调用:

```go
package 别名：
// 为fmt起别名为fmt2
import fmt2 "fmt"
```

省略调用(不建议使用):
```go
// 调用的时候只需要Println()，而不需要fmt.Println()
import . "fmt"
```
前
面加个点表示省略调用，那么调用该模块里面的函数，可以不用写模块名称了:
```go
import . "fmt"
func main (){
    Println("hello,world")
}
```
}

通过 const 关键字来进行常量的定义。

通过在函数体外部使用 var 关键字来进行全局变量的声明和赋值。

通过 type 关键字来进行结构(struct)和接口(interface)的声明。

通过 func 关键字来进行函数的声明。

# 3 可见性规则

Go语言中，使用大小写来决定该常量、变量、类型、接口、结构或函数是否可以被外部包所调用。

函数名首字母小写即为 private :

    func getId() {}

函数名首字母大写即为 public :

    func Printf() {}
