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
	for token := iter.Next(); token != nil; token = iter.Next() {
		fmt.Printf("%s/%s ", token.Text, token.Type)
	}
}
