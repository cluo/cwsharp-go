// Copyright (c) CWSharp. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package cwsharp

//提供过滤词功能的分词器,扩展类
type StopwordTokenizer struct {
	defTokenizer Tokenizer
	data         map[string]bool
}

func (this *StopwordTokenizer) Traverse(text string) TokenIterator {
	var defaultIterator = this.defTokenizer.Traverse(text)

	var iterator TokenIterator

	iterator = func() (Token, TokenIterator) {
		for token, next := defaultIterator(); next != nil; token, next = defaultIterator() {
			if _, stopWord := this.data[token.Text]; stopWord {
				continue
			}
			return token, iterator
		}
		return token_empty, nil

	}
	return iterator
}

//default_tokenizer : 默认采用的
func NewStopwordTokenizer(default_tokenizer Tokenizer, stopwords map[string]bool) *StopwordTokenizer {
	if stopwords == nil {
		stopwords = make(map[string]bool, 0)
	}
	return &StopwordTokenizer{
		defTokenizer: default_tokenizer,
		data:         stopwords,
	}
}
