// Copyright (c) CWSharp. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package cwsharp

import (
	"testing"
)

func TestMaximumMatchTokenBreaker(t *testing.T) {
	dawg := CreateDawg()
	test := "中国世博园,EXPO 2008"
	reader := newStringReader(test, false)
	tokenizer := maximumMatchTokenBreaker{whiteSpaceTokenBreaker: whiteSpaceTokenBreaker{reader: reader}, dawg: dawg}
	for token := tokenizer.Next(); !token.IsNull(); token = tokenizer.Next() {
		t.Log(token)
	}
}
