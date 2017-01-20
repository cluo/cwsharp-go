package cwsharp

import "io"

// Token iterator.
type Iterator func() (Token, bool)

// Tokenizer is an interface that divides text into a
// sequence of tokens.
type Tokenizer interface {
	// Tokenize reads a text stream and divides into a
	// sequence of tokens.
	Tokenize(io.Reader) Iterator
}

func whitespaceTokenize(r *reader) (Token, bool) {
	for {

	}
	return Token{}, false
}
