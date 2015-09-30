// 二元分词包
package ngram

import (
	"fmt"
	"github.com/zhengchun/cwsharp-go/cwsharp"
	"github.com/zhengchun/cwsharp-go/cwsharp/base"
	"unicode"
)

type Tokenizer struct {
	inst cwsharp.Tokenizer
}

const (
	KPunct   cwsharp.TokenType = iota //标点符号
	KNumber                           //数字，包含小数
	KLetter                           //包含字母或数字混合
	KUnicode                          //其它字符
	KCjk                              //中文
)

type Token struct {
	text string
	kind cwsharp.TokenType
}

type TextBreaker struct {
	cur    cwsharp.Token
	p      *Tokenizer
	reader cwsharp.Reader
	wrap   *base.TextBreaker
	cwsharp.TextBreaker
}

func New() *Tokenizer {
	return &Tokenizer{inst: base.New()}
}

func (t *Tokenizer) Traverse(r cwsharp.Reader) cwsharp.TokenIterator {
	wrap := t.inst.Traverse(r).(base.TextBreaker)
	return &TextBreaker{
		p:      t,
		reader: r,
		wrap:   wrap,
	}
}

func (t *Token) Text() string {
	return t.text
}

func (t *Token) Kind() cwsharp.TokenType {
	return t.kind
}

func (t *Token) String() string {
	k := "Unicode"
	switch t.kind {
	case KPunct:
		{
			k = "Punct"
		}
	case KLetter:
		{
			k = "Letter"
		}
	case KNumber:
		{
			k = "Number"
		}
	case KCjk:
		{
			k = "Cjk"
		}
	}
	return fmt.Sprintf("%s, %s", t.text, k)
}

func (b *TextBreaker) NextToken() (cwsharp.Token, error) {
	var err error
	r := b.reader.Peek()
	if unicode.Is(unicode.Scripts["Han"], r) {
		return &Token{}, nil
	} else {
		//交给默认的分词包处理
		return b.NextToken()
	}
}

func (b *TextBreaker) Next() bool {
	t, err := b.NextToken()
	if err != nil || err == cwsharp.DONE {
		return false
	}
	b.cur = t
	return true
}

func (b *TextBreaker) Cur() cwsharp.Token {
	return b.cur
}
