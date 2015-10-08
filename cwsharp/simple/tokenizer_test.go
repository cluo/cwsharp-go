package simple

import (
	"fmt"
	"github.com/zhengchun/cwsharp-go/cwsharp"
	"testing"
)

func Test(t *testing.T) {
	var input string = "Hello,World!你好，世界!"
	tokenizer := New()
	iter := tokenizer.Traverse(cwsharp.NewStringReader(input))
	for iter.Next() {
		fmt.Println(iter.Cur())
	}

}
