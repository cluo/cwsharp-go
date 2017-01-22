package cwsharp

/*
import (
	"fmt"
	"io"
	"os"

	"github.com/zhengchun/cwsharp-go/dawg"
)


type mmsegTokenizer struct {
	dawg *dawg.Dawg
}

func (t *mmsegTokenizer) tokenize(r *bufferedReader) Iterator {
	return func() (Token, bool) {
		c, err := r.Peek()
		if err == io.EOF {
			return eof()
		}

		if node := t.dawg.Root.Next(c); node == nil || !node.HasChilds() {
			return whitespaceTokenize(r)
		}
		return eof()
	}
}

func (t *mmsegTokenizer) Tokenize(r io.Reader) Iterator {
	//return t.tokenize(newReader(r))
}

func loadDawg(file string) (*dawg.Dawg, error) {
	r, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return dawg.NewDecoder(r).Decode()
}

// New returns a standard tokenizer using a specified
// lexicon file.
func New(file string) (Tokenizer, error) {
	dawg, err := loadDawg(file)
	if err != nil {
		return nil, fmt.Errorf("cwsharp:load dawg file got error(%v)", err)
	}
	return &mmsegTokenizer{dawg: dawg}, nil
}
*/
