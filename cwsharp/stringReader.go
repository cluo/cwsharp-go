// Copyright (c) CWSharp. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package cwsharp

import (
	"errors"
)

type stringReader struct {
	s          []rune
	ignoreCase bool
	length     int
	cursor     int
}

func (this *stringReader) Read() rune {
	if this.cursor >= this.length {
		return 0
	}
	ch := rune(this.s[this.cursor])
	this.cursor++
	if !this.ignoreCase {
		return this.Normalize(ch)
	}
	return ch
}

func (this *stringReader) ReadCount(count int) []rune {
	if this.cursor+count > this.length {
		panic(errors.New("The read count is arrived the length of this stream reader."))
	}
	array := make([]rune, count)
	for i := 0; i < count; i++ {
		array[i] = this.Read()
	}
	return array
}

func (this *stringReader) Seek(offset int) {
	this.cursor = offset
}

func (this *stringReader) Peek() rune {
	if this.cursor >= this.length {
		return 0
	}
	ch := rune(this.s[this.cursor])
	if !this.ignoreCase {
		return ch
	}
	return this.Normalize(ch)
}

func (this *stringReader) Normalize(ch rune) rune {
	if isUpperCase(ch) {
		return ch + 32
	} else if ch >= 0xff01 && ch <= 0xff5d {
		return ch - 0xFEE0
	}
	return ch
}

func (this *stringReader) Position() int {
	return this.cursor
}

func newStringReader(text string, ignoreCase bool) *stringReader {
	var s = []rune(text)
	return &stringReader{s: []rune(text), length: len(s), ignoreCase: ignoreCase, cursor: 0}
}
