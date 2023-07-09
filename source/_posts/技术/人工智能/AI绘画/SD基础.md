---
title: Stable Diffusion 基础
date: 2023-07-09 14:25:50
tags:
  - AI
categories:
  - AI
---

#  Stable Diffu 基础

## 什么是大模型

[stable diffusion 常用大模型解释和推荐（持续更新ing）](https://zhuanlan.zhihu.com/p/631941039)


![知乎上面说的](/gallery/2023-07-09-15-08-06.png)


![我们的界面](/gallery/2023-07-09-15-08-57.png)


也就是说选择的这个模型就是决定了图片的基地风格,是二次元风,还是写实风.


![比如说这种网站,可以就下载绘画AI模型了](/gallery/2023-07-09-15-12-14.png)

## 什么是VAE

[Stable Diffusion高级教程 - VAE](https://www.dongwm.com/post/stable-diffusion-vae/)

![就是微调滤镜](/gallery/2023-07-09-15-23-50.png)

也就是说这个VAE就是在大模型输出图片之后,然后进行对最终图片的处理,有点类似美化这张图片的效果,但是不会对风格产生很大的影响

![我们的界面](/gallery/2023-07-09-15-25-46.png)


# 出图参数调整

![CFG 和 迭代](/gallery/2023-07-09-15-27-01.png)


![描述](/gallery/2023-07-09-15-28-09.png)


# 提示词

顾名思义,就是用提示词去提供信息给AI模型,我们需要什么样的效果,不希望什么样的效果

![图片描述](/gallery/2023-07-09-15-30-19.png)


# 采样器


![根据这个按照类别来测试](/gallery/2023-07-09-15-38-49.png)

![我们的界面](/gallery/2023-07-09-15-38-27.png)


![抄作业](/gallery/2023-07-09-15-53-57.png)

# 附加功能介绍


[探索【Stable-Diffusion WEBUI】的附加功能：图片缩放&amp;抠图_若苗瞬的博客-CSDN博客](https://blog.csdn.net/ddrfan/article/details/130341905)

![我们的界面](/gallery/2023-07-09-16-07-39.png)


这个东西就是用各种算法进行图片放大缩小但是又保证图片质量的应该东西。


# 模型合并

![我们的界面](/gallery/2023-07-09-16-09-41.png)

看样子就是把两个模型混合起来以达到风格融合的这种感觉，这个是针对SD大模型进行的操作


# Embedding

[在Stable Diffusion中如何使用embeddings](https://stable-diffusion.org.cn/t/topic/31)

[一分钟学会：手把手教会AI画手](https://zhuanlan.zhihu.com/p/622013773)

[Stable Diffusion用embedding模型避免画崩](https://zhuanlan.zhihu.com/p/627500143)

![图片描述](/gallery/2023-07-09-16-15-22.png)

本身也是应该模型，其作用就是提示，引导AI达到我们想要的效果，比如说保证不要绘画坏手这种，加载方式就是在提示词那里写入模型的名字

![知乎解释](/gallery/2023-07-09-16-17-59.png)


## 下载的embedding 触发词

verybadimagenegative_v1.3
badhandv4
EasyNegative
ng_deepnegative_v1_75t
rev2-badprompt `rev2-badprompt,lowres, missing finger, extra digits, fewer digits`


# HyperNetwork

[保姆级的Stable-Diffusion Hypernetwork训练教程](https://www.bilibili.com/read/cv19174085/)

我们主打一个实操，我看了一下，好像是有点类似VAE的东西，也就是基本风格不变，但是会改变一写出图效果

[【AI绘画】矢车菊V5模型发布](https://www.bilibili.com/read/cv20400815)

```
(art by yaguru magiku, (001glitch-core, dreamcore), 8sconception, alberto-mielgo, anime-background-style-v2, Super Quality, Vaporwave,Synth Wave,weirdcore, dreamy, full body
((#在这里放入你想要的内容#))
Visual effects:movie lights, ,((masterpiece)),best quality,high resolution illustrations,particle effect,((very detailed CG)),((masterpiece))absurd,intricate details,((8k_wallpaper)),
[[red, blue, green, yellow, black, white, pink, purple, cowboy shot, ((((colorful)))), ink and wash painting, [white hair], delicate and beautiful girl, ((illustration)), ((floating hair)), ((chromatic aberration))]]

Negative prompt: {Multiple people},lowres,bad anatomy,bad hands, text, error, missing fingers,extra digit, fewer digits, cropped, worstquality, low quality, normal quality,jpegartifacts,signature, watermark, username,blurry,bad feet,cropped,poorly drawn hands,poorly drawn face,mutation,deformed,worst quality,low quality,normal quality,jpeg artifacts,signature,watermark,extra fingers,fewer digits,extra limbs,extra arms,extra legs,malformed limbs,fused fingers,too many fingers,long neck,mutated hands,polar lowres,bad body,bad proportions,gross proportions,text,error,missing fingers,missing arms,missing legs,extra digit, ((Intricate clothes, too much lace, too much dacoration on clothes)) 

作者：Toooajk https://www.bilibili.com/read/cv20400815 出处：bilibili
```

# LoRA

[什么是LoRA模型，如何使用和训练LoRA模型？你想要的都在这！](https://zhuanlan.zhihu.com/p/624230991)


看了一下基本就是说，能够在不修改SD的情况下，通过少量的数据训练出来的插件模型，配合SD模型ckpt，使用，能够达到风格修改的一个程度

![使用模型](/gallery/2023-07-09-16-57-10.png)


