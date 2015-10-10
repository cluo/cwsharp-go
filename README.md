分词包(golang)，支持中英文，混合词组，自定义字典，具有良好的自定义分词扩展。

C#版本：[CWSharp-C#](https://github.com/yamool/CWSharp)

安装&运行
====
> go get github.com/zhengchun/cwsharp-go

> go run test.go

分词包列表
====
目前提供三种分词包，不同的分词包采用不同的分词算法

- simple - 简单的分词包,提供基本的字母或数字的分词功能，输出单个中文字符(一元分词)

- bigram - 二元分词包

- mmseg -  基于词典的分词包,支持自定义字典和中英文混合。字典采用[DAFSA](https://en.wikipedia.org/wiki/Deterministic_acyclic_finite_state_automaton)

示例
====
```go
package main

import (
	"fmt"
	"github.com/zhengchun/cwsharp-go/cwsharp"
	"github.com/zhengchun/cwsharp-go/cwsharp/mmseg"
)

func main() {
	tokenizer := mmseg.New("data//cwsharp.dawg")
	for iter := tokenizer.Traverse(cwsharp.ReadString(text)); iter.Next(); {
		token:=iter.Cur()
		fmt.Printf("%s:%s\n", token.Text(), token.Kind())
	}
}
```

更新日志
====
- 1.0 - C#版本的移植
- 1.1 - 重构架构方面的设计，实现了自定义分词扩展。