// Copyright (c) CWSharp. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.
package cwsharp

import (
	"testing"
)

func TestEqual(t *testing.T) {
	token_1 := NewToken("hello", TokenType_ALPHANUM)
	token_2 := NewToken("hello", TokenType_ALPHANUM)
	if token_1 != token_2 {
		t.Fail()
	}
}

func TestLength(t *testing.T) {
	token := NewToken("hello", TokenType_ALPHANUM)
	if token.Length() != 5 {
		t.Fail()
	}
}

func TestChangeToken(t *testing.T) {
	token := NewToken("hello", TokenType_ALPHANUM)
	token.SetText("10").SetType(TokenType_NUM)
	if !(token.Length() == 2 || token.Type == TokenType_NUM) {
		t.Fail()
	}
}
