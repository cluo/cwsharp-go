// 英文分词包，提供对包含英文、数字或字符分隔符组成的字符串分词
package base

import (
	"fmt"
	"github.com/zhengchun/cwsharp-go/cwsharp"
	"unicode"
)

const (
	KPunct   cwsharp.TokenType = iota //标点符号
	KNumber                           //数字，包含小数
	KLetter                           //包含字母或数字混合
	KUnicode                          //其它字符
)

type Tokenizer struct {
	cwsharp.Tokenizer
}

type Token struct {
	text string
	kind cwsharp.TokenType
}

type TextBreaker struct {
	cur    cwsharp.Token
	p      *Tokenizer
	reader cwsharp.Reader
	// 自定义字符组合的策略
	CheckContinue func(r cwsharp.Reader, ch rune, via []rune) (cwsharp.TokenType, bool)
}

type TokenIterator struct {
	b *TextBreaker
}

func (t *Tokenizer) Traverse(r cwsharp.Reader) cwsharp.TokenIterator {
	return &TextBreaker{
		p:      t,
		reader: r,
	}
}

func New() *Tokenizer {
	return &Tokenizer{}
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
	}
	return fmt.Sprintf("%s, %s", t.text, k)
}

func DefaultCheckContinue(r cwsharp.Reader, ch rune, via []rune) (cwsharp.TokenType, bool) {
	if unicode.IsSpace(ch) {
		return KPunct, false
	} else if unicode.IsPunct(ch) {
		/*
			U.S.A
			22.3
			abc-hello
			name@domain.com
		*/
		return KPunct, false
	} else if unicode.IsLower(ch) || unicode.IsUpper(ch) {
		return KLetter, true
	} else if unicode.IsNumber(ch) {
		return KNumber, true
	}
	return KUnicode, false
}

func (b *TextBreaker) NextToken() (cwsharp.Token, error) {
	var err error
	t := &Token{}
	if b.CheckContinue == nil {
		b.CheckContinue = DefaultCheckContinue
	}
	var via []rune
	r := b.reader
	for ch, d := r.ReadRule(), 0; ch != cwsharp.EOF; ch = r.ReadRule() {
		ch = normalizeRule(ch)
		k, ok := b.CheckContinue(r, ch, via)
		if ok {
			via = append(via, ch)
			t.kind = k
		} else {
			if d == 0 {
				via = append(via, ch)
				t.kind = k
			} else {
				r.Seek(r.Pos() - 1)
			}
			break
		}
		d++
	}
	if len(via) == 0 {
		err = cwsharp.DONE
	}
	t.text = string(via)
	return t, err
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

// 大写字母转化小化、全角转半角
func normalizeRule(r rune) rune {
	if unicode.IsUpper(r) {
		return r + 32
	} else if r >= 0xff01 && r <= 0xff5d {
		return r - 0xFEE0
	}
	return r
}
