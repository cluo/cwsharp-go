package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zhengchun/cwsharp-go"
)

func main() {
	tokenizer, err := cwsharp.New("../data/cwsharp.dawg")
	if err != nil {
		panic(err)
	}
	arg := os.Args[1]
	iter := tokenizer.Tokenize(strings.NewReader(arg))
	for token, ok := iter(); ok; token, ok = iter() {
		fmt.Printf("%s/%s ", token.Text, token.Type)
	}
}
