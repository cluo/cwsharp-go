package cwsharp

import "io"

func bigramTokenize(w writer, r reader) *Token {
	w.Reset()
	ch, _, err := r.PeekRune()
	if ch == -1 || err == io.EOF {
		return nil
	}

	if isCjk(ch) {
		w.WriteRune(ch)
		r.ReadRune()
		c2, _, _ := r.PeekRune()
		if isCjk(ch) {
			w.WriteRune(c2)
			// make sure the next chinese following chinese,
			// if follow not chinese,ignored it.
			r.ReadRune()
			if c3, _, _ := r.PeekRune(); isCjk(c3) {
				r.UnreadRune()
			}
		}
		return &Token{Text: w.String(), Type: WORD}
	}
	return whitespaceTokenize(w, r)
}

// WhitespaceTokenize tokenizes a specified text reader
// on the N-grams token algorithms.
func BigramTokenize(r io.Reader) Iterator {
	buf := make([]byte, 8)
	w := newWriter(buf)
	rr := newReader(r)
	return IteratorFunc(func() *Token {
		return bigramTokenize(w, rr)
	})
}
