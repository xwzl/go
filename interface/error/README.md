Go语言的错误处理思想及设计包含以下特征：
- 一个可能造成错误的函数，需要返回值中返回一个错误接口（error）。如果调用是成功的，错误接口将返回 nil，否则返回错误。
- 在函数调用后需要检查错误，如果发生错误，进行必要的错误处理。

Go语言没有类似 Java 或 .NET 中的异常处理机制，虽然可以使用 defer、panic、recover 模拟，但官方并不主张这样做。Go语言的设计者认为其他语言的异常机制已被过度使用，上层逻辑需要为函数发生的异常付出太多的资源。同时，如果函数使用者觉得错误处理很麻烦而忽略错误，那么程序将在不可预知的时刻崩溃。

Go语言希望开发者将错误处理视为正常开发必须实现的环节，正确地处理每一个可能发生错误的函数。同时，Go语言使用返回值返回错误的机制，也能大幅降低编译器、运行时处理错误的复杂度，让开发者真正地掌握错误的处理。

### net 包中的例子

net.Dial() 是Go语言系统包 net 即中的一个函数，一般用于创建一个 Socket 连接。

net.Dial 拥有两个返回值，即 Conn 和 error。这个函数是阻塞的，因此在 Socket 操作后，会返回 Conn 连接对象和 error；如果发生错误，error 会告知错误的类型，Conn 会返回空。

根据Go语言的错误处理机制，Conn 是其重要的返回值。因此，为这个函数增加一个错误返回，类似为 error。参见下面的代码：

    func Dial(network, address string) (Conn, error) {
        var d Dialer
        return d.Dial(network, address)
    }
    
在 io 包中的 Writer 接口也拥有错误返回，代码如下：

    type Writer interface {
        Write(p []byte) (n int, err error)
    }
    
io 包中还有 Closer 接口，只有一个错误返回，代码如下：
  
    type Closer interface {
        Close() error
    }
    
### 错误接口的定义格式

error 是 Go 系统声明的接口类型，代码如下：

    type error interface {
        Error() string
    }
    
所有符合 Error()string 格式的方法，都能实现错误接口。

Error() 方法返回错误的具体描述，使用者可以通过这个字符串知道发生了什么错误。

### 自定义一个错误

返回错误前，需要定义会产生哪些可能的错误。在Go语言中，使用 errors 包进行错误的定义，格式如下：

    var err = errors.New("this is an error")
    
错误字符串由于相对固定，一般在包作用域声明，应尽量减少在使用时直接使用 errors.New 返回。

#### errors 包

Go语言的 errors 中对 New 的定义非常简单，代码如下：

    // 创建错误对象
    func New(text string) error {
        return &errorString{text}
    }
    
    // 错误字符串
    type errorString struct {
        s string
    }
    
    // 返回发生何种错误
    func (e *errorString) Error() string {
        return e.s
    }
    
代码说明如下：

- 第 2 行，将 errorString 结构体实例化，并赋值错误描述的成员。
- 第 7 行，声明 errorString 结构体，拥有一个成员，描述错误内容。
- 第 12 行，实现 error 接口的 Error() 方法，该方法返回成员中的错误描述。

#### 在代码中使用错误定义

下面的代码会定义一个除法函数，当除数为 0 时，返回一个预定义的除数为 0 的错误。

    package main
    
    import (
        "errors"
        "fmt"
    )
    
    // 定义除数为0的错误
    var errDivisionByZero = errors.New("division by zero")
    
    func div(dividend, divisor int) (int, error) {
        // 判断除数为0的情况并返回
        if divisor == 0 {
            return 0, errDivisionByZero
        }
        // 正常计算，返回空错误
        return dividend / divisor, nil
    }
    func main() {
        fmt.Println(div(1, 0))
    }
    
有一些函数总是成功返回的。比如，strings.Contains 和 strconv.FormatBool 对所有可能的参数变量都有定义好的返回结果，不会调用失败——尽管还有灾难性的和不可预知的场景，像内存耗尽，这类错误的表现和起因相差甚远而且恢复的希望也很渺茫。

其他的函数只要符合其前置条件就能够成功返回。比如函数始终会利用年、月等构成 time.Time，但是如果最后一个参数（表示时区）为 nil 则会导致宕机。这个宕机标志着这是一个明显的 bug，应该避免这样调用代码。

对于许多其他函数，即使在高质量的代码中，也不能保证一定能够成功返回，因为有些因素并不受程序设计者的掌控。比如任何操作 I/O 的函数都一定会面对可能的错误，只有没有经验的程序员会认为一个简单的读或写不会失败。事实上，这些地方是我们最需要关注的，很多可靠的操作都可能会毫无征兆地发生错误。

