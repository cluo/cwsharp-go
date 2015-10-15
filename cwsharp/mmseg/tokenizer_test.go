package mmseg

import (
	"fmt"
	"github.com/zhengchun/cwsharp-go/cwsharp"
	"testing"
)

func Test1(t *testing.T) {
	input := "研究生命起源.一次性交100元"
	tokenizer := New("https://github.com/zhengchun/cwsharp-go/raw/master/data/cwsharp.dawg")
	iter := tokenizer.Traverse(cwsharp.NewStringReader(input))
	for iter.Next() {
		fmt.Println(iter.Cur())
	}
}
