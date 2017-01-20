package cwsharp

import (
	"fmt"
	"unicode"
)

// Token represents a word text and with its kind of type.
type Token struct {
	Text string
	Kind Kind
}

func (t Token) String() string {
	return fmt.Sprintf("%s: %s", t.Text, t.Kind)
}

// Kind is a word type.
type Kind int

const (
	// 标点标号
	PUNCT Kind = 1 << iota
	// 数字
	NUMERAL = 1 << iota
	// 字母
	LETTER = 1 << iota
	// 中文
	CJK = 1 << iota
)

func (k Kind) String() (s string) {
	s = "UNKNOW"
	if k&NUMERAL == NUMERAL && k&LETTER == LETTER {
		s = "ALPHANUM" // 字母+数字混合
		return
	}
	switch {
	case k == PUNCT:
		s = "PUNCT"
	case k == NUMERAL:
		s = "NUMERAL"
	case k == LETTER:
		s = "LETTER"
	case k == CJK:
		s = "CJK"
	}
	return s
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

func DetermineKind(r rune) Kind {
	switch {
	case isPunct(r):
		return PUNCT
	case isLetter(r):
		return LETTER
	case isNumer(r):
		return NUMERAL
	case isCjk(r):
		return CJK
	}
	return 0
}
