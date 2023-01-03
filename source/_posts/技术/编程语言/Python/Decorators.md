---
title: Decorators
index_img: /gallery/2022-01-14-13-01-48.png
date: 2022-01-16 12:51:41
updated: 2022-01-16 12:51:41
tags:
  - Python
categories:
  - Python
---

# Decorators

装饰器在Python中是一个非常强大和有用的工具，因为它允许程序员修改函数或类的行为。

装饰器允许我们包装另一个函数，以扩展包装函数的行为，而无需永久修改它。但是，在深入研究装饰器之前，让我们了解一些在学习装饰器时会派上用场的概念。


# 基础知识

>In Python, functions are first class objects that mean that functions in Python can be used or passed as arguments

结合之前的元类来看,可以知道 `function`对象其实就是 `type` 类的一个对象,所以在`Python`中可以做参数传递

1. 函数是个对象
2. 函数可以作为参数传入
3. 函数可以作为返回值

## 修饰器语法

在 Decorators 中，函数作为参数被引入到另一个函数中，然后在包装函数内部调用。

```python
@gfg_decorator
def hello_decorator():
    print("Gfg")

'''Above code is equivalent to -

def hello_decorator():
    print("Gfg")
    
hello_decorator = gfg_decorator(hello_decorator)'''
```
1. `gfg_decorator` 是可调用的函数,并且可以添加一些代码
2. `hello_decorator`会返回被包装后的函数

## 注意执行流程

![图片描述](/gallery/2022-01-16-13-06-29.png)

![图片描述](/gallery/2022-01-16-13-06-44.png)

这里有几个值得注意的点.

![图片描述](/gallery/2022-01-16-13-09-39.png)

在继续第三步的时候有个第二步

![图片描述](/gallery/2022-01-16-13-10-09.png)

这里的第二步可以看见有`func`,那么可以说在`inner1`
环境下是有 对应`hello_decorator`的变量环境的.

所以在 ![图片描述](/gallery/2022-01-16-13-11-35.png) 返回`inner1`指针,然后在![图片描述](/gallery/2022-01-16-13-12-17.png) 调用回来到 ![图片描述](/gallery/2022-01-16-13-12-36.png) 其实是这个时候 `inner1`是有 其外层函数的变量环境的..

所以说在 ![图片描述](/gallery/2022-01-16-13-13-16.png) 可以定位到 `func` 的指针,就是传入进来的![图片描述](/gallery/2022-01-16-13-13-54.png)
## 传递参数

```python
def hello_decorator(func):
    def inner1(*args, **kwargs):
         
        print("before Execution")
         
        # getting the returned value
        returned_value = func(*args, **kwargs)
        print("after Execution")
         
        # returning the value to the original frame
        return returned_value
         
    return inner1
 
 
# adding decorator to the function
@hello_decorator
def sum_two_numbers(a, b):
    print("Inside the function")
    return a + b
 
a, b = 1, 2
 
# getting the value through return of the function
print("Sum =", sum_two_numbers(a, b))
# sum_two_numbers(a, b) -> hello_decorator(sum_two_numbers)(a, b) -> inner1(a, b) -> func(*args, **kwargs)
                                                                                  #->  sum_two_number(a, b)
```

# `functools.wraps()`

functools 是用于高阶函数（作用于或返回其他函数的函数）的标准 `Python` 模块。 `wraps()` 是一个装饰器，应用于装饰器的包装函数。它通过复制 `__name__、__doc__`（文档字符串）等属性来更新包装函数，使其看起来像包装函数。

## `Example 1: Without functools.wraps()`

```python
def a_decorator(func):
    def wrapper(*args, **kwargs):
        """A wrapper function"""
        # Extend some capabilities of func
        func()
    return wrapper
 
@a_decorator
def first_function():
    """This is docstring for first function"""
    print("first function")
 
@a_decorator
def second_function(a):
    """This is docstring for second function"""
    print("second function")
 
print(first_function.__name__)
print(first_function.__doc__)
print(second_function.__name__)
print(second_function.__doc__)
---
Output:
wrapper
A wrapper function
wrapper
A wrapper function
```

可以看见,被包装器修饰后的函数 文档,名字,都变成包装器函数了
现在，如果我们编写`help（first_function）`和`help（second_function）`会发生什么

