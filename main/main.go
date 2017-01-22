package main

import (
	"fmt"
	"strings"

	"github.com/zhengchun/cwsharp-go"
)

func main() {
	iter := cwsharp.WhitespaceTokenize(strings.NewReader("Hello world!你好，世界！"))
	for t := iter(); t.Type != cwsharp.EOF; t = iter() {
		fmt.Println(t.Text)
	}
	fmt.Println("done.")
}
