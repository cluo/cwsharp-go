package ngram

import (
	"errors"
	"fmt"
	"github.com/zhengchun/cwsharp-go/cwsharp"
)

var Kind int

const (
	Unigram Kind = iota //一元
	Bigram              //二元
)

func New(k Kind) cwsharp.Tokenizer {
	switch k {
	case Bigram:
		{
			return &BigramTokenizer{}
		}
	}
	panic(errors.New(fmt.Sprintf("不支持指定类型的分词。%d", k)))
}
