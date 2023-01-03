---
title: MagicMethod
index_img: /gallery/2022-01-14-13-01-48.png
date: 2022-01-16 19:11:26
updated: 2022-01-16 19:11:26
tags:
  - Python
categories:
  - Python
---

# MagicMethod

reference :

{% note success %}
![图片描述](/gallery/2022-01-16-19-12-39.png)

https://stackoverflow.com/questions/2657627/why-does-python-use-magic-methods

https://stackoverflow.com/questions/16237659/python-how-to-implement-getattr

http://www.sefidian.com/2021/06/06/python-__getattr__-and-__getattribute__-magic-methods/

https://stackoverflow.com/questions/12660871/python-get-method

https://stackoverflow.com/questions/12660871/python-get-method

https://stackoverflow.com/questions/59328914/how-does-getattribute-fetch-a-value

https://www.geeksforgeeks.org/__call__-in-python/

https://www.geeksforgeeks.org/dunder-magic-methods-python/

https://pythongeeks.org/methods-in-python/
{% endnote %}

## `Why does Python use 'magic methods'?`

### 问题

I've been playing around with Python recently, and one thing I'm finding a bit odd is the extensive(广泛) use of 'magic methods', e.g. to make its length available, an object implements(施行) a method, def **len**(self), and then it is called when you write len(obj).

I was just wondering why objects don't simply define a len(self) method and have it called directly as a member of the object, e.g. obj.len()? I'm sure there must be good reasons for Python doing it the way it does, but as a newbie I haven't worked out what they are yet.

## 回答-1 (我觉得一般)

AFAIK, len is special in this respect(在这方面) and has historical roots(历史渊源。).

Here's a quote from the FAQ:

> Why does Python use methods for some functionality (e.g. list.index()) but functions for other (e.g. len(list))?

> The major reason is history. Functions were used for those operations that were generic for a group of types(一组类型的泛型) and which were intended to work even for objects(甚至可以用于对象) that didn’t have methods at all (e.g. tuples). It is also convenient to have a function that can readily be applied to an amorphous collection(无定形集合) of objects when you use the functional features of Python (map(), apply() et al).

> In fact, implementing len(), max(), min() as a built-in function is actually less code than implementing them as methods for each type.(把它们作为各种类型的方法实现) One can quibble about individual cases (可以对个别情况吹毛求疵)but it’s a part of Python, and it’s too late to make such fundamental(根本的) changes now. The functions have to remain to avoid massive code breakage.(必须保留这些功能以避免大量代码损坏)

主要原因是历史。函数用于那些对一组类型通用的操作，并且它们甚至可以用于根本没有方法的对象（例如元组）。
当您使用 Python 的功能特性（map()、apply() 等）时，拥有一个可以轻松应用于无定形对象集合的函数也很方便。

事实上，将 len()、max()、min() 实现为内置函数实际上比将它们实现为每种类型的方法所需的代码更少。
人们可以对个别案例争论不休，但它是 Python 的一部分，现在做出如此根本的改变为时已晚。必须保留这些功能以避免大量代码损坏。

> The other "magical methods" (actually called special method in the Python folklore(民俗学)) make lots of sense,(很有道理) and similar functionality exists(功能存在) in other languages. They're mostly used for code that gets called implicitly when special syntax is used.(它们是那种当特殊语法被使用后,隐式调用的那些代码)

For example:

- overloaded operators (exist in C++ and others)
- constructor/destructor
- hooks for accessing attributes
- tools for metaprogramming

## 回答 2 (我觉得很吊)

Python uses the word "magic methods", because those methods really performs magic for you program. One of the biggest advantages of using Python's magic methods is that they provide a simple way to make objects behave like built-in types. That means you can avoid ugly, counter-intuitive(反直觉的), and nonstandard ways of performing basic operators.

Consider a following example:

```python
dict1 = {1 : "ABC"}
dict2 = {2 : "EFG"}

dict1 + dict2
Traceback (most recent call last):
  File "python", line 1, in <module>
TypeError: unsupported operand type(s) for +: 'dict' and 'dict'
```

This gives an error, because the dictionary type doesn't support addition. Now, let's extend dictionary class and add "**add**" magic method:

```python
class AddableDict(dict):

    def __add__(self, otherObj):
        self.update(otherObj)
        return AddableDict(self)


dict1 = AddableDict({1 : "ABC"})
dict2 = AddableDict({2 : "EFG"})

print (dict1 + dict2)
```

Now, it gives following output.

`{1: 'ABC', 2: 'EFG'}`

Thus, by adding this method, suddenly magic has happened and the error you were getting earlier, has gone away.

I hope, it makes things clear to you. For more information, refer to:


