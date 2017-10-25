package cwsharp

import (
	"testing"
)

func TestWhitespaceTokenizer(t *testing.T) {
	s := `hello,world!你好，世界！`
	expected := []string{"hello", ",", "world", "!", "你", "好", ",", "世", "界", "!"}

	testTokenizer(t, TokenizerFunc(WhitespaceTokenize), s, expected)
}
