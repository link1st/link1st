## go 多模块工作区教程

## 导读
- 随着 2020 年 3 月 15 日 go 1.18 正式发布，新版本除了对性能的提升之外，还引入了很多新功能，其中就有 go 期盼已久的功能 泛型，同时还引入的 模糊测试和多模块工作区。

- go 模块工作区主要解决 go.mod `replace` 将工作中需要的代码仓库替换为本地仓库，有两个好处
> 1. 方便进行依赖的代码调试(打断点、修改代码)、排查依赖代码 bug
> 3. 方便同时进行两个依赖仓库进行同时开发调试

## 开发流程演示
- 在现在微服务盛行的年代，一个人会维护多个代码仓库，很多的时候是多个仓库进行同时开发
- 假设我们现在进行 **hello** 仓库开发，先建立 **hello** 仓库代码，并生成
> **hello** 仓库实现的功能是，实现将传入的字符串反转

```shell
mkdir hello
cd hello
# 代码仓库启动 go mod 依赖管理，生成 go.mod 文件
go mod init github.com/link1st/link1st/workspaces/hello
# 下载依赖包
go get github.com/link1st/example
go mod tidy
# 编写 main 文件
vim main.go
```

- **main.go** 代码

```go
// Package main main 文件，go 多模块工作区演示代码
// 实现将传入的字符串反转
package main

import (
"flag"
"fmt"

"github.com/link1st/example/stringutil"
)

var (
str = ""
)

func init() {
flag.StringVar(&str, "str", str, "输入字符")
flag.Parse()
}

func main() {
if str == "" {
fmt.Println("示例: go run main.go -str hello")
fmt.Println("str 参数必填")
flag.Usage()
return
}

// 调用公共仓库，进行字符串反转
str = stringutil.Reversal(str)
// 输出反转后的字符串
fmt.Println(str)
return
}


```

- 运行代码 `go run main.go -str hello` 或 `go run github.com/link1st/example/workspaces/hello -str "hello world"` 可以看到输出了 **hello world** 反转以后的字符串
```
> go run main.go -str "hello world"
dlrow olleh
```
- 到这里，最初的功能已近完成，但是后续需求变动，不仅需要输出反转以后的字符串，还需要将字符串大写
- 我们则需要去 common 包中开发新的功能。

```
# 回到工作根目录，将 common 代码下载到本地进行添加新的功能


```

## 多模块工作区
### 使用条件
- go 1.18 或更高版本 [go 安装](https://go.dev/doc/install)

```shell
# 查看 go 版本
> go version
go version go1.18 darwin/amd64
```

- 初始化模块

```shell
$ mkdir hello
$ cd hello
$ go mod init example.com/hello
go: creating new go.mod: module example.com/hello
```

- 运行代码

```shell
go run main.go
go run example.com/hello
```

## 9、参考文献

[Go 1.18 is released!](https://go.dev/blog/go1.18)

[Tutorial: Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces)
