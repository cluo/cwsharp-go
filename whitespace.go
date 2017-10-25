package cwsharp

import (
	"io"
	"unicode"
)

func whitespaceTokenize(w writer, r reader) *Token {
	w.Reset()
	var (
		ch  rune
		err error
		typ Type
	)
	ch, _, err = r.ReadRune()
	if ch == -1 || err == io.EOF {
		return nil
	}

	w.WriteRune(ch)

	switch determineType(ch) {
	case PUNCT:
		typ = PUNCT
	case NUMBER:
		typ = scanNumber(w, r)
	case ALPHA:
		typ = scanLetters(w, r)
	case WORD:
		typ = WORD
	}
	return &Token{Text: w.String(), Type: typ}
}

func scanLetters(w writer, r reader) (typ Type) {
	typ = ALPHA
	for {
		ch, _, err := r.ReadRune()
		if ch == -1 || err == io.EOF {
			goto exit
		}
		switch determineType(ch) {
		case PUNCT:
			r.UnreadRune()
			goto exit
		case NUMBER:
			w.WriteRune(ch)
			typ = WORD
		case ALPHA:
			w.WriteRune(ch)
		default:
			r.UnreadRune()
			goto exit
		}
	}
exit:
	return
}

func scanNumber(w writer, r reader) (typ Type) {
	typ = NUMBER
	for {
		ch, _, err := r.ReadRune()
		if ch == -1 || err == io.EOF {
			goto exit
		}
		switch determineType(ch) {
		case PUNCT:
			// check number is float type?
			if c2, _, _ := r.PeekRune(); ch == '.' && unicode.IsNumber(c2) {
				w.WriteRune(ch)
			} else {
				r.UnreadRune()
				goto exit
			}
		case ALPHA:
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

	return IteratorFunc(func() *Token {
		return whitespaceTokenize(w, rr)
	})
}
