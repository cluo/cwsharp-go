//

package simpletokenizer

import (
	"bufio"
	"io"

	"github.com/zhengchun/cwsharp-go/cwsharp"
)

type tokenizer struct{}

type iterator struct {
	r   *bufio.Reader
	buf []rune

	len, pos int
}

func (t *tokenizer) Tokenize(r io.Reader) cwsharp.Iterator {
	return &iterator{
		r:   bufio.NewReader(r),
		buf: make([]rune, 8),
	}
}

// resize the buffer capacity.
func (iter *iterator) resize(size int) {
	if size == len(iter.buf) {
		buf := make([]rune, size*2)
		copy(buf, iter.buf)
		iter.buf = buf
	}
}

func (iter *iterator) Next() (cwsharp.Token, bool) {
	if iter.pos < iter.len {
		r := iter.buf[iter.pos]
		token := cwsharp.Token{Text: string(r), Kind: cwsharp.DetermineKind(r)}
		iter.pos++
		return token, true
	}

	var kind cwsharp.Kind
	var length, size int
	for {
		r, _, err := iter.r.ReadRune()
		if err == io.EOF {
			goto exit
		}

		iter.resize(length)
		iter.buf[length] = r
		length++

		switch cwsharp.DetermineKind(r) {
		case cwsharp.PUNCT:
			goto exit
		case cwsharp.LETTER:
			kind |= cwsharp.LETTER
			size++
		case cwsharp.NUMERAL:
			kind |= cwsharp.NUMERAL
			size++
		case cwsharp.CJK:
			goto exit
		default:
			goto exit
		}
	}

exit:

	if length == 0 {
		return cwsharp.Token{}, false
	}
	if length == 1 {
		size = 1
		kind = cwsharp.DetermineKind(iter.buf[0])
	}
	token := cwsharp.Token{Text: string(iter.buf[:size]), Kind: kind}
	iter.pos = size
	iter.len = length
	return token, true
}

func New() cwsharp.Tokenizer {
	return &tokenizer{}
}
