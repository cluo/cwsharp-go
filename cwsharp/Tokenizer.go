package cwsharp

type Tokenizer interface {
	Traverse(r Reader) TokenIterator
}

type TokenIterator interface {
	Next() bool
	Cur() Token
}
