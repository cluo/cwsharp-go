package cwsharp

import (
	"bytes"
	"io"
)

var tokenEOF = Token{Type: EOF}

func whitespaceTokenize(b *bytes.Buffer, r reader) (Token, bool) {
	b.Reset()
	var (
		ch  rune
		err error
		typ Type
	)
	if ch, _, err = r.ReadRune(); ch == -1 || err == io.EOF {
		return tokenEOF, false
	}
	b.WriteRune(ch)
	switch determineType(ch) {
	case PUNC:
		typ = PUNC
	case NUMBER:
		typ = scanNumber(b, r)
	case latinAlpha:
		scanLetters(b, r)
		typ = WORD
	case WORD, cjk:
		typ = WORD
	}
	return Token{Text: b.String(), Type: typ}, true
}

func scanLetters(b *bytes.Buffer, r reader) {
	for {
		ch, _, err := r.ReadRune()
		if ch == -1 || err == io.EOF {
			goto exit
		}
		switch determineType(ch) {
		case PUNC:
			r.UnreadRune()
			goto exit
		case NUMBER, latinAlpha:
			b.WriteRune(ch)
		case WORD:
			r.UnreadRune()
			goto exit
		}
	}
exit:
}

func scanNumber(b *bytes.Buffer, r reader) (typ Type) {
	typ = NUMBER
	for {
		ch, _, err := r.ReadRune()
		if ch == -1 || err == io.EOF {
			goto exit
		}
		switch determineType(ch) {
		case PUNC:
			// check number  is float type?
			r.UnreadRune()
			goto exit
		case latinAlpha:
			b.WriteRune(ch)
			typ = WORD
		case NUMBER:
			b.WriteRune(ch)
		case WORD:
			r.UnreadRune()
			goto exit
		}
	}
exit:
	return
}

// WhitespaceTokenize tokenizes a specified text reader
// on the whitespace token algorithms.
func WhitespaceTokenize(r io.Reader) Iterator {
	buff := make([]byte, 8)
	b := bytes.NewBuffer(buff)
	rr := newReader(r)
	return func() (Token, bool) {
		return whitespaceTokenize(b, rr)
	}
}
