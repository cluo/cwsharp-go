package main

import (
	"fmt"
	"github.com/zhengchun/cwsharp-go/cwsharp"
	"github.com/zhengchun/cwsharp-go/cwsharp/mmseg"
)

func main() {
	file := "data//cwsharp.dawg"
	tokenizer := cwsharp.NewStopwordFilter(mmseg.New(file))
	w := map[string]bool{"the": true}
	tokenizer.CheckIgnore = func(t cwsharp.Token) bool {
		_, ok := w[t.Text()]
		if t.Kind() == cwsharp.PUNCT || ok {
			return true
		}
		return false
	}
	for _, text := range []string{"长春市长春药店", "The quick brown fox jumps over the lazy dog"} {
		for iter := tokenizer.Traverse(cwsharp.NewStringReader(text)); iter.Next(); {
			fmt.Print(iter.Cur())
			fmt.Print(" / ")
		}
		fmt.Println()
	}
}
