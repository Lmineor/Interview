<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Python语言特性](#python%E8%AF%AD%E8%A8%80%E7%89%B9%E6%80%A7)
  - [1 Python的函数参数传递](#1-python%E7%9A%84%E5%87%BD%E6%95%B0%E5%8F%82%E6%95%B0%E4%BC%A0%E9%80%92)
  - [2 Python中的元类(metaclass)](#2-python%E4%B8%AD%E7%9A%84%E5%85%83%E7%B1%BBmetaclass)
  - [3 @staticmethod和@classmethod](#3-staticmethod%E5%92%8Cclassmethod)
  - [4 类变量和实例变量](#4-%E7%B1%BB%E5%8F%98%E9%87%8F%E5%92%8C%E5%AE%9E%E4%BE%8B%E5%8F%98%E9%87%8F)
  - [5 Python自省](#5-python%E8%87%AA%E7%9C%81)
  - [6 字典推导式](#6-%E5%AD%97%E5%85%B8%E6%8E%A8%E5%AF%BC%E5%BC%8F)
  - [7 Python中单下划线和双下划线](#7-python%E4%B8%AD%E5%8D%95%E4%B8%8B%E5%88%92%E7%BA%BF%E5%92%8C%E5%8F%8C%E4%B8%8B%E5%88%92%E7%BA%BF)
  - [8 字符串格式化:%和.format](#8-%E5%AD%97%E7%AC%A6%E4%B8%B2%E6%A0%BC%E5%BC%8F%E5%8C%96%25%E5%92%8Cformat)
  - [9 迭代器和生成器](#9-%E8%BF%AD%E4%BB%A3%E5%99%A8%E5%92%8C%E7%94%9F%E6%88%90%E5%99%A8)
  - [10 `*args` and `**kwargs`](#10-args-and-kwargs)
  - [11 面向切面编程AOP和装饰器](#11-%E9%9D%A2%E5%90%91%E5%88%87%E9%9D%A2%E7%BC%96%E7%A8%8Baop%E5%92%8C%E8%A3%85%E9%A5%B0%E5%99%A8)
  - [12 鸭子类型](#12-%E9%B8%AD%E5%AD%90%E7%B1%BB%E5%9E%8B)
  - [13 Python中重载](#13-python%E4%B8%AD%E9%87%8D%E8%BD%BD)
  - [14 新式类和旧式类](#14-%E6%96%B0%E5%BC%8F%E7%B1%BB%E5%92%8C%E6%97%A7%E5%BC%8F%E7%B1%BB)
  - [15 `__new__`和`__init__`的区别](#15-__new__%E5%92%8C__init__%E7%9A%84%E5%8C%BA%E5%88%AB)
  - [16 单例模式](#16-%E5%8D%95%E4%BE%8B%E6%A8%A1%E5%BC%8F)
    - [1 使用`__new__`方法](#1-%E4%BD%BF%E7%94%A8__new__%E6%96%B9%E6%B3%95)
    - [2 共享属性](#2-%E5%85%B1%E4%BA%AB%E5%B1%9E%E6%80%A7)
    - [3 装饰器版本](#3-%E8%A3%85%E9%A5%B0%E5%99%A8%E7%89%88%E6%9C%AC)
    - [4 import方法](#4-import%E6%96%B9%E6%B3%95)
    - [5 线程安全的单例模式](#5-%E7%BA%BF%E7%A8%8B%E5%AE%89%E5%85%A8%E7%9A%84%E5%8D%95%E4%BE%8B%E6%A8%A1%E5%BC%8F)
  - [17 Python中的作用域](#17-python%E4%B8%AD%E7%9A%84%E4%BD%9C%E7%94%A8%E5%9F%9F)
  - [18 GIL线程全局锁](#18-gil%E7%BA%BF%E7%A8%8B%E5%85%A8%E5%B1%80%E9%94%81)
  - [19 协程](#19-%E5%8D%8F%E7%A8%8B)
  - [20 闭包](#20-%E9%97%AD%E5%8C%85)
  - [21 lambda函数](#21-lambda%E5%87%BD%E6%95%B0)
  - [22 Python函数式编程](#22-python%E5%87%BD%E6%95%B0%E5%BC%8F%E7%BC%96%E7%A8%8B)
  - [23 Python里的拷贝](#23-python%E9%87%8C%E7%9A%84%E6%8B%B7%E8%B4%9D)
  - [24 Python垃圾回收机制](#24-python%E5%9E%83%E5%9C%BE%E5%9B%9E%E6%94%B6%E6%9C%BA%E5%88%B6)
    - [1 引用计数](#1-%E5%BC%95%E7%94%A8%E8%AE%A1%E6%95%B0)
    - [2 标记-清除机制](#2-%E6%A0%87%E8%AE%B0-%E6%B8%85%E9%99%A4%E6%9C%BA%E5%88%B6)
    - [3 分代技术](#3-%E5%88%86%E4%BB%A3%E6%8A%80%E6%9C%AF)
  - [25 Python的List](#25-python%E7%9A%84list)
  - [26 Python的is](#26-python%E7%9A%84is)
  - [27 read,readline和readlines](#27-readreadline%E5%92%8Creadlines)
  - [28 Python2和3的区别](#28-python2%E5%92%8C3%E7%9A%84%E5%8C%BA%E5%88%AB)
  - [29 super init](#29-super-init)
  - [30 range and xrange](#30-range-and-xrange)
  - [31 Python Web](#31-python-web)
    - [Nginx、WSGI、 uWSGI、 uwsgi的区别](#nginxwsgi-uwsgi-uwsgi%E7%9A%84%E5%8C%BA%E5%88%AB)
      - [WSGI](#wsgi)
      - [uwsgi 和 uWSGI](#uwsgi-%E5%92%8C-uwsgi)
      - [Nginx](#nginx)
      - [uWSGI 和 Nginx 的关系](#uwsgi-%E5%92%8C-nginx-%E7%9A%84%E5%85%B3%E7%B3%BB)
  - [31 Python 内存管理机制](#31-python-%E5%86%85%E5%AD%98%E7%AE%A1%E7%90%86%E6%9C%BA%E5%88%B6)
    - [引用计数](#%E5%BC%95%E7%94%A8%E8%AE%A1%E6%95%B0)
      - [引用计数增加](#%E5%BC%95%E7%94%A8%E8%AE%A1%E6%95%B0%E5%A2%9E%E5%8A%A0)
      - [引用计数减少](#%E5%BC%95%E7%94%A8%E8%AE%A1%E6%95%B0%E5%87%8F%E5%B0%91)
    - [垃圾回收](#%E5%9E%83%E5%9C%BE%E5%9B%9E%E6%94%B6)
    - [内存池机制](#%E5%86%85%E5%AD%98%E6%B1%A0%E6%9C%BA%E5%88%B6)
    - [Python的内存池(金字塔)](#python%E7%9A%84%E5%86%85%E5%AD%98%E6%B1%A0%E9%87%91%E5%AD%97%E5%A1%94)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Python语言特性

## 1 Python的函数参数传递

看两个例子:

```python
a = 1
def fun(a):
    a = 2
fun(a)
print(a)  # 1
```

```python
a = []
def fun(a):
    a.append(1)
fun(a)
print(a)  # [1]
```

所有的变量都可以理解是内存中一个对象的“引用”，或者，也可以看似c中void*的感觉。

通过`id`来看引用`a`的内存地址可以比较理解：

```python
a = 1
def fun(a):
    print("func_in",id(a))   # func_in 41322472
    a = 2
    print("re-point",id(a), id(2))   # re-point 41322448 41322448
print("func_out",id(a), id(1))  # func_out 41322472 41322472
fun(a)
print(a)  # 1
```

注：具体的值在不同电脑上运行时可能不同。

可以看到，在执行完`a = 2`之后，`a`引用中保存的值，即内存地址发生变化，由原来`1`对象的所在的地址变成了`2`这个实体对象的内存地址。

而第2个例子`a`引用保存的内存值就不会发生变化：

```python
a = []
def fun(a):
    print("func_in",id(a))  # func_in 53629256
    a.append(1)
print("func_out",id(a))     # func_out 53629256
fun(a)
print(a)  # [1]
```

这里记住的是类型是属于对象的，而不是变量。而对象有两种,“可更改”（mutable）与“不可更改”（immutable）对象。在python中，strings, tuples, 和numbers是不可更改的对象，而 list, dict, set 等则是可以修改的对象。(这就是这个问题的重点)

当一个引用传递给函数的时候,函数自动复制一份引用,这个函数里的引用和外边的引用没有半毛关系了.所以第一个例子里函数把引用指向了一个不可变对象,当函数返回的时候,外面的引用没半毛感觉.而第二个例子就不一样了,函数内的引用指向的是可变对象,对它的操作就和定位了指针地址一样,在内存里进行修改.

如果还不明白的话,这里有更好的解释: http://stackoverflow.com/questions/986006/how-do-i-pass-a-variable-by-reference

## 2 Python中的元类(metaclass)

这个非常的不常用,但是像ORM这种复杂的结构还是会需要的,详情请看:http://stackoverflow.com/questions/100003/what-is-a-metaclass-in-python

## 3 @staticmethod和@classmethod

Python其实有3个方法,即静态方法(staticmethod),类方法(classmethod)和实例方法,如下:

```python
def foo(x):
    print("executing foo(%s)"%(x))

class A(object):
    def foo(self,x):
        print("executing foo(%s,%s)"%(self,x))

    @classmethod
    def class_foo(cls,x):
        print("executing class_foo(%s,%s)"%(cls,x))

    @staticmethod
    def static_foo(x):
        print("executing static_foo(%s)"%x)

a=A()

```

这里先理解下函数参数里面的self和cls.这个self和cls是对类或者实例的绑定,对于一般的函数来说我们可以这么调用`foo(x)`,这个函数就是最常用的,它的工作跟任何东西(类,实例)无关.对于实例方法,我们知道在类里每次定义方法的时候都需要绑定这个实例,就是`foo(self, x)`,为什么要这么做呢?因为实例方法的调用离不开实例,我们需要把实例自己传给函数,调用的时候是这样的`a.foo(x)`(其实是`foo(a, x)`).类方法一样,只不过它传递的是类而不是实例,`A.class_foo(x)`.注意这里的self和cls可以替换别的参数,但是python的约定是这俩,还是不要改的好.

对于静态方法其实和普通的方法一样,不需要对谁进行绑定,唯一的区别是调用的时候需要使用`a.static_foo(x)`或者`A.static_foo(x)`来调用.

| \\      | 实例方法     | 类方法            | 静态方法            |
| :------ | :------- | :------------- | :-------------- |
| a = A() | a.foo(x) | a.class_foo(x) | a.static_foo(x) |
| A       | 不可用      | A.class_foo(x) | A.static_foo(x) |

更多关于这个问题:
1. http://stackoverflow.com/questions/136097/what-is-the-difference-between-staticmethod-and-classmethod-in-python
2. https://realpython.com/blog/python/instance-class-and-static-methods-demystified/
## 4 类变量和实例变量

**类变量：**

> ​	是可在类的所有实例之间共享的值（也就是说，它们不是单独分配给每个实例的）。例如下例中，num_of_instance 就是类变量，用于跟踪存在着多少个Test 的实例。

**实例变量：**

> 实例化之后，每个实例单独拥有的变量。

```python
class Test(object):  
    num_of_instance = 0  
    def __init__(self, name):  
        self.name = name  
        Test.num_of_instance += 1  
  
if __name__ == '__main__':  
    print(Test.num_of_instance)   # 0
    t1 = Test('jack')  
    print(Test.num_of_instance)  # 1
    t2 = Test('lucy')  
    print(t1.name , t1.num_of_instance)  # jack 2
    print(t2.name , t2.num_of_instance)  # lucy 2
```

> 补充的例子

```python
class Person:
    name="aaa"

p1=Person()
p2=Person()
p1.name="bbb"
print(p1.name)  # bbb
print(p2.name)  # aaa
print(Person.name)  # aaa
```

这里`p1.name="bbb"`是实例调用了类变量,这其实和上面第一个问题一样,就是函数传参的问题,`p1.name`一开始是指向的类变量`name="aaa"`,但是在实例的作用域里把类变量的引用改变了,就变成了一个实例变量,self.name不再引用Person的类变量name了.

可以看看下面的例子:

```python
class Person:
    name=[]

p1=Person()
p2=Person()
p1.name.append(1)
print(p1.name)  # [1]
print(p2.name)  # [1]
print(Person.name)  # [1]
```

参考:http://stackoverflow.com/questions/6470428/catch-multiple-exceptions-in-one-line-except-block

## 5 Python自省与反射

### 自省

这个也是python彪悍的特性.

自省就是面向对象的语言所写的程序在运行时,所能知道对象的类型.简单一句就是运行时能够获得对象的类型.比如type(),dir(),getattr(),hasattr(),isinstance().

```python
a = [1,2,3]
b = {'a':1,'b':2,'c':3}
c = True
print(type(a),type(b),type(c)) # <type 'list'> <type 'dict'> <type 'bool'>
print(isinstance(a,list))  # True
```

### 反射

反射机制就是在运行时，动态的确定对象的类型，并可以通过字符串调用对象属性、方法、导入模块，是一种基于字符串的事件驱动

## 6 字典推导式

可能你见过列表推导时,却没有见过字典推导式,在2.7中才加入的:

```python
d = {key: value for (key, value) in iterable}
```

## 7 Python中单下划线和双下划线

```python
>>> class MyClass():
...     def __init__(self):
...             self.__superprivate = "Hello"
...             self._semiprivate = ", world!"
...
>>> mc = MyClass()
>>> print(mc.__superprivate)
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
AttributeError: myClass instance has no attribute '__superprivate'
>>> print(mc._semiprivate)
, world!
>>> print(mc.__dict__)
{'_MyClass__superprivate': 'Hello', '_semiprivate': ', world!'}
```

`__foo__`:一种约定,Python内部的名字,用来区别其他用户自定义的命名,以防冲突，就是例如`__init__()`,`__del__()`,`__call__()`这些特殊方法

`_foo`:一种约定,用来指定变量私有.程序员用来指定私有变量的一种方式.不能用from module import * 导入，其他方面和公有一样访问；

`__foo`:这个有真正的意义:解析器用`_classname__foo`来代替这个名字,以区别和其他类相同的命名,它无法直接像公有成员一样随便访问,通过对象名._类名__xxx这样的方式可以访问.

详情见:http://stackoverflow.com/questions/1301346/the-meaning-of-a-single-and-a-double-underscore-before-an-object-name-in-python

或者: http://www.zhihu.com/question/19754941

## 8 字符串格式化:%和.format

.format在许多方面看起来更便利.对于`%`最烦人的是它无法同时传递一个变量和元组.你可能会想下面的代码不会有什么问题:

```
"hi there %s" % name
```

但是,如果`name`恰好是`(1,2,3)`,它将会抛出一个TypeError异常.为了保证它总是正确的,你必须这样做:

```
"hi there %s" % (name,)   # 提供一个单元素的数组而不是一个参数
```

但是有点丑`.format`就没有这些问题.你给的第二个问题也是这样,`.format`好看多了.

你为什么不用它?

* 不知道它(在读这个之前)
* 为了和Python2.5兼容(譬如logging库建议使用`%`([issue #4](https://github.com/taizilongxu/interview_python/issues/4)))

http://stackoverflow.com/questions/5082452/python-string-formatting-vs-format

## 9 迭代器和生成器

### 生成器
这个是stackoverflow里python排名第一的问题,值得一看: http://stackoverflow.com/questions/231767/what-does-the-yield-keyword-do-in-python

这是中文版: http://taizilongxu.gitbooks.io/stackoverflow-about-python/content/1/README.html

这里有个关于生成器的创建问题面试官有考：
问：  将列表生成式中[]改成() 之后数据结构是否改变？ 
答案：是，从列表变为生成器

```python
>>> L = [x*x for x in range(10)]
>>> L
[0, 1, 4, 9, 16, 25, 36, 49, 64, 81]
>>> g = (x*x for x in range(10))
>>> g
<generator object <genexpr> at 0x03A33840>
>>> next(g)
0
>>> next(g)
1
>>> next(g)
4
```
通过列表生成式，可以直接创建一个列表。但是，受到内存限制，列表容量肯定是有限的。而且，创建一个包含百万元素的列表，不仅是占用很大的内存空间，如：我们只需要访问前面的几个元素，后面大部分元素所占的空间都是浪费的。因此，没有必要创建完整的列表（节省大量内存空间）。在Python中，我们可以采用生成器：边循环，边计算`next`的机制—>generator

```py
# 斐波那契数列
def fib(max):
    n, a, b = 0, 0, 1
    while n < max:
        yield a
        a, b = b, a+b
        n = n + 1
    return 'Done'

fi = fib(10)
try:
    for i in range(15):
        print(next(fi))
except StopIteration:
    pass
```

### 迭代器

## 10 `*args` and `**kwargs`

用`*args`和`**kwargs`只是为了方便并没有强制使用它们.

当你不确定你的函数里将要传递多少参数时你可以用`*args`.例如,它可以传递任意数量的参数:

```python
>>> def print_everything(*args):
        for count, thing in enumerate(args):
...         print('{0}. {1}'.format(count, thing))
...
>>> print_everything('apple', 'banana', 'cabbage')
0. apple
1. banana
2. cabbage
```

相似的,`**kwargs`允许你使用没有事先定义的参数名:

```python
>>> def table_things(**kwargs):
...     for name, value in kwargs.items():
...         print('{0} = {1}'.format(name, value))
...
>>> table_things(apple = 'fruit', cabbage = 'vegetable')
cabbage = vegetable
apple = fruit
```

你也可以混着用.命名参数首先获得参数值然后所有的其他参数都传递给`*args`和`**kwargs`.命名参数在列表的最前端.例如:

```
def table_things(titlestring, **kwargs)
```

`*args`和`**kwargs`可以同时在函数的定义中,但是`*args`必须在`**kwargs`前面.

当调用函数时你也可以用`*`和`**`语法.例如:

```python
>>> def print_three_things(a, b, c):
...     print('a = {0}, b = {1}, c = {2}'.format(a,b,c))
...
>>> mylist = ['aardvark', 'baboon', 'cat']
>>> print_three_things(*mylist)

a = aardvark, b = baboon, c = cat
```

就像你看到的一样,它可以传递列表(或者元组)的每一项并把它们解包.注意必须与它们在函数里的参数相吻合.当然,你也可以在函数定义或者函数调用时用*.

http://stackoverflow.com/questions/3394835/args-and-kwargs

## 11 面向切面编程AOP和装饰器

装饰器是一个很著名的设计模式，经常被用于有切面需求的场景，较为经典的有插入日志、性能测试、事务处理等。装饰器是解决这类问题的绝佳设计，有了装饰器，我们就可以抽离出大量函数中与函数功能本身无关的雷同代码并继续重用。概括的讲，**装饰器的作用就是为已经存在的对象添加额外的功能。**
```py
from time import ctime, sleep
def timefun(func):
    def wrappedfunc():
        print("%s called at %s"%(func.__name__, ctime()))
        return func()
    return wrappedfunc
@timefun
def foo():
    print("I am foo")
foo()
sleep(2)
foo()
"""
foo called at Sat Jul 25 20:06:32 2020
I am foo
foo called at Sat Jul 25 20:06:34 2020
I am foo
"""
```
这个问题比较大,推荐: http://stackoverflow.com/questions/739654/how-can-i-make-a-chain-of-function-decorators-in-python

中文: http://taizilongxu.gitbooks.io/stackoverflow-about-python/content/3/README.html

## 12 鸭子类型

“当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子。”

我们并不关心对象是什么类型，到底是不是鸭子，只关心行为。

比如在python中，有很多file-like的东西，比如StringIO,GzipFile,socket。它们有很多相同的方法，我们把它们当作文件使用。

又比如list.extend()方法中,我们并不关心它的参数是不是list,只要它是可迭代的,所以它的参数可以是list/tuple/dict/字符串/生成器等.

鸭子类型在动态语言中经常使用，非常灵活，使得python不想java那样专门去弄一大堆的设计模式。

## 13 Python中重载

引自知乎:http://www.zhihu.com/question/20053359

函数重载主要是为了解决两个问题。

1. 可变参数类型。
2. 可变参数个数。

另外，一个基本的设计原则是，仅仅当两个函数除了参数类型和参数个数不同以外，其功能是完全相同的，此时才使用函数重载，如果两个函数的功能其实不同，那么不应当使用重载，而应当使用一个名字不同的函数。

好吧，那么对于情况 1 ，函数功能相同，但是参数类型不同，python 如何处理？答案是根本不需要处理，因为 python 可以接受任何类型的参数，如果函数的功能相同，那么不同的参数类型在 python 中很可能是相同的代码，没有必要做成两个不同函数。

那么对于情况 2 ，函数功能相同，但参数个数不同，python 如何处理？大家知道，答案就是缺省参数。对那些缺少的参数设定为缺省参数即可解决问题。因为你假设函数功能相同，那么那些缺少的参数终归是需要用的。

好了，鉴于情况 1 跟 情况 2 都有了解决方案，**python 自然就不需要函数重载了**。

## 14 新式类和旧式类

这个面试官问了,我说了老半天,不知道他问的真正意图是什么.

[stackoverflow](http://stackoverflow.com/questions/54867/what-is-the-difference-between-old-style-and-new-style-classes-in-python)

这篇文章很好的介绍了新式类的特性: http://www.cnblogs.com/btchenguang/archive/2012/09/17/2689146.html

新式类很早在2.2就出现了,所以旧式类完全是兼容的问题,Python3里的类全部都是新式类.这里有一个MRO问题可以了解下(新式类继承是根据C3算法,旧式类是深度优先),<Python核心编程>里讲的也很多.

> 一个旧式类的深度优先的例子

```python
class A():
    def foo1(self):
        print("A")
class B(A):
    def foo2(self):
        pass
class C(A):
    def foo1(self):
        print("C")
class D(B, C):
    pass

d = D()
d.foo1()

# A
```

**按照经典类的查找顺序`从左到右深度优先`的规则，在访问`d.foo1()`的时候,D这个类是没有的..那么往上查找,先找到B,里面没有,深度优先,访问A,找到了foo1(),所以这时候调用的是A的foo1()，从而导致C重写的foo1()被绕过**



## 15 `__new__`和`__init__`的区别

```py
class TestClass(object):
    def __new__(cls, *args, **kwargs):
        print("__new__")
        return object.__new__(cls)

    def __init__(self, name):
        print("__init__")
        self.name = name

tc = TestClass('TC')
print(tc.name)
"""
__new__
__init__
TC
"""
```
1. `__new__`是一个静态方法,而`__init__`是一个实例方法.
2. `__new__`方法会返回一个创建的实例,而`__init__`什么都不返回.
3. 只有在`__new__`返回一个cls的实例时后面的`__init__`才能被调用.
4. 当创建一个新实例时调用`__new__`,初始化一个实例时用`__init__`.

[stackoverflow](http://stackoverflow.com/questions/674304/pythons-use-of-new-and-init)

ps: `__metaclass__`是创建类时起作用.所以我们可以分别使用`__metaclass__`,`__new__`和`__init__`来分别在类创建,实例创建和实例初始化的时候做一些小手脚.

## 16 单例模式

> ​	单例模式是一种常用的软件设计模式。在它的核心结构中只包含一个被称为单例类的特殊类。通过单例模式可以保证系统中一个类只有一个实例而且该实例易于外界访问，从而方便对实例个数的控制并节约系统资源。如果希望在系统中某个类的对象只能存在一个，单例模式是最好的解决方案。
>
> `__new__()`在`__init__()`之前被调用，用于生成实例对象。利用这个方法和类的属性的特点可以实现设计模式的单例模式。单例模式是指创建唯一对象，单例模式设计的类只能实例
**这个绝对常考啊.绝对要记住1~2个方法,当时面试官是让手写的.**

### 1 使用`__new__`方法

```python
class Singleton(object):
    def __new__(cls, *args, **kw):
        if not hasattr(cls, '_instance'):
            orig = super(Singleton, cls)
            cls._instance = orig.__new__(cls, *args, **kw)
        return cls._instance

class MyClass(Singleton):
    a = 1
```

### 2 共享属性

创建实例时把所有实例的`__dict__`指向同一个字典,这样它们具有相同的属性和方法.

```python

class Borg(object):
    _state = {}
    def __new__(cls, *args, **kw):
        ob = super(Borg, cls).__new__(cls, *args, **kw)
        ob.__dict__ = cls._state
        return ob

class MyClass2(Borg):
    a = 1
```

### 3 装饰器版本

```python
def singleton(cls):
    instances = {}
    def getinstance(*args, **kw):
        if cls not in instances:
            instances[cls] = cls(*args, **kw)
        return instances[cls]
    return getinstance

@singleton
class MyClass:
  ...
```

### 4 import方法

作为python的模块是天然的单例模式

```python
# mysingleton.py
class My_Singleton(object):
    def foo(self):
        pass

my_singleton = My_Singleton()

# to use
from mysingleton import my_singleton

my_singleton.foo()

```

### 5 线程安全的单例模式
```py
import threading
class Singleton(object):
    _instance_lock = threading.Lock()
    def __init__(self):
        pass
    def __new__(cls, *args, **kwargs):
        if not hasattr(Singleton, "_instance"):
            with Singleton._instance_lock:
                if not hasattr(Singleton, "_instance"):
                    Singleton._instance = object.__new__(cls)
        return Singleton._instance
obj1 = Singleton()
obj2 = Singleton()
print(obj1,obj2)
def task(arg):
    obj = Singleton()
    print(obj)
for i in range(10):
    t = threading.Thread(target=task,args=[i,])
    t.start()

"""
打印结果如下：
<__main__.Singleton object at 0x000001DF5CB72D68> <__main__.Singleton object at 0x000001DF5CB72D68>
<__main__.Singleton object at 0x000001DF5CB72D68>
<__main__.Singleton object at 0x000001DF5CB72D68>
<__main__.Singleton object at 0x000001DF5CB72D68>
<__main__.Singleton object at 0x000001DF5CB72D68>
<__main__.Singleton object at 0x000001DF5CB72D68>
<__main__.Singleton object at 0x000001DF5CB72D68>
<__main__.Singleton object at 0x000001DF5CB72D68>
<__main__.Singleton object at 0x000001DF5CB72D68>
<__main__.Singleton object at 0x000001DF5CB72D68>
<__main__.Singleton object at 0x000001DF5CB72D68>

模式，以后实例化对象时，和平时实例化对象的方法一样 obj = Singleton() 
"""
```
**[单例模式伯乐在线详细解释](http://python.jobbole.com/87294/)**

## 17 Python中的作用域

Python 中，一个变量的作用域总是由在代码中被赋值的地方所决定的。

当 Python 遇到一个变量的话他会按照这样的顺序进行搜索：

本地作用域（Local）→当前作用域被嵌入的本地作用域（Enclosing locals）→全局/模块作用域（Global）→内置作用域（Built-in）

## 18 GIL线程全局锁

### 先说下python解释器

**解释器**是能够执行用其他计算机语言编写的程序的系统软件，它是一种翻译程序。它的执行方式是一边翻译一边执行，因此其执行效率一般偏低，但是解释器的实现较为简单，而且编写源程序的高级语言可以使用更加灵活和富于表现力的语法。

**解释器**由一个**编译器**和一个**虚拟机**构成，编译器负责将源代码转换成**字节码**文件，而虚拟机负责执行字节码。
所以，解释型语言其实也有编译过程，只不过这个编译过程并不是直接生成目标代码，而是中间代码（字节码），然后再通过虚拟机来逐行解释执行字节码

有很多种python解释器

#### CPython

当我们从Python官方网站下载并安装好Python 3.x后，我们就直接获得了一个官方版本的解释器：CPython。这个解释器是用C语言开发的，所以叫CPython。在命令行下运行python就是启动CPython解释器。

CPython是使用最广的Python解释器。

#### IPython

IPython是基于CPython之上的一个交互式解释器，也就是说，IPython只是在交互方式上有所增强，但是执行Python代码的功能和CPython是完全一样的。好比很多国产浏览器虽然外观不同，但内核其实都是调用了IE。

CPython用>>>作为提示符，而IPython用In [序号]:作为提示符。

#### PyPy

PyPy是另一个Python解释器，它的目标是执行速度。PyPy采用JIT技术，对Python代码进行动态编译（注意不是解释），所以可以显著提高Python代码的执行速度。

绝大部分Python代码都可以在PyPy下运行，但是PyPy和CPython有一些是不同的，这就导致相同的Python代码在两种解释器下执行可能会有不同的结果。如果你的代码要放到PyPy下执行，就需要了解PyPy和CPython的不同点。

#### Jython

Jython是运行在Java平台上的Python解释器，可以直接把Python代码编译成Java字节码执行。

#### IronPython

IronPython和Jython类似，只不过IronPython是运行在微软.Net平台上的Python解释器，可以直接把Python代码编译成.Net的字节码。

#### 小结

Python的解释器很多，但使用最广泛的还是CPython。如果要和Java或.Net平台交互，最好的办法不是用Jython或IronPython，而是通过网络调用来交互，确保各程序之间的独立性。

参考：https://www.liaoxuefeng.com/wiki/1016959663602400/1016966024263840

### 线程全局锁

官方解释：

> In CPython, the global interpreter lock, or GIL, is a mutex that prevents multiple native threads from executing Python bytecodes at once. This lock is necessary mainly because CPython’s memory management is not thread-safe. (However, since the GIL exists, other features have grown to depend on the guarantees that it enforces.)

线程全局锁(Global Interpreter Lock),即Python为了保证线程安全而采取的独立线程运行的限制,说白了就是一个核只能在同一时间运行一个线程.**对于io密集型任务，python的多线程起到作用，但对于cpu密集型任务，python的多线程几乎占不到任何优势，还有可能因为争夺资源而变慢。**

需要明确的一点是**GIL并不是Python的特性**，它是在实现Python解析器(**CPython**)时所引入的一个概念。

见[Python 最难的问题](http://www.oschina.net/translate/pythons-hardest-problem)

解决办法就是多进程和下面的协程(协程也只是单CPU,但是能减小切换代价提升性能).

## 19 协程

知乎被问到了,呵呵哒,跪了

简单点说协程是进程和线程的升级版,进程和线程都面临着内核态和用户态的切换问题而耗费许多切换时间,而协程就是用户自己控制切换的时机,不再需要陷入系统的内核态.

Python里最常见的yield就是协程的思想!可以查看第九个问题.


## 20 闭包

好难╯︿╰ https://www.cnblogs.com/yssjun/p/9887239.html
闭包(closure)是函数式编程的重要的语法结构。闭包也是一种组织代码的结构，它同样提高了代码的可重复使用性。

当一个内嵌函数引用其外部作作用域的变量,我们就会得到一个闭包. 总结一下,创建一个闭包必须满足以下几点:

1. 必须有一个内嵌函数
2. 内嵌函数必须引用外部函数中的变量
3. 外部函数的返回值必须是内嵌函数

感觉闭包还是有难度的,几句话是说不明白的,还是查查相关资料.

重点是函数运行后并不会被撤销,就像16题的instance字典一样,当函数运行完后,instance并不被销毁,而是继续留在内存空间里.这个功能类似类里的类变量,只不过迁移到了函数上.

闭包就像个空心球一样,你知道外面和里面,但你不知道中间是什么样.

## 21 lambda函数

其实就是一个匿名函数,为什么叫lambda?因为和后面的函数式编程有关.

推荐: [知乎](http://www.zhihu.com/question/20125256)


## 22 Python函数式编程

这个需要适当的了解一下吧,毕竟函数式编程在Python中也做了引用.

推荐: [酷壳](http://coolshell.cn/articles/10822.html)

python中函数式编程支持:

filter 函数的功能相当于过滤器。调用一个布尔函数`bool_func`来迭代遍历每个seq中的元素；返回一个使`bool_seq`返回值为true的元素的序列。

```python
>>>a = [1,2,3,4,5,6,7]
>>>b = filter(lambda x: x > 5, a)
>>>print(b)
>>>[6,7]
```

map函数是对一个序列的每个项依次执行函数，下面是对一个序列每个项都乘以2：

```python
>>> a = map(lambda x:x*2,[1,2,3])
>>> list(a)
[2, 4, 6]
```

reduce函数是对一个序列的每个项迭代调用函数，下面是求3的阶乘：

```python
>>> reduce(lambda x,y:x*y,range(1,4))
6
```

## 23 Python里的拷贝

引用和copy(),deepcopy()的区别

```python
import copy
a = [1, 2, 3, 4, ['a', 'b']]  #原始对象

b = a                       #赋值，传对象的引用
c = copy.copy(a)            #对象拷贝，浅拷贝
d = copy.deepcopy(a)        #对象拷贝，深拷贝

a.append(5)                 #修改对象a
a[4].append('c')            #修改对象a中的['a', 'b']数组对象

print('a = ', a)
print('b = ', b)
print('c = ', c)
print('d = ', d)

输出结果：
a =  [1, 2, 3, 4, ['a', 'b', 'c'], 5]
b =  [1, 2, 3, 4, ['a', 'b', 'c'], 5]
c =  [1, 2, 3, 4, ['a', 'b', 'c']]
d =  [1, 2, 3, 4, ['a', 'b']]
```

## 24 Python垃圾回收机制

Python GC主要使用引用计数（reference counting）来跟踪和回收垃圾。在引用计数的基础上，通过“标记-清除”（mark and sweep）解决容器对象可能产生的循环引用问题，通过“分代回收”（generation collection）以空间换时间的方法提高垃圾回收效率。

### 1 引用计数

PyObject是每个对象必有的内容，其中`ob_refcnt`就是做为引用计数。当一个对象有新的引用时，它的`ob_refcnt`就会增加，当引用它的对象被删除，它的`ob_refcnt`就会减少.引用计数为0时，该对象生命就结束了。

源码如下
```c
typedef struct _object {
    _PyObject_HEAD_EXTRA // 活跃堆对象的指针信息
    Py_ssize_t ob_refcnt;
    struct _typeobject *ob_type;
} PyObject;
```

优点:

1. 简单
2. 实时性

缺点:

1. 维护引用计数消耗资源
2. 循环引用

### 2 标记-清除机制

基本思路是先按需分配，等到没有空闲内存的时候从寄存器和程序栈上的引用出发，遍历以对象为节点、以引用为边构成的图，把所有可以访问到的对象打上标记，然后清扫一遍内存空间，把所有没标记的对象释放。

### 3 分代技术

分代回收的整体思想是：将系统中的所有内存块根据其存活时间划分为不同的集合，每个集合就成为一个“代”，垃圾收集频率随着“代”的存活时间的增大而减小，存活时间通常利用经过几次垃圾回收来度量。

Python默认定义了三代对象集合，索引数越大，对象存活时间越长。

举例：
当某些内存块M经过了3次垃圾收集的清洗之后还存活时，我们就将内存块M划到一个集合A中去，而新分配的内存都划分到集合B中去。当垃圾收集开始工作时，大多数情况都只对集合B进行垃圾回收，而对集合A进行垃圾回收要隔相当长一段时间后才进行，这就使得垃圾收集机制需要处理的内存少了，效率自然就提高了。在这个过程中，集合B中的某些内存块由于存活时间长而会被转移到集合A中，当然，集合A中实际上也存在一些垃圾，这些垃圾的回收会因为这种分代的机制而被延迟。

## 25 Python的List

推荐: http://www.jianshu.com/p/J4U6rR

## 26 Python的is

is是对比地址,==是对比值

## 27 read,readline和readlines

* read        读取整个文件
* readline    读取下一行,使用生成器方法
* readlines   读取整个文件到一个迭代器以供我们遍历

## 28 Python2和3的区别(很少问了)
推荐：[Python 2.7.x 与 Python 3.x 的主要差异](http://chenqx.github.io/2014/11/10/Key-differences-between-Python-2-7-x-and-Python-3-x/)

## 29 super init
super() lets you avoid referring to the base class explicitly, which can be nice. But the main advantage comes with multiple inheritance, where all sorts of fun stuff can happen. See the standard docs on super if you haven't already.

Note that the syntax changed in Python 3.0: you can just say super().`__init__`() instead of super(ChildB, self).`__init__`() which IMO is quite a bit nicer.

http://stackoverflow.com/questions/576169/understanding-python-super-with-init-methods

[Python2.7中的super方法浅见](http://blog.csdn.net/mrlevo520/article/details/51712440)

## 30 Python 的dict
python是用hash表实现的dict对象
C中的描述
```c
PyDictKeysObject //实现字典的hash表
```
## 31 Python Web
### Nginx、WSGI、 uWSGI、 uwsgi的区别

#### WSGI
全称是Web Server Gateway Interface，WSGI不是服务器、python模块、框架、API或者任何软件，只是一种**规范**，描述web server如何与web application通信的规范。它是 Python 语言定义出来的 Web 服务器和 Web 应用程序之间的简单而通用的接口，基于现存的 CGI 标准设计，后来在很多其他语言中也出现了类似的接口。 

按照 web 组件分类，WSGI 内部可以分为三类：
- web 应用程序
- web 服务器
- web 中间件。

**web 应用程序端**的部分通过Python 语言的各种 Web 框架实现，比如 Flask，Django这些，有了框架，开发者就不需要处理 WSGI，框架会帮忙解决这些，开发者只需处理 HTTP 请求和响应。
**web 服务器**的部分就要复杂一点，可以通过 uWSGI 实现，也可以用最常见的 Web 服务器，比如 Apache、Nginx，但这些 Web 服务器没有内置 WSGI 的实现，是通过扩展完成的。如 Apache，通过扩展模块 mod_wsgi 来支持WSGI，Nginx可以通过代理的方式，将请求封装好，交给应用服务器，比如 uWSGI。uWSGI 可以完成 WSGI 的服务端，进程管理以及对应用的调用。
**WSGI 中间件**的部分可以这样理解：我们把 WSGI 看做桥，这个桥有两个桥墩，一个是应用程序端，另一个是服务器端，那么桥面就是 WSGI 中间件，中间件同时具备服务器、应用程序端两个角色，当然也需要同时遵守 WSGI 服务器和 WSGI 应用程序两边的限制和需要。

#### uwsgi 和 uWSGI
> uWSGI 是一个 Web 服务器程序；

> WSGI，上面已经谈到，是一种协议；

> uwsgi 也是一种协议，uWSGI 实现了 uwsgi、WSGI、http 等协议。 

*uwsgi 的介绍可以看这里：uwsgi 是 uWSGI 使用的一个自有的协议，它用4个字节来定义传输数据类型描述。尽管都是协议，uwsgi 和 WSGI 并没有联系，我们需要区分这两个词。*

#### Nginx
Nginx 是高效的 Web 服务器和反向代理服务器，可以用作负载均衡（当有 n 个用户访问服务器时，可以实现分流，分担服务器的压力），与 Apache 相比，Nginx 支持高并发，可以支持百万级的 TCP 连接，十万级别的并发连接，部署简单，内存消耗少，成本低，但 Nginx 的模块没有 Apache 丰富。Nginx 支持 uWSGI 的 uwsgi 协议，因此我们可以将 Nginx 与 uWSGI 结合起来，Nginx 通过 uwsgi_pass 将动态内容交给 uWSGI 处理。

#### uWSGI 和 Nginx 的关系

从上面的讲解中，我们知道，uWSGI 可以起到 Web 服务器的作用，那么为什么有了 uWSGI 还需要 Nginx 呢？
最普遍的说法是 Nginx 对于处理静态文件更有优势，性能更好。其实如果是小网站，没有静态文件需要处理，只用 uWSGI 也是可以的，但加上 Nginx 这一层，优势可以很具体：
1. 对于运维来说比较方便，如果服务器被某个 IP 攻击，在 Nginx 配置文件黑名单中添加这个 IP 即可，如果只用 uWSGI，那么就需要在代码中修改了。另一方面，Nginx 是身经百战的 Web 服务器了，在表现上 uWSGI 显得更专业，比如说 uWSGI 在早期版本里是不支持 https 的，可以说 Nginx 更安全。
2. Nginx 的特点是能够做负载均衡和 HTTP 缓存，如果不止一台服务器，Nginx 基本就是必选项了，通过 Nginx，将资源可以分配给不同的服务器节点，只有一台服务器，也能很好地提高性能，因为 Nginx 可以通过 headers 的Expires or E-Tag，gzip 压缩等方式很好地处理静态资源，毕竟是 C 语言写的，调用的是 native 的函数，针对 I/O做了优化，对于动态资源来说，Nginx 还可以实现缓存的功能，配合 CDN 优化（这是 uWSGI 做不到的）。Nginx 支持epoll/kqueue 等高效网络库，能够很好地处理高并发短连接请求，性能比 uWSGI 不知道高到哪里去了。
3. 如果服务器主机上运行了PHP，Python 等语言写的多个应用，都需要监听80端口，这时候 Nginx 就是必选项了。因为我们需要一个转发的服务。

## 31 Python 内存管理机制
### 引用计数
查看对象的引用计数：`sys.getrefcount()`

#### 引用计数增加

- 1、对象被创建
- 2、另外的别人被创建
- 3、作为容器对象的一个元素
- 4、被作为参数传递给函数：foo(x)

#### 引用计数减少
- 1、对象的别名被显式的销毁
`del m`
- 2、对象的一个别名被赋值给其他对象
- 3、对象从一个窗口对象中移除，或，窗口对象本身被销毁
`a.remove(123)`
- 4、一个本地引用离开了它的作用域，比如上面的foo(x)函数结束时，x指向的对象引用减1。
### 垃圾回收
[垃圾回收](#24-python垃圾回收机制)
### 内存池机制
Python中有分为大内存和小内存：（256K为界限分大小内存）
- 1、大内存使用malloc进行分配
- 2、小内存使用内存池进行分配
- 3、Python的内存池(金字塔)

### Python的内存池(金字塔)
- 第3层：最上层，用户对Python对象的直接操作
- 第1层和第2层：内存池，有Python的接口函数PyMem_Malloc实现-----若请求分配的内存在1~256K字节之间就使用内存池管理系统进行分配，调用malloc函数分配内存，但是每次只会分配一块大小为256K的大块内存，不会调用free函数释放内存，将该内存块留在内存池中以便下次使用。
- 第0层：大内存-----若请求分配的内存大于256K，malloc函数分配内存，free函数释放内存。
- 第-1，-2层：操作系统进行操作

## 32 Python的序列类型
**list, tuple, range**


### 不可变序列类型

不可变序列类型普遍实现而可变序列类型未实现的唯一操作就是对 hash() 内置函数的支持。

这种支持允许不可变类型，例如 tuple 实例被用作 dict 键，以及存储在 set 和 frozenset 实例中。

尝试对包含有不可哈希值的不可变序列进行哈希运算将会导致 TypeError。

有：
- tuple
- range：通常用于在 for 循环中循环指定的次数。


### 可变序列类型

顾名思义，该序列是可变的，不支持hash()

- list

## 3 面向对象的特点

本文介绍面向对象的3个特点：

- 1 封装
封装是从业务逻辑中抽象对象时，要赋予对象相关数据与操作，将一些数据和操作打包在一起的过程。封装是使用对象的主要魅力之一，它提供了一个简单方法来创建复杂方案，解决了世界是如何工作的这一问题，我们自然的认为周围的世界是由相互作用的对象组成，每个对象都有自己相关的数据，并能完成一定的功能，从设计的角度来看，封装还提供了一个重要的服务，它分开了是什么和怎么做这两个问题。对象的实现与使用是相互独立的，封装的另外一个优势是支持代码复用，它可以将常用功能以组件方式打包起来。

- 2 多态
多态意味着多种形式，当用面向对象时，它是指对象是怎么回应一个依赖于对象类型或种类的消息。多态的作用是让程序在不同情况下用一个函数名启用不同的方法。
多态举例：
示例代码
```py
class Person(object):
    def __init__(self,name,sex):
        self.name = name
        self.sex = sex

    def print_title(self):
        if self.sex == "male":
            print("man")
        elif self.sex == "female":
            print("woman")

class Child(Person):                            # Child 继承 Person
    pass

May = Child("May","female")
Peter = Person("Peter","male")

print(May.name,May.sex,Peter.name,Peter.sex)    # 子类继承父类方法及属性
May.print_title()
Peter.print_title()
```
在示例代码的 Child 类中重写 print_title() 方法：若为male，print boy；若为female，print girl
```py
class Child(Person):                # Child 继承 Person
    def print_title(self):
        if self.sex == "male":
            print("boy")
        elif self.sex == "female":
            print("girl")

May = Child("May","female")
Peter = Person("Peter","male")

print(May.name,May.sex,Peter.name,Peter.sex)
May.print_title()
Peter.print_title()
```
当子类和父类都存在相同的 `print_title()`方法时，子类的 `print_title()` 覆盖了父类的 `print_title()`，在代码运行时，会调用子类的 `print_title()`。
这样，我们就获得了继承的另一个好处：**多态**。 

多态给予了面向对象系统极大的灵活性，对象可以用该对象应该用的方式来执行动作，如果没有面向对象，这种灵活性很难实现。当我们需要传入更多的子类，例如新增 Teenagers、Grownups 等，我们只需要继承 Person 类型就可以了，而print_title()方法既可以不重写（即使用Person的），也可以重写一个特有的。这就是多态的意思。调用方只管调用，不管细节，而当我们新增一种Person的子类时，只要确保新方法编写正确，而不用管原来的代码。这就是著名的**开闭**原则：

    对扩展开放（Open for extension）：允许子类重写方法函数
    对修改封闭（Closed for modification）：不重写，直接继承父类方法函数

- 3 继承
一个类（subclass,基类）可以继承另一个类（superclass，超类）.
举例：
建立一个系统以记录员工信息，需要一个Employee类，它包含所有员工都具有的一般信息，其中一个方法是homeAddress()，该方法可返回员工的住址信息。员工分为正式员工(按月发工资)和临时员工（按天发工资），为正式员工定义一个MonthlyEmployee类（Employee类的子类），为临时员工定义一个DaylyEmployee类（Employee类的子类），这两个子类都继承了Employee类，因此都有homeAddress()方法，然而，不同类型的员工，发工资的方法不同，正式员工有monthlyPay()方法，临时员工有daylyPay()方法。
**继承的优点：**
1、建造系统中的类，避免重复操作，例如，我们不必为两个子类分别写一个homeAddress()方法，从Employee类继承即可。
2、新类经常是基于已经存在的类，这样就可以提升代码的复用程度。
