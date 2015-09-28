package cwsharp

import (
	"errors"
)

var DONE error = errors.New("No more token to read.")

type Tokenizer interface {
	Traverse(r Reader) TokenIterator
}

type Breaker interface {
	Next() (Token, error)
}

type TokenIterator interface {
	Next() bool
	Cur() Token
}

// TokenIterator的实现
type TokenIteratorWrap struct {
	b   Breaker
	cur Token
}

func (t *TokenIteratorWrap) Next() bool {
	ch, err := t.b.Next()
	if err != nil {
		if err == DONE {
			return false
		}
		panic(err) //异常触发
	}
	t.cur = ch
	return true
}

func (t *TokenIteratorWrap) Cur() Token {
	return t.cur
}
