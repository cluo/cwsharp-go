package cwsharp

import "unicode"

// a token type.
type Type int

const (
	EOF    Type = iota
	PUNC        // .,| []
	NUMBER      // 12345 12.34
	WORD        // abc 中文 ABC123 wi-fi

	latinAlpha // [a-z,A-Z]
	cjk        // 你好，世界
)

// Token represents a word text and with its kind of type.
type Token struct {
	Text string
	// A type
	Type Type
	// An arbitrary source position location.
	Pos int
}

func determineType(r rune) Type {
	switch {
	case unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r):
		return PUNC
	case unicode.IsNumber(r):
		return NUMBER
	case unicode.IsUpper(r) || unicode.IsLower(r):
		return latinAlpha
	case unicode.Is(unicode.Scripts["Han"], r):
		return cjk
	}
	return WORD
}
