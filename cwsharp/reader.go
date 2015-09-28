package cwsharp

import (
	"errors"
	"fmt"
)

var EOF rune = -1

type Reader interface {
	ReadRule() rune
	ReadRules(count int) []rune
	Peek() rune
	Seek(offset int)
	Pos() int
}

type buffReader struct {
	src []rune
	off int

	Reader
}

func (r *buffReader) Init(src []rune) {
	r.src = src
	r.off = -1
}

func (r *buffReader) ReadRule() rune {
	ch := r.Peek()
	if ch != EOF {
		r.off++
	}
	return ch
}

func (r *buffReader) Peek() rune {
	if r.off+1 >= len(r.src) {
		return EOF
	}
	return r.src[r.off+1]
}

func (r *buffReader) ReadRules(count int) []rune {
	s := make([]rune, count)
	for i := 0; i < count; i++ {
		s[i] = r.ReadRule()
	}
	return s
}

func (r *buffReader) Seek(offset int) {
	if offset < 0 || offset >= len(r.src) {
		panic(errors.New(fmt.Sprintf("offset<0 || offset>=%d", len(r.src))))
	}
	r.off = offset
}

func (r *buffReader) Pos() int {
	return r.off
}

func NewStringReader(src string) Reader {
	r := &buffReader{}
	r.Init([]rune(src))
	return r
}

type streamReader struct {
	Reader
}
