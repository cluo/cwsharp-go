// Copyright (c) CWSharp. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package cwsharp

import (
	"testing"
)

func TestwhiteSpaceTokenBreaker_1(t *testing.T) {
	reader := newStringReader("Hello,World!", true)
	tokenizer := whiteSpaceTokenBreaker{reader: reader}
	token := tokenizer.Next()
	if token.Text != "hello" {
		t.Fail()
	}
	token = tokenizer.Next()
	if token.Text != "," {
		t.Fail()
	}
	token = tokenizer.Next()
	if token.Text != "world" {
		t.Fail()
	}
	token = tokenizer.Next()
	if token.Text != "!" {
		t.Fail()
	}
	//end
	token = tokenizer.Next()
	if !token.IsNull() {
		t.Fail()
	}
}

func TestwhiteSpaceTokenBreaker_2(t *testing.T) {
	reader := newStringReader("this is a cwsharp.", true)
	tokenizer := whiteSpaceTokenBreaker{reader: reader}
	token := tokenizer.Next()
	if token.Text != "this" {
		t.Fail()
	}
	token = tokenizer.Next()
	//space
	if token.Text != " " {
		t.Fail()
	}
	token = tokenizer.Next()
	if token.Text != "is" {
		t.Fail()
	}
}
