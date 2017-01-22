package main

import (
	"fmt"
	"strings"

	"github.com/zhengchun/cwsharp-go"
)

func main() {
	iter := cwsharp.WhitespaceTokenize(strings.NewReader("Hello,world!你好，世界！"))
	for tok, ok := iter(); ok; tok, ok = iter() {
		fmt.Println(tok.Text)
	}
	fmt.Println("done.")
}
