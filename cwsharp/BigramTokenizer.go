// Copyright (c) CWSharp. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package cwsharp

//二元分词
type BigramTokenizer struct {
	OutputOriginalCase bool //是否保留大小写
}

type bigramTokenBreaker struct {
	whiteSpaceTokenBreaker
	beginState bool
}

func (this *bigramTokenBreaker) Next() Token {
	code := this.reader.Read()
	if code == 0 {
		return token_empty
	}
	if isLetterCase(code) || isNumeralCase(code) {
		this.reader.Seek(this.reader.cursor - 1)
		//called by parent class.
		return this.whiteSpaceTokenBreaker.Next()
	} else if isCjkCase(code) {
		nextCode := this.reader.Read()
		if nextCode == 0 {
			//结束位置
			if this.beginState {
				return NewToken(string(code), TokenType_CJK)
			}
			return token_empty
		}
		if isCjkCase(nextCode) {
			this.beginState = false
			if isCjkCase(this.reader.Peek()) {
				this.reader.Seek(this.reader.cursor - 1)
			}
			return NewToken(string([]rune{code, nextCode}), TokenType_CJK)
		}
		//may be code is a one of letter&numeral&punc.
		this.reader.Seek(this.reader.cursor - 2)
		return this.whiteSpaceTokenBreaker.Next()
	}
	this.beginState = true
	return NewToken(string(code), TokenType_PUNC)
}

func (this *BigramTokenizer) Traverse(text string) TokenIterator {
	var iterator TokenIterator
	reader := newStringReader(text, this.OutputOriginalCase)
	//二元分词器
	breaker := &bigramTokenBreaker{
		whiteSpaceTokenBreaker: whiteSpaceTokenBreaker{reader: reader},
	}
	iterator = func() (Token, TokenIterator) {
		token := breaker.Next()
		if token.IsNull() {
			return token, nil
		}
		return token, iterator
	}
	return iterator
}

func NewBigramTokenizer(ignoreCase bool) *BigramTokenizer {
	var tokenizer BigramTokenizer
	tokenizer.OutputOriginalCase = !ignoreCase
	return &tokenizer
}
