package mmseg

import (
	"fmt"
	"github.com/zhengchun/cwsharp-go/cwsharp"
	"testing"
)

func Test1(t *testing.T) {
	input := "研究生命起源."
	tokenizer := New("D:\\gorepos\\src\\github.com\\zhengchun\\cwsharp-go\\data\\cwsharp.dawg")
	iter := tokenizer.Traverse(cwsharp.NewStringReader(input))
	for iter.Next() {
		fmt.Println(iter.Cur())
	}
}
