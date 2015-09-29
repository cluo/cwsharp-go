package cwsharp

import (
	"errors"
)

// 指定的内容已经遍历完成结束
var DONE error = errors.New("No more token to read.")

type Tokenizer interface {
	Traverse(r Reader) TokenIterator
}

type TextBreaker interface {
	NextToken() (Token, error)
}

type TokenIterator interface {
	Next() bool
	Cur() Token
}
