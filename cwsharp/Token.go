package cwsharp

import (
	"fmt"
)

type TokenType int32

type Token struct {
	text string
	kind TokenType
}

func (t *Token) String() string {
	return fmt.Sprintf("%s:%s", t.text, t.kind)
}

func (t *Token) Length() int {
	return len(t.text)
}

func (t *Token) Text() string {
	return t.text
}

func (t *Token) Kind() TokenType {
	return t.kind
}

func (t *Token) SetText(text string) {
	t.text = text
}

func (t *Token) SetType(k TokenType) {
	t.kind = k
}