```python
First Function
Help on function wrapper in module __main__:

wrapper(*args, **kwargs)
    A wrapper function


Second Function
Help on function wrapper in module __main__:

wrapper(*args, **kwargs)
    A wrapper function

```
虽然上面的代码在逻辑上可以正常工作，但是如果您正在编写API或库，并且有人想知道您的函数的作用以及它的名称或只是键入`help（yourFunction）`，请考虑这一点，它将 <font color="#FF0000">**始终显示包装函数的名称和文档字符串**</font> 。如果您 <font color="#FF0000">**为不同的函数使用相同的包装器函数**</font> ，这将变得更加混乱，因为它将为 <font color="#FF0000">**每个函数显示相同的详细信息**</font> 。
理想情况下，它应该显示包装函数的名称和文档字符串，而不是包装函数。手动解决方案是在返回包装函数之前在包装函数中分配`__name__，__doc__`属性。


![图片描述](/gallery/2022-01-16-13-33-35.png)

基本的函数名字和函数文档的问题可以解决了但是,函数的签名(就是函数接受参数的描述) ![图片描述](/gallery/2022-01-16-13-36-04.png)

还是和包装器函数是一样的,所以为了解决这些问题并且保证可读性和时间,我们可以使用 `functools.wraps() as decorator to wrapper function.`

![图片描述](/gallery/2022-01-16-13-37-53.png)

可以看到问题解决了

![图片描述](/gallery/2022-01-16-13-38-13.png)

# 链式修饰符

{% note success %}
[可视化](https://pythontutor.com/visualize.html#code=def%20decor1%28func%29%3A%0A%20%20%20%20def%20inner%28%29%3A%0A%20%20%20%20%20%20%20%20x%20%3D%20func%28%29%0A%20%20%20%20%20%20%20%20return%20x%20*%20x%0A%20%20%20%20return%20inner%0A%20%0Adef%20decor%28func%29%3A%0A%20%20%20%20def%20inner%28%29%3A%0A%20%20%20%20%20%20%20%20x%20%3D%20func%28%29%0A%20%20%20%20%20%20%20%20return%202%20*%20x%0A%20%20%20%20return%20inner%0A%20%0A%40decor1%0A%40decor%0Adef%20num%28%29%3A%0A%20%20%20%20return%2010%0A%20%0Aprint%28num%28%29%29&cumulative=false&curInstr=24&heapPrimitives=nevernest&mode=display&origin=opt-frontend.js&py=3&rawInputLstJSON=%5B%5D&textReferences=false)
{% endnote %}

## 运行流程分析

代码:
```python
def decor1(func):
    def inner():
        x = func()
        return x * x
    return inner


def decor(func):
    def inner():
        x = func()
        return 2 * x
    return inner

@decor1
@decor
def num():
    return 10

print(num())#400
```

---

流程

```python
@decor
num()
    -> decor(num) -> decor : func = num -> return decor.inner

@decor1
decor.inner
    -> decor1(decir.inner) -> decor1: func = decor.inner -> retun decor1.inner
    
num() -> decor1.inner() = 400  # function called
# run decor1.inner() => 400
    ->x = decor.inner() = 20
        -> in decor.inner: 
            -> x = num() = 10
            -> return 10*2
    -> return 20 * 20
--- equal to ---
func = decor1(decor(num))
print(func())#400
```

# 类修饰符

这里可以参考一下 
[PEP 3129 -- Class Decorators](https://www.python.org/dev/peps/pep-3129/#abstract)

结合之前的理解,本质上类就是对象,所以可以作为函数参数传入,同时也可以做链式

```python
class A:
  pass
A = foo(bar(A))


@foo
@bar
class A:
  pass
```

# 实战

1. 修饰符基础

![图片描述](/gallery/2022-01-16-13-25-21.png)

2. `function.tools`的使用

![图片描述](/gallery/2022-01-16-13-49-35.png)

![图片描述](/gallery/2022-01-16-13-49-09.png)

3. 链式修饰符

![图片描述](/gallery/2022-01-16-14-24-17.png)

4. 类修饰符

![图片描述](/gallery/2022-01-16-14-35-00.png)

5. 保持链式修饰后的函数文档

![图片描述](/gallery/2022-01-16-14-46-01.png)
