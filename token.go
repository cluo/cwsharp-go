package cwsharp

import "unicode"

// a token type.
type Type int

const (
	// The end of file.
	EOF    Type = iota
	PUNC        // .,| []
	NUMBER      // 12345 12.34
	WORD        // abc 中文 ABC123 wi-fi
	alphabet
)

var TokenEOF = Token{Type: EOF}

// Token represents a word text and with its kind of type.
type Token struct {
	Text string
	// A type
	Type Type
	// An arbitrary source position location.
	Pos int
}

func DetermineType(r rune) Type {
	switch {
	case unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r):
		return PUNC
	case unicode.IsNumber(r):
		return NUMBER
	case unicode.IsLower(r) || unicode.IsUpper(r):
		return alphabet
		//case unicode.Is(unicode.Scripts["Han"], r):
	}
	return WORD
}
