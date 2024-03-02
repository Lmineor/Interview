

# Channel介绍
channel 是 goroutine 之间通信的一种方式，可以类比成 Unix 中的进程的通信方式管道。

## CSP 模型

在讲 channel 之前，有必要先提一下 CSP 模型，传统的并发模型主要分为 Actor 模型和 CSP 模型，CSP 模型全称为 communicating sequential processes，CSP 模型由并发执行实体(进程，线程或协程)，和消息通道组成，实体之间通过消息通道发送消息进行通信。和 Actor 模型不同，CSP 模型关注的是消息发送的载体，即通道，而不是发送消息的执行实体。关于 CSP 模型的更进一步的介绍，有兴趣的同学可以阅读论文 Communicating Sequential Processes，Go 语言的并发模型参考了 CSP 理论，其中执行实体对应的是 goroutine， 消息通道对应的就是 channel。

## Channel
channel 提供了一种通信机制，通过它，一个 goroutine 可以想另一 goroutine 发送消息。channel 本身还需关联了一个类型，也就是 channel 可以发送数据的类型。例如: 发送 int 类型消息的 channel 写作 chan int 。

### channel 创建

channel 使用内置的 make 函数创建，下面声明了一个 chan int 类型的 channel:

`ch := make(chan int)`

ch和 map 类似，make 创建了一个底层数据结构的引用，当赋值或参数传递时，只是拷贝了一个 channel 引用，指向相同的 channel 对象。和其他引用类型一样，channel 的空值为 nil 。使用 == 可以对类型相同的 channel 进行比较，只有指向相同对象或同为 nil 时，才返回 true。

### channel 的读写操作
```
ch := make(chan int)
// write to channel
ch <- x
// read from channel
x <- ch
// another way to read
x = <- ch
```
channel 一定要初始化后才能进行读写操作，否则会永久阻塞。
### 关闭 channel
golang 提供了内置的 close 函数对 channel 进行关闭操作。
```
ch := make(chan int)
close(ch)
```
有关 channel 的关闭，你需要注意以下事项:

- 关闭一个未初始化(nil) 的 channel 会产生 panic
- 重复关闭同一个 channel 会产生 panic
- 向一个已关闭的 channel 中发送消息会产生 panic
- 从已关闭的 channel 读取消息不会产生 panic，且能读出 channel 中还未被读取的消息，若消息均已读出，则会读到类型的零值。从一个已关闭的 channel 中读取消息永远不会阻塞，并且会返回一个为 false 的 ok-idiom，可以用它来判断 channel 是否关闭
- 关闭 channel 会产生一个广播机制，所有向 channel 读取消息的 goroutine 都会收到消息
```
ch := make(chan int, 10)  
ch <- 11  
ch <- 12

close(ch)
for x := range ch {  
    fmt.Println(x)  
}

x, ok := <- ch  
fmt.Println(x, ok)
```
-----  
output:

```
11  
12  
0 false
```
### channel 的类型

channel 分为不带缓存的 channel 和带缓存的 channel。

#### 无缓存的 channel

从无缓存的 channel 中读取消息会阻塞，直到有 goroutine 向该 channel 中发送消息；同理，向无缓存的 channel 中发送消息也会阻塞，直到有 goroutine 从 channel 中读取消息。

通过无缓存的 channel 进行通信时，接收者收到数据 happens before 发送者 goroutine 唤醒

#### 有缓存的 channel

有缓存的 channel 的声明方式为指定 make 函数的第二个参数，该参数为 channel 缓存的容量

`ch := make(chan int, 10)`

有缓存的 channel 类似一个阻塞队列(采用环形数组实现)。当缓存未满时，向 channel 中发送消息时不会阻塞，当缓存满时，发送操作将被阻塞，直到有其他 goroutine 从中读取消息；相应的，当 channel 中消息不为空时，读取消息不会出现阻塞，当 channel 为空时，读取操作会造成阻塞，直到有 goroutine 向 channel 中写入消息。
```
ch := make(chan int, 3)

// blocked, read from empty buffered channel  
<- ch

ch := make(chan int, 3)  
ch <- 1  
ch <- 2  
ch <- 3

// blocked, send to full buffered channel  
ch <- 4
```
通过 len 函数可以获得 chan 中的元素个数，通过 cap 函数可以得到 channel 的缓存长度。

### channel 的用法

goroutine 通信

看一个 effective go 中的例子：
```
c := make(chan int)  // Allocate a channel.

// Start the sort in a goroutine; when it completes, signal on the channel.  
go func() {  
    list.Sort()  
    c <- 1  // Send a signal; value does not matter.  
}()

doSomethingForAWhile()  
<-c
```
主 goroutine 会阻塞，直到执行 sort 的 goroutine 完成。

range 遍历

channel 也可以使用 range 取值，并且会一直从 channel 中读取数据，直到有 goroutine 对改 channel 执行 close 操作，循环才会结束。
```
// consumer worker  
ch := make(chan int, 10)  
for x := range ch{  
    fmt.Println(x)  
}
```

等价于
```
for {
    x, ok := <- ch
    if !ok {
        break
    }
    fmt.Println(x)
}
```
配合 select 使用

select 用法类似与 IO 多路复用，可以同时监听多个 channel 的消息状态，看下面的例子
```
select {  
    case <- ch1:  
    ...  
    case <- ch2:  
    ...  
    case ch3 <- 10;  
    ...  
    default:  
    ...  
}
```

- select 可以同时监听多个 channel 的写入或读取
- 执行 select 时，若只有一个 case 通过(不阻塞)，则执行这个 case 块
- 若有多个 case 通过，则随机挑选一个 case 执行
- 若所有 case 均阻塞，且定义了 default 模块，则执行 default 模块。若未定义 default 模块，则 select 语句阻塞，直到有 case 被唤醒。
- 使用 break 会跳出 select 块。

1. 设置超时时间
```
ch := make(chan struct{})

// finish task while send msg to ch  
go doTask(ch)

timeout := time.After(5 * time.Second)  
select {  
    case <- ch:  
        fmt.Println("task finished.")  
    case <- timeout:  
        fmt.Println("task timeout.")  
}
```

2. quite channel

有一些场景中，一些 worker goroutine 需要一直循环处理信息，直到收到 quit 信号

```
msgCh := make(chan struct{})  
quitCh := make(chan struct{})  
for {  
    select {  
    case <- msgCh:  
        doWork()  
    case <- quitCh:  
        finish()  
        return  
}
```

单向 channel

即只可写入或只可读的channel，事实上 channel 只读或只写都没有意义，所谓的单向 channel 其实只是声明时用，比如

`func foo(ch chan<- int) <-chan int {...}`

`chan<- int` 表示一个只可写入的 channel，`<-chan int` 表示一个只可读取的 channel。上面这个函数约定了 foo 内只能从向 ch 中写入数据，返回只一个只能读取的 channel，虽然使用普通的 channel 也没有问题，但这样在方法声明时约定可以防止 channel 被滥用，这种预防机制发生再编译期间。
## Channel是同步的还是异步的？
Channel是异步进行的, channel存在3种状态：

1. `nil`，未初始化的状态，只进行了声明，或者手动赋值为nil
2. `active`，正常的channel，可读或者可写
3. `closed`，已关闭，千万不要误认为关闭channel后，channel的值是nil，对已关闭channel写会panic，读不会

## Channel死锁场景
1. 非缓存channel只写不读
2. 非缓存channel读在写后面
3. 缓存channel写入超过缓冲区数量
4. 空读
5. 多个协程相互等待