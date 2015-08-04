CWSharp-Go
====
Go中文分词库，支持中英文，混合词组，自定义字典。[CWSharp](https://github.com/yamool/CWSharp)的Golang版本.

安装&运行
====
> go get github.com/zhengchun/cwsharp-go

> go run test.go

说明
====

- StandardTokenizer.go - 基于词典的分词类

- BigramTokenizer.go - 二元分词类

- StopwordTokenizer.go - 扩展类，提供过滤词的分词

- WordUtil.go - 提供字典管理帮助类

使用帮助可以查看对应的测试用例。