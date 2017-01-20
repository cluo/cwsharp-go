package cwsharp

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/zhengchun/cwsharp-go/dawg"
)

type stdTokenizer struct {
	dawg *dawg.Dawg
}

func (t *stdTokenizer) tokenize(r *reader) Iterator {
	eof := func() (Token, bool) { return Token{}, false }

	return func() (Token, bool) {
		c, err := r.Read()
		if err == io.EOF {
			return eof()
		}

		r.buf = append(r.buf, c)
		if node := t.dawg.Root.Next(c); node == nil || !node.HasChilds() {
			return whitespaceTokenize(r)
		}
		return eof()
	}
}

func (t *stdTokenizer) Tokenize(r io.Reader) Iterator {
	r2 := &reader{
		rd:  bufio.NewReader(r),
		buf: make([]rune, 0),
	}
	return t.tokenize(r2)
}

func loadDawg(file string) (*dawg.Dawg, error) {
	r, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return dawg.NewDecoder(r).Decode()
}

// NewStandard returns a standard tokenizer using a specified
// lexicon file.
func NewStandard(file string) (Tokenizer, error) {
	dawg, err := loadDawg(file)
	if err != nil {
		return nil, fmt.Errorf("cwsharp:load dawg file got error(%v)", err)
	}
	return &stdTokenizer{dawg: dawg}, nil
}
