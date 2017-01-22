package cwsharp

import (
	"bufio"
	"bytes"
	"io"
)

/*
type scanner struct {
	rd     *bufio.Reader
	buf    []rune
	bufPos int
	bufLen int
}

// checkBuff check the buffer capacit and expand size.
func (s *scanner) checkBuff() {
	if size := s.bufLen; size == len(s.buf) {
		buf := make([]rune, size<<1)
		copy(buf, s.buf)
		s.buf = buf
	}
}

func (s *scanner) read() (r rune, err error) {
	if s.bufPos < s.bufLen {
		r = s.buf[s.bufPos]
		s.bufPos++
		return
	}
	r, _, err = s.rd.ReadRune()
	s.checkBuff()

	s.buf[s.bufLen] = r
	s.bufLen++
	s.bufPos++
	return
}

func (s *scanner) peek() (r rune, err error) {
	if s.bufPos < s.bufLen {
		r = s.buf[s.bufPos]
		return
	}
	r, _, err = s.rd.ReadRune()
	s.checkBuff()

	s.buf[s.bufLen] = r
	s.bufLen++
	return
}


func scanDigitalOrAlphabets(s *scanner) string {
	var (
		size int
	)
	for {
		r, err := s.read()
		if err == io.EOF {
			break
		}
		if kind := DeterminedKind(r); kind != 1 {
			break
		}
		size++
	}
	text := string(s.buf[:size])
	s.buf[0] = s.buf[s.bufLen-1]
	s.bufPos = 0
	s.bufLen = s.bufLen - size
	return text
}

*/

func whitespaceTokenize(b *bytes.Buffer, r *bufio.Reader) Token {
	b.Reset()
	var (
		ch  rune
		err error
		typ Type
	)
	if ch, _, err = r.ReadRune(); err == io.EOF {
		return TokenEOF
	}
	b.WriteRune(ch)
	switch DetermineType(ch) {
	case PUNC:
		typ = PUNC
	case NUMBER:
		typ = scanDigitals(b, r)
	case alphabet:
		scanAlphabets(b, r)
		typ = WORD
	case WORD:
		typ = WORD
	}
	return Token{Text: b.String(), Type: typ}
}

func scanAlphabets(b *bytes.Buffer, r *bufio.Reader) {
	for {
		ch, _, err := r.ReadRune()
		if err == io.EOF {
			break
		}
		switch DetermineType(ch) {
		case PUNC:
			r.UnreadRune()
			return
		case alphabet:
			b.WriteRune(ch)
		case WORD:
			r.UnreadRune()
			return
		}
	}
}

func scanDigitals(b *bytes.Buffer, r *bufio.Reader) (typ Type) {
	typ = NUMBER
	for {
		ch, _, err := r.ReadRune()
		if err == io.EOF {
			break
		}
		switch DetermineType(ch) {
		case PUNC:
			r.UnreadRune()
			return
		case alphabet:
			b.WriteRune(ch)
			typ = WORD
		case WORD:
			r.UnreadRune()
			return
		}
	}
	return
}

// WhitespaceTokenize tokenizes a specified text reader
// on the whitespace token algorithms.
func WhitespaceTokenize(r io.Reader) Iterator {
	buf := make([]byte, 6)
	return func() Token {
		return whitespaceTokenize(bytes.NewBuffer(buf), bufio.NewReader(r))
	}
}
