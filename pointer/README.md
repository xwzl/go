<div id="arc-body">不像 <a href="/java/" target="_blank">Java</a> 和 .NET，<a href="/golang/" target="_blank">Go语言</a>为程序员提供了控制<a href="/data_structure/" target="_blank">数据结构</a>的指针的能力；但是，并不能进行指针运算。通过给予程序员基本内存布局，Go语言允许你控制特定集合的数据结构、分配的数量以及内存访问模式，这些对构建运行良好的系统是非常重要的：指针对于性能的影响是不言而喻的，而如果你想要做的是系统编程、操作系统或者网络应用，指针更是不可或缺的一部分。<br>
<br>
指针（pointer）在Go语言中被拆分为两个核心概念：
<ul>
<li>
类型指针，允许对这个指针类型的数据进行修改。传递数据使用指针，而无须拷贝数据。类型指针不能进行偏移和运算。</li>
<li>
切片，由指向起始元素的原始指针、元素数量和容量组成。</li>
</ul>
<br>
受益于这样的约束和拆分，Go语言的指针类型变量拥有指针的高效访问，但又不会发生指针偏移，从而避免非法修改关键性数据的问题。同时，垃圾回收也比较容易对不会发生偏移的指针进行检索和回收。<br>
<br>
切片比原始指针具备更强大的特性，更为安全。切片发生越界时，运行时会报出宕机，并打出堆栈，而原始指针只会崩溃。
<h4>
C/<a href="/cplus/" target="_blank">C++</a>中的指针</h4>
说到 C/C++ 中的指针，会让许多人“谈虎色变”，尤其对指针偏移、运算、转换都非常恐惧。<br>
<br>
其实，指针是使 C/C++ 语言有极高性能的根本，在操作大块数据和做偏移时方便又便捷。因此，操作系统依然使用<a href="/c/" target="_blank">C语言</a>及指针特性进行编写。<br>
<br>
C/C++ 中指针饱受诟病的根本原因是指针运算和内存释放。<br>
<br>
C/C++ 语言中的裸指针可以自由偏移，甚至可以在某些情况下偏移进入操作系统核心区域。我们的计算机操作系统经常需要更新、修复漏洞的本质，是为解决指针越界访问所导致的“缓冲区溢出”。<br>
<br>
要明白指针，需要知道几个概念：指针地址、指针类型和指针取值，下面将展开细说。
<h2>
认识指针地址和指针类型</h2>
一个指针变量可以指向任何一个值的内存地址，它指向那个值的内存地址，在 32 位机器上占用 4 个字节，在 64 位机器上占用 8 个字节，并且与它所指向的值的大小无关。当然，可以声明指针指向任何类型的值来表明它的原始性或结构性；你可以在指针类型前面加上 * 号（前缀）来获取指针所指向的内容，这里的 * 号是一个类型更改器。使用一个指针引用一个值被称为间接引用。<br>
<br>
当一个指针被定义后没有分配到任何变量时，它的值为 nil。一个指针变量通常缩写为 ptr。<br>
<br>
每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。Go语言中使用<code style="font-size: 14px;">&amp;</code>操作符放在变量前面对变量进行“取地址”操作。<br>
<br>
格式如下：
<p class="info-box">
ptr := &amp;v &nbsp; &nbsp;// v的类型为T</p>
其中 v 代表被取地址的变量，被取地址的 v 使用 ptr 变量进行接收，ptr 的类型就为<code style="font-size: 14px;">*T</code>，称做 T 的指针类型。<code style="font-size: 14px;">*</code>代表指针。<br>
<br>
指针实际用法，通过下面的例子了解：
<div class="snippet-container" style="undefined;"><div class="sh_bright snippet-wrap"><div class="snippet-menu sh_sourceCode" style="display: none;"><pre><a class="snippet-copy sh_url" href="#" style="display: none;">复制</a><a class="snippet-text sh_url" href="#">纯文本</a><a class="snippet-window sh_url" href="#">复制</a></pre></div><pre class="go sh_go snippet-formatted"><ol class="snippet-num"><li>package main</li><li><span style="display:none;">&nbsp;</span></li><li>import (</li><li>    "fmt"</li><li>)</li><li><span style="display:none;">&nbsp;</span></li><li>func main() {</li><li>    var cat int = 1</li><li>    var str string = "banana"</li><li>    fmt.Printf("%p %p", &amp;cat, &amp;str)</li><li>}</li></ol></pre><pre class="snippet-textonly sh_sourceCode" style="display:none;">package main

