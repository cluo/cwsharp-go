package ngram

import (
	"github.com/zhengchun/cwsharp-go/cwsharp"
	"io"
)

type BigramTokenizer struct {
	cwsharp.Tokenizer
}

func (t *BigramTokenizer) Traverse(r cwsharp.Reader) cwsharp.Traverser {
	return nil
}

type bigramTokenIterator struct {
	src cwsharp.Reader
}
