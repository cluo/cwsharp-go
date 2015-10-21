package simple

import (
	"fmt"
	"github.com/zhengchun/cwsharp-go/cwsharp"
	"testing"
)

func Test(t *testing.T) {
	var input string = "Hello,World!你好，世界!14.4%,abc1233,11a-b,2015.2.1"
	tokenizer := New()
	for iter := tokenizer.Traverse(cwsharp.NewStringReader(input)); iter.Next(); {
		fmt.Println(iter.Cur())
	}
}
