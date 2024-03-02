<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [1. 引擎及区别](#1-%E5%BC%95%E6%93%8E%E5%8F%8A%E5%8C%BA%E5%88%AB)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# 1. 引擎及区别
主要 MyISAM 与 InnoDB 两个引擎，其主要区别如下：
一、InnoDB 支持事务，MyISAM 不支持，这一点是非常之重要。事务是一种高级的处理方式，如在一些列增删改中只要哪个出错还可以回滚还原，而 MyISAM就不可以了；
二、MyISAM 适合查询以及插入为主的应用，InnoDB 适合频繁修改以及涉及到安全性较高的应用；
三、InnoDB 支持外键，MyISAM 不支持；
四、MyISAM 是默认引擎，InnoDB 需要指定；
五、InnoDB 不支持 FULLTEXT 类型的索引；
六、InnoDB 中不保存表的行数，如 `select count(*) from table` 时，InnoDB；需要扫描一遍整个表来计算有多少行，但是 MyISAM 只要简单的读出保存好的行数即可。注意的是，当 `count(*)`语句包含 where 条件时 MyISAM 也需要扫描整个表；
七、对于自增长的字段，InnoDB 中必须包含只有该字段的索引，但是在 MyISAM表中可以和其他字段一起建立联合索引；
八、清空整个表时，InnoDB 是一行一行的删除，效率非常慢。MyISAM 则会重建表；
九、InnoDB 支持行锁（某些情况下还是锁整表，如 `update table set a=1 where user like '%lee%'`

## 1.1 简述InnoDB存储引擎

InnoDB 是 MySQL 的默认事务型引擎，支持事务，表是基于聚簇索引建立的。支持表级锁和行级锁，支持外键，适合数据增删改查都频繁的情况。

InnoDB 采用 MVCC 来支持高并发，并且实现了四个标准的隔离级别。其默认级别是 REPEATABLE READ，并通过间隙锁策略防止幻读，间隙锁使 InnoDB 不仅仅锁定查询涉及的行，还会对索引中的间隙进行锁定防止幻行的插入。

## 1.2 简述MyISAM存储引擎

MySQL5.1及之前，MyISAM 是默认存储引擎。MyISAM不支持事务，Myisam支持表级锁，不支持行级锁，表不支持外键，该存储引擎存有表的行数，count运算会更快。适合查询频繁，不适合对于增删改要求高的情况