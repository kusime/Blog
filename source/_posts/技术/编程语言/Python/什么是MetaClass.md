---
title: 什么是MetaClasss
index_img: /gallery/2022-01-14-13-01-48.png
date: 2022-01-14 13:02:32
updated: 2022-01-14 13:02:32
tags:
  - Python
categories:
  - Python
---

# 什么是MetaClasss

{% note success %}

[原文链接](https://stackoverflow.com/questions/100003/what-are-metaclasses-in-python)

{% endnote %}

## 一切皆类

在了解元类之前你需要掌握Python中的类,而且Python类的组织有十分特别的点子,从 Smalltalk 语言借鉴来的

在很多语言中,类只是一段描述如何产生对象的代码,打个比方

```python
>>> class ObjectCreator(object):
...       pass
...

>>> my_object = ObjectCreator()
>>> print(my_object)
<__main__.ObjectCreator object at 0x8974f2c>
```

但是类在Python中,类其实也是对象,是的,对象.

只要你使用了关键词 `class`,`Python`执行它然后创建一个对象.


下面的代码会在内存中创建一个名字为 `ObjectCreator` 对象
```python
>>> class ObjectCreator(object):
...       pass
...
```

>`This object (the class) is itself capable of creating objects (the instances), and this is why it's a class.
`

因为(类)是对象,所以,你可以:
- 你可以把它赋值到变量中
- 你可以复制它
- 你可以给他添加属性
- 你可以作为函数参数传递

下面是例子

```python
>>> print(ObjectCreator) # you can print a class because it's an object
<class '__main__.ObjectCreator'>
>>> def echo(o):
...       print(o)
...
>>> echo(ObjectCreator) # you can pass a class as a parameter
<class '__main__.ObjectCreator'>
>>> print(hasattr(ObjectCreator, 'new_attribute'))
False
>>> ObjectCreator.new_attribute = 'foo' # you can add attributes to a class
>>> print(hasattr(ObjectCreator, 'new_attribute'))
True
>>> print(ObjectCreator.new_attribute)
foo
>>> ObjectCreatorMirror = ObjectCreator # you can assign a class to a variable
>>> print(ObjectCreatorMirror.new_attribute)
foo
>>> print(ObjectCreatorMirror())
<__main__.ObjectCreator object at 0x8997b4c>
```

## 动态的创建类

由于类是对象，因此您可以像任何对象一样动态创建它们。
首先,你能在函数中使用 `class`

```python
>>> def choose_class(name):
...     if name == 'foo':
...         class Foo(object):
...             pass
...         return Foo # return the class, not an instance
...     else:
...         class Bar(object):#class code
...             pass#class code
...         return Bar
...
>>> MyClass = choose_class('foo')
>>> print(MyClass) # the function returns a class, not an instance
<class '__main__.Foo'>
>>> print(MyClass()) # you can create an object from this class
<__main__.Foo object at 0x89c6d4c>
```

但是这也不是那么的动态,因为你还是需要自己写整个类的代码
因为类是对象,他们必须会生成些东西.
当年使用`class`关键词的时候,`Python` 会自动的创建一个对象,
但与Python中的大多数东西一样，它为您提供了一种手动执行此操作的方法。

还记得`type`函数吗?,这是一个好的老函数,可以让你知道对象是什么类型

```python
>>> print(type(1))
<type 'int'>
>>> print(type("1"))
<type 'str'>
>>> print(type(ObjectCreator))
<type 'type'>
>>> print(type(ObjectCreator()))
<class '__main__.ObjectCreator'>
```

好嘛,`type`有着完全不同的能力,它可以动态的创建类,`type`能把
类的秒说作为其参数传入.然后返回一个类.

`type` works this way:

```python
type(name, bases, attrs)

name: name of the class
bases: tuple of the parent class (for inheritance, can be empty) # empty like object
attrs: dictionary containing attributes names and values

```

所以一下的代码

```python
>>> class MyShinyClass(object):
...       pass
```

可以通过下面的代码手工创建

```python
>>> MyShinyClass = type('MyShinyClass', (), {}) # returns a class object
>>> print(MyShinyClass)
<class '__main__.MyShinyClass'>
>>> print(MyShinyClass()) # create an instance with the class
<__main__.MyShinyClass object at 0x8997cec>
```

你可以注意到,我们使用 `MyShinyClass` 作为类的名称和保存类引用的变量。它们可以不同，但没有理由使事情复杂化。

`type` 接受一个字典来定义类的属性

```python
>>> class Foo(object):
...       bar = True
```

可以被翻译为

```python
>>> Foo = type('Foo', (), {'bar':True})
```
然后一样的可以正藏使用类

```python
>>> print(Foo)
<class '__main__.Foo'>
>>> print(Foo.bar)
True
>>> f = Foo()
>>> print(f)
<__main__.Foo object at 0x8a9b84c>
>>> print(f.bar)
True
```

当然,也一样的继续继承什么的,so

```python
>>>   class FooChild(Foo):
...         pass
```

就像

```python
>>> FooChild = type('FooChild', (Foo,), {})
>>> print(FooChild)
<class '__main__.FooChild'>
>>> print(FooChild.bar) # bar is inherited from Foo
True
```

最后你会想给你的类添加一些方法.你只要定义好函数然后注册为对象的
属性就好了

```python
>>> def echo_bar(self):
...       print(self.bar)
...
>>> FooChild = type('FooChild', (Foo,), {'echo_bar': echo_bar})
>>> hasattr(Foo, 'echo_bar')
False
>>> hasattr(FooChild, 'echo_bar')
True
>>> my_foo = FooChild()
>>> my_foo.echo_bar()
True
```

你甚至还可以添加更多的方法,在你动态的创建类.为对象添加方法一样普通

```python
>>> def echo_bar_more(self):
...       print('yet another method')
...
>>> FooChild.echo_bar_more = echo_bar_more
>>> hasattr(FooChild, 'echo_bar_more')
True
```

就像我们之前讲的: 在 `python` 中,类就是对象,
然后你可以动态的创建类,这也是你用`class` 关键词,
它也是用元类来实现的


## 什么是元类?(最终)

元类是创建类的`东西` ,你定义类就是为了创建对象,是吧?
但是我们之前知道了在Python中,类也是对象.

好嘛,元类就是创造这些对象的,他们是 最类的类 ,
你可以用下面的代码来理解

```python
MyClass = MetaClass()# MyClass is object of MetaClass
my_object = MyClass()
```

你之前也知道了 `type` 让你如以下代码做些事情

```python
MyClass = type('MyClass',(),{})
```
这是因为,`type`这个函数其实就是元类
`type`就是`Python`背后用来创建类的元类

现在你可能好奇,为什么这个全部都是小写的,
而不是`Type`?

好嘛,我想这是一个与`str`一致性的问题，
`str` 创建字符串对象的类，`int`创建整数对象的类。
`type` 只是创建类对象的类。你可以通过`__class__`属性来查看属性来理解

---

一切，我的意思是一切，都是Python中的一个对象。这包括整数、字符串、函数和类。它们都是对象。所有这些都是从一个类创建的

```python
>>> age = 35
>>> age.__class__
<type 'int'>
>>> name = 'bob'
>>> name.__class__
<type 'str'>
>>> def foo(): pass
>>> foo.__class__
<type 'function'>
>>> class Bar(object): pass
>>> b = Bar()
>>> b.__class__
<class '__main__.Bar'>
```

`Now, what is the __class__ of any __class__ ?`

```python
>>> age.__class__.__class__
<type 'type'>
>>> name.__class__.__class__
<type 'type'>
>>> foo.__class__.__class__
<type 'type'>
>>> b.__class__.__class__
<type 'type'>
```

所以说,元类就是创造所有类对象的东西,如果你想的话,你也可以把其成为 `class factory`,

`type` 是`Python`默认的元类,当让你也可以创建自己的元类

## `The __metaclass__ attribute`

在 Python 2 中，您可以在编写类时添加`__metaclass__`属性（请参阅下一节 来了解 Python 3 语法）：

```python
class Foo(object):
    __metaclass__ = something...
    [...]
```

如果你如上做的话,那么`Python`就会是用你定义的元类来创建`Foo`

<font color="#FF0000">**小心，这很棘手。**</font>

你首先写了 `class Foo(object)` ,但是 类对象 `Foo`在没有在内存中创建

`Python` 将会在类的定义中查找 `__metaclass__`,如果找到了他就
会使用这个指定的类创建,如果没有,那么就会是用 `type` 来创建.

把以上多读几次,当你输入

```python
class Foo(Bar):
    pass
```


`Foo` 中是否有`__metaclass__`属性？

如果是，使用`__metaclass__`中的内容创建内存中的类对象（我说的是类对象，请留这里）。名称为 `Foo`

如果`Python`找不到`__metaclass__`，它将在`MODULE`级别寻找`__metaclass__`，并尝试做同样的事情（但仅限于不继承任何东西的类，基本上是旧式类）。

然后，如果它根本找不到任何__metaclass__，它将使用 Bar（第一个父级）自己的元类（可能是默认类型）来创建类对象。

请注意，`__metaclass__`属性不会被继承，父级（`Bar.__class__`）的元类将被继承。如果 `Bar` 使用`__metaclass__`属性创建了带有 `type（）` 的 `Bar`（而不是 `type.__new__（）`），则子类将不会继承该行为。

现在最大的问题是，你能在`__metaclass__`放什么？
答案是,能够创建类的东西.

什么可以创建一个类？
1. `type`，
2. 或任何子类或使用它的东西。
>And what can create a class? type, or anything that subclasses or uses it.

## `Metaclasses in Python 3`


设置元类的语法在 Python 3 中已更改：

```python
class Foo(object, metaclass=something):
    ...
```

> 不再使用__metaclass__属性，而是在基类列表中使用关键字参数。
然而，元类的行为在很大程度上保持不变。

在Python 3中添加到元类中的一件事是，您还可以将属性作为关键字参数传递到元类中，如下所示：

```python
class Foo(object, metaclass=something, kwarg1=value1, kwarg2=value2):
    ...
```

请阅读以下部分，了解 Python 如何处理此问题。

## `Custom metaclasses`

元类的主要目的是在创建类时自动更改类。

通常对 `API` 执行此操作，在 `API` 中，您希望创建与当前上下文匹配的类。

想象一个愚蠢的例子，你决定模块中的所有类都应该用大写字母编写它们的属性。有几种方法可以做到这一点，但一种方法是在模块级别设置`__metaclass__`。

这样，<font color="#FF0000">**此模块的所有类都将使用此元类创建**</font>，我们只需要告诉元类将所有属性转换为大写即可。

幸运的是，`__metaclass__`实际上可以是任何可调用的，它不需要是一个正式的类（我知道，它的名字中带有'class'的东西不一定是一个类......但它很有帮助）。

因此，我们将从一个简单的示例开始，通过使用一个函数。

```python
# the metaclass will automatically get passed the same argument
# that you usually pass to `type`
def upper_attr(future_class_name, future_class_parents, future_class_attrs):
    """
      Return a class object, with the list of its attribute turned
      into uppercase.
    """
    # pick up any attribute that doesn't start with '__' and uppercase it
    uppercase_attrs = {
        attr if attr.startswith("__") else attr.upper(): v
        for attr, v in future_class_attrs.items()
    }

    # let `type` do the class creation
    return type(future_class_name, future_class_parents, uppercase_attrs)

__metaclass__ = upper_attr # this will affect all classes in the module

class Foo(): # global __metaclass__ won't work with "object" though
    # but we can define __metaclass__ here instead to affect only this class
    # and this will work with "object" children
    bar = 'bip'
```

现在，让我们执行完全相同的操作，但对元类使用真实类：

```python
# remember that `type` is actually a class like `str` and `int`
# so you can inherit from it
class UpperAttrMetaclass(type):
    # __new__ is the method called before __init__
    # it's the method that creates the object and returns it
    # while __init__ just initializes the object passed as parameter
    # you rarely use __new__, except when you want to control how the object
    # is created.
    # here the created object is the class, and we want to customize it
    # so we override __new__
    # you can do some stuff in __init__ too if you wish
    # some advanced use involves overriding __call__ as well, but we won't
    # see this
    def __new__(upperattr_metaclass, future_class_name,
                future_class_parents, future_class_attrs):
        uppercase_attrs = {
            attr if attr.startswith("__") else attr.upper(): v
            for attr, v in future_class_attrs.items()
        }
        return type(future_class_name, future_class_parents, uppercase_attrs)
```

让我们重写上面的内容，但现在我们知道它们的含义，因此使用更短，更现实的变量名称：

```python
class UpperAttrMetaclass(type):
    def __new__(cls, clsname, bases, attrs):
        uppercase_attrs = {
            attr if attr.startswith("__") else attr.upper(): v
            for attr, v in attrs.items()
        }
        return type(clsname, bases, uppercase_attrs)
```

您可能已经注意到了额外的参数 `cls` 。它没有什么特别之处：`__new__`总是接收定义它的类，作为第一个参数。就像你有普通方法的 `self` ，它们接收实例作为第一个参数，或者类方法的定义类。

但这不是正确的 `OOP` 。我们在直接调用`type`
我们没有(`aren't`)覆盖或调用父级的 `__new__`,

>`But this is not proper OOP. We are calling type directly and we aren't overriding or calling the parent's __new__. Let's do that instead:`

{% note warning %}
刚刚我做了个测试代码.
![图片描述](/gallery/2022-01-15-01-00-11.png)
可以看到不管我是在 `Chiled.__new__`中使用哪种模式的方法.
最后结果都是一样的,那就是`Partent.__new__`方法会被调用.(相对的也会返回一个父级的对象(实例))
所以说,这里不是简单的重写父级,比如下面的示例

![图片描述](/gallery/2022-01-15-01-06-58.png)

可以看见这里的机制其实是比简单函数的重写要复杂一些的.

所以我认为作者是 <font color="#FF0000">**建议**</font> 一个面对象对象 <font color="#FF0000">**开发的规范**</font>

然后那个规范就是具体的规范就是:
如果子类定义了与父级函数有相同的函数名,然后我们推荐
- <font color="#FF0000">**重写这个函数(参考`ChildNormal`)**</font>
- <font color="#FF0000">**在这个函数中去调用那个和在父类定义的这个函数**</font>

{% endnote %}

{% note danger %}
本质上,综合我学了这么多来看,我认为本质上重写这个概念本质上就是
`__dict__`中对应函数指针的重新赋值,可以理解一下下面代码的证明

![图片描述](/gallery/2022-01-15-01-18-00.png)
{% endnote %}

让我们用下面的代码代替




让我们改为这样做`:`



```python
class UpperAttrMetaclass(type):
    def __new__(cls, clsname, bases, attrs):
        uppercase_attrs = {
            attr if attr.startswith("__") else attr.upper(): v
            for attr, v in attrs.items()
        }
        return type.__new__(cls, clsname, bases, uppercase_attrs)
```

我们可以通过使用 `super` 使它更加干净，这将简化继承（因为是的，你可以有元类，从元类继承，从`type`继承）`:`


```python
class UpperAttrMetaclass(type):
    def __new__(cls, clsname, bases, attrs):
        uppercase_attrs = {
            attr if attr.startswith("__") else attr.upper(): v
            for attr, v in attrs.items()
        }
        return super(UpperAttrMetaclass, cls).__new__(
            cls, clsname, bases, uppercase_attrs)
```

哦，在Python 3中，如果你用关键字参数进行这种调用，就像这样：

```python
class Foo(object, metaclass=MyMetaclass, kwarg1=value1):
    ...
```

它在元类中转换为使用它：

```python
class MyMetaclass(type):
    def __new__(cls, clsname, bases, dct, kwargs1=default):
        ...
```

就是这样。实际上，元类仅此而已。

使用元类的代码复杂性背后的原因不是因为元类，而是因为你通常使用元类来做扭曲的东西，依赖于内省，操纵继承，`vars`，如`__dict__`等。

事实上，元类对于执行黑魔法特别有用，因此是复杂的东西。但就其本身而言，它们很简单：

- 截获类创建
- 修改类
- 返回修改后的类

## `Why would you use metaclasses classes instead of functions?`

既然`__metaclass__`可以接受任何可调用的，既然它显然更复杂，为什么还要使用类呢？

这样做有几个原因：

意图是明确的。当你阅读 `UpperAttrMetaclass（type）` 时，你知道接下来会发生什么
- 您可以使用 `OOP`。元类可以从元类继承，重写父方法。元类甚至可以使用元类。
- 如果指定了元类类，则类的子类将是其元类的实例，但未指定元类函数。
- 您可以更好地构建代码。你永远不会将元类用于像上面的例子这样微不足道的事情。它通常用于复杂的事情。能够制作多个方法并将它们分组到一个类中，这对于使代码更易于阅读非常有用。
- 你可以迷上`__new__`，`__init__`和`__call__`。这将允许你做不同的事情，即使通常你可以`__new__`完成所有的事情，有些人只是更舒服地使用`__init__`。
- 这些被称为元类，该死的！它一定意味着什么！

## 为什么要使用元类？

现在是一个大问题。为什么要使用一些容易出错的晦涩难懂的功能？

好吧，通常你不会：

>Metaclasses are deeper magic that 99% of users should never worry about it. If you wonder whether you need them, you don't (the people who actually need them to know with certainty that they need them and don't need an explanation about why).

