package dawg

import (
	"encoding/binary"
	"errors"
	"io"
)

type Decoder struct {
	r io.Reader
}

func (d *Decoder) Decode() (*Dawg, error) {
	var version float32
	binary.Read(d.r, binary.LittleEndian, &version)
	if version != 1 {
		return nil, errors.New("this verison of this dawg file is not compatible")
	}

	var count int32
	binary.Read(d.r, binary.LittleEndian, &count)
	nodes := make([]*Node, count)
	for i := int32(0); i < count; i++ {
		n := &Node{eow: false}
		var char uint16
		binary.Read(d.r, binary.LittleEndian, &char)
		n.char = rune(char)
		binary.Read(d.r, binary.LittleEndian, &n.frequency)
		binary.Read(d.r, binary.LittleEndian, &n.depth)
		var eow byte
		binary.Read(d.r, binary.LittleEndian, &eow)
		if eow == 1 {
			n.eow = true
		}
		var size int32
		binary.Read(d.r, binary.LittleEndian, &size)
		n.childs = make(map[rune]*Node, size)
		nodes[i] = n
	}

	//build a dawg.
	binary.Read(d.r, binary.LittleEndian, &count)
	root := &Node{}
	root.childs = make(map[rune]*Node, count)
	for i := int32(0); i < count; i++ {
		var j int32
		binary.Read(d.r, binary.LittleEndian, &j)
		root.addChild(nodes[j])
	}
	binary.Read(d.r, binary.LittleEndian, &count)

	for i := int32(0); i < count; i++ {
		var label int32
		binary.Read(d.r, binary.LittleEndian, &label)
		n := nodes[label]
		var childNodeCount int32
		binary.Read(d.r, binary.LittleEndian, &childNodeCount)
		for j := int32(0); j < childNodeCount; j++ {
			binary.Read(d.r, binary.LittleEndian, &label)
			n.addChild(nodes[label])
		}
	}
	return &Dawg{version, root}, nil
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r}
}
