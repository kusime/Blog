---
title: How to code
date: 2022-06-25 14:44:27
tags:
  - Web
categories:
  - Web
top: true
---

# How to code

1. set a goal , make specfic goal ...
   -- prepare for the framework ...

2. do not copy code , please really understand how code work ...

3. reinforce the code

learning a new feature and immediately writing code for traning .. and finish challenge ..

codewars..

4. particing code by own , no matter big or small picture ..

dont stay in the confort zone ...

5. do not get stuck in how to write perfect code ..

6. enbrace you dont know anything ...

7. not learn isolate... learn by share ..

8. just finishd course not means you can be a good web developer ...

# how to solving problem

1.  make 100% understand the problem ask right question ..
    exp: "we need a funciton that reverse whatever we pass into it" 1. what is "whatever" 2. what to do if someting else is passed in 3. what should be return ? 4. how to recognize whether the argument is a number 5. how to reverse a number ..?

2.  devide and conquer : Break a bit problem into a smaller sub-problem
    exp: "we need a funciton that reverse whatever we pass into it" 1. check if argument is a number or a string or an array 2. implement reversing a number 3. implement reversing a string 4. implement reversing an array 5. return reverse value ..

3.  not afarid to do as much research as you have to
    exp: use stackoverflow , google , document 1. how to check if a value is a sting 1. how to check if a value is a number 1. how to check if a value is a ....

4.  For bigger problem . write pseudo-code before writing the actual code ..
    exp :
    code`if value type !string && !number && !arrary return value if type of sring: return reverse string `

## how to solving problem ..

```js
"use strict";
// problem

// we work for a company buiding a smart home thermometer .
// our most recent task is this : giving an array of temperatures of
// one day , caculate the temperature amplitude keep in mind that sometimes
// there might be a sensor error..

const temperatures = [1, -2, -5, 1, "error", 0, 12, 33];

// 1> understanding the porblem
// - what is tem ampliture ? : Answer : difference between highest and lowest temperature
// - how to compute max and min temperature ?
// - what is the sensor error ? and What to do ?

// 2> breaking up into sub-problem
// - ignore the sensor ...  but how to ingnore
// - find max value in temp array
// - find min value in temp array
// - subtract min from max (amplitude and return it)
```

# how to debugging

1. identify

   1. during the development
   2. testing software
   3. user reports during production
   4. context browsers . users

2. find

   1. developer console
   2. debugger (complex code)

3. fix

   1. replace wrong solution with new correct solution

4. prevent
   1. searching for the same bug in similar code
   2. writing test using testing software

now