`Python Guru Tim Peters`

元类的主要用例是创建 `API`。一个典型的例子是`Django ORM`。它允许您定义如下内容：

```python
class Person(models.Model):
    name = models.CharField(max_length=30)
    age = models.IntegerField()
```

但是，如果您这样做：
```python
person = Person(name='bob', age='35')
print(person.age)
```

它不会返回 `IntegerField` 对象。它将返回一个`int`，甚至可以直接从数据库中获取它。

这是可能的，因为模型。模型定义了`__metaclass__`它使用了一些魔术，将你刚刚用简单语句定义的人员变成一个复杂的数据库字段钩子。

`Django`通过公开一个简单的`API`并使用元类，从这个`API`重新创建代码来做幕后的真正工作，使复杂的东西看起来很简单。

## The last word

首先，您知道类是可以创建实例的对象。
好吧，实际上，类本身就是实例。的元类。

```python
>>> class Foo(object): pass
>>> id(Foo)
142630324
```

在`Python`中，一切都是一个对象，它们要么是类的实例，要么是元类的实例。

`type`除外。

`type` 实际上是它自己的元类.
这不是你可以在纯`Python`中复制的东西.
而是通过在实现级别作弊来完成的.

其次，元类很复杂。您可能不想将它们用于非常简单的类更改。
您可以使用两种不同的技术更改类：

- 猴子修补
- 类装饰器
- `99%`的时间你需要改变类，你最好使用这些.
- 但是在`98%`的情况下，你根本不需要改变阶级。