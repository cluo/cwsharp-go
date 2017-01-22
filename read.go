package cwsharp

import (
	"io"

	"github.com/SteelSeries/bufrr"
)

type reader interface {
	ReadRune() (rune, int, error)
	PeekRune() (rune, int, error)
	UnreadRune() error
}

func newReader(r io.Reader) reader {
	return bufrr.NewReader(r)
}
