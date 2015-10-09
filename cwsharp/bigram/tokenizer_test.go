package bigram

import (
	"fmt"
	"github.com/zhengchun/cwsharp-go/cwsharp"
	"testing"
)

func Test(t *testing.T) {
	input := "一次性交100元"
	tokenizer := New()
	iter := tokenizer.Traverse(cwsharp.ReadString(input))
	for iter.Next() {
		fmt.Println(iter.Cur())
	}
}
