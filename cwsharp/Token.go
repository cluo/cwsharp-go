package cwsharp

import (
	"fmt"
)

type TokenType int

const (
	PUNC     TokenType = (iota + 1) //标号
	ALPHANUM                        //混合，包含字母、数字或其它连接符
	NUMERAL                         //数字
	CJK                             //中文
)

type Token struct {
	text string
	kind TokenType
}

func (t TokenType) String() string {
	tokenType := "UNKNOW"
	switch t {
	case PUNC:
		tokenType = "PUNC"
	case ALPHANUM:
		tokenType = "ALPHANUM"
	case NUMERAL:
		tokenType = "NUMERAL"
	case CJK:
		tokenType = "CJK"
	}
	return tokenType
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

func (t *Token) Type() TokenType {
	return t.kind
}

// 更改Token的内容
func (t *Token) SetText(text string) *Token {
	t.text = text
	return t
}

// 更改Token类型
func (t *Token) SetType(k TokenType) *Token {
	t.kind = k
	return t
}
