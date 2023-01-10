---
title: Flutter 学习总结
index_img: /img/Flutter Logo.png
date: 2023-01-06 05:54:59
tags:
  - Flutter
categories:
  - Flutter
---

# Flutter 学习总结

这里主要是对我之前学习的 Flutter 知识进行总结。因为截止最后一次学习 已经过了一段时间了 可能有些东西我都已经忘记了。

基本结构

```dart
main() => runApp(MyApp()); # running the app
class MyApp extends StatefulWidget {
  @override
  State<MyApp> createState() => _MyAppState();
}
class _MyAppState extends State<MyApp> {

  @override
  Widget build(BuildContext context) {

    return MaterialApp(
      home: Scaffold(
        appBar: ...
        body:...,
       floatingActionButtonLocation:...,
        floatingActionButton:...,
      )
    )
  }

}
```

# MaterialApp

这个是 FLutter App 的最高的上下文,然后下面的所有组件都要在这个构造器下面

## Builder

这个可以结一次上下文包,在使用如 `showModalBottomSheet` 这个需要上下文,但是找不到的情况下还是很好用的.

## Scaffold

这个是定义 APP 的大体框架, 把一个应用分为三段

![三段](/gallery/2023-01-06-06-32-03.png)

第一段就是标题,然后一些按钮
第二段是引用本体
第三段就是放导航栏,和一些浮动按钮这些

### appBar: AppBar

在 `Scaffold` 下面的东西,导航栏这个部分

![error_loading](/gallery/2023-01-06-06-36-35.png)

```dart
          appBar: AppBar(
            title: Text("Expense Tracer"), // 配置标题
            actions: [ // 配置动作部分
              IconButton(
                onPressed: () => showModal(bCTX),
                icon: Icon(
                  Icons.zoom_in,
                ),
              )
            ],
          )
```

### 浮动按钮,在第三段

{% note success %}
Flutter 的命名还是可以的,找到对应的属性,然后看对面的累一般就是差不多的名字
{% endnote %}

```dart
          floatingActionButtonLocation:
              FloatingActionButtonLocation.centerFloat,
          floatingActionButton: FloatingActionButton(
            onPressed: () => showModal(bCTX),
            child: Icon(
              Icons.add,
            ),
          ),
```

这上面就是按钮,后面会介绍

## Container

Container 主要是拿来布局的,可以比较方便的调整位置,大小,对齐等.
而且下面的孩子会自动继承这个 Container 父级的位置

```dart
 Container(
            width: double.infinity, // 容器全宽
            child: Container(
              margin: EdgeInsets.only(bottom: 30),
              child: Column(
                mainAxisAlignment: MainAxisAlignment.start,
                children: [TransList(transactions)],
              ),
            ),
          ),
```

- `margin: EdgeInsets.only(bottom: 30),` 来自于 `EdgeInsets` .
  ![然后这个东西会提供很多可选构造器](/gallery/2023-01-06-06-41-41.png)

- `padding: EdgeInsets.all(10),` 来自于 `EdgeInsets` .

###  decoration:

就是容器的描边 相当于 border

```dart
Container(
                      margin:
                          EdgeInsets.symmetric(vertical: 10, horizontal: 15),
                        // border
                      decoration: BoxDecoration(
                          border: Border.all(
                              width: 2,
                              color: Color.fromARGB(255, 255, 6, 143))),
                        // border
                      padding: EdgeInsets.all(10),
                      child: Text(
                        '\$${tx.Amount.toString()}',
                        style: TextStyle(
                            fontWeight: FontWeight.bold,
                            fontSize: 26,
                            color: Colors.lightBlue),
                      ),
                    ),
```

## Column

这个是 flex 布局的一个老朋友了.对其方式和老朋友 CSS 里面的 flex 布局概念是一样的.主轴和 Cross 轴这些.
我个人不喜欢背记,所以直接试一下就好了.然后这边控制 mainAxis 和 crossAxis 是用了两个不同的类来控制的

![error_loading](/gallery/2023-01-06-06-44-36.png)

