// 基于词典的分词，关于此算法详细请查看：http://technology.chtsai.org/mmseg/

package mmseg

import (
	"fmt"
	"github.com/zhengchun/cwsharp-go/cwsharp"
	"github.com/zhengchun/cwsharp-go/cwsharp/simple"
	"io"
	"net/http"
	"os"
	"strings"
)

type Tokenizer struct {
	dawg *dawg
}

type Token struct {
	text string
	kind cwsharp.Kind
}

type TokenizerIterator struct {
	p      *Tokenizer
	inner  cwsharp.TokenIterator
	reader cwsharp.Reader
	cur    cwsharp.Token
}

func (t *Tokenizer) init(r io.Reader) {
	c := dawgCoder{DawgFileVersion}
	t.dawg = c.Decode(r)
}

func (t *Tokenizer) Traverse(r cwsharp.Reader) cwsharp.TokenIterator {
	return &TokenizerIterator{
		p:      t,
		inner:  simple.New().Traverse(r),
		reader: r,
	}
}

func (t *Token) Text() string {
	return t.text
}

func (t *Token) Kind() cwsharp.Kind {
	return t.kind
}

func (t *Token) String() string {
	return fmt.Sprintf("%s : %s", t.text, t.kind)
}

var filters = []func([]chunk, int) int{lawlFunc, svwlFunc, lawlFunc}

func (iter *TokenizerIterator) Next() bool {
	offset := iter.reader.Pos()
	c := iter.reader.Peek()
	if c == cwsharp.EOF {
		return false
	}
	dg := iter.p.dawg
	if node, ok := dg.root.Next(c); !ok || !node.HasChilds() {
		return iter.nextInternal()
	}
	firstOfNodes := iter.MatchedNodes(offset)
	if len(firstOfNodes) == 0 {
		iter.reader.Seek(offset)
		return iter.nextInternal()
	}
	maxLength := 0
	chunks := make([]chunk, 0)
	for i := len(firstOfNodes) - 1; i >= 0; i-- {
		//
		offset1 := offset + int(firstOfNodes[i].depth) + 1
		secondOfNodes := iter.MatchedNodes(offset1)
		if len(secondOfNodes) > 0 {
			for j := len(secondOfNodes) - 1; j >= 0; j-- {
				offset2 := offset1 + int(secondOfNodes[j].depth) + 1
				thirdOfNodes := iter.MatchedNodes(offset2)
				if len(thirdOfNodes) > 0 {
					for k := len(thirdOfNodes) - 1; k >= 0; k-- {
						offset3 := offset2 + int(thirdOfNodes[k].depth) + 1
						length := offset3 - offset
						if length >= maxLength {
							maxLength = length
							chunk := chunk{length, []wordPoint{
								wordPoint{offset, offset1 - offset, int(firstOfNodes[i].frequency)},
								wordPoint{offset1, offset2 - offset1, int(secondOfNodes[j].frequency)},
								wordPoint{offset2, offset3 - offset2, int(thirdOfNodes[k].frequency)},
							}}
							chunks = append(chunks, chunk)
						}
					}
				} else {
					length := offset2 - offset
					if length > maxLength {
						maxLength = length
						chunk := chunk{length, []wordPoint{
							wordPoint{offset, offset1 - offset, int(firstOfNodes[i].frequency)},
							wordPoint{offset1, offset2 - offset1, int(secondOfNodes[j].frequency)},
						}}
						chunks = append(chunks, chunk)
					}
				}
			}
		} else {
			length := offset1 - offset
			if length > maxLength {
				maxLength = length
				chunk := chunk{length, []wordPoint{
					wordPoint{offset, offset1 - offset, int(firstOfNodes[i].frequency)},
				}}
				chunks = append(chunks, chunk)
			}
		}
	}
	if len(chunks) > 1 {
		count := len(chunks)
		for _, filter := range filters {
			count = filter(chunks, count)
			if count == 1 {
				break
			}
		}
	}
	bestChunk := chunks[0]
	iter.reader.Seek(bestChunk.WordPoints[0].Offset)
	runes := make([]rune, bestChunk.WordPoints[0].Length)
	for i := 0; i < len(runes); i++ {
		runes[i] = iter.reader.ReadRule()
	}
	iter.cur = &Token{string(runes), cwsharp.CJK}
	return true
}

func (iter *TokenizerIterator) Cur() cwsharp.Token {
	return iter.cur
}

func (iter *TokenizerIterator) nextInternal() bool {
	if iter.inner.Next() {
		iter.cur = iter.inner.Cur()
		return true
	}
	return false
}

// uri参数可以是本地的文件路径或者是绝对的HTTP地址
func New(uri string) cwsharp.Tokenizer {
	t := &Tokenizer{}
	if strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://") {
		resp, err := http.Get(uri)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		t.init(resp.Body)
	} else {
		f, err := os.Open(uri)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		t.init(f)
	}
	return t
}

func (iter *TokenizerIterator) MatchedNodes(offset int) []*dawgNode {
	nodes := make([]*dawgNode, 0)
	iter.reader.Seek(offset)
	node := iter.p.dawg.root
	for code := iter.reader.ReadRule(); code != cwsharp.EOF; code = iter.reader.ReadRule() {
		node, _ = node.Next(code)
		if node == nil {
			break
		}
		if node.eow {
			nodes = append(nodes, node)
		}
	}
	return nodes

}
