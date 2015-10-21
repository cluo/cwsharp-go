// 英文分词包，提供对包含英文、数字或字符分隔符组成的字符串分词
// 支持下列混合格式:12-ab,ak49,12.1,P.M?

package simple

import (
	"fmt"
	"github.com/zhengchun/cwsharp-go/cwsharp"
	"unicode"
)

type Tokenizer struct {
	//CheckContinue func(r cwsharp.Reader, ch rune, via []rune) (cwsharp.Kind, bool)
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
	//if t.CheckContinue == nil {
	//	t.CheckContinue = DefaultCheckContinue
	//}
	return &TokenIterator{
		reader: r,
		p:      t,
	}
}

func New() *Tokenizer {
	return &Tokenizer{}
}

func isPunct(c rune) bool {
	return unicode.IsSpace(c) || unicode.IsPunct(c) || unicode.IsSymbol(c)
}

func isLetter(c rune) bool {
	return unicode.IsLower(c) || unicode.IsUpper(c)
}

func isCjk(c rune) bool {
	return unicode.Is(unicode.Scripts["Han"], c)
}

func isNumer(c rune) bool {
	return unicode.IsNumber(c)
}

func (iter *TokenIterator) Next() bool {
	if iter.reader.Peek() == cwsharp.EOF {
		return false
	}		
	ch := iter.reader.ReadRule()
	if isCjk(ch) {
		iter.ch = &Token{string(ch), cwsharp.CJK}
		return true
	}
	var kind cwsharp.Kind
	var via []rune
	if isLetter(ch){
		for ch != cwsharp.EOF {
			if isLetter(ch) {
				via = append(via, ch)
				kind |= cwsharp.LETTER
			} else if isNumer(ch) {
				via = append(via, ch)
				kind |= cwsharp.NUMERAL
			} else {
				nextch := iter.reader.Peek()
				if ch == '-' && (isNumer(nextch) || isLetter(nextch)) {
					via = append(via, ch)
					kind |= cwsharp.LETTER
				} else {
					break
				}
			}
			ch = iter.reader.ReadRule()
		}
		iter.ch = &Token{string(via), kind}
		return true
	} else if isNumer(ch) {
		times := 0
		endIndex := 0
		offset:=iter.reader.Pos()
		for i := 0; ch != cwsharp.EOF; i++ {
			if isNumer(ch) {
				via = append(via, ch)
				kind |= cwsharp.NUMERAL
			} else if isLetter(ch) {
				via = append(via, ch)
				kind |= cwsharp.LETTER
			} else if ch == '.'{
				if kind&cwsharp.LETTER == cwsharp.LETTER{
					via = via[:endIndex]
					iter.reader.Seek(iter.reader.Pos() - 1)
					break
				}else if times>0{
					//BUG 2015.1.1
					via = via[:endIndex]
					iter.reader.Seek(offset+endIndex-1)
					break
				}
				via = append(via, ch)
				endIndex = i
				times++
			} else {				
				iter.reader.Seek(iter.reader.Pos() - 1)
				break
			}
			ch = iter.reader.ReadRule()
		}
		iter.ch = &Token{string(via), kind}
		return true
	}
	iter.ch = &Token{string(ch), cwsharp.PUNCT}
	return true
}

func (iter *TokenIterator) Cur() cwsharp.Token {
	return iter.ch
}
