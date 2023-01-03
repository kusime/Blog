---
title: Solts in Python
index_img: /gallery/2022-01-14-13-01-48.png
date: 2022-01-16 17:09:37
updated: 2022-01-16 17:09:37
tags:
  - Python
categories:
  - Python
---

{% note success %}
[slots in python](https://www.geeksforgeeks.org/slots-in-python/)

[Use of __slots__](https://www.geeksforgeeks.org/python-use-of-__slots__/)
{% endnote %}

# `Solts in Python`


`Python` 中的插槽是一种特殊机制(`mechanism `)，用于减少对象的内存。在 `Python` 中，所有对象都使用动态字典来添加属性。 `Slots` 是一种静态类型方法，在此分配属性时不需要动态字典

```python
# defining the class.
class gfg:
      
    # defining the slots.
    __slots__ =('course', 'price')
      
    def __init__(self):
          
        # initializing the values
        self.course ='DSA Self Paced'
        self.price = 3999
  
# create an object of gfg class
a = gfg()
  
# print the slot
print(a.__slots__)
  
# print the slot variable
print(a.course, a.price)
Output

('course', 'price')
DSA Self Paced 3999
```

## 理解

如果定义了`__slots__`之后,那么这个类就不能使用动态字典了.这也就意味着,不能够动态的添加属性了

![使用](/gallery/2022-01-16-17-30-02.png)

![不使用](/gallery/2022-01-16-17-37-25.png)

所以说是,相当于把初始一个动态字典的内存变为初始一个元组的内存,所以说就是节省了内存.


# `Use of __slots__`

当我们为类创建对象时，它需要内存，并且属性以字典的形式存储。如果我们需要分配数千个对象，则会占用大量内存空间。
插槽提供了一种特殊的机制来减小对象的大小。它是对对象进行内存优化的一个概念

```python
class GFG(object):
      __slots__=['a', 'b']
      def __init__(self, *args, **kwargs):
                self.a = 1
                self.b = 2
  
if __name__ == "__main__":
     instance = GFG()
     print(instance.__slots__)
Output :

['a', 'b']
```

```python
class GFG(object):
      __slots__=['a', 'b']
      def __init__(self, *args, **kwargs):
                self.a = 1
                self.b = 2
  
if __name__ == "__main__":
     instance = GFG()
     print(instance.__dict__)
Output :

AttributeError: 'GFG' object has no attribute '__dict__'
```

# `Why Uses it ?`

`Result of using __slots__:`

1. `Fast access to attributes`
2. `Saves memory space`