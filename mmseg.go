package cwsharp

import (
	"fmt"
	"io"
	"math"
	"os"

	"github.com/zhengchun/cwsharp-go/dawg"
)

// mmseg algorithms at http://technology.chtsai.org/mmseg/

type mmsegTokenizer struct {
	dawg *dawg.Dawg
}

type runeBuf struct {
	buf []rune
}

type runeCacheReader struct {
	rr  reader
	buf []rune
}

func (m *mmsegTokenizer) cjkTokenize(b *runeBuf, w writer, r reader) (Token, bool) {
	nodes_1 := m.matchedNodes(b, 0, r)
	if len(nodes_1) == 0 {
		return whitespaceTokenize(w, r)
	}
	var (
		maxWordLength int
		chunks        []chunk
		offset        int
	)
	for i := len(nodes_1) - 1; i >= 0; i-- {
		offset_1 := offset + int(nodes_1[i].Depth()) + 1
		nodes_2 := m.matchedNodes(b, offset_1, r)
		if len(nodes_2) > 0 {

			for j := len(nodes_2) - 1; j >= 0; j-- {
				offset_2 := offset_1 + nodes_2[j].Depth() + 1
				nodes_3 := m.matchedNodes(b, offset_2, r)
				if len(nodes_3) > 0 {

					for k := len(nodes_3) - 1; k >= 0; k-- {
						offset_3 := offset_2 + nodes_3[k].Depth() + 1
						wordLength := offset_3 - offset
						if wordLength >= maxWordLength {
							maxWordLength = wordLength
							chunk := chunk{wordLength, []wordPoint{
								wordPoint{offset, offset_1 - offset, nodes_1[i].Freq()},
								wordPoint{offset_1, offset_2 - offset_1, nodes_2[j].Freq()},
								wordPoint{offset_2, offset_3 - offset_2, nodes_3[k].Freq()},
							}}
							chunks = append(chunks, chunk)
						}
					}
				} else {
					wordLength := offset_2 - offset
					if wordLength > maxWordLength {
						maxWordLength = wordLength
						chunk := chunk{wordLength, []wordPoint{
							wordPoint{offset, offset_1 - offset, nodes_1[i].Freq()},
							wordPoint{offset_1, offset_2 - offset_1, nodes_2[j].Freq()},
						}}
						chunks = append(chunks, chunk)
					}
				}
			}
		} else {
			wordLength := offset_1 - offset
			if wordLength > maxWordLength {
				maxWordLength = wordLength
				chunk := chunk{wordLength, []wordPoint{
					wordPoint{offset, offset_1 - offset, nodes_1[i].Freq()},
				}}
				chunks = append(chunks, chunk)
			}
		}
	}

	chunk := selectChunk(chunks)
	//fmt.Println(string(b.buf))

	length := chunk.WordPoints[0].Length
	tok := Token{Text: string(b.buf[:length]), Type: WORD}
	b.buf = b.buf[length:]
	return tok, true
}

func (m *mmsegTokenizer) tokenize(b *runeBuf, w writer, r reader) (Token, bool) {
	w.Reset()

	// if b still have a any rune we can skip into cjk tokenize.
	if len(b.buf) == 0 {
		ch, _, err := r.PeekRune()
		if ch == -1 || err == io.EOF {
			return tokenEOF, false
		}

		// checks specified char has in the lexicon table.
		if node := m.dawg.Root.Next(ch); node == nil || !node.HasChilds() {
			return whitespaceTokenize(w, r)
		}
	}

	return m.cjkTokenize(b, w, r)
}

var filters = []func([]chunk) []chunk{lawlFunc, svwlFunc, lawlFunc}

func selectChunk(s []chunk) chunk {
	for _, fn := range filters {
		if len(s) > 1 {
			s = fn(s)
		}
	}
	return s[0]
}

func (m *mmsegTokenizer) matchedNodes(b *runeBuf, bufPos int, r reader) []*dawg.Node {
	var nodes []*dawg.Node
	node := m.dawg.Root
	for i := bufPos; i < len(b.buf); i++ {
		ch := b.buf[i]
		node = node.Next(ch)
		if node == nil {
			return nodes
		}
		if node.EOW() {
			nodes = append(nodes, node)
		}
	}
	// if the runebuf all matched,try read a rune from reader .
	for {
		ch, _, err := r.ReadRune()
		if ch == -1 || err == io.EOF {
			break
		}
		node = node.Next(ch)
		if node == nil {
			r.UnreadRune()
			break
		}
		b.buf = append(b.buf, ch)
		if node.EOW() {
			nodes = append(nodes, node)
		}
	}
	return nodes
}

var lawlFunc = func(s []chunk) []chunk {
	var max float64
	b := s[:0]
	for _, c := range s {
		v := c.WordAverageLength()
		switch {
		case v > max:
			b = s[:0]
			b = append(b, c)
			max = v
		case v == max:
			b = append(b, c)
		}
	}
	return b
}

var svwlFunc = func(s []chunk) []chunk {
	var min float64 = -1
	b := s[:0]
	for _, c := range s {
		v := c.Variance()
		switch {
		case min == -1 || v < min:
			b := s[:0]
			b = append(b, c)
			min = v
		case v == min:
			b = append(b, c)
		}
	}
	return b
}

var lsdmfocwFunc = func(s []chunk) []chunk {
	var max float64
	b := s[:0]
	for _, c := range s {
		v := c.Degree()
		switch {
		case v > max:
			b := s[:0]
			b = append(b, c)
			max = v
		case v == max:
			b = append(b, c)
		}
	}
	return b
}

type wordPoint struct {
	Offset, Length, Freq int
}

type chunk struct {
	Length     int // word length
	WordPoints []wordPoint
}

func (c *chunk) Get(index int) wordPoint {
	return c.WordPoints[index]
}

func (c *chunk) WordAverageLength() float64 {
	return float64(c.Length) / float64(len(c.WordPoints))
}

func (c *chunk) Variance() float64 {
	var sum float64
	averageLength := c.WordAverageLength()
	for _, wordPoint := range c.WordPoints {
		sum = sum + (math.Pow(float64(wordPoint.Length)-averageLength, 2))
	}
	return math.Sqrt(sum / float64(len(c.WordPoints)))
}

func (c *chunk) Degree() float64 {
	var sum float64
	for _, wordPoint := range c.WordPoints {
		sum = sum + math.Log10(float64(wordPoint.Freq))
	}
	return sum
}

// Tokenize is tokenizes the specifed Reader into individ tokens.
func (t *mmsegTokenizer) Tokenize(r io.Reader) Iterator {
	buf := make([]byte, 8)
	w := newWriter(buf)
	b := new(runeBuf)
	rr := newReader(r)
	return func() (Token, bool) {
		return t.tokenize(b, w, rr)
	}
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
