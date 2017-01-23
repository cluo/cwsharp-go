package dawg

import (
	"encoding/binary"
	"io"
	"sort"
)

type Encoder struct {
	dawg *Dawg
}

func (e *Encoder) Encode(w io.Writer) error {
	var count int32
	var nodeLabels map[*Node]int32
	iter := e.dawg.Root.descendants(false)
	for node := iter(); node != nil; node = iter() {
		if _, ok := nodeLabels[node]; ok {
			continue
		}
		nodeLabels[node] = count
		count++
	}

	p := make(nodeSortList, len(nodeLabels))
	i := 0
	for node, num := range nodeLabels {
		p[i] = nodeSort{node, num}
		i++
	}
	sort.Sort(p)

	binary.Write(w, binary.LittleEndian, e.dawg.Version)
	binary.Write(w, binary.LittleEndian, count)
	count = 0
	for _, k := range p {
		node := k.node
		binary.Write(w, binary.LittleEndian, uint16(node.char))
		binary.Write(w, binary.LittleEndian, node.frequency)
		binary.Write(w, binary.LittleEndian, node.depth)
		var eow byte
		if eow = 0; node.eow {
			eow = 1
		}
		binary.Write(w, binary.LittleEndian, eow)
		binary.Write(w, binary.LittleEndian, int32(len(node.childs)))
		if node.HasChilds() {
			count = count + 1
		}
	}
	//the top node of dawg.
	binary.Write(w, binary.LittleEndian, int32(len(e.dawg.Root.childs)))
	for _, node := range e.dawg.Root.childs {
		label, _ := nodeLabels[node]
		binary.Write(w, binary.LittleEndian, label)
	}
	//the child node of dawg.
	binary.Write(w, binary.LittleEndian, count)
	for _, k := range p {
		node := k.node
		label := k.num
		if !node.HasChilds() {
			continue
		}
		binary.Write(w, binary.LittleEndian, label)
		binary.Write(w, binary.LittleEndian, int32(len(node.childs)))
		for _, sub_node := range node.childs {
			label, _ := nodeLabels[sub_node]
			binary.Write(w, binary.LittleEndian, label)
		}
	}
	return nil
}

type nodeSort struct {
	node *Node
	num  int32
}

type nodeSortList []nodeSort

func (list nodeSortList) Len() int {
	return len(list)
}

func (list nodeSortList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list nodeSortList) Less(i, j int) bool {
	return list[i].num < list[j].num
}

func NewEncoder(dawg *Dawg) *Encoder {
	return &Encoder{dawg}
}
