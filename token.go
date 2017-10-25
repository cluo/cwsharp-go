package cwsharp

import "unicode"

// A token type.
type Type int

const (
	PUNCT  = iota // .,| []
	NUMBER        // 12345 12.34
	ALPHA         // [a-z]
	WORD          // abc 中文 ABC123 wi-fi
)

func (typ Type) String() string {
	switch typ {
	case PUNCT:
		return "punct"
	case NUMBER:
		return "number"
	case ALPHA:
		return "alpha"
	case WORD:
		return "word"
	}
	return ""
}

// Token represents a word text and with its kind of type.
type Token struct {
	// A token text.
	Text string
	// A token type.
	Type Type
	// An arbitrary source position location.
	Pos int
}

func isNumber(r rune) bool {
	return unicode.IsNumber(r)
}

func isCjk(r rune) bool {
	return unicode.Is(unicode.Scripts["Han"], r)
}

func determineType(r rune) Type {
	switch {
	case unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r):
		return PUNCT
	case unicode.IsNumber(r):
		return NUMBER
	case unicode.IsUpper(r) || unicode.IsLower(r):
		return ALPHA
		//case unicode.Is(unicode.Scripts["Han"], r):
		//	return cjk
	}
	return WORD
}
