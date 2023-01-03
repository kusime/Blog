---
title: A Guide to Python's Magic Methods (Rafe Kettler, 2012)
index_img: /gallery/2022-01-14-13-01-48.png
date: 2022-01-17 00:48:02
updated: 2022-01-17 00:48:02
top: true
tags:
  - Python
categories:
  - Python
---


{% note danger %}
[A Guide to Python's Magic Methods (Rafe Kettler, 2012)](http://web.archive.org/web/20161024123835/http://www.rafekettler.com/magicmethods.html)
This translation version Licensed under Creative Commons CC--NC-BY-SA (see http://creativecommons.org/licenses/by-nc-sa/3.0/).
这篇翻译版本遵循 [CC--NC-BY-SA 3.0](http://creativecommons.org/licenses/by-nc-sa/3.0/) .转载请标明出处
{% endnote %}

# A Guide to Python's Magic Methods

## 1 Introduction

This guide is the culmination (高潮) of a few months’ worth of blog posts. The subject(主题) is magic methods.

什么是 magic methods? , 它们是 Python 中面向对象的一切,它们是很特别的方法,你可以通过定义它们来给你的类添加"魔法",它们都是由 `__` 双下划线开头和结尾的(e.g. `__init__` or `__lt__`).,然后关于这些 "魔法" 的文档也很少,尽管它们是很重要的方法(They’re also not as well documented as they need to be)所有 Python 中的魔法方法都出现在其文档的(同一部分)(same section),但是它们是分散的,并且 (组织松散)(loosely organized),而且其文档几乎也没有几个例子.(也许是被设计的十分好了,因为它们是属于`(语言引用)(language reference)`并且也是伴随着十分无聊的语法解释,等)

所以,为了修复我认为是 Python 文档中的缺陷,我开始提供一些(更加直白的表达)(more plain-English),并且(结合具体的实例)(example-driven)来解释 Python 中的魔法方法,我开始了每周发布博客文章,(终于我完成了这些工作)(I’ve finished with those),并且(我统一的整理到了这篇指南中)(I’ve put together this guide.)

希望你能享受这篇文章,(或是作为教程,或是作为复习,或是作为文献)(Use it as a tutorial, a refresher, or a reference;),总之我的指南只是希望去提供一个(友好的方式)( user-friendly guide)来解释 Python 中的魔法方法.


## 2 (Construction and Initialization)(构造和初始化)

大家都知道最基本的魔法方法`__init__`,我们可以通过这个方式来定义对象初始化的行为,然而,当我尝试去调用 `x = SomeClass()` 时,` __init__ `不是一个被调用的方法,实际上,这有一个叫做` __new__`的方法,实际上实例是它创建的,然后在(创建时)(at creation),(传递任何参数)(passes any arguments)(到初始化器)(on to the initializer),然后在(对象的生命周期结束的时候)(end of the
object’s lifespan),会有一个` __del__.`方法,(让我们来仔细瞧瞧这)( Let’s take a closer look at these) 3 个魔法方法.

`__new__(cls, [...)` , `__new__`是在(类的初始化)(object’s instantiation.)的时候第一个被调用的方法,它接收类,然后其它的参数将会将传递给`__init__`,`__new__`(几乎很少)(fairly rarely)使用,(但其确实由它的用途)(but it does have its purposes),(特别是)(particularly )在(子类化)(subclassing )(不可变类型)(immutable type)[如元组或字符串]时,我不想在继续深究其细节,因为这不是那么的有用处,但它在 Python 文档中(有详细介绍)(covered in great detail)

`__init__(self, [...) `是类的(初始化器)(initializer),不管(主构造器)(primary constructor)调用什么,它都会被传递,[举个例子,如果我们调用了`x = SomeClass(10, ’foo’),`,`__init__`会接收到 `10 and ’foo’ as arguments.`]`__init__`是在Python类定义中使用的最多的方法

`__del__(self)`如果说`__new__ and __init__`是构造对象的函数,那么,`__del__`就是(析构函数)(destructor),它和`del x`的(执行方式)(implement behavior)不一样(因此该代码不会转换为` x.__del__()`),(相反)(Rather)，它定义了对象被垃圾回收时的行为。这对于对象(在删除时)( upon deletion)需要额外的清理工作来说是十分有用的,比如说网络编程或者文件对象,注意了,在解释器退出的时候,如果对象仍然存在,则没有保证`__del__ `会被执行,所以 `__del__ ` (不能被视为)(can’t serve as a replacement)代替(良好的编码习惯)(good coding practices)[就比如说在你使用完一个链接的时候你要关闭它],实际上`__del__`应该尽可能的避免使用,因为它被调用的话会存在(潜在的隐患)(precarious[不安的]
circumstances[情况]),谨慎行之!

(综上)(Putting it all together),here’s an example of `__init__ and __del__` in action:


```python
from os . path import join
class FileObject :
    ’’’ Wrapper for file objects to make sure the file gets closed
    on deletion . ’’’
    def __init__ ( self , filepath = ’~ ’ , filename = ’ sample . txt ’ ):
    # open a file filename in filepath in read and write mode
    self . file = open ( join ( filepath , filename ) , ’r + ’)
def __del__ ( self ):
    self . file . close ()
    del self . file
```

## 3 Making (Operators)(操作符) Work on Custom Classes

使用 Python 的魔法方法(最大的优势)(biggest advantages)是其提供了一种简单的方式来让对象行为的就像内置的类型一样,这就意味着,你可以避免丑陋,并反直觉,且不标准的,方式来执行基本的操作符,就比如说在一些语言中,通常你需要如下进行编写

```python
if instance . equals ( other_instance ):
    # do something
```
当然了,你在Python中也可以这么做,但是这就(增加了复杂度)( adds confusion),和(不必要的冗余)(unnecessarily verbose).不同的库(可能对相同的操作使用不同的名称)(might use different names for the same operations),(让使用者花更多的功夫在不必要的事情上)(making the client do way more work than necessary),有了魔法方法,(比如)(however),我们可以定义一个方法,(`__eq__, in this case`),(这个时候我们就可以用以下代码来代替上面代码的含义)(and say what we mean instead):

```python
if instance == other_instance :
    # do something
```

(那是魔法方法的一部分功能)(That’s part of the power of magic methods),其绝大多数允许我们定义操作符的含义,以至于我们可以在自己的类中使用它们(操作符),就好像他们(自定义类)是内置类型一样.

### 3.1 Comparison (magic)(比较) methods

python设计了一大堆的魔法方法,用来在对象之间使用操作符来(进行直观的比较)(implement intuitive comparisons),(而不是傻乎乎的方法调用)(not awkward method calls),也提供了修改Python默认对对象的比较行为的方式,这儿有一系列的方法,以及其作用

`__cmp__(self, other) :` `__cmp__ ` 是最基本的比较魔法方法,(它实际上实现了)(actually implements behavior)所有比较符(` ==, !=, etc`)的行为,但是其可能不总是以你想要的方式行动,(例如,如果一个实例与另一个实例相等是由一个(指标)(criterion)决定的,或者一个实例是否大于另一个实例,是由其他东西(决定的)( determined by)),如果`self < other`,`__cmp__`应当返回一个负值,如果`self == other` 则为零，如果 `self > other` 则为正,相比一次定义他们,通常来说,分别定义会更好,但是当你需要以相同的标准来进行比较的时候,`__cmp__`能恰到好处的减少重复,并且更加明了

{% note danger %}

{% note success %}
[Why can't I use the method __cmp__ in Python 3 as for Python 2?](https://stackoverflow.com/questions/8276983/why-cant-i-use-the-method-cmp-in-python-3-as-for-python-2)

{% endnote %}

在Python3 中,`__cmp__`方法被废除了,需要单独指定

![图片描述](/gallery/2022-01-17-15-13-45.png)

---

![图片描述](/gallery/2022-01-17-15-04-21.png)

因为这个指南是有一定年纪了(2014年写的),所以有些变动是可以理解的,我的想法就是在我翻译完这个文档的的时候,或者过程中我去实现那些代码,然后做一个可行性的表格,然后相当于是对这个手册的一个更新吧.


```doc
__eq__(self, other) Defines behavior for the equality operator, ==.
__ne__(self, other) Defines behavior for the inequality operator, !=.
__lt__(self, other) Defines behavior for the less-than operator, <.
__gt__(self, other) Defines behavior for the greater-than operator, >.
__le__(self, other) Defines behavior for the less-than-or-equal-to operator, <=.
__ge__(self, other) Defines behavior for the greater-than-or-equal-to operator, >=
```

{% endnote %}


例如，考虑一个类来建模一个词,我们想要我们的单词能(按字典顺序)(lexicographically),那么默认的方法就是比较字符串,(但是我们想让其比较基于一些别的指标)(but we also might want to do it based on some other criterion),比如说长度,或者(音节)(syllables)的个数,(这里有个实践)( Here’s an implementation:)

{% note warning %}
原文是Python2 的代码,我重写了一份

![图片描述](/gallery/2022-01-17-15-54-58.png)

![图片描述](/gallery/2022-01-17-15-53-50.png)
{% endnote %}

现在我们可以创建两个 `Words`对象,然后基于长度进行比较了,注意,我们没有没有定义 `__eq__ and __ne__`,这是因为这会导致一些古怪的行为 (notably that Word(’foo’) == Word(’bar’) would evaluate to true) ,更具长度来测试相等(是没有意义的)(It wouldn’t make sense to),所以让我们退回到 `str` 其原本的相等含义(注:就是两个字符串完全一样才相等的意思)

现在是个好时候来注意你不必为每个比较方法来定义其魔法方法,在标准库模组`functiontoos`中,友好的为我们提供了一个类修饰符,其使我们只需要定义 ` __eq__ ` 和任一一个魔法方法`(e.g. __gt__, __lt__, etc.) `就能实现各种比较方法,这个特点智能在 Python 2.7 及其之后的版本中使用,如果你能有机会使用的话,这将会节省大量的时间和精力,你可以通过在你的类定义前面放置一个`@total_ordering `

{% note success %}
可以使用,但是我估计需要保证传出来的是 bool 值,然后用一些逻辑学推导出来的.
![图片描述](/gallery/2022-01-17-16-19-18.png)
{% endnote %}

### 3.2 Numeric Magic Methods

就像你可以创造一些方法来使你类的实例能通过比较操作符来比较一样,你可以定义数值运算符的行为,(系好你的安全带老伙计我们要发车了:D)(Buckle your seat belts, folks),这儿还有很多东西呢,我把数字运算方法分为5个类别.(一元运算符、普通算术运算符、反射算术运算符,增强赋值和类型转换。)(unary operators, normal arithmetic operators, reflected arithmetic operators)稍后会做详细介绍.

#### 3.2.1 Unary operators and functions

一元运算符和函数只有一个操作数,否定、绝对值等.

`__pos__(self) ` : 执行一元正符的行为. (e.g. +some_object)

`__neg__(self) `: 执行定义负符的行为  (e.g. -some_object)

`__abs__(self)` : 定义调用内置方法`abs()`时的行为

`__invert__(self) `: 执行定义`~`的行为

`__round__(self, n) `: 定义调用内置方法`round()`时的行为,`n` 是一个十进制数字,用来传入给`round`

`__floor__(self) `: 定义调用`math.floor`方法,即 rounding down to the nearest integer

`__ceil__(self)` : Implements behavior for `math.ceil()`, i.e., rounding up to the nearest integer.

`__trunc__(self)` : Implements behavior for `math.trunc()`, i.e., (truncating)(截断) to an integral.

#### 3.3 Normal arithmetic operators

现在我们来学习一下经典的二元操作符,`+,-,*,...` 和于此类似的,这些在(大多数情况下其意义都是十分显而易见的)(for the most part, pretty self-explanatory.)

`__add__(self, other) ` Implements addition.

`__sub__(self, other)` Implements subtraction.

`__mul__(self, other)` Implements multiplication.

`__div__(self, other)` Implements division using the` /` operator.

`__mod__(self, other)` Implements modulo using the `%` operator.

`__pow__` Implements behavior for exponents using the `**` operator.

`__lshift__(self, other)` Implements left bitwise shift using the `<<` operator.

`__rshift__(self, other) `Implements right bitwise shift using the `>>` operator.

`__and__(self, other)` Implements bitwise and using the `&` operator.

`__or__(self, other)` Implements bitwise or using the `|` operator.

`__xor__(self, other)` Implements bitwise xor using the `^` operator

`__floordiv__(self, other)` Implements integer division using the `// `operator.

下面把一些不常用的(我看不懂).的单独收起来
```doc
`__truediv__(self, other)` Implements true division. Note that this only works when from `__future__` import division is in effect.
`__divmod__(self, other)` Implements behavior for long division using the `divmod()` built in function.
`__future__` import division is in effect.
```

##### 3.3.1 Reflected arithmetic operators

{% note warning %}
参考这里会更加的明了一点(有例子,有细节)

{% note success %}
[python魔法方法-反射运算和增量运算](https://www.cnblogs.com/scolia/p/5686267.html)
{% endnote %}

关于其博客原文说的其实有问题的..我提出来一下

![图片描述](/gallery/2022-01-17-21-38-18.png)

如果按照其逻辑来说这样直接copy下来代码当让不能运行,但是那个博主显然没有深度思考:D

![图片描述](/gallery/2022-01-17-21-37-16.png)

不妨把父级对象换为 `int`

![图片描述](/gallery/2022-01-17-21-42-34.png)

可以看见是能够正常运行的,所以说其逻辑是错误的,具体为什么我做了解释,应该是能理解的


{% endnote %}

(你知道为啥我单独把反射算数符单独放一点呢?)(You know how I said I would get to reflected arithmetic in a bit?),(你们可能会开始想,啊,这肯定是个大的,吓人的,陌生的概念.实际上这还是很简单的,这里有个例子)(Some of you might think it’s some big, scary, foreign concept. It’s actually quite simple. Here’s an example:)

`some_object + other`

那是个"正常"的加法,反射其实和这大抵相似,除了操作数是反向的

`other + some_object`

所以,所有的这些魔术方法实际上和没反射的方法做的是同样的事情,(除了在没反射的`self`其变成了反射中的`other`,反之亦然)(except the perform the operation with other as the first operand and self as the second, rather than the other way around),在绝大多数情况,反射操作符的结果和正常的结果是相等的,所以,你可以不用定义 `__radd__` 与之代替的是定义 `__add__` ,注意,在运算符的左侧的对象一定不能定义其非反射版本(或者返回`NotImplemented`),例如，在示例中, `some_object.__radd__` 一定要在 `other` 没有定义 `__add__` 的情况下才会被调用

{% note success %}
还是结合其具体的引用的博主的那个逻辑,它举得最后一个例子相当的好!

![图片描述](/gallery/2022-01-17-21-49-36.png)

这里直接使用它的说法

![图片描述](/gallery/2022-01-17-21-50-13.png)

以及结合我自己的实践总结出来的

![图片描述](/gallery/2022-01-17-21-59-37.png)

综上:
1. <font color="#FF0000">**反向方法总是在右操作对象中定义调用**</font>
2. <font color="#FF0000">**正向方法的优先级大于其反向**</font>
3. <font color="#FF0000">**寻找属性的逻辑符合向上传递的逻辑**</font>
{% endnote %}

下面是反射方法的总结:

`__radd__(self, other)` Implements reflected addition.

`__rsub__(self, other)` Implements reflected subtraction.

`__rmul__(self, other)` Implements reflected multiplication.

`__rfloordiv__(self, other)` Implements reflected integer division using the` //` operator.

`__rdiv__(self, other)` Implements reflected division using the `/` operator.

`__rmod__(self, other)` Implements reflected modulo using the` % `operator.

`__rpow__` Implements behavior for reflected exponents using the `**` operator.

`__rlshift__(self, other)` Implements reflected left bitwise shift using the `<<` operator.

`__rrshift__(self, other)` Implements reflected right bitwise shift using the `>>` operator.

`__rand__(self, other)` Implements reflected bitwise and using the `&` operator.

`__ror__(self, other)` Implements reflected bitwise or using the `|` operator.

`__rxor__(self, other)` Implements reflected bitwise xor using the `^` operator.

一下是不常用的
```doc
`__rtruediv__(self, other)` Implements reflected true division. Note that this only works when from `__future__` import division is in effect.
`__rdivmod__(self, other)` Implements behavior for long division using the `divmod()` built in function, when `divmod(other, self)` is called.
```

##### 3.3.2 Augmented assignment(自增符)

Python 还提供了广泛并且多样的魔法方法来自定义自增符的行为,你可能早就熟悉自增符了,它就是"普通"运算符和赋值符的组合,如果你还是不晓得我在说啥,这有个例子给你

```python
x = 5
x += 1 # in other words x = x + 1
```

所有的这些方法的返回值应该会被分配到左手边的变量中去, (for instance, for `a += b, __iadd__` might return` a + b`, which would be assigned to `a`). Here’s the list:

`__iadd__(self, other)` Implements addition with assignment.

`__isub__(self, other)` Implements subtraction with assignment.

`__imul__(self, other) `Implements multiplication with assignment.

`__ifloordiv__(self, other)` Implements integer division with assignment using the `//=` operator.

`__idiv__(self, other) Implements` division with assignment using the `/= `operator

`__imod__(self, other) `Implements modulo with assignment using the `%=` operator.

`__ipow__` Implements behavior for exponents with assignment using the `**=` operator.

`__ilshift__(self, other)` Implements left bitwise shift with assignment using the `<<=` operator.

`__irshift__(self, other)` Implements right bitwise shift with assignment using the `>>=` operator.

`__iand__(self, other)` Implements bitwise and with assignment using the `&=` operator.

`__ior__(self, other) `Implements bitwise or with assignment using the `|=` operator.

`__ixor__(self, other)` Implements bitwise xor with assignment using the` ^= `operator.


##### 3.3.3 Type conversion magic methods

Python 还设计了一系列用来进行类型转化的魔法方法,比如说` float().`,这儿是:

`__int__(self)` Implements type conversion to `int`.

`__long__(self)` Implements type conversion to `long`.

`__float__(self)` Implements type conversion to `float`.

`__complex__(self) `Implements type conversion to `complex`.

`__oct__(self) `Implements type conversion to `octal`.

`__hex__(self) `Implements type conversion to `hexadecimal`.

`__index__(self)` Implements type conversion to an int when the object is used in a <font color="#FF0000">**slice expression**</font> . If you define a custom numeric type that might be used in slicing, you should define `__index__`. [参考](https://stackoverflow.com/questions/27449310/python-index-special-method)

`__trunc__(self)` Called when `math.trunc(self)` is called. `__trunc__` should return the value of self (truncated)(截断 ) to an integral type (usually a long).

`__coerce__(self, other)` Method to implement mixed mode arithmetic. `__coerce__` should return None if type conversion is impossible. Otherwise, it should return a pair (2-tuple) of self and other, manipulated to have the same type.

## 4 Representing your Classes

通常一个类有其字符串表示是十分有用的,在Python中,有几个方法,你可以通过实现他们来定义当内置函数被调用的时候如何返回类的表示

`__str__(self)` Defines behavior for when `str()` is called on an instance of your class.

`__repr__(self)` Defines behavior for when `repr() `is called on an instance of your class. The major difference between str() and repr() is (intended audience)(目标听众). repr() is intended to produce output that is mostly machine-readable (in many cases, it could be (valid)(有效的) Python code even), whereas str() is intended to be human-readable.

`__unicode__(self)` Defines behavior for when `unicode()` is called on an instance of your class. `unicode()` is like `str()`, but it returns a unicode string. Be wary: if a client calls `str()` on an instance of your class and you’ve only defined `__unicode__()`, <font color="#FF0000">**it won’t work**</font>. You should always try to define `__str__()` as well in case someone doesn’t have the luxury of using unicode.

`__format__(self, formatstr)` Defines behavior for when an instance of your class is used in new-style string formatting. For instance, " `Hello, 0:abc!".format(a)`would lead to the call `a.__format__("abc")`. This can be useful for defining your own numerical or string types that you might like to give special formatting options.

`__hash__(self)` Defines behavior for when hash() is called on an instance of your class. It has to return an integer, and its result is used for quick key comparison in dictionaries. Note that this usually (entails implementing)(需要实施) `__eq__` as well. Live by the following rule:` a== b `implies `hash(a) == hash(b)`.

`__nonzero__(self)` Defines behavior for when `bool()` is called on an instance of your class. Should return `True` or `False`, depending on whether you would want to consider the instance to be True or False.

`__dir__(self)` : Defines behavior for when `dir()` is called on an instance of your class. This method should return a list of attributes for the user. (Typically)(通常), implementing `__dir__` is unnecessary, but it can be (vitally important)(非常重要) for (interactive)(交互的) use of your classes if you redefine `__getattr__` or `__getattribute__` (which you will see in the next section) or are otherwise dynamically generating attributes.

我们差不多要结束魔法方法指南中无聊并且(无实例)( example-free)的部分了,截至目前,(我们已经介绍了一些)(we’ve covered some of)相对来说很基础的魔法方法,现在是时候来看(更加高级的材料了)(more advanced material.).


## 5 Controlling Attribute Access

许多从别的语言转来Python的抱怨道,这里面的类缺少真正的封装,(比如说没有办法去定义一个私有属性,然后和一个公共的获取器和设置其),这个说法差不多是事实:相比于在方法或字段使用显式的修饰符来说,Python(通过"魔法"来封装)(encapsulation through “magic”)上(大有作为)( accomplishes a great deal),让我们瞧吧.

`__getattr__(self, name) `You can define behavior for when a user attempts to access an attribute that doesn’t exist (either at all or yet). This can be useful for catching and redirecting common misspellings, giving warnings about using (deprecated)(已弃用) attributes (you can still choose to compute and return that attribute, if you wish), or (deftly)(灵巧地) handing an `AttributeError`.  <font color="#FF0000">**It only gets called when a nonexistent attribute is accessed**</font> , however, so it isn’t a true (encapsulation solution.)(封装解决方案。)

`__setattr__(self, name, value)` Unlike `__getattr__`, `__setattr__` is an encapsulation solution. It allows you to define behavior for assignment to an attribute regardless of whether or not that attribute exists, meaning you can define custom rules for <font color="#FF0000">**any changes in the values of attributes.**</font>  However, you have to be careful with how you use `__setattr__`, as the example at the end of the list will show

`__delattr__(self, name)` This is the exact same as `__setattr__`, but for (deleting)(删除) attributes instead of setting them. The same (precautions)(预防措施) need to be taken as with `__setattr__` as well in order to prevent infinite recursion (calling del self.name in the (implementation)(执行) of `__delattr__` would cause infinite recursion).


`__getattribute__(self, name)` (After all this)(毕竟), `__getattribute__` fits in pretty well with its companions `__setattr__` and `__delattr__`. However, I don’t recommend you use it. `__getattribute__` can only be used with new-style classes (all classes are new-style in the newest versions of Python, and in older versions you can make a class new-style by sub classing object. It allows you to define rules for (whenever)(每当) an attribute’s value is accessed. It suffers from some similar infinite recursion problems as its (partners-in-crime)(同谋) (this time you call the base class’s `__getattribute__` method to prevent this). It also mainly (obviates)(避免) the need for `__getattr__`, which only gets called when `__getattribute__` is implemented if it is called explicitly or an AttributeError is raised. This method can be used (after all, it’s your choice), but I don’t recommend it because it has a small use case (it’s far more rare that we need special behavior to retrieve a value than to assign to it) and because it can be really difficult to implement bug-free

---

因为比较难理解我下面翻译成中文


`__getattr__(self, name)`  你可以定义当用户访问不存在(目前或根本)的属性时的行为,这个捕捉并重定向普通的拼写错误是十分有用的,在尝试引用已经被废弃的属性时给出写警告(如果你想的话,你依然可以选择计算并返回其属性),或者灵活的处理`AttributeError`, <font color="#FF0000">**这个方法只会在尝试接触不存在的属性时会被调用**</font>,然而,这也不是真封装解决方案

`__setattr__(self, name, value) ` : 与 `__getattr__` 不同， `__setattr__` 是一个封装解决方案,其允许你定义给属性赋值时的行为,不论次属性是否存在,意味这你可以客制化改变属性的规则,然而,你在使用 `__setattr__`时还是得留意,之后我会给个例子给你的

`__delattr__(self, name) ` 这和 `__setattr__` 是差不多的,只不过就是删除其属性而不是设置他们,它需要和 `__setattr__`一样采取相同的预防无限递归的措施（(在 __delattr__ 的实现中调用 del self.name 会导致无限递归)(calling `del self.name` in the implementation of `__delattr__` would cause infinite recursion))）

`__getattribute__(self, name)`,毕竟,`__getattribute__(self, name)`和其同伴 `__setattr__` and `__delattr__`(十分类似.)(fits in pretty well) ,但是我还是推荐你别使用,`__getattribute__ ` 只能在新式类中使用,(在最新版本中的所有类都是新式类,在老版本,你可以通过让你的老版本类继承`object`类使其变成新式类),它允许你定义 <font color="#FF0000">**(每当属性被获取到时)(访问属性值时)**</font> 的规则,它遇到了一些与其犯罪伙伴类似的无限递归问题,(这时,你可以通过调用基类的`__getattribute__`来预防这个问题),同时很大程度上避免了对`__getattr__`的需要,因为其只有当 `__getattribute__` 被显式调用或引发 `AttributeError` 时才被调用。这个方法可以使用,但是我不推荐你使用,因为其适用面太小了(it’s far more rare that we need special behavior to retrieve a value than to assign to it),并且因为实现无错误可能真的很困难

{% note warning %}
emm,`(it’s far more rare that we need special behavior to retrieve a value than to assign to it` 这句话我纠结了好久..我干脆再去学一下比较级了..

下面这个我想出来的翻译我觉得相对来说还是比较完美的 :P

![图片描述](/gallery/2022-01-18-12-43-39.png)

assign:

![图片描述](/gallery/2022-01-18-12-46-46.png)

{% endnote %}

---

在你定义任何控制属性访问的方法时,你很容易引发问题,思考一下下面这个例子

```python
def __setattr__(self,name,value):
    self.name = value
"""
因为每次分配属性的时候, `__setattr__ ()` 就会被调用,这就是一个递归,所有,实际上`self.name = value`相当于就是在调用`self . __setattr__ ( ’ name ’, value )`然后由于其不断调用自生,这就导致递归不断执行,直到崩溃
"""

def __setattr__(self,name,value):
    self.__dict__[name] = value
    # 把其值赋值到类字典中去
    #....


```

再者,Python的魔法方法有着难以置信的力量,与此同时,巨大的能力就意味这巨大的责任,了解魔法方法的正确使用姿势是十分重要的事情,以至于你不会让你的代码崩溃

所有,关于控制Python属性控制我们学到了些啥?(这些东西不是轻易就能使用的)(It’s not to be used lightly.)实际上,其有种趋向与一种绝对的力量并且非好理解,(但是其存在的原因好像微不足道)(. But the reason why it exists is to scratch a (certain)(确定) (itch)(痒)):Python 不会寻求一种方法让坏的事情变得不可能,相对的,只会让他更加的难以出现,自由是至高无上的,所以你确实你可以为所欲为,下面有个例子,在一些特殊的属性被访问的时候会做出一些行动的方法.(注意,我们在代码中使用 `super`函数是因为不是所有的类都有`__dict__`这个属性)

{% note warning %}

![图片描述](/gallery/2022-01-18-13-39-04.png)

原文代码已经失效:

![图片描述](/gallery/2022-01-18-13-42-29.png)

简单的理解就是调用父方法来实现设置属性失败了,所以我们可以对应的对照一个修改

![图片描述](/gallery/2022-01-18-14-01-17.png)

{% endnote %}

## 6 Making Custom Sequences

这儿有许多方式能让你的类像内置序列(字典,元组,列表,字符串)一样,(这些也是我在Python中最喜欢的魔法方法,因为其提供了绝对的控制权,并且让你类的实例和一系列内置方法完美运行的一种方式,)(These are by far my favorite magic methods in Python because of the absurd degree of control they give you and the way that they magically make a whole array of global functions work beautifully on instances of your class.)(但是在我开始将这些好东西之前我们们需要带过一下前置知识)(But before we get down to the good stuff, a quick word on requirements.)

### 6.1 Requirements

现在我们要讨论的是在Python中创建你自己的序列,现在是时候来谈谈`协议`,协议和其他语言中的(接口)(interfaces)类似,他们要求你必须定义一些方法,然而,在Python中的协议是(松散)(informal)的并且(不需要明确声明)(require no explicit declarations)(来实现的)(to implement),换句话说,他们更像是一种 (方针)(guidelines)

为什么现在要讨论协议呢?(因为在Python中实现客制化容器类型涉及其中的一些协议)(Because implementing custom container types in Python involves using some of these protocols),首先,(有定义不可变容器的协议)(there’s the protocol for defining immutable containers):为了创建一个不可变容器,你只需要定义 `__len__ and __getitem_ _ (more on these later)`,对应的,可变类型协议则需要在所有不可变类型所需的定义之上添加对`__setitem__ and __delitem__`的定义,最后,如果你想要你创建的对象是(可迭代的)(iterable),那么你需要去定义`__iter__`,它会返回一个迭代器,然后其迭代器必须遵守迭代器协议,其需要迭代器调用`__iter__`(返回其自生),和`next`方法

### 6.2 The magic behind containers

`__len__(self)` 返回容器的长度,他们是 `可变` 和 `不可变` 协议的一个部分

`__getitem__(self, key) ` 使用标记`self[key]`,定义当`item`被访问时的行为,这个方法应该(合理地触发异常)(raise appropriate exceptions):如果键的类型错误，则为 `TypeError` ，如果该键没有对应的值，则为 `KeyError`.

`__setitem__(self, key, value)` 定义当进行`item` 赋值是的行为,使用标记符` self[key] = value`,这个是可变类型协议的一部分,再者,(你应该也要合理的处理异常)(you should raise KeyError and TypeError where appropriate)

`__delitem__(self, key)` 定义其`item`被删除时候的行为,这仅是可变类型容器的协议的一部分,当不存在的`key`被使用的时候,你也需要合理的触发异常

`__iter__(self)` 应该返回容器的迭代器,迭代器返回上下文数字,(最值得注意的是 )(most notably)iter() 内置函数和使用`for x in container`对迭代器进行轮询的时:迭代器是他们自己的对象,(他们)( own objects)也必须定义返回`self`的`__iter__`方法

`__reversed__(self) `定义内置函数`reversed()`被调用执行的行为.其应该返回序列的反向版本,其只对有序序列执行,比如说 列表或元组

`__contains__(self, item)` 其定义当使用 `in / not in.` 的行为,你问,为什么这个不是序列协议的一个部分?因为,当 `__contains__`没有定义的时候,Python就会迭代这个序列,然后如果找到在寻找的值就返回`True`

`__missing__(self, key)` 其是`dict`的子类,它定义了每当使用不在字典中存在的`key`来访问值的时候的行为(so,打个比方,如果我有一个字典`d`,然后使用`d["george"]`来访问其值,但是在字典中不存在这个`"george"`,`d.__missing__("george")`就会被调用)

### 6.3 An example

举一个例子,让我们看一个列表,其实现了一些你可能会在其他语言中使用的功能结构

{% note warning %}

![图片描述](/gallery/2022-01-18-18-22-02.png)`

对源代码进行了修改
```python
class KusimeList:  # oriented from object ,so the sequence function will defince by hand

    def __new__(cls):
        print("Instance Creat Success")
        return super().__new__(cls)

    def __init__(self):
        print("Start Init the Instace")
        self.value = []
        pass

    def __delattr__(self,name):# class attribute ,should not be called at __delitem__
        print("__delattr__ called")
        pass
    
    def __len__(self):
        return len(self.value)

    def __getitem__(self,key):
        return self.value[key]
    
    def __setitem__(self, key, value):
        self.value[key] = value
    
    def __delitem__(self,key):
        del self.value[key]

    def __iter__(self):
        return iter(self.value)
    
    def __reversed__(self):
        return reversed(self.value)# here reverse will not call self.__reversed__
    
    def append(self,value):
        self.value.append(value)
    
    def head(self):
        return self.value[0]

    def tail(self):
        return self.value[-1]
    
    #...
    
A = KusimeList()
A.append(1)
print (A[0])
```
![图片描述](/gallery/2022-01-18-18-20-07.png)
{% endnote %}

现在你有了实现你自己序列,相对来说有用的例子,当然了,客制化序列应用面会更广,但是少数的几个在标准化库(batteries included, right?)已经被实现了.比如说 `Counter,OrderedDict, and NamedTuple`

## 7 Reflection

你也可以通过定义一些魔法方法来控制内置函数`isinstance and issubclass()`反射的行为,这些魔法方法是

`__instancecheck__(self, instance)` Checks if an instance is an instance of the class you defined (`e.g. isinstance(instance, class).`)

`__subclasscheck__(self, subclass)` Checks if a class subclasses the class you defined (`e.g. issubclass(subclass, class)`).

这些魔法方法的使用面似乎很小,但是他们十分有用是真的,我不会花太多的事件来解释反射魔法方法,因为他们不是十分的有用,但是他们反应了一些Python面向对象编程和Python的哲学:总会有更简单的方法来完成某些事情,即使他们几乎没有必要,这些魔法方法看起来不是特别的有用,但是如果你需要他们的时候它们就在那,相比你也会是很开心的(本指南也是!)

## 8 Abstract Base Classes

{% note warning %}

{% note success %}
[参考-0](http://docs.python.org/2/library/abc.html)
[参考-1](https://docs.python.org/3.9/library/abc.html)
[参考-2](https://docs.python.org/3.9/glossary.html#term-abstract-base-class)
{% endnote %}

![图片描述](/gallery/2022-01-18-18-57-08.png)

为了整合起见,我去参考了Python3 的这个部分的文档.然后再做一个引用

![图片描述](/gallery/2022-01-18-19-01-54.png)
{% endnote %}


## 9 Callable Objects

也许你以及知道了,在Python中,(函数是一等对象)(functions are first-class objects),这就意味着,他们能被函数和方法传递,就好像其像任何其他对象一样,这是一个十分有力的特性

Python 中的一种特殊魔法方法允许您的类的实例(表现得好像它们是函数一样)(behave as if they were functions) ,以至于你能 "call" 他们,传递给函数,接收参数值,等等..这是另一个有用并且方便的一个特性,这使得在Python中编程变得(更加甜蜜了)(much sweeter)

`__call__(self, [args...]) ` 允许类的实例像函数一样被调用.本质上,这就意味着`x()` 和 `x.__call__()`是一样的,注意,`__self__` 会接收可变数量的参数,这就意味着你可以定义 `__call__` 就像定义其他函数一样,(你想要定义多少个参数就定义多少个)(taking however many arguments you’d like it to.)

`__call__` 在实例经常需要改变其状态的时候非常有用,"调用"这个实例是一种十分符合直觉并且优雅的一种改变对象状态的一种方式,作为一个可能表示实体在平面上的位置的例子类来看:

{% note danger %}
![图片描述](/gallery/2022-01-18-20-05-50.png)

![图片描述](/gallery/2022-01-18-20-05-40.png)
{% endnote %}

## 10 (Context Managers)(上下文管理器)

在Python 2.5 中,有一个新的关键词被引入到了Python中,一个新的代码重用方式,`with`陈述,上下文管理器这个概念其实说不上新了,(它早就作为库的一部分被实现了),但是知道 PEP 343 被接受后,他才成为了(第一类语言结构,)(as a first class language construct),你以前也许见过一下陈述:

```python
with open ( ’ foo . txt ’) as bar :
    # perform some action with bar
```

上下文管理器允许在对象的创建用 with 语句包装时对其进行设置和清理操作,上下文管理器的行为由两个魔术方法决定

{% note success %}
![图片描述](/gallery/2022-01-18-20-40-42.png)
{% endnote %}

`__enter__(self)` 定义在执行`with`创建的代码块之前上下文管理器应该做的事情,注意这个函数的返回值是和 `with` 的陈述的目标是绑定的,或者说是`as`之后的东西

`__exit__(self, exception_type, exception_value, traceback)` 定义在其代码块完成(或者中断)的时候上下文管理器应该做的一些事情,其可以用来处理异常,执行清理,或者是做一些在代码块结束就需要立即做的事情,如果代码块成功的执行了,`exception_type, exception_value, and traceback` 的值都会是`None`,否则,你就要选择是在这个函数中处理这个异常或者让用户自己处理这个异常.如果你不想处理,那么就保证`__exit__`返回值为 `True` ,(毕竟其已经完成了)(all is said and done),如果你不想让上下文管理器来处理异常就那么做就可以了.

`__enter__ and __exit__ ` 在有着明确定义的 设置和结束行为的时候十分的有用,或者你也可以使用这些方法来创建一个通用的上下文管理器,(其把请求转移到其他对象[包装其他对象])(that wrap other objects) 这里有个例子

{% note warning %}
原文举得例子是关闭Ftp链接的,但是我做了个差不多的
![图片描述](/gallery/2022-01-18-21-11-54.png)
{% endnote %}

看到了我们的包装器如何(优雅的)(gracefully)(处理合适与不合适的操作了吗?)(handled both proper and improper uses?),这就是上下文管理器和魔术方法的力量,注意下,Python的标准库中有个模组叫`contextlib`,其包含了上下文管理器,`contextlib.closing(),`这个方法做差不多一样的事情(除了说不能处理对象没有`close()`方法这种情况)

## 11 Building (Descriptor)(描述符) Objects

描述符是一种类,通过获取、设置或删除访问时,其能影响另一个类,描述符不是独立的,对应的,他们(主要是)(meant to)被主类所(掌管)(held),(当构建一个面向对象的数据库或者类,其属性的值取决于其他[的属性]的情况下是十分有用的.)(. Descriptors can be useful when building object-oriented databases or classes
that have attributes whose values are dependent on each other),(例如原始类中表达的是距离,其需要转化为平面上的坐标)((like distance from the origin in a class to represent a point on a grid)

{% note warning %}
![图片描述](/gallery/2022-01-18-21-46-10.png)
{% endnote %}

为了实现一个描述符,其需要至少实现 `__get__, __set__, and __delete__`三者之一,然我们看看那些魔法方法吧.

`__get__(self, instance, owner) ` 定义当描述符值被搜检时的行为, `instance` 是主类对象,owner就是主类自身

`__set__(self, instance, value)`定义当描述符值被改变的时的行为,`instance` 是主类对象,`value` (是要设置到修饰符的值)(value is the value to set the descriptor to)

`__delete__(self, instance)` 定义当描述符值被删除的时的行为, `instance` 是主类对象,owner就是主类自身

{% note warning %}
原文给的例子是做单位转换的,但是我之前发现了,Python默认初始化实例会给每个实例分配一个 `__dict__` ,然后反查其类就是在Python中自己实现了一个描述符,用来存储每个实例的属性值的,所以我像自己实现一个
原文例子

![图片描述](/gallery/2022-01-18-21-56-22.png)

我编写的例子,其实这个个就算是基本Python控制实例属性的一个基本模型了

![图片描述](/gallery/2022-01-18-22-46-41.png)

{% endnote %}


## 12 Copying

有时候,尤其是处理可变类型对象的情况下,你可能会像复制这个对象,以至于对副本修改不会影响到原来的对象,这也是Python `copy` 的用武之处,然而(幸运的是),(Python 模块没有感知能力)(Python modules are not sentient,)所以我们不必担心基于 Linux 的(机器人起义)(robot uprising),但是我们确实要告诉Python如何高效的复制东西:

`__copy__(self)` 定义为你类的实例使用`copy.copy()`的行为,`copy.copy()`(返回你对象的浅副本)(returns a shallow copy of your object),这就意味着,实例本身是一个新实例，但它的所有数据都被引用,举个例子,对象自身被复制了,但是其数据仍然被引用[它这里的意思是数据指针还是没变],因此,改变浅层复制的副本可能导致原始数据的改变

`__deepcopy__(self, memodict=) :`定义为你类的实例使用`copy.deepcopy()`的行为,`copy.deepcopy()` 返回你对象的深层副本,其对象本身和它自生的数据都会被复制, `memodict` 是为之前复制的对象做一个缓存--当复制递归数据结构的时候,其优化复制过程,并且预防无限递归,当你想要为你的个人属性进行深层复制的时候,调用`copy.deepcopy()`,以`memodict`作为第一个参数,


{% note warning %}
这里是一个浅层复制的例子:
![图片描述](/gallery/2022-01-18-23-00-03.png)
所以相对应的,`deepcopy`要做的就是把里面的列表也全部重新复制一份
{% endnote %}

使用这些魔术方法的情况是那些?通常来说,在你需要比默认行为更细粒度的控制的情况下. For instance, if you are attempting to copy an object that stores a cache as a dictionary (which might be large), it might not make sense to copy the cache as well – if the cache can be shared in memory between instances, then it should be


{% note warning %}
![图片描述](/gallery/2022-01-18-23-29-17.png)
{% endnote %}



## 13 Pickling


如果你和一些其他Python 爱好者相处过一些事件,您可能至少听说过pickling,picking是Python数据结构序列化的过程,当你需要存储一个对象并且之后要对其进行搜检的时候尤其有用,(当然这也是担忧和困扰的主要来源)(It’s also a major source of worries and confusion)
[因为酸洗不好听也难以理解,我结合上文的意思我把它翻译为序列化算了]
就算是没有其自己的模组(pickle)其也是相当重要的,但是因为其定义了自己的协议,所以对应的,我们定义的魔法方法也要遵循它们,但首先,我想简要的说明如何对存在的类型进行(序列化)(pickle)[如果你已经了解过这个部分了那么你可以跳过这个部分]

### 13.1 Pickling: A Quick Soak in the Brine

现在让我们介绍一下pickling,假设你有一个字典,你想要存储一个值,并且等一下就取出.你可以把其内容写进文件,确保你认真的书写了正确的语法,然后通过执行`exec()`或者对文件输入进行加工,但这充其量是不稳定的: 如果你在纯文本中存贮了一些重要的数据,其可能会被很多种方法给破坏或者修改,导致你的程序崩溃,再者,更糟糕的情况会在你的电脑上运行一些恶意代码,作为代替的解决方案,我们可以 (pickle)(腌制) 它.

![图片描述](/gallery/2022-01-19-00-06-12.png)

现在,几个小时过去了,(我们现在想取回我们的数据)(作者现在是在幽默 pickle 这个词),现在我们只需要对其解释(unpickle)

![图片描述](/gallery/2022-01-19-00-07-43.png)

发生了什么? 就像你所期待的一样,(我们成功的取出了我们之前保存的数据)(It’s just like we had data all along.)

现在,这里还是得提个醒,pickling 不是完美的,Pickle 文件可以还是会有意的或无意的被损坏,相对于纯文本来说,pickling 可能是个相对安全的,但是它依然可以被利用来运行恶意代码,同时其跨版本兼容性也不好(这里说的应该是 pickle 这个库),所以不要对分配腌制对象和人们能够打开它们保有太大的希望(作者的意思应该是:不要把安全和稳定寄托与文件和用户),但是，它也可以是(缓存)(caching)和其他常见序列化任务的强大工具

### 13.2 Pickling your own Objects

Pickling 不是内置的数据类型,它适用于任何遵循 pickle 协议的类。pickle 协议为 Python 对象提供了四种可选方法来自定义它们的行为方式（C 扩展有点不同，但这不在我们的范围内）

`__getinitargs__(self)` : 如果你想在你的类被 unpickle 时调用 `__init__` ，你可以定义 `__getinitargs__` ，它应该返回一个你想传递给 `__init__` 的参数元组。<font color="#FF0000">**请注意，此方法仅适用于旧式类。**</font>

`__getnewargs__(self)` 对于新式类，您可以影响在 unpickling 时传递给 `__new__` 的参数。此方法还应该返回一个参数元组，然后将其传递给 `__new__。`

`__getstate__(self)` 当对象被腌制时，您可以返回要存储的自定义状态，而不是存储对象的 `__dict__` 属性。 `__setstate__` 在对象未腌制时将使用该状态

`__setstate__(self, state)` 当对象被 unpickle 时，如果定义了 `__setstate__` ，对象的状态将被传递给它，而不是直接应用于对象的 `__dict__` 。这与 `__getstate__` 齐头并进：当两者都被定义时，您可以随心所欲地表示对象的腌制状态。

上面两者结合具体实例会更加的好理解一点...

`__reduce__(self) `在定义扩展类型（即使用 Python 的 C API 实现的类型）时，如果你想让它们腌制它们，你必须告诉 Python 如何腌制它们。` __- reduce__()` 在定义它的对象被腌制时被调用。它可以返回一个表示 Python 将查找和腌制的全局名称的字符串，也可以返回一个元组。
元组包含 2 到 5 个元素：调用以重新创建对象的可调用对象、该可调用对象的参数元组、要传递给 `__setstate__` 的状态（可选）、产生要腌制的列表项的迭代器（可选） , 和一个迭代器产生要腌制的字典项目（可选）。

`__reduce_ex__(self)` `__reduce_ex__` 的存在是为了兼容。如果已定义，则 `__reduce__ex__` 将在酸洗时通过 `__reduce__` 调用。也可以为不支持 `__reduce_ex__` 的旧版本酸洗 API 定义 `__reduce__` 。

### 13.3 An Example

我们的例子是 `Slate` (石板?) ,其会记得何值,以及何时,被写入的,但是，每次腌制时，这个特定的石板都会变成空白：不会保存当前值。

{% note warning %}

{% note success %}
[参考](https://www.cnblogs.com/zhouyixian/p/11129347.html#_label5)
上面那个博主写的还是不错的,可以去参考一下,因为考虑到
pickling 十分重要的话,我打算单独写一个文章来做一个实现
{% endnote %}

文章的示例代码

![图片描述](/gallery/2022-01-19-00-40-54.png)

{% endnote %}

## 14 Conclusion

本指南的目标是为任何阅读它的人带来一些东西，无论他们是否有 Python 或面向对象编程的经验。如果您刚刚开始使用 Python，那么您已经获得了编写功能丰富、优雅且易于使用的类的基础知识。如果你是一名中级 Python 程序员，您可能已经掌握了一些巧妙的新概念和策略以及一些减少您和客户编写的代码量的好方法。如果你是 Python 专家，你可能已经对一些你可能已经忘记的东西有了新的认识，并且可能在此过程中学到了一些新技巧。无论您的经验水平如何，我希望这次 Python 特殊方法的旅程真的很神奇（我无法抗拒最后的双关语）。

## 15 Appendix 1: How to Call Magic Methods

Python中的一些魔术方法直接映射到内置函数；在这种情况下，如何调用它们是相当明显的。但是，在其他情况下，调用远不那么明显。本附录致力于公开导致调用魔术方法的非显而易见的语法。

![图片描述](/gallery/2022-01-19-00-44-19.png)

## 16 Appendix 2: Changes in Python 3

我们罗列了一些就其对象模型而言,Python 3 和 Python 2.x 的改变

![图片描述](/gallery/2022-01-19-00-45-41.png)
