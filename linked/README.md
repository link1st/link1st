# Go实现双向链表


本文介绍什么是链表，常见的链表有哪些，然后介绍链表这种数据结构会在哪些地方可以用到，以及 Redis 队列是底层的实现，通过一个小实例来演示 Redis 队列有哪些功能，最后通过 Go 实现一个双向链表。

![链表](https://img.mukewang.com/5d820e2100014a2d20360992.png)

## 目录
- [1、链表](#1链表)
    - [1.1 说明](#11-说明)
    - [1.2 单向链表](#12-单向链表)
    - [1.3 循环链表](#13-循环链表)
    - [1.4 双向链表](#14-双向链表)
- [2、redis队列](#2redis队列)
    - [2.1 说明](#21-说明)
    - [2.2 应用场景](#22-应用场景)
    - [2.3 演示](#23-演示)
- [3、Go双向链表](#3Go双向链表)
    - [3.1 说明](#31-说明)
    - [3.2 实现](#32-实现)
- [4、总结](#总结)
- [5、参考文献](#5参考文献)


## 1、链表
### 1.1 说明

![链表](https://img.mukewang.com/5d8191630001181a16320862.png)

链表（Linked list）是一种常见的基础数据结构，是一种线性表，但是并不会按线性的顺序存储数据，而是在每一个节点里存到下一个节点的指针(Pointer)。由于不必须按顺序存储，链表在插入的时候可以达到O(1)的复杂度，比另一种线性表顺序表快得多，但是查找一个节点或者访问特定编号的节点则需要O(n)的时间，而顺序表相应的时间复杂度分别是O(logn)和O(1)。

链表有很多种不同的类型：单向链表，双向链表以及循环链表。

- 优势:

可以克服数组链表需要预先知道数据大小的缺点，链表结构可以充分利用计算机内存空间，实现灵活的内存动态管理。链表允许插入和移除表上任意位置上的节点。


- 劣势:

由于链表增加了节点指针，空间开销比较大。链表一般查找数据的时候需要从第一个节点开始每次访问下一个节点，直到访问到需要的位置，查找数据比较慢。

- 用途:

常用于组织检索较少，而删除、添加、遍历较多的数据。

如：文件系统、LRU cache、Redis 列表、内存管理等。


### 1.2 单向链表

链表中最简单的一种是单向链表，

一个单向链表的节点被分成两个部分。它包含两个域，一个信息域和一个指针域。第一个部分保存或者显示关于节点的信息，第二个部分存储下一个节点的地址，而最后一个节点则指向一个空值。单向链表只可向一个方向遍历。

单链表有一个头节点head，指向链表在内存的首地址。链表中的每一个节点的数据类型为结构体类型，节点有两个成员：整型成员（实际需要保存的数据）和指向下一个结构体类型节点的指针即下一个节点的地址（事实上，此单链表是用于存放整型数据的动态数组）。链表按此结构对各节点的访问需从链表的头找起，后续节点的地址由当前节点给出。无论在表中访问那一个节点，都需要从链表的头开始，顺序向后查找。链表的尾节点由于无后续节点，其指针域为空，写作为NULL。


### 1.3 循环链表

循环链表是与单向链表一样，是一种链式的存储结构，所不同的是，循环链表的最后一个结点的指针是指向该循环链表的第一个结点或者表头结点，从而构成一个环形的链。

循环链表的运算与单链表的运算基本一致。所不同的有以下几点：

1、在建立一个循环链表时，必须使其最后一个结点的指针指向表头结点，而不是象单链表那样置为NULL。此种情况还使用于在最后一个结点后插入一个新的结点。

2、在判断是否到表尾时，是判断该结点链域的值是否是表头结点，当链域值等于表头指针时，说明已到表尾。而非象单链表那样判断链域值是否为NULL。

### 1.4 双向链表

![双向链表](https://img.mukewang.com/5d81a0a80001247f18540700.png)

双向链表其实是单链表的改进，当我们对单链表进行操作时，有时你要对某个结点的直接前驱进行操作时，又必须从表头开始查找。这是由单链表结点的结构所限制的。因为单链表每个结点只有一个存储直接后继结点地址的链域，那么能不能定义一个既有存储直接后继结点地址的链域，又有存储直接前驱结点地址的链域的这样一个双链域结点结构呢？这就是双向链表。

在双向链表中，结点除含有数据域外，还有两个链域，一个存储直接后继结点地址，一般称之为右链域（当此“连接”为最后一个“连接”时，指向空值或者空列表）；一个存储直接前驱结点地址，一般称之为左链域（当此“连接”为第一个“连接”时，指向空值或者空列表）。

## 2、redis队列
### 2.1 说明

Redis列表是简单的字符串列表，按照插入顺序排序。你可以添加一个元素到列表的头部（左边）或者尾部（右边）


Redis 列表使用两种数据结构作为底层实现：双端列表(linkedlist)、压缩列表(ziplist)

通过配置文件中(list-max-ziplist-entries、list-max-ziplist-value)来选择是哪种实现方式

在数据量比较少的时候，使用双端链表和压缩列表性能差异不大，但是使用压缩列表更能节约内存空间

redis 链表的实现源码 [redis src/adlist.h](https://github.com/antirez/redis/blob/unstable/src/adlist.h)

### 2.2 应用场景

消息队列，秒杀项目

秒杀项目:

提前将需要的商品码信息存入 Redis 队列，在抢购的时候每个用户都从 Redis 队列中取商品码，由于 Redis 是单线程的，同时只能有一个商品码被取出，取到商品码的用户为购买成功，而且 Redis 性能比较高，能抗住较大的用户压力。

### 2.3 演示

如何通过 Redis 队列中防止并发情况下商品超卖的情况。

假设：

网站有三件商品需要卖，我们将数据存入 Redis 队列中

1、 将三个商品码(10001、10002、10003)存入 Redis 队列中

```
# 存入商品
RPUSH commodity:queue 10001 10002 10003
```

2、 存入以后，查询数据是否符合预期

```
# 查看全部元素
LRANGE commodity:queue 0 -1

# 查看队列的长度
LLEN commodity:queue
```

3、 抢购开始，获取商品码，抢到商品码的用户则可以购买(由于 Redis 是单线程的，同一个商品码只能被取一次
)

```
# 出队
LPOP commodity:queue
```

## 3、Go双向链表
### 3.1 说明


用Go语言实现一个双向链表，并实现在链表右边插入数据，在链表左边取数据实现 Redis 列表的 (RPUSH、LRANGE、LPOP、LLEN)等功能。

### 3.2 实现

![golang 双向链表](https://img.mukewang.com/5d820e2100014a2d20360992.png)

- 节点定义

```
// 链表的一个节点
type ListNode struct {
    prev  *ListNode // 前一个节点
    next  *ListNode // 后一个节点
    value string    // 数据
}

// 创建一个节点
func NewListNode(value string) (listNode *ListNode) {
    listNode = &ListNode{
        value: value,
    }

    return
}

// 当前节点的前一个节点
func (n *ListNode) Prev() (prev *ListNode) {
    prev = n.prev

    return
}

// 当前节点的前一个节点
func (n *ListNode) Next() (next *ListNode) {
    next = n.next

    return
}

// 获取节点的值
func (n *ListNode) GetValue() (value string) {
    if n == nil {

        return
    }
    value = n.value

    return
}
```

- 定义一个链表

```
// 链表
type List struct {
    head *ListNode // 表头节点
    tail *ListNode // 表尾节点
    len  int       // 链表的长度
}


// 创建一个空链表
func NewList() (list *List) {
    list = &List{
    }
    return
}

// 返回链表头节点
func (l *List) Head() (head *ListNode) {
    head = l.head

    return
}

// 返回链表尾节点
func (l *List) Tail() (tail *ListNode) {
    tail = l.tail

    return
}

// 返回链表长度
func (l *List) Len() (len int) {
    len = l.len

    return
}
```

- 在链表的右边插入一个元素

```
// 在链表的右边插入一个元素
func (l *List) RPush(value string) {

    node := NewListNode(value)

    // 链表未空的时候
    if l.Len() == 0 {
        l.head = node
        l.tail = node
    } else {
        tail := l.tail
        tail.next = node
        node.prev = tail

        l.tail = node
    }

    l.len = l.len + 1

    return
}
```

- 从链表左边取出一个节点

```
// 从链表左边取出一个节点
func (l *List) LPop() (node *ListNode) {

    // 数据为空
    if l.len == 0 {

        return
    }

    node = l.head

    if node.next == nil {
        // 链表未空
        l.head = nil
        l.tail = nil
    } else {
        l.head = node.next
    }
    l.len = l.len - 1

    return
}
```

- 通过索引查找节点

通过索引查找节点，如果索引是负数则从表尾开始查找

自然数和负数索引分别通过两种方式查找节点，找到指定索引或者是链表全部查找完则结束

```
// 通过索引查找节点
// 查不到节点则返回空
func (l *List) Index(index int) (node *ListNode) {

    // 索引为负数则表尾开始查找
    if index < 0 {
        index = (-index) - 1
        node = l.tail
        for true {
            // 未找到
            if node == nil {

                return
            }

            // 查到数据
            if index == 0 {

                return
            }

            node = node.prev
            index--
        }
    } else {
        node = l.head
        for ; index > 0 && node != nil; index-- {
            node = node.next
        }
    }

    return
}
```

- 返回指定区间的元素

```
// 返回指定区间的元素
func (l *List) Range(start, stop int) (nodes []*ListNode) {
    nodes = make([]*ListNode, 0)

    // 转为自然数
    if start < 0 {
        start = l.len + start
        if start < 0 {
            start = 0
        }
    }

    if stop < 0 {
        stop = l.len + stop
        if stop < 0 {
            stop = 0
        }
    }

    // 区间个数
    rangeLen := stop - start + 1
    if rangeLen < 0 {

        return
    }

    startNode := l.Index(start)
    for i := 0; i < rangeLen; i++ {
        if startNode == nil {
            break
        }

        nodes = append(nodes, startNode)
        startNode = startNode.next
    }

    return
}
```

## 4、总结

到这里关于链表的使用已经结束，介绍链表是有哪些(单向链表，双向链表以及循环链表)，也介绍了链表的应用场景(Redis 列表使用的是链表作为底层实现)，最后用 Go 实现了双向链表，演示了链表在 Go 语言中是怎么使用的，大家可以在项目中更具实际的情况去使用。


## 5、参考文献

[维基百科 链表](https://zh.wikipedia.org/wiki/%E9%93%BE%E8%A1%A8)

[github redis](https://github.com/antirez/redis)

项目地址:[go 实现队列](https://github.com/link1st/link1st/tree/master/linked)

https://github.com/link1st/link1st/tree/master/linked

