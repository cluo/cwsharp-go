package main

import (
	"fmt"
	"github.com/zhengchun/cwsharp-go/cwsharp"
)

func main() {
	file := "data//cwsharp.dawg"
	tokenizer := cwsharp.NewStandardTokenizer(file, true)
	for _, text := range []string{"长春市长春药店", "研究生命起源", "Hello,World!"} {
		for token, next := tokenizer.Traverse(text)(); next != nil; token, next = next() {
			fmt.Printf(token.String())
			fmt.Print(" / ")
		}
		fmt.Println()
	}
}