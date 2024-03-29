<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [需要了解的词](#%E9%9C%80%E8%A6%81%E4%BA%86%E8%A7%A3%E7%9A%84%E8%AF%8D)
  - [DMA](#dma)
  - [MMU](#mmu)
  - [虚拟内存](#%E8%99%9A%E6%8B%9F%E5%86%85%E5%AD%98)
  - [NFS](#nfs)
  - [copy-on-write](#copy-on-write)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

存储器是计算机的核心部件之一，在完全理想的状态下，存储器应该要同时具备以下三种特性：

- 速度足够快：存储器的存取速度应当快于 CPU 执行一条指令，这样 CPU 的效率才不会受限于存储器；
- 容量足够大：容量能够存储计算机所需的全部数据；
- 价格足够便宜：价格低廉，所有类型的计算机都能配备。

但是现实往往是残酷的，我们目前的计算机技术无法同时满足上述的三个条件，于是现代计算机的存储器设计采用了一种分层次的结构：

![Description](../src/storage_hier.jpg)

从顶至底,现代计算机里的存储器类型分别有:**寄存器**,**高速缓存**,**主存**和**硬盘**.这些存储器的速度逐级递减而容量主机递增.

存储速度最快的**寄存器**,速度快,价格贵,容量小,一般32位的cpu配置32*32Bit,64位cpu配置64*64,寄存器容量都小于1KB,且寄存器也必须通过软件自行管理.

第二层是**高速缓存**,也即我们平时了解的CPU高速缓存L1,L2,L3,一般L1是每个CPU独享,L3全部CPU共享,而L2则根据不同的架构设计会被设计成独享或者共享两种模式之一,比如Intel的多核心CPU采用共享L2而AMD采用独享L2.

第三层则是主存,RAM(Random Access Memory),是与CPU直接进行数据交换的内部存储器.它可以随时读写(刷新时除外),而且速度很快,通常作为操作系统或其他正在运行中的程序的临时资料存储介质.

磁盘:离cpu最远的一层,读写速度相差内存上百倍;针对磁盘操作的优化也非常多,如**零拷贝**,**direct I/O**,**异步I/O**等,都是为了提高系统的吞吐量;另外操作系统内核中也有**磁盘高速缓存区**,**PageCache**,**TLB**等,可以有效的减少磁盘的访问次数.

现实情况中,大部分系统在由小变大的过程中,最先出现瓶颈的就是I/O,尤其是在现代网络应用从 CPU 密集型转向了 I/O 密集型的大背景下，I/O越发成为大多数应用的性能瓶颈.

传统的Linux操作系统的标准I/O接口是基于数据拷贝操作的,即I/O 操作会导致数据在操作系统内核地址空间的缓冲区和用户进程地址空间定义的缓冲区之间进行传输.**设置缓冲区最大的好处是可以减少磁盘 I/O 的操作**，如果所请求的数据已经存放在操作系统的高速缓冲存储器中，那么就不需要再进行实际的物理磁盘 I/O 操作；然而传统的 Linux I/O 在数据传输过程中的数据拷贝操作深度依赖 CPU，也就是说 I/O 过程需要 CPU 去执行数据拷贝的操作，因此导致了极大的系统开销，限制了操作系统有效进行数据传输操作的能力。
这篇文章就从文件传输场景以及零拷贝技术深究 Linux I/O的发展过程、优化手段以及实际应用。

# 需要了解的词

## DMA
全称 Direct Memory Access，即直接存储器访问，是为了避免 CPU 在磁盘操作时承担过多的中断负载而设计的；在磁盘操作中，CPU 可将总线控制权交给 DMA 控制器，由 DMA 输出读写命令，直接控制 RAM 与 I/O 接口进行 DMA 传输，无需 CPU 直接控制传输，也没有中断处理方式那样保留现场和恢复现场过程，使得 CPU 的效率大大提高。
## MMU 
Memory Management Unit—内存管理单元，主要实现：
   - 竞争访问保护管理需求：需要严格的访问保护，动态管理哪些内存页/段或区，为哪些应用程序所用。这属于资源的竞争访问管理需求；
   - 高效的翻译转换管理需求：需要实现快速高效的映射翻译转换，否则系统的运行效率将会低下；
   - 高效的虚实内存交换需求：需要在实际的虚拟内存与物理内存进行内存页/段交换过程中快速高效。
- Page Cache为了避免每次读写文件时，都需要对硬盘进行读写操作，Linux 内核使用 页缓存（Page Cache） 机制来对文件中的数据进行缓存。
![Description](../src/page_cache.jpg)

此外,由于读取磁盘数据的时候,需要找到数据所在的位置，但是对于机械磁盘来说，就是通过磁头旋转到数据所在的扇区，再开始「顺序」读取数据，但是旋转磁头这个物理动作是非常耗时的，为了降低它的影响，PageCache 使用了「预读功能」。
比如，假设 read 方法每次只会读 32 KB 的字节，虽然 read 刚开始只会读 0 ～ 32 KB 的字节，但内核会把其后面的 32 ～ 64 KB 也读取到 PageCache，这样后面读取 32 ～ 64 KB 的成本就很低，如果在 32 ～ 64 KB 淘汰出 PageCache 前，有进程读取到它了，收益就非常大。

## 虚拟内存

在计算机领域有一句如同摩西十诫般神圣的哲言："计算机科学领域的任何问题都可以通过增加一个间接的中间层来解决"，从内存管理、网络模型、并发调度甚至是硬件架构，都能看到这句哲言在闪烁着光芒，而虚拟内存则是这一哲言的完美实践之一。
虚拟内存为每个进程提供了一个一致的、私有且连续完整的内存空间；所有现代操作系统都使用虚拟内存，使用虚拟地址取代物理地址，主要有以下几点好处：

利用上述的第一条特性可以优化，可以把内核空间和用户空间的虚拟地址映射到同一个物理地址，这样在 I/O 操作时就不需要来回复制了。
- 多个虚拟内存可以指向同一个物理地址
- 虚拟内存空间可以远远大于物理内存空间
- 应用层面可以管理连续的内存空间,减少出错

## NFS

网络文件系统是FreeBSD支持的文件系统的一种,也被称为NFS,NFS允许一个系统在网络上与它人共享目录和文件,通过使用NFS,用户和程序可以像访问本地文件一样访问远端系统上的文件.

## copy-on-write

写入时复制,cow,是一种计算机程序设计领域的优化策略.其核心思想是,如果有多个调用者(callers)同时请求相同的资源(如内存或磁盘上的数据存储),他们会共同获取相同的指针指向相同的资源,直到某个调用者试图修改资源的内容时,系统才会真正复制一份专用副本(private copy)给该调用者,而其他调用者所见到的最初的资源任然保持不变.这过程对其他的调用者都是透明的。此作法主要的优点是如果调用者没有修改该资源，就不会有副本（private copy）被创建，因此多个调用者只是读取操作时可以共享同一份资源。