![crossAxisAlignment](/gallery/2023-01-06-06-44-58.png)

## Row

这个也是老 flex 朋友了.

和 Colum 一样,接收多个子集元素

# 按钮

这里主要介绍一些配置和样子,以及需要的一些回调

```dart
 OutlinedButton(
                style: OutlinedButton.styleFrom(
                    padding:
                        EdgeInsets.symmetric(vertical: 10, horizontal: 20)),
                onPressed: () =>
                    {widget.cb(ReasonController.text, AmountController.text)},
                child: Text(
                  "Submit",
                  style: TextStyle(fontSize: 15, color: Colors.purpleAccent),
                ))
```

![OutlinedButton](/gallery/2023-01-06-06-48-58.png)

## 样式配置`style:`

就是上面所述的,推荐的配置方式就是 对应的按钮类名字.styleFrom(styling)

## 回调 `onPressed :`

要求是没有返回,没有参数的回调函数.点击后会调用

## `child`

这个就是在按钮上面显示的东西.可以是 Icon 也可以是文字

# 显示

这里主要介绍一些视觉元素,非交互的东西

## Text

```dart
Text(
                  "Submit",
                  style: TextStyle(fontSize: 15, color: Colors.purpleAccent),
                )
```

第一个参数是文字,后面是字体的样式.

### TextStyle

- `fontWeight: FontWeight.bold,`
- `fontSize: 26,`
- `color: Colors.lightBlue`

## 图标

```dart
Icon(Icons.zoom_in),
```

一般和按钮配合使用

- 单独使用一个图标按钮注册

```dart
IconButton(
                onPressed: () => showModal(bCTX),
                icon: Icon(
                  Icons.zoom_in,
                ),
              )
```

- 按照普通的视觉元素注册,注册到 child 上面

```dart
FloatingActionButton(
            onPressed: () => showModal(bCTX),
            child: Icon(
              Icons.add,
            ),
          ),
```

## Card

emm,可以简单的理解为就是带了一定背景效果的容器
大概这种感觉?

```html
<div class="flex shadow-sm rounded-xl p-5"></div>
```

```dart
Card(
  child:... )

```

{% note danger %}
注意,这个之接受一个子元素,所以我们要在一个卡片里面做一些复杂的内容展现,我们要使用到 Row Colum 这些东西
{% endnote %}

# 滚动组件

这里就是对应css里面的overflow-auto 这种感觉.做一个可以滚动的视觉窗口

##  ListView.builder

这个是升级版的,会自动的提高性能,推荐使用.但是有注意事项,你可以看见用 `Container` 抱起来了.
因为这个 需要指定滚动视窗的高度.如果没有指定会报错的

```dart
 Container(
            // ListView must have the height
            height: 400,
            child: ListView.builder(
              itemCount: transactions.length,
              itemBuilder: (context, index) {
                final tx = transactions[index];
                return // 要构造的单个元素
                }
```

### itemCount:

指定列表,元素的长度

### itemBuilder 
提供构建上下文 (这个)`itemBuilder: (context, index) {` 会自动提供给我们.
{% note danger %}
注意,这个我们在这里只是定义回调,不是我们去调用这个回调,而且这个上下文是flutter管理的,而不是上面的`build` 函数提供的.
{% endnote %}



# 日期格式化

需要用到的包

```yml
dependencies:
  flutter:
    sdk: flutter


  # The following adds the Cupertino Icons font to your application.
  # Use with the CupertinoIcons class for iOS style icons.
  cupertino_icons: ^1.0.2
  intl: ^0.18.0

dev_dependencies:
  flutter_test:
    sdk: flutter

  # The "flutter_lints" package below contains a set of recommended lints to
  # encourage good coding practices. The lint set provided by the package is
  # activated in the `analysis_options.yaml` file located at the root of your
  # package. See that file for information about deactivating specific lint
  # rules and activating additional ones.
  flutter_lints: ^2.0.0
  lints: ^2.0.1
```

顺便附送一下 analysis_options的配置

