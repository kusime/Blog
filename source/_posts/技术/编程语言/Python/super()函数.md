---
title: super()函数
index_img: /gallery/2022-01-14-13-01-48.png
date: 2022-01-15 23:06:00
updated: 2022-01-15 23:06:00
tags:
  - Python
categories:
  - Python
---

{% note success %}
[What exactly does super() return in Python 3?](https://stackoverflow.com/questions/44725483/what-exactly-does-super-return-in-python-3)
[What does `super()` mean in `__new__`](https://stackoverflow.com/questions/9056955/what-does-super-mean-in-new)
{% endnote %}

# `What exactly does super() return in Python 3?`

问题翻译:

从 `Python3` 的文档 `super()` “返回一个代理对象，该对象将方法调用委托给类型的父类或同级类。”这意味着什么？假设我们现在有下面的代码

```python
class SuperClass():
    def __init__(self):
        print("__init__ from SuperClass.")
        print("self object id from SuperClass: " + str(id(self)))

class SubClass(SuperClass):
    def __init__(self):
        print("__init__ from SubClass.")
        print("self object id from SubClass: " + str(id(self)))
        super().__init__()


sc = SubClass()
```

我的到了如下输出

```python
__init__ from SubClass.
self object id from SubClass: 140690611849200
__init__ from SuperClass.
self object id from SuperClass: 140690611849200
```

这意味着在 `super().__init__()` 行中，`super()` 返回当前对象，然后隐式传递给超类的 `__init__()` 方法。这是准确的还是我在这里遗漏了什么？

简单地说，我想了解以下内容：

当 super().**init**() 运行时，

1. 究竟是什么传递给 `__init__()` 以及如何传递？我们在 `super()` 上调用它，所以根据我目前对 Python 的理解，无论返回什么都应该传递给 `__init__()` 方法。
2. 为什么我们不必将 `self` 传递给 `super().__init__()？`

## 回答

> `returns a proxy object that delegates method calls to a parent or sibling class of type.`

上面的 `proxy` 会充当父级类的方法调用部分,它不是类本身；<font color="#FF0000">**相反，它只是足够的信息**</font> ，以便您可以使用它来调用父类方法

如果您调用 `__init__()，`您将获得自己的本地子类 `__init__ `函数。当你调用 `super() `时，你会得到那个 <font color="#FF0000">**代理对象**</font> ，它将把你 <font color="#FF0000">**重定向**</font> 到父类的方法。

类似地，如果您要调用 `super().foo`，您将从父类获得 `foo` 方法——同样，由该代理重定向。

### 对一些评论的回应

> `But that must mean that this proxy object is being passed to init() when running super().init() right?`

错误的。代理对象就像一个包名，比如调用`math.sqrt()`。您没有将 `math` 传递给 `sqrt` ，而是使用它来表示您正在使用`哪个 sqrt`。如果你想将代理传递给 `init` ，调用将是 `init(super())`。当然，这个调用在语义上是荒谬的

> `When we have to actually pass in self which is the sc object in my example.`

> 上文意思是:当我们必须实际传递 `self` 时，这就是我的示例中的 `sc` 对象。

不，您 <font color="#FF0000">**没有**</font> 通过 `sc` ；这是对象创建调用（内部方法 <font color="#FF0000">****new****</font> ）的结果，其中包括调用 init

`self` 对象是 `Python` 运行时系统为您创建的新`item`。
对于大多数类方法，第一个参数（不按惯例称为 `self` ，在其他语言中称为 `this` ）是调用该方法的对象。

# What does `super()` mean in `__new__`

注意: 涉及到一些 `Python` 实现 蝇量级

```python
import weakref # 12.15没学习 标记

class CarModel:
    _models = weakref.WeakValueDictionary()

    def __new__(cls, model_name, *args, **kwargs):
        model = cls._models.get(model_name)
        if not model:
            model = super().__new__(cls)
            cls._models[model_name] = model
        return model

    def __init__(self, model_name, air=False):
        if not hasattr(self, "initted"):
          self.model_name = model_name
          self.air = air
          self.initted=True
```

> `Question 1> what does super() mean? Does it mean the parent class of CarModel?`

> `model = super().__new__(cls)` 是什么意思?

## 回答

在开始了解`super`函数之前你要知道`super()` 是 `super(A,B)` 的缩写,其中 A 是这段代码发生的类名,然后 B 是运行`super` 的第一个参数,所以结合你的代码片段来看的话就是可以把 `super()` 扩展为
`super(CarModel,cls)`

然后依次来看,`super(T,O)` 返回一个 超级对象,为了理解什么是超级的对象,你需要理解在`Python`中.属性在实例和类中是如何工作的.

假设没有 `__getattr__ ` 和 `__getattribute__ ` 方法的调用,那么在 O 对象中引用 A 对象就是 `O.A` 或者`getattr(O, "A")` ,这个会有以下步骤组成.

1. 如果 A 在 `O.__dict__` 中,那么就返回其对应的值,正是这样(`precisely as it is.`)
2. 否则,在 O 中找对应的方法中查找其字典,`If found, call the value O.`
3. 反过来，如果 O 未定义`__get__`，则按原样返回。但是，如果定义了，则 O 称为'描述符'，并且调用其`__get__`方法，其中 O 作为第一个参数，键入 type（O） 作为第二个参数。

类上的属性引用的工作方式大致相同，用实例引用来代理类，但有以下区别：

1. 步骤 1 不适用。
2. 调用`__get__`方法时，None 作为第一个参数，引用类作为第二个参数。

那么，使用`super（T，O）`创建的超级对象是一个（内置）对象，其中包含一个`__getattribute__`方法，该方法在其上的每个属性引用上调用，并在 O 的 MRO(维护表)中 T 之后的唯一类的字典中查找属性。然后，它找到的值，它会像往常一样调用**get**。

该过程有点复杂，因此作为示例，以下是它如何针对您的特定情况进行操作。由于 `CarModel` 的定义是按原样定义的，因此其 `MRO` 是 [CarModel，object]。

对 `super().__new__(cls)` 的解释

1. `super()` 扩展为`super(CarModel,cls)`
2. `super` 会产生一个超级对象 `S`
3. Python 会在超级对象中查找 `"__new__"` 属性,等同于 `getattr(S, "__new__")`
4. 因为 `S` 是 在 `CarModel` 类中创建的.这会考虑到在 CarModel 的 MRO 中考虑 CarModel 后面的类，并在`object`类本身的命令中找到'**new**'。它的值（一个静态方法）具有一个**get**方法，该方法使用参数 None 和 cls 进行调用。由于**new**是静态方法，因此其**get**方法只是按原样返回函数，
5. 在最后一步（即 object.**new**）中获得的函数是使用 cls 参数调用的，其中 cls 可能是 CarModel，最后是 CarModel 类的新实例

> 为了完整起见，应该提到的是，对象类上的实际**new**函数实际上不是静态方法，而是一个特殊的内置函数，它根本没有**get**方法，但是由于静态方法上的**get**方法只是返回它们所定义的函数，因此效果是相同的

{% note danger %}
您可能已经注意到了额外的参数 cls 。它没有什么特别之处：**new**总是接收定义它的类，作为第一个参数。就像你有普通方法的 self ，它们接收实例作为第一个参数，或者类方法的定义类。
{% endnote %}

# `'super' object has no attribute '__dict__'`

![图片描述](/gallery/2022-01-16-17-51-05.png)

# 总结

`super`函数就是会返回目前运行这个函数的父级代理对象,然后可以理解为是一种代理(代理地址不和真正被代理对象的地址相同),然后这个代理一般和后面具体的方法来使用
这边的只是好像设计到 MRO.之后学了会补上的

> 该对象将方法调用委托给类型的父类或同级类
> `returns a proxy object that delegates method calls to a parent or sibling class of type.`

{% note success %}
下面是一些简单的测试代码..
![图片描述](/gallery/2022-01-16-12-20-41.png)

```vim

    P
  /   \
C1     C2

C1 C2 is same Level Call

```

但是不管怎么样,都会使用`super`后,其代理的对象都是相对于输入到`super(A,B)` A 的上一级类(夫类)

![图片描述](/gallery/2022-01-16-12-43-35.png)
{% endnote %}

```
super() is used to reference the superclass (i.e. the parent class from which you are subclassing).

__new__ is a method that is invoked to create a new instance of the class, if it is defined.
```
