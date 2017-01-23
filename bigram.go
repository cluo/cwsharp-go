package cwsharp

import "io"

func bigramTokenize(w writer, r reader) (Token, bool) {
	w.Reset()
	ch, _, err := r.PeekRune()
	if ch == -1 || err == io.EOF {
		return tokenEOF, false
	}

	if determineType(ch) == cjk {
		w.WriteRune(ch)
		r.ReadRune()
		c2, _, _ := r.PeekRune()
		if determineType(c2) == cjk {
			w.WriteRune(c2)
			// make sure the next chinese following chinese,
			// if follow not chinese,ignored it.
			r.ReadRune()
			if c3, _, _ := r.PeekRune(); determineType(c3) == cjk {
				r.UnreadRune()
			}
		}
		return Token{Text: w.String(), Type: WORD}, true
	}
	return whitespaceTokenize(w, r)
}

// WhitespaceTokenize tokenizes a specified text reader
// on the N-grams token algorithms.
func BigramTokenize(r io.Reader) Iterator {
	buf := make([]byte, 8)
	w := newWriter(buf)
	rr := newReader(r)
	return func() (Token, bool) {
		return bigramTokenize(w, rr)
	}
}
