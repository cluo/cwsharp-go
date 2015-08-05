// Copyright (c) CWSharp. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package cwsharp

import (
	"bufio"
	"net/http"
	"os"
	"strings"
)

//标准的中文分词器
type StandardTokenizer struct {
	dawg               *dawg
	OutputOriginalCase bool
}

func (this *StandardTokenizer) Traverse(text string) TokenIterator {
	reader := newStringReader(text, this.OutputOriginalCase)
	breaker := &maximumMatchTokenBreaker{
		whiteSpaceTokenBreaker: whiteSpaceTokenBreaker{reader: reader},
		dawg: this.dawg,
	}
	var iterator TokenIterator
	iterator = func() (Token, TokenIterator) {
		token := breaker.Next()
		if token.IsNull() {
			return token, nil
		}
		return token, iterator
	}
	return iterator
}

func (this *StandardTokenizer) init(file string, ignoreCase bool) {
	file = strings.ToLower(file)
	var r *bufio.Reader
	if strings.Index(file, "http://") >= 0 || strings.Index(file, "https://") >= 0 {
		response, err := http.Get(file)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		r = bufio.NewReader(response.Body)

	} else {
		f, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		r = bufio.NewReader(f)
	}
	coder := dawgCoder{DawgFileVersion}
	dawg := coder.Decode(r)
	this.dawg = dawg
	this.OutputOriginalCase = !ignoreCase
}

func NewStandardTokenizer(file string, ignoreCase bool) *StandardTokenizer {
	var tokenizer StandardTokenizer
	tokenizer.init(file, ignoreCase)
	return &tokenizer
}
