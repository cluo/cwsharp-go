cwsharp-go
====
cwsharp-go是Golang实现的中文分词库，支持多种分词模式，支持自定义字典和扩展。

.NET版：[CWSharp-C#](https://github.com/yamool/CWSharp)

Python版: [CWSharp-Python](https://github.com/zhengchun/cwsharp-python)

安装&测试
====
```
$ go get github.com/zhengchun/cwsharp-go
$ cd main
$ go run main.go Hello,World!你好，世界!
```

分词算法
====
cwsharp-go支持多种分词算法，你可以根据需求选择适合自己的或者自定义新的分词算法。

## mmseg-tokenizer

标准的基于词典的分词方法。

**tips: 建议使用单一实例，避免每次分词都需重新加载字典**

```go
tokenizer, err := cwsharp.New("../data/cwsharp.dawg") //加载字典
iter := tokenizer.Tokenize(strings.NewReader("Hello,world!你好,世界!"))
for tok := iter.Next(); tok != nil; tok = iter.Next() {
	fmt.Printf("%s/%s ", tok.Text, tok.Type)
}
>> hello/w ,/p world/w !/p 你好/w ,/p 世界/w !/p
```

## bigram-tokenizer

二元分词方法，无需字典，速度快，支持完整的英文和数字切分。

```go
iter := cwsharp.BigramTokenize(strings.NewReader("世界人民大团结万岁!"))
for token := iter.Next(); token != nil; token = iter.Next() {
	fmt.Printf("%s/%s ", token.Text, token.Type)
}
>> 世界/w 界人/w 人民/w 民大/w 大团/w 团结/w 结万/w 万岁/w !/p
```

## whitespace-tokenizer

标准的英文分词，无需字典，适合切分英文的内容，中文会被当做独立的字符输出。

```go
iter := cwsharp.WhitespaceTokenize(strings.NewReader("Hello,world!你好!"))
for token := iter.Next(); token != nil; token = iter.Next() {
	fmt.Printf("%s/%s ", token.Text, token.Type)
}
>> hello/w ,/p world/w !/p 你/w 好/w !/p
```

版本历史
====
- 2.0 [2017-01]

	重写了代码以及目录布局, 尽量将代码简化以及符合golang的使用.
		
	- bigram,mmseg,simple三个独立的包整合到一起.
	- golang标准库的io.Reader代替自定义cwsharp.Reader的实现.
	- WhitespaceTokenize取代simple包.
	- BigramTokenize取代bigram包.
	- Token.Type代替Token.Kind. (PUNC,NUMBER,WORD)
	- Tokenizer接口约束.
- 1.1 
	- 重构架构方面的设计，实现了自定义分词扩展。
- 1.0 
	- C#版本的移植

TODO-List
====
- 自定义Filter功能,比如StopwordFilter。
- 将早期版本的自定义字典功能移到新版本中。