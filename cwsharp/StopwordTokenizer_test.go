// Copyright (c) CWSharp. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package cwsharp

import (
	"testing"
)

func TestStopWordTokenizer_1(t *testing.T) {
	stopwords := map[string]bool{
		"的": true, "the": true, "。": true,
	}
	file := InitDawgFile()
	//默认采用标准的分词
	defaultTokenizer := NewStandardTokenizer(file, true)
	swTokenizer := NewStopwordTokenizer(defaultTokenizer, stopwords)
	iterator := swTokenizer.Traverse("你是我的小苹果。")
	for token, next := iterator(); next != nil; token, next = next() {
		t.Log(token.String())
	}
	//>> 你，是，我，小，苹果
}
