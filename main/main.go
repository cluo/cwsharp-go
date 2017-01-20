package main

import (
	"fmt"
	"strings"

	"github.com/zhengchun/cwsharp-go"
)

func main() {
	tokenizer, err := cwsharp.NewStandard(`D:\gorepos\src\github.com\zhengchun\cwsharp-go\data\cwsharp.dawg`)
	if err != nil {
		panic(err)
	}

	iter := tokenizer.Tokenize(strings.NewReader("Hello,world!你好，世界！"))
	for t, ok := iter(); ok; t, ok = iter() {
		fmt.Println(t.Text)
	}
	fmt.Println("done.")
}
