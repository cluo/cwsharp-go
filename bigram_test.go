package cwsharp

import (
	"testing"
)

func TestBigramTokenizer(t *testing.T) {
	s := `hello,world!你好，世界！`
	expected := []string{"hello", ",", "world", "!", "你好", ",", "世界", "!"}

	testTokenizer(t, TokenizerFunc(BigramTokenize), s, expected)
}
