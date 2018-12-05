# genpost

genpost 是markdown文件生成器

## 安装方法

```bash
go get -u -t github.com/hidevopsio/genpost
```

## 生成文章类别

在项目文件夹下面运行 `genpost -c` 按照提示输入文章类别的目录和标题即可

```bash

genpost -c

✔ 目录: articles
✔ 标题: 分享文章
✔ 排序 : 1
```

## 生成文章

在项目文件夹下面运行 `genpost` 按照提示选择文章类型，输入文章标题和作者即可

```bash

genpost

___  / / /__(_)__  /_______________  /_
__  /_/ /__  /__  __ \  __ \  __ \  __/
_  __  / _  / _  /_/ / /_/ / /_/ / /_     Hiboot Application Framework
/_/ /_/  /_/  /_.___/\____/\____/\__/     https://hidevops.io/hiboot

Use the arrow keys to navigate: ↓ ↑ → ←

? 选择类型:
    Hiboot云原生应用框架
  ▸ 分享文章
    代码阅读

✔ 标题 : 我的文章标题█
✔ 作者 : 邓冰寒

```

## 获取帮助

```bash
genpost -h

______  ____________             _____
___  / / /__(_)__  /_______________  /_
__  /_/ /__  /__  __ \  __ \  __ \  __/
_  __  / _  / _  /_/ / /_/ / /_/ / /_     Hiboot Application Framework
/_/ /_/  /_/  /_.___/\____/\____/\__/     https://hidevops.io/hiboot

Run genpost command

Usage:
  genpost [flags]

Examples:

1. 生成类别

✔ 目录: articles
✔ 标题: 分享文章
✔ 排序 : 1

2. 生成文章
Use the arrow keys to navigate: ↓ ↑ → ←

? 选择类型:
    Hiboot云原生应用框架
  ▸ 分享文章
    代码阅读

✔ 标题 : 我的文章标题█
✔ 作者 : 邓冰寒


Flags:
  -c, --category   --category=true or -c
  -h, --help       help for genpost

```