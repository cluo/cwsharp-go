package main

import (
	"fmt"
	"github.com/zhengchun/cwsharp-go/cwsharp"
	"github.com/zhengchun/cwsharp-go/cwsharp/mmseg"
)

func main() {
	file := "data//cwsharp.dawg"
	tokenizer := mmseg.New(file)
	for _, text := range []string{"长春市长春药店", "研究生命起源", "Hello,World!"} {
		for iter := tokenizer.Traverse(cwsharp.ReadString(text)); iter.Next(); {
			fmt.Print(iter.Cur())
			fmt.Print(" / ")
		}
		fmt.Println()
	}
}
