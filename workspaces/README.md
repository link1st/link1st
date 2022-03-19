## go 多模块工作区教程

## 导读
- 随着 2020 年 3 月 15 日 go 1.18 正式发布，新版本除了对性能的提升之外，还引入了很多新功能，其中就有 go 期盼已久的功能泛型，同时还引入的多模块工作区和模糊测试。

- 关于泛型网上已经有很多介绍的教程了，这里我介绍一个实用的功能，多模块工作区的使用方法和教程。
- Go 多模块工作区能够使开发者能够更容易地同时处理多个模块的工作，如: 
> 方便进行依赖的代码调试(打断点、修改代码)、排查依赖代码 bug  
> 方便同时进行多个个依赖仓库进行同时开发调试

## 目录
- [多模块工作区](#多模块工作区)
- [开发流程演示](#开发流程演示)
- [总结](#总结)
- [参考文献](#参考文献)

## 多模块工作区
### 说明
- go 使用的是多模块工作区，可以让开发者更容易同时处理多个模块的开发。在 Go 1.17 之前，只能使用 `go.mod replace` 指令来实现，如果你正巧是同时进行多个模块的开发，使用它可能是很痛苦的。每次当你想要提交代码的时候，都不得不删除掉 **go.mod** 中的 `replace` 才能使模块稳定发布版本。
- 在使用 go 1.18 多模块工作区功能的时候，就使用这项工作变得简单容易处理。下面我来介绍怎么使用这以功能。
- [Go 多模块工作区文档、代码示例](https://github.com/link1st/link1st/tree/master/workspaces)

### 使用条件

- 首先 我们需要 go 1.18 或更高版本 [go 安装](https://go.dev/doc/install)

```shell
# 查看 go 版本
> go version
go version go1.18 darwin/amd64
```

### `go work` 支持命令

- 通常情况下，建议不要提交 **go.work** 文件到 git 上，因为它主要用于本地代码开发。
- 推荐在: `$GOPATH` 路径下执行，生成 **go.work** 文件

- `go work init` 初始化工作区文件，用于生成 **go.work** 工作区文件
> 初始化并写入一个新的 **go.work** 到当前路径下，可以指定需要添加的代码模块  
> 示例: `go work init  ./hello` 将本地仓库 **hello** 添加到工作区  
> **hello** 仓库必须是 go mod 依赖管理的仓库(**./hello/go.mod** 文件必须存在)  

- `go work use` 添加新的模块到工作区
> 命令示例:  
> `go work use ./example` 添加一个模块到工作区  
> `go work use ./example ./example1` 添加多个模块到工作区  
> `go work use -r ./example` 递归 **./example** 目录到当前工作区  
> 删除命令使用 `go work edit -dropuse=./example` 功能  

- `go work edit` 用于编辑 **go.work** 文件
> 可以使用 `edit` 命令编辑和手动编辑 `go.work` 文件效果是相同的
> 示例:  
> `go work edit -fmt go.work` 重新格式化 **go.work** 文件  
> `go work edit -replace=github.com/link1st/example=./example go.work` 替换代码模块  
> `go work edit -dropreplace=github.com/link1st/example` 删除替换代码模块  
> `go work edit -use=./example go.work` 添加新的模块到工作区  
> `go work edit -dropuse=./example go.work` 从工作区中删除模块  

- `go work sync` 将工作区的构建列表同步回
工作区的模块

- `go env GOWORK`
> 查看环境变量，查看当前工作区文件路径
> 可以排查工作区文件是否设置正确，**go.work** path 找不可以使用 GOWORK 指定  

```
> go env GOWORK
$GOPATH/src/link1st/link1st/workspaces/go.work
```

### **go.work** 文件结构
- 文件结构和 **go.mod** 文件结构类似，支持的 Go 版本号、指定工作区和需要替换的仓库
- 文件结构示例:

```
go 1.18

use (
    ./hello
    ./example
)

replace (
    github.com/link1st/example => ./example1
)

```

#### `use` 指定使用的模块目录

- 可以使用 `go work use <command>` 添加模块，也可以手动修改 **go.work** 工作区添加新的模块
- 在工作区中添加了模块路径，编译的时候会自动使用 **use** 中的本地代码进行代码编译，和 `replaces` 功能类似。

```
# 单模块结构
use ./hello

# 多模块结构
use (
    ./hello
    ./example
)
```

#### `replaces` 替换依赖仓库地址

- `replaces` 命令与 **go.mod** 指令相同，用于替换项目中依赖的仓库地址
- 需要注意的是 `replaces` 和 `use`  不能同时指定相同的本地路径  


> 同时指定报错信息:   
> go: workspace module github.com/link1st/example is replaced at all versions in the go.work file. To fix, remove the replacement from the go.work file or specify the version at which to replace the module.  

- 错误示例   

> 同时在 `use` 和 `replace` 指定相同的本地路径  

```
go 1.18

use (
    ./hello
    ./example
)

replace (
    github.com/link1st/example => ./example
)

```

#### **go.work** 文件优先级高于 **go.mod** 中定义

- 在同时使用 **go.work** 和 **go.mod** `replace` 功能的的时候分别指定不同的代码仓库路径，**go.work** 优先级高于 **go.mod** 中定义

> **go.mod** 中定义替换为本地仓库 **example**

```
replace (
    github.com/link1st/example => ./example1
)
```

> **go.work** 中定义替换为本地仓库 **example1**

```
replace (
    github.com/link1st/example => ./example1
)
```

- 在代码构建时候使用的是 **go.work**  指定的 **example1** 仓库的代码，**go.work** 优先级别更高


### 如何使用
- 在 Go 1.18 `go run` 和 `go build` 都会默认使用工作区功能
- `GOWORK` 也可以指定配置 **go.work** 文件位置

```
export GOWORK="~/go/src/test/go.18/workspace/go.work"
```

### 如何禁用工作区
- Go 全局变量 `GOWORK` 设置 `off` 则可以禁用工作区功能
> `export GOWORK=off`


## 开发流程演示
- 演示如何使用多模块工作区功能。在现在微服务盛行的年代，一个人会维护多个代码仓库，很多的时候是多个仓库进行同时开发
- 假设我们现在进行 **hello** 仓库开发，实现的功能是，实现将输入的字符串反转并输出，字符串反转功能依赖于 **github.com/link1st/example** (下文统称 **example** 仓库)公共仓库实现

- 新建 **hello** 项目

```shell
mkdir hello
cd hello
# 代码仓库启动 go mod 依赖管理，生成 go.mod 文件
go mod init github.com/link1st/link1st/workspaces/hello
# 下载依赖包
go get github.com/link1st/example
# 编写 main 文件
vim main.go
```

- **main.go** 代码

```go
// Package main main 文件，go 多模块工作区演示代码
// 实现将输入的字符串反转输出并输出
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

- 运行代码 `go run main.go -str "hello world"` 或 `go run github.com/link1st/link1st/workspaces/hello -str "hello world"` 可以看到输出了 **hello world** 反转以后的字符串

```
> go run main.go -str "hello world"
dlrow olleh
```
- 到这里，最初的功能已近完成，但是后续需求变动，不仅需要输出反转以后的字符串，还需要将字符串大写
- 我们则需要去 **example** 仓库中添加开发 **将字符串大写的功能**

```
# 回到工作根目录，将 common 代码下载到本地进行添加新的功能
# 下载依赖的 example 包
git clone git@github.com:link1st/example.git
# 在 example 包中添加 字符串大学的功能

```

- **vim example/stringutil/to_upper.go** 代码如下

```
// Package stringutil stringutil
package stringutil

import (
    "unicode"
)

// ToUpper 将字符串进行大写
func ToUpper(s string) string {
    r := []rune(s)
    for i := range r {
        r[i] = unicode.ToUpper(r[i])
    }
    return string(r)
}

```

- 由于代码还在本地调试，未提交 git 仓库中，这个时候就需要用到 Go 多模块工作区的功能了。

- 进入项目根目录，初始化我们现在正在开发的啊领个模块

```
# 初始化 go.work 文件
go work init  ./hello ./example
# 查看 go.work 文件内容
cat go.work
```

- 文件结构如下

```
go 1.18

use (
    ./example
    ./hello
)
```

- 回到 **hello** 项目，`vim main.go` 将字符串大写的功能添加上。

```
func main() {
    ...

    // 调用公共仓库，进行字符串反转
    str = stringutil.Reversal(str)
    // 增加字符大写的功能
    str = stringutil.ToUpper(str)
    // 输出反转后的字符串
    fmt.Println(str)
    
    ...
}
```

- 运行代码
> 可以看到输出了反转并 **大写** 的字符串，大写的函数功能只在本地，未提交到 git 上

```
go run main.go -str "hello world"
DLROW OLLEH
```

- 到这里，演示的代码已经全部完成

## 总结
- 使用 Go 多模块工作区的功能，可以让我们轻松在多个模块之间切换工作，更能适应现代微服务架构开发。


## 参考文献

[Go 1.18 is released!](https://go.dev/blog/go1.18)

[Tutorial: Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces)

[go-1.18-features](https://sebastian-holstein.de/post/2021-11-08-go-1.18-features/)

[link1st workspaces](https://github.com/link1st/link1st/tree/master/workspaces)


