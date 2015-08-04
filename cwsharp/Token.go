// Copyright (c) CWSharp. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package cwsharp

type TokenType int

var TokenType_PUNC, TokenType_ALPHANUM, TokenType_NUM, TokenType_CJK TokenType = 1, 2, 3, 4

type Token struct {
	Text string
	Type TokenType
}

var token_empty = Token{"", TokenType_PUNC}

func (token Token) String() string {
	tokenType := "UNKNOW"
	switch token.Type {
	case TokenType_PUNC:
		tokenType = "PUNC"
	case TokenType_ALPHANUM:
		tokenType = "ALPHANUM"
	case TokenType_NUM:
		tokenType = "NUM"
	case TokenType_CJK:
		tokenType = "CJK"
	}
	return token.Text + ":" + tokenType
}

func (token *Token) Length() int {
	return len(token.Text)
}

func (token *Token) SetBuffer(text string) *Token {
	token.Text = text
	return token
}

func (token *Token) SetType(_type TokenType) *Token {
	token.Type = _type
	return token
}

func (token *Token) IsNull() bool {
	return token.Text == ""
}

func NewToken(text string, _type TokenType) Token {
	return Token{Text: text, Type: _type}
}
