package cwsharp

import (
	"bytes"
	"io"

	"github.com/SteelSeries/bufrr"
)

type writer interface {
	WriteRune(rune)
	String() string
	Reset()
}

type bufWriter struct {
	b *bytes.Buffer
}

func normalizeRune(r rune) rune {
	if r >= 0x41 && r <= 0x5A { //A-Z
		return r + 32
	} else if r >= 0xff01 && r <= 0xff5d {
		return r - 0xFEE0
	}
	return r
}

func (w *bufWriter) WriteRune(r rune) {
	w.b.WriteRune(normalizeRune(r))
}

func (w *bufWriter) String() string {
	return w.b.String()
}

func (w *bufWriter) Reset() {
	w.b.Reset()
}

func newWriter(buf []byte) writer {
	return &bufWriter{bytes.NewBuffer(buf)}
}

type reader interface {
	ReadRune() (rune, int, error)
	PeekRune() (rune, int, error)
	UnreadRune() error
}

func newReader(r io.Reader) reader {
	return bufrr.NewReader(r)
}