```yml
  rules:
    # avoid_print: false  # Uncomment to disable the `avoid_print` rule
    # prefer_single_quotes: true  # Uncomment to enable the `prefer_single_quotes` rule
    include: package:lints/recommended.yaml
```


## 创建一个时间

```dart
CreatedTime: DateTime.now() // -> DateTime
```

## 格式化一个时间

```dart
import 'package:intl/intl.dart';
DateFormat('EEE, M/d/y').format( DateTime )
```

# final 和 const 

final，和const和var的区别。
因为在这里，构建规则的特殊性所以延伸出来了这几个概念。
先说一下渲染流程，渲染这个概念在动态组件里面最常见，vue靠虚拟Dom，然后这个靠StatefulWidget。然后这里面触发rebuild的函数是靠set State，vue靠的是diff算法。也就是这里的触发是手动的。但是rebuild就意味着会全部重新构建一个新Widget。所以这边就涉及到运行时和编译时了。

这里把State做了一个抽离，单独抽象成为一个类。成为状态，一个状态类绑定一个状态Widget，这个部分是独立于视图的，可以理解为vue中的pinia组件这种感觉。

然后渲染是一个新widget就意味着重新构建一个对象。所以这个final就是所谓的运行时常量。因为在stateless的Widget中不存在状态修改。
const就是编译时常量。就是不管渲染不渲染，都是常量，这个是const。var就是可变。
也就是这个语言一切以视图为核心出发的.




# Stateful Widget vs Stateless Widget

emmm，就这么理解吧，Stateful Widget就是对Stateless Widget多了一层封装，然后因为这里把main这个widget从stateless变成了Stateful，所以导致上下文被外面一层Stateful包裹。但是我们showBottomModalSheet要用到的就是那一层stateless的，也就是被state 包装之前的。所以这里用到builder解包，拿到去掉state那一层的上下文，就可以正常进行查找了。这个是粗浅的理解，不过够用了。

![Stateful Widget vs Stateless Widget](/gallery/2023-01-06-07-22-47.png)




# showModalBottomSheet

```dart
      showModalBottomSheet(
          // we need the bCTX which comes from the MaterialApp Builder bCTX, since the showModalBottomSheet needs that bCTX and also will make some work on that !
          context: bCTX,
          builder: (_) {
            // the builder should return widget 
            return NewTrans(((Reason, Amount) {
              addTransaction(Reason, Amount, bCTX);
            }));
          });
```


# 输入


定义输入控制器,链接控制器,定义提交回调(optional)
```dart
final ReasonController = TextEditingController();
final AmountController = TextEditingController();


TextField(
              controller: ReasonController,
              decoration: InputDecoration(
                labelText: 'Reason',
              ),
              onSubmitted: (_) =>
                  {widget.cb(ReasonController.text, AmountController.text)},
            ),
```

## controller

这个会自动的处理用户输入,然后把他存储到 text对象成员里面去
## decoration: InputDecoration

修饰输入框样式,定义placeholder等
## onSubmitted
就是用户敲击回车后触发的,不是必选的

# Widget 参数传递


## StatelessWidget

- unnamed , order will take into a place
```dart
// declare 
class TransChart extends StatelessWidget {
  final List<Transaction> transactions;
  TransChart(this.transactions);
  ...

// passing
TransChart(
  transactions,
),
```

- named

```dart
// declare
class TransChart extends StatelessWidget {
  final List<Transaction> transactions;
  TransChart({required this.named}); // use required to avoid null check error
..

// passing

TransChart(
  named: transactions,
),
```


## StatefulWidget
在上面注册类上生命要用的参数,然后在下面使用widget来使用.
![error_loading](/gallery/2023-01-06-09-04-08.png)


# 列表数据映射到组件

- 使用 ListView.builder 这种内置的转换组件
```dart
 ListView.builder(
              itemCount: transactions.length,
              itemBuilder: (context, index) {
                final tx = transactions[index];
                // Widget
                return Card(
                  child: Row(children: [ 
                    ....
```

- 使用 Map 

```dart
List.map((){
  ...
}).toList()
```