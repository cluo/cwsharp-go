// Copyright (c) CWSharp. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package cwsharp

//空格,标点做为分隔符的分词器
type whiteSpaceTokenBreaker struct {
	reader *stringReader
}

var alphaNumStops = map[rune]bool{rune("#"[0]): false, rune("+"[0]): true, rune("-"[0]): true, rune("_"[0]): true}
var numStops = map[rune]bool{rune("."[0]): true}

func (this *whiteSpaceTokenBreaker) Next() Token {
	offset := this.reader.cursor
	code := this.reader.Read()
	if code == 0 {
		return token_empty
	}
	if isCjkCase(code) {
		return NewToken(string(code), TokenType_CJK)
	}
	if isLetterCase(code) {
		for code = this.reader.Read(); code != 0; code = this.reader.Read() {
			if isLetterCase(code) || isNumeralCase(code) {
				continue
			}
			if period, ok := alphaNumStops[code]; ok {
				if period {
					nextcode := this.reader.Peek()
					if _, ok = alphaNumStops[nextcode]; ok || isLetterCase(nextcode) || isNumeralCase(nextcode) {
						continue
					}
					break
				}
			}
			this.reader.Seek(this.reader.cursor - 1)
			break
		}
		length := this.reader.cursor - offset
		this.reader.Seek(offset)
		return NewToken(string(this.reader.ReadCount(length)), TokenType_ALPHANUM)
	} else if isNumeralCase(code) {
		mixed := false
		for code = this.reader.Read(); code != 0; code = this.reader.Read() {
			if isNumeralCase(code) || isLetterCase(code) || mixed {
				continue
			}
			if period, ok := numStops[code]; ok && period {
				nextCode := this.reader.Peek()
				if isNumeralCase(nextCode) {
					continue
				}
			}
			this.reader.Seek(this.reader.cursor - 1)
			break
		}
		length := this.reader.cursor - offset
		this.reader.Seek(offset)
		token_type := TokenType_NUM
		if mixed {
			token_type = TokenType_ALPHANUM
		}
		return NewToken(string(this.reader.ReadCount(length)), token_type)
	}
	return NewToken(string(code), TokenType_PUNC)
}
