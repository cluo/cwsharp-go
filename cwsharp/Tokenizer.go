package cwsharp

import "io"

// The text tokenizer.
type Tokenizer interface {
	Tokenize(io.Reader) Iterator
}

type Iterator interface {
	Next() (Token, bool)
}
