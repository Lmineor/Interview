
## 现象

先看一段代码示例：

```go
package main

import "fmt"

func main() {
    m := map[string]int{
        "apple":  1,
        "banana": 2,
        "orange": 3,
    }

    for k, v := range m {
        fmt.Printf("key=%s, value=%d\n", k, v)
    }
}
```

当我们多执行几次这段代码时，就会发现，输出的顺序是不同的。

## 原因

首先，Go 语言 map 的底层实现是哈希表，在进行插入时，会对 key 进行 hash 运算。这也就导致了数据不是按顺序存储的，和遍历的顺序也就会不一致。

第二，map 在扩容后，会发生 key 的搬迁，原来落在同一个 bucket 中的 key，搬迁后，有些 key 可能就到其他 bucket 了。

而遍历的过程，就是按顺序遍历 bucket，同时按顺序遍历 bucket 中的 key。

搬迁后，key 的位置发生了重大的变化，有些 key 被搬走了，有些 key 则原地不动。这样，遍历 map 的结果就不可能按原来的顺序了。

最后，也是最有意思的一点。

那如果说我已经初始化好了一个 map，并且不对这个 map 做任何操作，也就是不会发生扩容，那遍历顺序是固定的吗？

答：也不是。

Go 杜绝了这种做法，主要是担心程序员会在开发过程中依赖稳定的遍历顺序，因为这是不对的。

所以在遍历 map 时，并不是固定地从 0 号 bucket 开始遍历，每次都是从一个随机值序号的 bucket 开始遍历，并且是从这个 bucket 的一个随机序号的 cell 开始遍历。

## 如何顺序读取


如果希望按照特定顺序遍历 map，可以先将键或值存储到切片中，然后对切片进行排序，最后再遍历切片。

改造一下上面的代码，让它按顺序输出：

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    m := map[string]int{
        "apple":  1,
        "banana": 2,
        "orange": 3,
    }

    // 将 map 中的键存储到切片中
    keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }

    // 对切片进行排序
    sort.Strings(keys)

    // 按照排序后的顺序遍历 map
    for _, k := range keys {
        fmt.Printf("key=%s, value=%d\n", k, m[k])
    }
}
```

在上面的代码中，首先将 map 中的键存储到一个切片中，然后对切片进行排序。

最后，按照排序后的顺序遍历 map。这样就可以按照特定顺序输出键值对了。
