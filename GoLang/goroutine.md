<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [1. Goroutine](#1-goroutine)
- [2. Goroutine调度策略](#2-goroutine%E8%B0%83%E5%BA%A6%E7%AD%96%E7%95%A5)
- [3. Groutine的切换时机](#3-groutine%E7%9A%84%E5%88%87%E6%8D%A2%E6%97%B6%E6%9C%BA)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->


# 1. Goroutine

Go 为了提供更容易使用的并发方法，使用了 goroutine 和 channel。goroutine 来自协程的概念，让一组可复用的函数运行在一组线程之上，即使有协程阻塞，该线程的其他协程也可以被 runtime 调度，转移到其他可运行的线程上。最关键的是，程序员看不到这些底层的细节，这就降低了编程的难度，提供了更容易的并发。
Go 中，协程被称为 goroutine，它非常轻量，一个 goroutine 只占几 KB，并且这几 KB 就足够 goroutine 运行完，这就能在有限的内存空间内支持大量 goroutine，支持了更多的并发。虽然一个 goroutine 的栈只占几 KB，但实际是可伸缩的，如果需要更多内容，runtime 会自动为 goroutine 分配。

Goroutine 特点：
- 占用内存更小（2KB左右，系统线程需要1-8MB）
- 调度更灵活（runtime 调度）

# 2. Goroutine调度策略

1. 队列轮转：P会周期性的将G调度到M中执行，执行一段时间后，保存上下文，将G放到队列尾部，然后从队列中再取出一个G进行调度，P还会周期性的查看全局队列是否有G等待调度到M中执行
2. 系统调用：当G0即将进入系统调用时，M0将释放P，进而某个空闲的M1获取P，继续执行P队列中剩下的G。M1的来源有可能是M的缓存池，也可能是新建的。
3. 当G0系统调用结束后，如果有空闲的P，则获取一个P，继续执行G0。如果没有，则将G0放入全局队列，等待被其他的P调度。然后M0将进入缓存池睡眠。
# 3. Groutine的切换时机

1. select操作阻塞时  
2. io阻塞  
3. 阻塞在channel  
4. 程序员显示编码操作  
5. 等待锁  
6. 程序调用
