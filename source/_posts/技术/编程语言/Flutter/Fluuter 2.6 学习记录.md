---
title: Flutter 2023-1-6 学习记录
index_img: /img/Flutter Logo.png
date: 2023-01-06 07:59:07
tags:
  - Flutter
categories:
  - Flutter
---

# Flutter 2023-1-6 学习记录

## 统一主题色

{% note warning %}
[accentColor is deprecated](https://docs.flutter.dev/release/breaking-changes/theme-data-accent-properties#migration-guide)
{% endnote %}

```dart
    final ThemeData theme = ThemeData();
    return MaterialApp(
      title: "Expense Tracer", // control the app task title
      theme: theme.copyWith(
        colorScheme: theme.colorScheme.copyWith(
          primary: Colors.amber,
          secondary: Colors.pink,
        ),
      ),
```

调用

```dart
                      decoration: BoxDecoration(
                        border: Border.all(
                          width: 2,
                          color: Theme.of(context).colorScheme.primary,
                        ),
                      ),
```


## 导入图片

{% note success %}
[Adding assets and images](https://docs.flutter.dev/development/ui/assets-and-images)
{% endnote %}

定义要导入的图片,在 `pubspec.yaml`

```yml
# The following section is specific to Flutter packages.
flutter:

  # The following line ensures that the Material Icons font is
  # included with your application, so that you can use the icons in
  # the material Icons class.
  uses-material-design: true

  # To add assets to your application, add an assets section, like this:
  assets:
    - assets/images/ # include all images
```

- 使用图片,然后Flutter 会自动帮我们调整图片大小,太爽了.

```dart
Image(
image: AssetImage('assets/images/banner.jpg'),
),
```

#  日期格式化

[Format Dates in Flutter - GeeksforGeeks](https://www.geeksforgeeks.org/format-dates-in-flutter/)

![日期格式化](/gallery/2023-01-06-11-04-57.png)