import (
    "fmt"
)

func main() {
    var cat int = 1
    var str string = "banana"
    fmt.Printf("%p %p", &amp;cat, &amp;str)
}</pre></div></div>
运行结果：
<p class="info-box">
0xc042052088 0xc0420461b0</p>
代码说明如下：
<ul>
<li>
第 8 行，声明整型 cat 变量。</li>
<li>
第 9 行，声明字符串 str 变量。</li>
<li>
第 10 行，使用 fmt.Printf 的动词<code style="font-size: 14px;">%p</code>输出 cat 和 str 变量取地址后的指针值，指针值带有<code style="font-size: 14px;">0x</code>的十六进制前缀。</li>
</ul>
<br>
输出值在每次运行是不同的，代表 cat 和 str 两个变量在运行时的地址。<br>
<br>
在 32 位平台上，将是 32 位地址；64 位平台上是 64 位地址。<br>
<br>
提示：变量、指针和地址三者的关系是：每个变量都拥有地址，指针的值就是地址。
<h2>
从指针获取指针指向的值</h2>
在对普通变量使用<code style="font-size: 14px;">&amp;</code>操作符取地址获得这个变量的指针后，可以对指针使用<code style="font-size: 14px;">*</code>操作，也就是指针取值，代码如下。
<div class="snippet-container" style="undefined;"><div class="sh_bright snippet-wrap"><div class="snippet-menu sh_sourceCode" style="display: none;"><pre><a class="snippet-copy sh_url" href="#" style="display: none;">复制</a><a class="snippet-text sh_url" href="#">纯文本</a><a class="snippet-window sh_url" href="#">复制</a></pre></div><pre class="go sh_go snippet-formatted"><ol class="snippet-num"><li>package main</li><li><span style="display:none;">&nbsp;</span></li><li>import (</li><li>    "fmt"</li><li>)</li><li><span style="display:none;">&nbsp;</span></li><li>func main() {</li><li><span style="display:none;">&nbsp;</span></li><li>    // 准备一个字符串类型</li><li>    var house = "Malibu Point 10880, 90265"</li><li><span style="display:none;">&nbsp;</span></li><li>    // 对字符串取地址, ptr类型为*string</li><li>    ptr := &amp;house</li><li><span style="display:none;">&nbsp;</span></li><li>    // 打印ptr的类型</li><li>    fmt.Printf("ptr type: %T\n", ptr)</li><li><span style="display:none;">&nbsp;</span></li><li>    // 打印ptr的指针地址</li><li>    fmt.Printf("address: %p\n", ptr)</li><li><span style="display:none;">&nbsp;</span></li><li>    // 对指针进行取值操作</li><li>    value := *ptr</li><li><span style="display:none;">&nbsp;</span></li><li>    // 取值后的类型</li><li>    fmt.Printf("value type: %T\n", value)</li><li><span style="display:none;">&nbsp;</span></li><li>    // 指针取值后就是指向变量的值</li><li>    fmt.Printf("value: %s\n", value)</li><li><span style="display:none;">&nbsp;</span></li><li>}</li></ol></pre><pre class="snippet-textonly sh_sourceCode" style="display:none;">package main

import (
    "fmt"
)

func main() {

    // 准备一个字符串类型
    var house = "Malibu Point 10880, 90265"

    // 对字符串取地址, ptr类型为*string
    ptr := &amp;house

    // 打印ptr的类型
    fmt.Printf("ptr type: %T\n", ptr)

    // 打印ptr的指针地址
    fmt.Printf("address: %p\n", ptr)

    // 对指针进行取值操作
    value := *ptr

    // 取值后的类型
    fmt.Printf("value type: %T\n", value)

    // 指针取值后就是指向变量的值
    fmt.Printf("value: %s\n", value)

}</pre></div></div>
运行结果：
<p class="info-box">
ptr type: *string<br>
address: 0xc0420401b0<br>
value type: string<br>
value: Malibu Point 10880, 90265</p>
代码说明如下：
<ul>
<li>
第 10 行，准备一个字符串并赋值。</li>
<li>
第 13 行，对字符串取地址，将指针保存到 ptr 中。</li>
<li>
第 16 行，打印 ptr 变量的类型，类型为 *string。</li>
<li>
第 19 行，打印 ptr 的指针地址，每次运行都会发生变化。</li>
<li>
第 22 行，对 ptr 指针变量进行取值操作，value 变量类型为 string。</li>
<li>
第 25 行，打印取值后 value 的类型。</li>
<li>
第 28 行，打印 value 的值。</li>
</ul>
<br>
取地址操作符<code style="font-size: 14px;">&amp;</code>和取值操作符<code style="font-size: 14px;">*</code>是一对互补操作符，<code style="font-size: 14px;">&amp;</code>取出地址，<code style="font-size: 14px;">*</code>根据地址取出地址指向的值。<br>
<br>
变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
<ul>
<li>
对变量进行取地址（&amp;）操作，可以获得这个变量的指针变量。</li>
<li>
指针变量的值是指针地址。</li>
<li>
对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。</li>
</ul>
<h2>
使用指针修改值</h2>
通过指针不仅可以取值，也可以修改值。<br>
<br>
前面已经使用多重赋值的方法进行数值交换，使用指针同样可以进行数值交换，代码如下：
<div class="snippet-container" style="undefined;"><div class="sh_bright snippet-wrap"><div class="snippet-menu sh_sourceCode" style="display: none;"><pre><a class="snippet-copy sh_url" href="#" style="display: none;">复制</a><a class="snippet-text sh_url" href="#">纯文本</a><a class="snippet-window sh_url" href="#">复制</a></pre></div><pre class="go sh_go snippet-formatted"><ol class="snippet-num"><li>package main</li><li><span style="display:none;">&nbsp;</span></li><li>import "fmt"</li><li><span style="display:none;">&nbsp;</span></li><li>// 交换函数</li><li>func swap(a, b *int) {</li><li><span style="display:none;">&nbsp;</span></li><li>    // 取a指针的值, 赋给临时变量t</li><li>    t := *a</li><li><span style="display:none;">&nbsp;</span></li><li>    // 取b指针的值, 赋给a指针指向的变量</li><li>    *a = *b</li><li><span style="display:none;">&nbsp;</span></li><li>    // 将a指针的值赋给b指针指向的变量</li><li>    *b = t</li><li>}</li><li><span style="display:none;">&nbsp;</span></li><li>func main() {</li><li><span style="display:none;">&nbsp;</span></li><li>// 准备两个变量, 赋值1和2</li><li>    x, y := 1, 2</li><li><span style="display:none;">&nbsp;</span></li><li>    // 交换变量值</li><li>    swap(&amp;x, &amp;y)</li><li><span style="display:none;">&nbsp;</span></li><li>    // 输出变量值</li><li>    fmt.Println(x, y)</li><li>}</li></ol></pre><pre class="snippet-textonly sh_sourceCode" style="display:none;">package main

import "fmt"

// 交换函数
func swap(a, b *int) {

    // 取a指针的值, 赋给临时变量t
    t := *a

    // 取b指针的值, 赋给a指针指向的变量
    *a = *b

    // 将a指针的值赋给b指针指向的变量
    *b = t
}

func main() {

// 准备两个变量, 赋值1和2
    x, y := 1, 2

    // 交换变量值
    swap(&amp;x, &amp;y)

    // 输出变量值
    fmt.Println(x, y)
}</pre></div></div>
运行结果：
<p class="info-box">
2 1</p>
代码说明如下：
<ul>
<li>
第 6 行，定义一个交换函数，参数为 a、b，类型都为 *int，都是指针类型。</li>
<li>
第 9 行，将 a 指针取值，把值（int类型）赋给 t 变量，t 此时也是 int 类型。</li>
<li>
第 12 行，取 b 指针值，赋给 a 变量指向的变量。注意，此时<code style="font-size: 14px;">*a</code>的意思不是取 a 指针的值，而是“a指向的变量”。</li>
<li>
第 15 行，将 t 的值赋给 b 指向的变量。</li>
<li>
第 21 行，准备 x、y 两个变量，赋值 1 和 2，类型为 int。</li>
<li>
第 24 行，取出 x 和 y 的地址作为参数传给 swap() 函数进行调用。</li>
<li>
第 27 行，交换完毕时，输出 x 和 y 的值。</li>
</ul>
<br>
<code style="font-size: 14px;">*</code>操作符作为右值时，意义是取指针的值；作为左值时，也就是放在赋值操作符的左边时，表示 a 指向的变量。其实归纳起来，<code style="font-size: 14px;">*</code>操作符的根本意义就是操作指针指向的变量。当操作在右值时，就是取指向变量的值；当操作在左值时，就是将值设置给指向的变量。<br>
<br>
如果在 swap() 函数中交换操作的是指针值，会发生什么情况？可以参考下面代码：
<div class="snippet-container" style="undefined;"><div class="sh_bright snippet-wrap"><div class="snippet-menu sh_sourceCode" style="display:none;"><pre><a class="snippet-copy sh_url" href="#" style="display: none;">复制</a><a class="snippet-text sh_url" href="#">纯文本</a><a class="snippet-window sh_url" href="#">复制</a></pre></div><pre class="go sh_go snippet-formatted"><ol class="snippet-num"><li>package main</li><li><span style="display:none;">&nbsp;</span></li><li>import "fmt"</li><li><span style="display:none;">&nbsp;</span></li><li>func swap(a, b *int) {</li><li>    b, a = a, b</li><li>}</li><li><span style="display:none;">&nbsp;</span></li><li>func main() {</li><li>    x, y := 1, 2</li><li>    swap(&amp;x, &amp;y)</li><li>    fmt.Println(x, y)</li><li>}</li></ol></pre><pre class="snippet-textonly sh_sourceCode" style="display:none;">package main

import "fmt"

func swap(a, b *int) {
    b, a = a, b
}

func main() {
    x, y := 1, 2
    swap(&amp;x, &amp;y)
    fmt.Println(x, y)
}</pre></div></div>
运行结果：
<p class="info-box">
1 2</p>
结果表明，交换是不成功的。上面代码中的 swap() 函数交换的是 a 和 b 的地址，在交换完毕后，a 和 b 的变量值确实被交换。但和 a、b 关联的两个变量并没有实际关联。这就像写有两座房子的卡片放在桌上一字摊开，交换两座房子的卡片后并不会对两座房子有任何影响。
<h2>
示例：使用指针变量获取命令行的输入信息</h2>
Go语言的 flag 包中，定义的指令以指针类型返回。通过学习 flag 包，可以深入了解指针变量在设计上的方便之处。<br>
<br>
下面的代码通过提前定义一些命令行指令和对应变量，在运行时，输入对应参数的命令行参数后，经过 flag 包的解析后即可通过定义的变量获取命令行的数据。<br>
<br>
获取命令行输入：
<div class="snippet-container" style="undefined;"><div class="sh_bright snippet-wrap"><div class="snippet-menu sh_sourceCode" style="display:none;"><pre><a class="snippet-copy sh_url" href="#" style="display: none;">复制</a><a class="snippet-text sh_url" href="#">纯文本</a><a class="snippet-window sh_url" href="#">复制</a></pre></div><pre class="go sh_go snippet-formatted"><ol class="snippet-num"><li>package main</li><li><span style="display:none;">&nbsp;</span></li><li>// 导入系统包</li><li>import (</li><li>    "flag"</li><li>    "fmt"</li><li>)</li><li><span style="display:none;">&nbsp;</span></li><li>// 定义命令行参数</li><li>var mode = flag.String("mode", "", "process mode")</li><li><span style="display:none;">&nbsp;</span></li><li>func main() {</li><li><span style="display:none;">&nbsp;</span></li><li>    // 解析命令行参数</li><li>    flag.Parse()</li><li><span style="display:none;">&nbsp;</span></li><li>    // 输出命令行参数</li><li>    fmt.Println(*mode)</li><li>}</li></ol></pre><pre class="snippet-textonly sh_sourceCode" style="display:none;">package main

// 导入系统包
import (
    "flag"
    "fmt"
)

// 定义命令行参数
var mode = flag.String("mode", "", "process mode")

func main() {

    // 解析命令行参数
    flag.Parse()

    // 输出命令行参数
    fmt.Println(*mode)
}</pre></div></div>
将这段代码命名为main.go，然后使用如下命令行运行：
<p class="info-box">
$ go run flagparse.go --mode=fast</p>
命令行输出结果如下：
<p class="info-box">
fast</p>
代码说明如下：
<ul>
<li>
第 10 行，通过 flag.String，定义一个 mode 变量，这个变量的类型是 *string。后面 3 个参数分别如下：
<ul>
<li>
参数名称：在给应用输入参数时，使用这个名称。</li>
<li>
参数值的默认值：与 flag 所使用的函数创建变量类型对应，String 对应字符串、Int 对应整型、Bool 对应布尔型等。</li>
<li>
参数说明：使用 -help 时，会出现在说明中。</li>
</ul>
</li>
<li>
第 15 行，解析命令行参数，并将结果写入创建的指令变量中，这个例子中就是 mode 变量。</li>
<li>
第 18 行，打印 mode 指针所指向的变量。</li>
</ul>
<br>
由于之前使用 flag.String 已经注册了一个 mode 的命令行参数，flag 底层知道怎么解析命令行，并且将值赋给 mode*string 指针。在 Parse 调用完毕后，无须从 flag 获取值，而是通过自己注册的 mode 这个指针，获取到最终的值。代码运行流程如下图所示。<br>
<br>
<div style="text-align: center;">
<img alt="" src="/uploads/allimg/180813/1-1PQ311430K50.jpg"><br>
图：命令行参数与变量的关系</div>
<h2>
创建指针的另一种方法——new() 函数</h2>
Go语言还提供了另外一种方法来创建指针变量，格式如下：
<p class="info-box">
new(类型)</p>
一般这样写：
<div class="snippet-container" style="undefined;"><div class="sh_bright snippet-wrap"><div class="snippet-menu sh_sourceCode" style="display:none;"><pre><a class="snippet-copy sh_url" href="#" style="display: none;">复制</a><a class="snippet-text sh_url" href="#">纯文本</a><a class="snippet-window sh_url" href="#">复制</a></pre></div><pre class="go sh_go snippet-formatted"><ol class="snippet-num"><li>str := new(string)</li><li>*str = "ninja"</li><li><span style="display:none;">&nbsp;</span></li><li>fmt.Println(*str)</li></ol></pre><pre class="snippet-textonly sh_sourceCode" style="display:none;">str := new(string)
*str = "ninja"

fmt.Println(*str)</pre></div></div>
new() 函数可以创建一个对应类型的指针，创建过程会分配内存。被创建的指针指向的值为默认值。</div>