因此错误处理是包的 API 设计或者应用程序用户接口的重要部分，发生错误只是许多预料行为中的一种而已。这就是Go语言处理错误的方法。

如果当函数调用发生错误时返回一个附加的结果作为错误值，习惯上将错误值作为最后一个结果返回。如果错误只有一种情况，结果通常设置为布尔类型，就像下面这个查询缓存值的例子里面，往往都返回成功，只有不存在对应的键值的时候返回错误：

    value, ok := cache.Lookup(key)
    if !ok {
        // ...cache[key]不存在...
    }

更多时候，尤其对于 I/O 操作，错误的原因可能多种多样，而调用者则需要一些详细的信息。在这种情况下，错误的结果类型往往是 error。

error 是内置的接口类型。目前我们已经了解到，一个错误可能是空值或者非空值，空值意味着成功而非空值意味着失败，且非空的错误类型有一个错误消息字符串，可以通过调用它的 Error 方法或者通过调用 fmt.Println(err) 或 fmt.Printf("%v", err) 直接输出错误消息：

一般当一个函数返回一个非空错误时，它其他的结果都是未定义的而且应该忽略。然而，有一些函数在调用出错的情况下会返回部分有用的结果。比如，如果在读取一个文件的时候发生错误，调用 Read 函数后返回能够成功读取的字节数与相对应的错误值。正确的行为通常是在调用者处理错误前先处理这些不完整的返回结果。因此在文档中清晰地说明返回值的意义是很重要的。

与许多其他语言不同，Go语言通过使用普通的值而非异常来报告错误。尽管Go语言有异常机制，但是Go语言的异常只是针对程序 bug 导致的预料外的错误，而不能作为常规的错误处理方法出现在程序中。

这样做的原因是异常会陷入带有错误消息的控制流去处理它，通常会导致预期外的结果：错误会以难以理解的栈跟踪信息报告给最终用户，这些信息大都是关于程序结构方面的而不是简单明了的错误消息。

相比之下，Go 程序使用通常的控制流机制（比如 if 和 return 语句）应对错误。这种方式在错误处理逻辑方面要求更加小心谨慎，但这恰恰是设计的要点。

### 错误处理策略

当一个函数调用返回一个错误时，调用者应当负责检查错误并采取合适的处理应对。根据情形，将有许多可能的处理场景。接下来我们看 5 个例子。

首先也最常见的情形是将错误传递下去，使得在子例程中发生的错误变为主调例程的错误。《函数的多返回值》的一节中讨论过 findLinks 函数的示例。如果调用 http.Get 失败，findLinks 不做任何操作立即向调用者返回这个 HTTP 错误。

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }

对比之下，如果调用 html.Parse 失败，findLinks 将不会直接返回 HTML 解析的错误，因为它缺失两个关键信息：解析器的出错信息与被解析文档的 URL。在这种情况下，findLinks 构建一个新的错误消息，其中包含我们需要的所有相关信息和解析的错误信息：

    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        return nil, fmt.Errorf("parsing %s as HTML: .%v", url, err)
    }

fmt.Errorf 使用 fmt.Sprintf 函数格式化一条错误消息并且返回一个新的错误值。我们为原始的错误消息不断地添加额外的上下文信息来建立一个可读的错误描述。当错误最终被程序的 main 函数处理时，它应当能够提供一个从最根本问题到总体故障的清晰因果链，这让我想到 NASA 的事故调查有这样一个例子：

    genesis: crashed: no parachute: G.switch failed: bad relay orientation

因为错误消息频繁地串联起来，所以消息字符串首字母不应该大写而且应该避免换行。错误结果可能会很长，但能够使用 grep 这样的工具找到我们需要的信息。

设计一个错误消息的时候应当慎重，确保每一条消息的描述都是有意义的，包含充足的相关信息，并且保持一致性，不论被同一个函数还是同一个包下面的一组函数返回时，这样的错误都可以保持统一的形式和错误处理方式。

比如，OS 包保证每一个文件操作（比如 os.Open 或针对打开的文件的 Read、Write 或 Close 方法）返回的错误不仅包括错误的信息（没有权限、路径不存在等）还包含文件的名字，因此调用者在构造错误消息的时候不需要再包含这些信息。

一般地，f(x) 调用只负责报告函数的行为 f 和参数值 x，因为它们和错误的上下文相关。调用者负责添加进一步的信息，但是 f(x) 本身并不会，就像上面函数中 URL 和 html.Parse 的关系。

