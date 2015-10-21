package cwsharp

import (
	"unicode"
	"bufio"
	"fmt"
	"io"
)

type bufReader struct {
	offset int
	buf    []rune

	src *bufio.Reader
}

func (b *bufReader) init(src io.Reader) {
	b.src = bufio.NewReader(src)
	b.offset = 0
	b.fill()
}

func NewReader(src io.Reader) Reader {
	b := &bufReader{}
	b.init(src)
	return b
}

func (b *bufReader) ReadRule() rune {
	if b.offset == len(b.buf) && !b.fill() {
		return EOF
	}
	r := b.buf[b.offset]
	b.offset += 1
	if r == 65279 { //BOM-utf8
		r= b.ReadRule()
	}
	return normalize(r)
}

func (b *bufReader) Peek() rune {
	if b.offset == len(b.buf) && !b.fill() {
		return EOF
	}
	r := b.buf[b.offset]
	if r == 65279 {
		b.ReadRule() //ignore next rune
		r= b.Peek()
	}
	return unicode.ToLower(r)
}

func (b *bufReader) Seek(offset int) {
	b.offset = offset
}

func (b *bufReader) Pos() int {
	return b.offset
}

func (b *bufReader) fill() bool {
	line, _ := b.src.ReadString('\n')
	if len(line) == 0 {
		return false
	}
	b.offset = 0
	b.buf = []rune(line)
	fmt.Println(b.buf)
	return true
}

func normalize(r rune) rune {
	if r >= 0x41 && r <= 0x5A { //A-Z
		return r + 32
	} else if r >= 0xff01 && r <= 0xff5d {
		return r - 0xFEE0
	}
	return r
}