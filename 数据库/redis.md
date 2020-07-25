<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Redis 专题](#redis-%E4%B8%93%E9%A2%98)
  - [0 Redis是什么？](#0-redis%E6%98%AF%E4%BB%80%E4%B9%88)
    - [Redis数据库](#redis%E6%95%B0%E6%8D%AE%E5%BA%93)
    - [Redis优点](#redis%E4%BC%98%E7%82%B9)
    - [Redis缺点](#redis%E7%BC%BA%E7%82%B9)
  - [1 五种数据类型](#1-%E4%BA%94%E7%A7%8D%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8B)
    - [数据类型应用场景总结：](#%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8B%E5%BA%94%E7%94%A8%E5%9C%BA%E6%99%AF%E6%80%BB%E7%BB%93)
  - [2 Redis 持久化](#2-redis-%E6%8C%81%E4%B9%85%E5%8C%96)
    - [RDB 是怎么工作的](#rdb-%E6%98%AF%E6%80%8E%E4%B9%88%E5%B7%A5%E4%BD%9C%E7%9A%84)
    - [AOF 是怎么工作的](#aof-%E6%98%AF%E6%80%8E%E4%B9%88%E5%B7%A5%E4%BD%9C%E7%9A%84)
    - [AOF和RDB选择](#aof%E5%92%8Crdb%E9%80%89%E6%8B%A9)
    - [redis的持久化开启了RDB和AOF下重启服务是如何加载的](#redis%E7%9A%84%E6%8C%81%E4%B9%85%E5%8C%96%E5%BC%80%E5%90%AF%E4%BA%86rdb%E5%92%8Caof%E4%B8%8B%E9%87%8D%E5%90%AF%E6%9C%8D%E5%8A%A1%E6%98%AF%E5%A6%82%E4%BD%95%E5%8A%A0%E8%BD%BD%E7%9A%84)
  - [3 Redis 应用](#3-redis-%E5%BA%94%E7%94%A8)
  - [4 Redis为什么这么快](#4-redis%E4%B8%BA%E4%BB%80%E4%B9%88%E8%BF%99%E4%B9%88%E5%BF%AB)
  - [5 Redis 和 Memcached 的区别](#5-redis-%E5%92%8C-memcached-%E7%9A%84%E5%8C%BA%E5%88%AB)
  - [6 Redis 的淘汰策略](#6-redis-%E7%9A%84%E6%B7%98%E6%B1%B0%E7%AD%96%E7%95%A5)
  - [Redis部署](#redis%E9%83%A8%E7%BD%B2)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Redis 专题
## 0 Redis是什么？

1. 是一个完全开源免费的key-value内存数据库 
2. 通常被认为是一个数据结构服务器，主要是因为其有着丰富的数据结构 strings、map、 list、sets、 sorted sets

### Redis数据库

> ​	通常局限点来说，Redis也以消息队列的形式存在，作为内嵌的List存在，满足实时的高并发需求。在使用缓存的时候，redis比memcached具有更多的优势，并且支持更多的数据类型，把redis当作一个中间存储系统，用来处理高并发的数据库操作

### Redis优点
- **速度快**：使用标准C写，所有数据都在内存中完成，读写速度分别达到10万/20万 
- **持久化**：对数据的更新采用Copy-on-write技术，可以异步地保存到磁盘上，主要有两种策略，一是根据时间，更新次数的快照（save 300 10 ）二是基于语句追加方式(Append-only file，aof) 
- **自动操作**：对不同数据类型的操作都是自动的，很安全 
- **快速的主-从复制**，官方提供了一个数据，Slave在21秒即完成了对Amazon网站10G key set的复制。 
- **Sharding技术**： 很容易将数据分布到多个Redis实例中，数据库的扩展是个永恒的话题，在关系型数据库中，主要是以添加硬件、以分区为主要技术形式的纵向扩展解决了很多的应用场景，但随着web2.0、移动互联网、云计算等应用的兴起，这种扩展模式已经不太适合了，所以近年来，像采用主从配置、数据库复制形式的，Sharding这种技术把负载分布到多个特理节点上去的横向扩展方式用处越来越多。
### Redis缺点

- 是数据库容量受到物理内存的限制,不能用作海量数据的高性能读写,因此Redis适合的场景主要局限在较小数据量的高性能操作和运算上。
- Redis较难支持在线扩容，在集群容量达到上限时在线扩容会变得很复杂。为避免这一问题，运维人员在系统上线时必须确保有足够的空间，这对资源造成了很大的浪费。

## 1 五种数据类型

但是在说之前，我觉得有必要先来了解下 Redis 内部内存管理是如何描述这 5 种数据类型的。

![类型示意](../src/redis.jpg)

首先 Redis 内部使用一个 `redisObject` 对象来表示所有的 `key` 和 `value`。

`redisObject` 最主要的信息如上图所示：`type` 表示一个 `value` 对象具体是何种数据类型，`encoding` 是不同数据类型在 Redis 内部的存储方式。
比如：`type=string` 表示 `value` 存储的是一个普通字符串，那么 `encoding` 可以是 `raw` 或者 `int`。
下面简单说下 5 种数据类型：
①**String** 是 Redis 最基本的类型，可以理解成与 Memcached一模一样的类型，一个 Key 对应一个 Value。Value 不仅是 String，也可以是数字。
String 类型是二进制安全的，意思是 Redis 的 String 类型可以包含任何数据，比如 jpg 图片或者序列化的对象。String 类型的值最大能存储 512M。
②**Hash**是一个键值（key-value）的集合。Redis 的 Hash 是一个 String 的 Key 和 Value 的映射表，Hash 特别适合存储对象。常用命令：`hget`，`hset`，`hgetall` 等。
③**List**列表是简单的字符串列表，按照插入顺序排序。可以添加一个元素到列表的头部（左边）或者尾部（右边） 常用命令：`lpush`、`rpush`、`lpop`、`rpop`、`lrange`（获取列表片段）等。
> 应用场景：List 应用场景非常多，也是 Redis 最重要的数据结构之一，比如 Twitter 的关注列表，粉丝列表都可以用 List 结构来实现。

>数据结构：List 就是链表，可以用来当消息队列用。Redis 提供了 List 的 Push 和 Pop 操作，还提供了操作某一段的 API，可以直接查询或者删除某一段的元素。

>实现方式：Redis List 的是实现是一个双向链表，既可以支持反向查找和遍历，更方便操作，不过带来了额外的内存开销。

④**Set** 是 String 类型的无序集合。集合是通过 hashtable 实现的。Set 中的元素是没有顺序的，而且是没有重复的。常用命令：`sdd`、`spop`、`smembers`、`sunion` 等。

> 应用场景：Redis Set 对外提供的功能和 List 一样是一个列表，特殊之处在于 Set 是自动去重的，而且 Set 提供了判断某个成员是否在一个 Set 集合中。

⑤**Zset** 和 Set 一样是 String 类型元素的集合，且不允许重复的元素。常用命令：`zadd`、`zrange`、`zrem`、`zcard` 等。

> 使用场景：Sorted Set 可以通过用户额外提供一个优先级（score）的参数来为成员排序，并且是插入有序的，即自动排序。
当你需要一个有序的并且不重复的集合列表，那么可以选择 Sorted Set 结构。
和 Set 相比，Sorted Set关联了一个 Double 类型权重的参数 Score，使得集合中的元素能够按照 Score 进行有序排列，Redis 正是通过分数来为集合中的成员进行从小到大的排序。

> 实现方式：Redis Sorted Set 的内部使用 HashMap 和跳跃表（skipList）来保证数据的存储和有序，HashMap 里放的是成员到 Score 的映射。
而跳跃表里存放的是所有的成员，排序依据是 HashMap 里存的 Score，使用跳跃表的结构可以获得比较高的查找效率，并且在实现上比较简单。

### 数据类型应用场景总结：
|   类型    |   简介    |   特性    |    场景   |
|   -       |   -       |   -       |   -       |
|string(字符串)|二进制安全|可以包含任何数据，比如jpg图片或者序列化对象| ... |
|Hash(字典)|键值对集合，（dict in Python）|适合存储对象，并且可以像数据库中的update一个属性一样只修改某项属性值| 存储、读取、修改用户属性 |
|List(列表)|链表(双向链表)|增删快，提供了操作某一元素的api| 最新消息排行；消息队列 |
|Set(集合)|hash表实现，元素不重复|添加、删除、查找的复杂度都是O(1)，提供了求交集、并集、差集的操作| 共同好友；利用唯一性，统计访问网站的ip |
|Zset(有序集合)|将set中的元素增加一个权重参数score，元素按score有序排列|数据插入集合时，已经进行了天然排序| 排行榜；带权重的消息队列 |



## 2 Redis 持久化

- RDB（Redis DataBase）：快照形式是直接把内存中的数据保存到一个 dump 的文件中，定时保存，保存策略。
- AOF（Append Only File）：把所有的对 Redis 的服务器进行修改的命令都存到一个文件里，命令的集合。Redis 默认是快照 RDB 的持久化方式。

### RDB 是怎么工作的

默认 Redis 是会以快照"RDB"的形式将数据持久化到磁盘的一个二进制文件 dump.rdb。

工作原理简单说一下：当 Redis 需要做持久化时，Redis 会 fork 一个子进程，子进程将数据写到磁盘上一个临时 RDB 文件中。当子进程完成写临时文件后，将原来的 RDB 替换掉，这样的好处是可以 copy-on-write。

- RDB 的优点是：这种文件非常适合用于备份：比如，你可以在最近的 24 小时内，每小时备份一次，并且在每个月的每一天也备份一个 RDB 文件。这样的话，即使遇上问题，也可以随时将数据集还原到不同的版本。RDB 非常适合灾难恢复。
- RDB 的缺点是：如果你需要尽量避免在服务器故障时丢失数据，那么RDB不合适你。

### AOF 是怎么工作的

使用 AOF 做持久化，每一个写命令都通过 write 函数追加到 appendonly.aof 中，配置方式如下：

```shell
appendfsyncyesappendfsync always    #每次有数据修改发生时都会写入AOF文件。
appendfsync everysec                #每秒钟同步一次，该策略为AOF的缺省策略。
```

AOF 可以做到全程持久化，只需要在配置中开启 appendonly yes。这样 Redis 每执行一个修改数据的命令，都会把它添加到 AOF 文件中，当 Redis 重启时，将会读取 AOF 文件进行重放，恢复到 Redis 关闭前的最后时刻。
- 优点是会让 Redis 变得非常耐久。可以设置不同的 Fsync 策略，AOF的默认策略是每秒钟 Fsync 一次，在这种配置下，就算发生故障停机，也最多丢失一秒钟的数据。
- 缺点是对于相同的数据集来说，AOF 的文件体积通常要大于 RDB 文件的体积。根据所使用的 Fsync 策略，AOF 的速度可能会慢于 RDB。

### AOF和RDB选择
- 如果你非常关心你的数据，但仍然可以承受数分钟内的数据丢失，那么可以只使用 RDB 持久。
- AOF 将 Redis 执行的每一条命令追加到磁盘中，处理巨大的写入会降低Redis的性能
- 数据库备份和灾难恢复：定时生成 RDB 快照非常便于进行数据库备份，并且 RDB 恢复数据集的速度也要比 AOF 恢复的速度快。

当然了，Redis 支持同时开启 RDB 和 AOF，系统重启后，Redis 会优先使用 AOF 来恢复数据，这样丢失的数据会最少。
### redis的持久化开启了RDB和AOF下重启服务是如何加载的

1） AOF持久化开启且存在AOF文件时，优先加载AOF文件，
2） AOF关闭或者AOF文件不存在时，加载RDB文件，
3） 加载AOF/RDB文件成功后，Redis启动成功。
4） AOF/RDB文件存在错误时，Redis启动失败并打印错误信息

## 3 Redis 应用
**bitmap** 简介及相关应用
1 统计某用户登录天数
2 查看活跃用户总数

https://blog.csdn.net/ctwctw/article/details/105013817

## 4 Redis为什么这么快

- **Redis 完全基于内存**，绝大部分请求是纯粹的内存操作，非常迅速，数据存在内存中，类似于 HashMap，HashMap 的优势就是查找和操作的时间复杂度是 O(1)。
- **数据结构简单**，对数据操作也简单。
- **采用单线程**，避免了不必要的上下文切换和竞争条件，不存在多线程导致的 CPU 切换，不用去考虑各种锁的问题，不存在加锁释放锁操作，没有死锁问题导致的性能消耗。
- 使用**多路复用 IO**模型，非阻塞 IO。
- 使用**底层模型不同**，它们之间底层实现方式以及与客户端之间通信的应用协议不一样，Redis直接自己构建了VM 机制 ，因为一般的系统调用系统函数的话，会浪费一定的时间去移动和请求；

[参考资料](https://juejin.im/entry/5b7cfe976fb9a01a13366d95)

> 多路 I/O 复用模型

多路I/O复用模型是利用 select、poll、epoll 可以同时监察多个流的 I/O 事件的能力，在空闲的时候，会把当前线程阻塞掉，当有一个或多个流有 I/O 事件时，就从阻塞态中唤醒，于是程序就会轮询一遍所有的流（epoll 是只轮询那些真正发出了事件的流），并且只依次顺序的处理就绪的流，这种做法就避免了大量的无用操作。

这里“多路”指的是多个网络连接，“复用”指的是复用同一个线程。

采用多路 I/O 复用技术可以让单个线程高效的处理多个连接请求（尽量减少网络 IO 的时间消耗），且 Redis 在内存中操作数据的速度非常快，也就是说内存内的操作不会成为影响Redis性能的瓶颈，主要由以上几点造就了 Redis 具有很高的吞吐量。

## 5 Redis 和 Memcached 的区别
- 存储方式上：Memcache 会把数据全部存在内存之中，断电后会挂掉，数据不能超过内存大小。Redis 有部分数据存在硬盘上，这样能保证数据的持久性。
- 数据支持类型上：Memcache 对数据类型的支持简单，只支持简单的 key-value，，而 Redis 支持五种数据类型。
- 使用底层模型不同：它们之间底层实现方式以及与客户端之间通信的应用协议不一样。Redis 直接自己构建了 VM 机制，因为一般的系统调用系统函数的话，会浪费一定的时间去移动和请求。
- Value 的大小：Redis 可以达到 1GB，而 Memcache 只有 1MB。

## 6 Redis 的淘汰策略

1. **volatile-lru**：从设置过期时间的数据集(`server.db[i].expires`)中挑选出最近最少使用的数据淘汰。没有设置过期时间的key不会被淘汰，这样就可以在增加内存空间的同时保证需要持久化的数据不会丢失。

2. **volatile-ttl**：除了淘汰机制采用LRU，策略基本上与volatile-lru相似，从设置过期时间的数据集(`server.db[i].expires`)中挑选将要过期的数据淘汰，ttl值越大越优先被淘汰。

3. **volatile-random**：从已设置过期时间的数据集(`server.db[i].expires`)中任意选择数据淘汰。当内存达到限制无法写入非过期时间的数据集时，可以通过该淘汰策略在主键空间中随机移除某个key。

4. **allkeys-lru**：从数据集(`server.db[i].dict`)中挑选最近最少使用的数据淘汰，该策略要淘汰的key面向的是全体key集合，而非过期的key集合。

5. **allkeys-random**：从数据集(`server.db[i].dict`)中选择任意数据淘汰。

6. **no-enviction**：禁止驱逐数据，也就是当内存不足以容纳新入数据时，新写入操作就会报错，请求可以继续进行，线上任务也不能持续进行，采用no-enviction策略可以保证数据不被丢失，这也是系统默认的一种淘汰策略。

[Ref](https://stor.51cto.com/art/201904/594773.htm)

7. Redis 4.0 加入了 `LFU`（least frequency use）淘汰策略，包括 **volatile-lfu** 和 **allkeys-lfu**，通过统计访问频率，将访问频率最少，即最不经常使用的 KV 淘汰。

## 7 Redis部署
- redis单例
- 主从模式
- sentinel
- 集群

详见：https://cloud.tencent.com/developer/article/1440480