我们接下来看一下第二种错误处理策略。对于不固定或者不可预测的错误，在短暂的间隔后对操作进行重试是合乎情理的，超出一定的重试次数和限定的时间后再报错退出。

    // WaitForServer 尝试连接URL对应的服务器
    //在一分钟内使用指数退避策略进行重试
    //所有的尝试失败后返回错误
    func WaitForServer(url string) error {
        const timeout = 1 * time.Minute
        deadline := time.Now().Add(timeout)
        for tries := 0; time.Now().Before(deadline); tries++ {
            _, err := http.Head(url)
            if err == nil {
                return nil // 成功
            }
            log.Printf("server not responding (%s); retrying...", err)
            time.Sleep(time.Second << uint(tries))   //指数退避策略
        }
        return fmt.Errorf("server %s failed to respond after %s", url, timeout)
    }
第三，如果依旧不能顺利进行下去，调用者能够输出错误然后优雅地停止程序，但一般这样的处理应该留给主程序部分。通常库函数应当将错误传递给调用者，除非这个错误表示一个内部一致性错误，这意味着库内部存在 bug。

    // (In function main.)
    if err := WaitForServer(url); err != nil {
        fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
        os.Exit(1)
    }

一个更加方便的方法是通过调用 log.Fatalf 实现相同的效果。就和所有的日志函数一样，它默认会将时间和日期作为前缀添加到错误消息前。

    if err := WaitForServer(url); err != nil {
        log.Fatalf("Site is down: %v\n", err)
    }

默认的格式有助于长期运行的服务器，而对于交互式的命令行工具则意义不大：

    2006/01/02 15:04:05 Site is down: no such domain: bad.gopl.io

一种更吸引人的输岀方式是自己定义命令的名称作为 log 包的前缀，并且将日期和时间略去。

    log.SetPrefix("wait: ")
    log.SetFlags(0)

第四，在一些错误情况下，只记录下错误信息然后程序继续运行。同样地，可以选择使用 log 包来增加日志的常用前缀：

    if err := Ping(); err != nil {
        log.Printf("ping failed: %v; networking disabled", err)
    }

并且直接输出到标准错误流：

    if err := Ping(); err != nil {
        fmt.Fprintf(os.Stderr, "ping failed: %v; networking disabled\n", err)
    }

所有 log 函数都会为缺少换行符的日志补充一个换行符。

第五，在某些罕见的情况下我们可以直接安全地忽略掉整个日志：

    dir, err := ioutil.TempDir("", "scratch")
    if err != nil {
        return fmt.Errorf("failed to create temp dir: %v", err)
    }
    
    //...使用临时目录...
    os.RemoveAll(dir) //忽略错误，$TMPDIR 会被周期性删除

调用 os.RemoveAll 可能会失败，但程序忽略了这个错误，原因是操作系统会周期性地清理临时目录。在这个例子中，我们有意地抛弃了错误，但程序的逻辑看上去就和我们忘记去处理了一样。要习惯考虑到每一个函数调用可能发生的出错情况，当有意地忽略一个错误的时候，清楚地注释一下你的意图。

Go语言的错误处理有特定的规律。进行错误检查之后，检测到失败的情况往往都在成功之前。如果检测到的失败导致函数返回，成功的逻辑一般不会放在 else 块中而是在外层的作用域中。函数会有一种通常的形式，就是在开头有一连串的检查用来返回错误，之后跟着实际的函数体一直到最后。

### 文件结束标识

通常，最终用户会对函数返回的多种错误感兴趣而不是中间涉及的程序逻辑。偶尔，一个程序必须针对不同各种类的错误采取不同的措施。考虑如果要从一个文件中读取 n 个字节的数据。如果 n 是文件本身的长度，任何错误都代表操作失败。

另一方面，如果调用者反复地尝试读取固定大小的块直到文件耗尽，调用者必须把读取到文件尾的情况区别于遇到其他错误的操作。为此，io 包保证任何由文件结束引起的读取错误，始终都将会得到一个与众不同的错误 io.EOF，它的定义如下：
   
    package io
    import "errors"
    //当没有更多输入时，将会返回 EOF
    var EOF = errors.New("EOF")

调用者可以使用一个简单的比较操作来检测这种情况，在下面的循环中，不断从标准输入中读取字符。

    in := bufio.NewReader(os.Stdin)
    for {
        r, _, err := in.ReadRune()
        if err == io.EOF {
            break //结束读取
        }
        if err != nil {
            return fmt.Errorf("read failed: %v", err)
        }
        //...使用 r...
    }
    
除了反映这个实际情况外，因为文件结束的条件没有其他信息，所以 io.EOF 有一条固定的错误消息“EOF”。对于其他错误，我们可能需要同时得到错误相关的本质原因和数量信息，因此一个固定的错误值并不能满足我们的需求。