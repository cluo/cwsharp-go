package cwsharp

type Tokenizer interface {
	Traverse(r Reader) Traverser
}

type Breaker interface {
	Next() (Token, error)
}

type Traverser interface {
	Next() bool
	Cur() Token
}
