package cwsharp

import (
	"strings"
	"testing"
)

func testTokenizer(t *testing.T, tokenizer Tokenizer, s string, expected []string) {
	var result []string
	iter := tokenizer.Tokenize(strings.NewReader(s))
	for tok := iter.Next(); tok != nil; tok = iter.Next() {
		result = append(result, tok.Text)
	}

	if g, e := len(result), len(expected); g != e {
		t.Fatalf("expected token numbers is %d,but got %d", e, g)
	}

	for i := 0; i < len(result); i++ {
		if g, e := result[i], expected[i]; g != e {
			t.Fatalf("expected token[%d] is %s,but got is %s", i, e, g)
		}
	}
}

func TestMMSEGTest(t *testing.T) {
	s := `hello,world!你好，世界！`
	expected := []string{"hello", ",", "world", "!", "你好", ",", "世界", "!"}

	tokenizer, err := New(`./data/cwsharp.dawg`)
	if err != nil {
		t.Error(err)
	}
	testTokenizer(t, tokenizer, s, expected)
}
