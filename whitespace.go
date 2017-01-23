package cwsharp

import (
	"io"
	"unicode"
)

var tokenEOF = Token{Type: EOF}

func whitespaceTokenize(w writer, r reader) (Token, bool) {
	w.Reset()
	var (
		ch  rune
		err error
		typ Type
	)
	if ch, _, err = r.ReadRune(); ch == -1 || err == io.EOF {
		return tokenEOF, false
	}
	w.WriteRune(ch)
	switch determineType(ch) {
	case PUNC:
		typ = PUNC
	case NUMBER:
		typ = scanNumber(w, r)
	case latinAlpha:
		scanLetters(w, r)
		typ = WORD
	case WORD, cjk:
		typ = WORD
	}
	return Token{Text: w.String(), Type: typ}, true
}

func scanLetters(w writer, r reader) {
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
			w.WriteRune(ch)
		default:
			r.UnreadRune()
			goto exit
		}
	}
exit:
}

func scanNumber(w writer, r reader) (typ Type) {
	typ = NUMBER
	for {
		ch, _, err := r.ReadRune()
		if ch == -1 || err == io.EOF {
			goto exit
		}
		switch determineType(ch) {
		case PUNC:
			// check number is float type?
			if c2, _, _ := r.PeekRune(); ch == '.' && unicode.IsNumber(c2) {
				w.WriteRune(ch)
			} else {
				r.UnreadRune()
				goto exit
			}
		case latinAlpha:
			w.WriteRune(ch)
			typ = WORD
		case NUMBER:
			w.WriteRune(ch)
		default: // WORD, cjk
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
	buf := make([]byte, 8)
	w := newWriter(buf)
	rr := newReader(r)

	return func() (Token, bool) {
		return whitespaceTokenize(w, rr)
	}
}
