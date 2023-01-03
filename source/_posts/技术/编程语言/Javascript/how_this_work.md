---
title: this in JavaScript
date: 2022-06-26 16:22:21
tags:
  - JavaScript
categories:
  - JavaScript
---

# "this" in JavaScript

![this in JavaScript](https://www.geeksforgeeks.org/this-in-javascript/?ref=gcse)

**_this_** 这个关键词经常然很多开始学习这个语言的人感到困惑，这个的根源是因为 这个 this 对比其他语言比如说 Java 和 Python 中的 self 是不一样的，所以理解 this 绝对是一个读写更高级的 js 的一个提高，所以这就是为什么我们要花时间去理解掌握， this 到底是个什么东西。

我们有大量篇幅介绍 this 其在函数中的引用，这就是说，首先看看函数是个什么东西能帮助我们理解

## this and Functions

函数，在 javascript 中，本质上是个对象，就像对象他们能够赋值给变量一样，或者是做为其他函数的参数传入，和函数返回。而且和对象一样，函数也有着他们自己的属性，然后其中一个属性就是 <font color="#FF0000">**this**</font>。

那个 this 存储的值就是执行这个函数的上下文。因为，所以在函数中，this 这个值会取决于那个函数是如何定义的，以及如何被调用和默认执行上下文。

NOTE:

this 的值一直会是对象，然后那个对象描述了执行当前函数的上下文，
在我们进一步研究 this 在函数中的行为，让我们看看在他们之外的一些行为

## Global Context

在函数外面写的代码就是属于 global context ，然后这个 this 的值就会等于这个全局上下文、

举个例子，如果你在浏览器打开 控制台然后输入接下来的代码 `console.log(this)` ，然后你就会看见 Windows 对象，这是因为，全局对象，在浏览器运行时，比如说 Chrome 的运行是就是 Windows 对象
然而，在函数内部，全局上下文可能不再存在，函数可能有自己定义的上下文，因此 this 的值不同

JavaScript 中的函数可以通过多种方式调用：1.函数调用 2.方法调用 3.构造函数调用

## this with function invocation

函数调用是指使用函数的名称或表达式调用函数的过程，该表达式计算为函数对象，后跟一组打开和关闭的第一个括号（包含括号表示我们正在要求 JavaScript 引擎立即执行该函数）。例如：

```js
function doSomething() {
  // do something here
  // example
  console.log(this);
}

// function invocation
doSomething();
```

doSomething 函数里面的 this，如果是通过上面的函数调用来调用的话，它的值就是全局对象，也就是浏览器环境中的 window 对象：

```js
<!DOCTYPE html>
<html>
<body>
<script>
    function doSomething(a, b) {

       // adds a propone property to the Window object
        this.propone = "test value";
    }

// function invocation
doSomething();
document.write(window.propone);
</script>
</body>
</html>

Output:

test value
```

然而，这并非总是如此。如果 doSomething() 函数在严格模式下运行，它将记录未定义而不是全局窗口对象。这是因为，在严格模式下（由行：'use strict'; 表示），this 的默认值，对于任何函数对象都设置为未定义而不是全局对象。例如 ：

```js
<!DOCTYPE html>
<html>
<body>
<script>
    function doSomething() {
        // enable the strict mode
        'use strict';

       // logs undefined
        document.write(this + '<br>')
            function innerFunction() {
              // Also logs undefined, indicating that
              // strict mode permeates to inner function scopes
                document.write(this)
            }
        innerFunction();
    }

// function invocation
doSomething();
</script>
</body>
</html>


Output:

undefined
undefined
```

## this with method invocation:

函数，当定义为对象的字段或属性时，称为方法。

```javascript

<!DOCTYPE html>
<html>
<body>
<script>
    let person = {
        name : "John",
        age : 31,
        logInfo : function() {
            document.write(this.name + " is " + this.age + " years old ");
        }
    }
       // logs John is 31 years old
       person.logInfo()
                 </script>
</body>
</html>

Output:

John is 31 years old
```

在上面的代码用例，`loginfo()`是 `person object` 对象的一个成员,然后我们用对象访问的语法来调用他.我们用了属性访问操作符,这样的调用需要使用一个表达式来计算我们的方法所属的对象，以及一个属性访问器（例如：person.logInfo()），后跟一组左括号和右括号。

理解 function invocation 和 method invocation 的不用是有必要的。

这样能够帮助我们理解，在给定的函数中的 this 的上下文是什么，因为不同的调用方式 this 是不同的。
在使用属性访问器调用的此类方法中，this 将具有调用对象的值，即 this 将指向与属性访问器一起使用以进行调用的对象。

```javascript


<!DOCTYPE html>
<html>
<body>
<script>
    let add = {
        num : 0,
        calc : function() {

            // logs the add object
            document.write(this + ' ')
                this.num
                += 1;
            return this.num;
        }
    };

// logs 1
document.write(add.calc() + '<br>');
// logs 2
document.write(add.calc());
</script>
</body>
</html>

Output:

[object Object] 1
[object Object] 2

```

在上面的例子中，calc() 是 add 对象的一个 ​​ 方法，因此使用第 9 行和第 10 行的方法调用规则来调用。我们知道，当使用方法调用模式时，this 的值设置为调用目的。在这个 calc() 方法中，this 的值被设置为调用对象，在我们的例子中是 add。
因此我们可以成功访问 add 的 num 属性。

## What happens to this in a function nested inside a method of an object?

```js
<!DOCTYPE html>
<html>
<body>
<script>
    let add = {
        num : 0,
        calc : function() {

        // logs the add object
        document.write(this + ' ')

        function innerfunc() {
            this.num += 1;

        // logs the window object
        document.write(this + ' ');

        return this.num}

    return innerfunc();
    }
};

// logs NaN
document.write(add.calc() + '<br>');

// logs NaN
document.write(add.calc());
</script>
</body>
</html>
```

Let’s try to understand what just happened.
When we call calc() in lines 14 and 15 we are using method invocation which sets this to add in calc(). This can be verified using the log statement in line 4.
However, innerfunc() is called from within the calc() method using a simple function invocation(line 11 ). This means, inside innerfunc() this is set to the global object, which does not have a num property, and hence the NaN outputs are obtained.
How do we solve this issue?How can we retain the value of this from the outer method inside the nested function?
One solution is to assign the this value from the outer function to a variable to be used in the nested function like so:

```javascript
JavaScript
<!DOCTYPE html>
<html>
<body>
<script>
    let add = {
        num : 0,
        calc : function() {

            // logs the add object
            document.write(this + ' ')

           // using thisreference variable to
           // store the value of this
           thisreference = this;

            function innerfunc()
            {

             // using the variable to access the
             // context of the outer function
                thisreference.num += 1;

               // logs the add object
                document.write(thisreference + ' ');
                return thisreference.num;
            }
            return innerfunc();
        }
    };

// logs 1
document.write(add.calc() + '<br>');

// logs 2
document.write(add.calc());
</script>
</body>
</html>
Output:

[object Object] [object Object] 1
[object Object] [object Object] 2

```
