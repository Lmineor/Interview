<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [GMP分别是什么，分别有多少数量？](#gmp%E5%88%86%E5%88%AB%E6%98%AF%E4%BB%80%E4%B9%88%E5%88%86%E5%88%AB%E6%9C%89%E5%A4%9A%E5%B0%91%E6%95%B0%E9%87%8F)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# GMP分别是什么，分别有多少数量？
G：goroutine，go的协程，每个go关键字都会创建一个协程
M：machine，工作线程，在Go中称为Machine，数量对应真实的CPU数
P：process，包含运行Go代码所需要的必要资源，用来调度G和M之间的关联关系，其数量可以通过GOMAXPROCS来设置，默认为核心数

线程想运行任务就得获取 P，从 P 的本地队列获取 G，当 P 的本地队列为空时，M 也会尝试从全局队列或其他 P 的本地队列获取 G。M 运行 G，G 执行之后，M 会从 P 获取下一个 G，不断重复下去。


详细见：[GMP | mineor](https://www.mineor.xyz/posts/Go/gmp/)
