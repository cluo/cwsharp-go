// CWSharp is a text segmentation package for chinese.
//
package cwsharp

import "io"

// Token iterator.
type Iterator interface {
	Next() *Token
}

type IteratorFunc func() *Token

func (f IteratorFunc) Next() *Token {
	return f()
}

// Tokenizer is an interface that divides text into a
// sequence of tokens.
type Tokenizer interface {
	// Tokenize reads a text stream and divides into a
	// sequence of tokens.
	Tokenize(io.Reader) Iterator
}

// TokenizerFunc is the Tokenizer utility that help
// wrappered a specified tokenize function as Tokenizer.
type TokenizerFunc func(io.Reader) Iterator

func (f TokenizerFunc) Tokenize(r io.Reader) Iterator {
	return f(r)
}
