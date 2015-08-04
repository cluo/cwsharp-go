// Copyright (c) CWSharp. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package cwsharp

import (
	"testing"
)

func TestBigramTokenizer(t *testing.T) {
	var tokenizer = NewBigramTokenizer(true)
	iterator := tokenizer.Traverse("红酒柜")
	for token, next := iterator(); next != nil; token, next = iterator() {
		t.Log(token.Text)
	}
	//红酒,酒柜
	iterator = tokenizer.Traverse("一次性交一百元")
	for token, next := iterator(); next != nil; token, next = iterator() {
		t.Log(token.Text)
	}
}
