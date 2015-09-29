// base包提供对包含英文、数字或字符分隔符组成的字符串分词
package base

import (
	"github.com/zhengchun/cwsharp-go/cwsharp"
	"unicode"
)

const (
	KPunct  cwsharp.TokenType = (iota + 1) //标点符号
	KNumber                                //数字，包含小数
	KLetter                                //包含字母或数字混合
	KOther                                 //其它字符
)

type Tokenizer struct {
}

type TextBreaker struct {
	cur    cwsharp.Token
	p      *Tokenizer
	reader cwsharp.Reader
	// 自定义字符组合的策略
	CheckContinue func(r rune, n rune, via []rune) (cwsharp.TokenType, bool)
	cwsharp.TextBreaker
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

func defaultCheckContinue(r rune, n rune, via []rune) (cwsharp.TokenType, bool) {
	if unicode.IsSpace(r) {
		return KPunct, false
	} else if unicode.IsPunct(r) {
		/*
			U.S.A
		*/
		return KPunct, false
	} else if unicode.IsLower(r) || unicode.IsUpper(r) {
		return KLetter, true
	} else if unicode.IsNumber(r) {
		return KNumber, true
	}
	return KOther, false
}

func (b *TextBreaker) NextToken() (t cwsharp.Token, err error) {
	if b.CheckContinue == nil {
		b.CheckContinue = defaultCheckContinue
	}
	var via []rune
	r := b.reader
	for ch, d := r.ReadRule(), 0; ch != cwsharp.EOF; ch = r.ReadRule() {
		ch = normalizeRule(ch)
		k, ok := b.CheckContinue(ch, r.Peek(), via)
		if ok {
			via = append(via, ch)
			t.SetType(k)
		} else {
			if d == 0 {
				via = append(via, ch)
				t.SetType(k)
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
	t.SetText(string(via))
	return
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
