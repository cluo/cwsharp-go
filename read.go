package cwsharp

import (
	"bufio"
	"io"
	"unicode"
)

type reader struct {
	buf  []rune
	rd   *bufio.Reader
	r, l int
}

func (b *reader) Read() (r rune, err error) {
	if i := r.r; i < len(r.buf) {
		c = r.buf[r.r]
		r.buf = r.buf[i:]
		r.r++
		return
	}
	r.r = 0
	r.buf = r.buf[:0]
	for {
		c, _, err = r.rd.ReadRune()
		if err == io.EOF {
			return
		}
		// ignored any chars \r \t \n
		if unicode.IsSpace(c) {
			continue
		}
		return
	}
}
