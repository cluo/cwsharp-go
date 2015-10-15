// 英文分词包，提供对包含英文、数字或字符分隔符组成的字符串分词
package simple

import (
	"errors"
	"fmt"
	"github.com/zhengchun/cwsharp-go/cwsharp"
	"unicode"
)

var done error = errors.New("No more token to read.")

// 默认实例对象
var Default *Tokenizer = &Tokenizer{DefaultCheckContinue}

type Tokenizer struct {
	// 自定义字符组合的策略
	CheckContinue func(r cwsharp.Reader, ch rune, via []rune) (cwsharp.Kind, bool)
}

type Token struct {
	text string
	kind cwsharp.Kind
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

type TokenIterator struct {
	p      *Tokenizer
	ch     cwsharp.Token
	reader cwsharp.Reader
}

func (t *Tokenizer) Traverse(r cwsharp.Reader) cwsharp.TokenIterator {
	if t.CheckContinue == nil {
		t.CheckContinue = DefaultCheckContinue
	}
	return &TokenIterator{
		reader: r,
		p:      t,
	}
}

func New() *Tokenizer {
	return &Tokenizer{}
}

func isPunct(c rune) bool {
	return unicode.IsSpace(c) || unicode.IsPunct(c)
}

func isLetter(c rune) bool {
	return unicode.IsLower(c) || unicode.IsUpper(c)
}

func isCjk(c rune) bool {
	return unicode.Is(unicode.Scripts["Han"], c)
}

func checkContinueWithPunct(r cwsharp.Reader, ch rune, via []rune) (cwsharp.Kind, bool) {
	if ch == ' ' || len(via) == 0 || r.Peek() == cwsharp.EOF {
		return cwsharp.PUNCT, false
	}
	return cwsharp.PUNCT, false
}

func DefaultCheckContinue(r cwsharp.Reader, ch rune, via []rune) (cwsharp.Kind, bool) {
	if isPunct(ch) {
		return checkContinueWithPunct(r, ch, via)
	} else if isLetter(ch) {
		for _, c := range via {
			if unicode.IsNumber(c) {
				return cwsharp.ALPHANUM, true
			}
		}
		return cwsharp.LETTER, true
	} else if unicode.IsNumber(ch) {
		for _, c := range via {
			if unicode.IsLetter(c) {
				return cwsharp.ALPHANUM, true
			}
		}
		return cwsharp.NUMERAL, true
	} else if isCjk(ch) {
		return cwsharp.CJK, false
	}
	return cwsharp.PUNCT, false

}

func (iter *TokenIterator) nextToken() (cwsharp.Token, error) {
	var err error
	t := &Token{}
	var via []rune
	r := iter.reader
	for ch, d := r.ReadRule(), 0; ch != cwsharp.EOF; ch = r.ReadRule() {
		ch = normalizeRule(ch)
		k, ok := iter.p.CheckContinue(r, ch, via)
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
		err = done
	}
	t.text = string(via)
	return t, err
}

func (iter *TokenIterator) Next() bool {
	t, err := iter.nextToken()
	if err != nil || err == done {
		return false
	}
	iter.ch = t
	return true
}

func (iter *TokenIterator) Cur() cwsharp.Token {
	return iter.ch
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
