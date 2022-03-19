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
	// 将反转以后的字符串进行大写
	str = stringutil.ToUpper(str)
	// 输出反转后的字符串
	fmt.Println(str)
	return
}
