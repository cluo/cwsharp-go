// 二元分词包
package bigram

import (
	"errors"
	"fmt"
	"github.com/zhengchun/cwsharp-go/cwsharp"
	"github.com/zhengchun/cwsharp-go/cwsharp/simple"
	"unicode"
)

var done error = errors.New("No more token to read.")

type Tokenizer struct {
}

type Token struct {
	text string
	kind cwsharp.Kind
}

type TokenIterator struct {
	state  bool
	inner  cwsharp.TokenIterator
	ch     cwsharp.Token
	reader cwsharp.Reader
}

func New() *Tokenizer {
	return &Tokenizer{}
}

func (t *Tokenizer) Traverse(r cwsharp.Reader) cwsharp.TokenIterator {
	return &TokenIterator{
		inner:  simple.New().Traverse(r),
		reader: r,
		state:  true,
	}
}

func (iter *TokenIterator) Next() bool {
	ch := iter.reader.Peek()
	if ch == cwsharp.EOF {
		return false
	}
	if isCjk(ch) {
		//二元分词包
		if via, ok := cjkToken(iter.reader, iter.state); ok {
			iter.ch = &Token{string(via), cwsharp.CJK}
			iter.state = false
			return true
		}
	}
	iter.state = true
	//交给默认的分词包处理
	if iter.inner.Next() {
		iter.ch = iter.inner.Cur()
		return true
	}
	return false
}

func (iter *TokenIterator) Cur() cwsharp.Token {
	return iter.ch
}

func cjkToken(r cwsharp.Reader, s bool) ([]rune, bool) {
	ch := r.ReadRule()
	if c := r.Peek(); c != cwsharp.EOF && isCjk(c) {
		return []rune{ch, r.Peek()}, true
	} else if s {
		return []rune{ch}, true
	}
	return nil, false
}

func isCjk(c rune) bool {
	return unicode.Is(unicode.Scripts["Han"], c)
}

func (t *Token) Text() string {
	return t.text
}

func (t *Token) Kind() cwsharp.Kind {
	return t.kind
}

func (t *Token) String() string {
	return fmt.Sprintf("%s : %s", t.text, t.kind)
}
