# Go实现双向链表


本文介绍什么时候链表，链表和双向链表的区别
链表在redis中是怎么使用的，再通过一个在线秒杀的小示例演示


## 目录
- [1、链表](#1链表)
    - [1.1 说明](#11-说明)
    - [1.2 单向链表](#12-单向链表)
    - [1.3 双向链表](#13-双向链表)
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


```
# commodity:queue

# 存入商品
RPUSH commodity:queue 001
RPUSH commodity:queue 002
RPUSH commodity:queue 003
RPUSH commodity:queue 004

# 查看全部元素
LRANGE commodity:queue 0 -1

# llen
llen commodity:queue

# 出队
LPOP commodity:queue

# 根据索引查找元素
LINDEX commodity:queue 0

# 删除队列
del  commodity:queue



```