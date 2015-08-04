// Copyright (c) CWSharp. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package cwsharp

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"sort"
)

var DawgFileVersion float32 = 1

type dawg struct {
	root *dawgNode
}

type dawgNode struct {
	char      rune
	childs    map[rune]*dawgNode
	depth     int32
	frequency int32
	eow       bool
	parent    *dawgNode
}

type dawgCoder struct {
	version float32
}

type dawgNodeSort struct {
	node *dawgNode
	num  int32
}

type dawgNodeSortList []dawgNodeSort

func (list dawgNodeSortList) Len() int {
	return len(list)
}

func (list dawgNodeSortList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list dawgNodeSortList) Less(i, j int) bool {
	return list[i].num < list[j].num
}

type dawgNodeIterator func() (node *dawgNode, next dawgNodeIterator)

//返回当前节点下包含指定字符的子节点
func (this *dawgNode) Next(char rune) (*dawgNode, bool) {
	value, ok := this.childs[char]
	return value, ok
}

func (this *dawgNode) HasChilds() bool {
	if len(this.childs) == 0 {
		return false
	}
	return true
}

//添加子节点到当前节点
func (this *dawgNode) AddChild(node *dawgNode) *dawgNode {
	if value, ok := this.childs[node.char]; ok {
		return value
	}
	if node.parent == nil {
		node.parent = this
	}
	this.childs[node.char] = node
	return node
}

func (this *dawgNode) RemoveChild(node *dawgNode) bool {
	if _, ok := this.childs[node.char]; !ok {
		return false
	}
	//delete this node
	delete(this.childs, node.char)
	return true
}

func (this *dawgNode) Descendants(self bool) dawgNodeIterator {
	type entry struct {
		node *dawgNode
	}
	pop := func(stack []entry) ([]entry, *dawgNode) {
		if len(stack) <= 0 {
			return stack, nil
		} else {
			e := stack[0]
			return stack[1:len(stack)], e.node
		}
	}
	stack := make([]entry, 0)
	if self {
		stack = append(stack, entry{node: this})
	}
	for _, item := range this.childs {
		stack = append(stack, entry{node: item})
	}
	var iterator dawgNodeIterator
	iterator = func() (*dawgNode, dawgNodeIterator) {
		var node *dawgNode
		stack, node = pop(stack)
		if node == nil {
			return nil, nil
		}
		for _, next_node := range node.childs {
			stack = append(stack, entry{node: next_node})
		}
		return node, iterator
	}
	return iterator
}

func (dawgNode dawgNode) String() string {
	output := fmt.Sprintf("%s[%d]:%d", string(dawgNode.char), dawgNode.frequency, dawgNode.depth)
	if dawgNode.eow {
		output = output + "[EOW]"
	}
	return output
}

//查找是否包含指定的词组
func (this *dawg) Contains(word string) bool {
	if len(word) == 0 {
		return false
	}
	runes := []rune(word)
	nextNode, ok := this.root.Next(runes[0])
	i := 1
	for ok && i < len(runes) {
		nextNode, ok = nextNode.Next(runes[i])
		i = i + 1
	}

	return nextNode != nil && nextNode.eow
}

//查找包含指定前缀起始的词组
func (this *dawg) MatchsPrefix(prefix string) map[string]int32 {
	result := make(map[string]int32)
	nextNode := this.root
	runes := []rune(prefix)
	for i := 0; nextNode != nil && i < len(runes); i++ {
		nextNode, _ = this.root.Next(runes[i])
	}
	if nextNode == nil {
		return result
	}
	iterateNodesString(prefix, nextNode, result)
	return result
}

func iterateNodesString(commonPrefix string, node *dawgNode, words map[string]int32) {
	for _, childNode := range node.childs {
		nextCommonPrefix := commonPrefix + string(childNode.char)
		//at end of word
		if childNode.eow {
			words[nextCommonPrefix] = childNode.frequency
		}
		iterateNodesString(nextCommonPrefix, childNode, words)
	}
}

func getDawgNodeId(node *dawgNode) uint32 {
	var buffer bytes.Buffer
	for node, next := node.Descendants(true)(); next != nil; node, next = next() {
		buffer.WriteRune(node.char)
		if node.eow {
			buffer.WriteRune('1')
		} else {
			buffer.WriteRune('0')
		}
	}
	return fnvHash(buffer.Bytes())
}

func newDawgNode(char rune) *dawgNode {
	return &dawgNode{char: char, childs: make(map[rune]*dawgNode)}
}

func buildDawg(wordBag map[string]int32) *dawg {
	root := newDawgNode(rune(0))
	levelNodeCollections := make(map[int][]*dawgNode, 0)
	for word, freq := range wordBag {
		nextNode := root
		var currNode *dawgNode
		var found bool
		var level int = 0
		for _, char := range word {
			if currNode, found = nextNode.Next(char); !found {
				currNode = newDawgNode(char)
				currNode.depth = int32(level)
				nextNode.AddChild(currNode)
			}
			var nodes []*dawgNode
			if nodes, found = levelNodeCollections[level]; !found {
				nodes = make([]*dawgNode, 0)
				levelNodeCollections[level] = nodes
			}
			nodes = append(nodes, currNode)
			nextNode = currNode
			level++
		}
		nextNode.eow = true
		nextNode.frequency = freq
	}
	trackingNodes := make(map[*dawgNode]bool, 0)
	for j := len(levelNodeCollections) - 1; j >= 0; j = j - 1 {
		uniqNodeTables := make(map[uint32]*dawgNode, 0)
		for _, node := range levelNodeCollections[j] {
			_, ok := trackingNodes[node]
			if node.eow || ok {
				nodeId := getDawgNodeId(node)
				if findNode, ok := uniqNodeTables[nodeId]; ok {
					if node != findNode {
						node.parent.RemoveChild(node)
						node.parent.AddChild(node)
					}
					findNode.eow = findNode.eow || node.eow
					trackingNodes[node.parent] = true
					trackingNodes[findNode.parent] = true
				} else {
					uniqNodeTables[nodeId] = node
				}

			}
		}
	}
	return &dawg{root}
}

//保存到文件
func (coder *dawgCoder) Encode(w *bufio.Writer, dawg *dawg) {
	defer w.Flush()
	var count int32
	nodeLabels := make(map[*dawgNode]int32, 0)
	for node, next := dawg.root.Descendants(false)(); next != nil; node, next = next() {
		if _, ok := nodeLabels[node]; ok {
			continue
		}
		nodeLabels[node] = count
		count++
	}
	p := make(dawgNodeSortList, len(nodeLabels))
	i := 0
	for node, num := range nodeLabels {
		p[i] = dawgNodeSort{node, num}
		i++
	}
	sort.Sort(p)
	//the header of dawg.
	binary.Write(w, binary.LittleEndian, coder.version)
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
	binary.Write(w, binary.LittleEndian, int32(len(dawg.root.childs)))
	for _, node := range dawg.root.childs {
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
}

//读取文件
func (coder *dawgCoder) Decode(r *bufio.Reader) *dawg {
	var fileVersion float32
	binary.Read(r, binary.LittleEndian, &fileVersion)
	if coder.version != fileVersion {
		panic(errors.New(fmt.Sprintf("DAWG文件版本不一致.原始:v%f,当前:v%f", fileVersion, coder.version)))
	}
	var count int32
	binary.Read(r, binary.LittleEndian, &count)
	//read header of dawg file
	allNodes := make([]*dawgNode, count)
	for i := int32(0); i < count; i++ {
		node := &dawgNode{eow: false}
		var char uint16
		binary.Read(r, binary.LittleEndian, &char)
		node.char = rune(char)
		binary.Read(r, binary.LittleEndian, &node.frequency)
		binary.Read(r, binary.LittleEndian, &node.depth)
		var eow byte
		binary.Read(r, binary.LittleEndian, &eow)
		if eow == 1 {
			node.eow = true
		}
		var size int32
		binary.Read(r, binary.LittleEndian, &size)
		node.childs = make(map[rune]*dawgNode, size)
		allNodes[i] = node
	}
	//build a dawg.
	count = 0
	binary.Read(r, binary.LittleEndian, &count)
	root := &dawgNode{}
	root.childs = make(map[rune]*dawgNode, count)
	for i := int32(0); i < count; i++ {
		var j int32
		binary.Read(r, binary.LittleEndian, &j)
		node := allNodes[j]
		root.AddChild(node)
	}
	binary.Read(r, binary.LittleEndian, &count)
	for i := int32(0); i < count; i++ {
		var label int32
		binary.Read(r, binary.LittleEndian, &label)
		node := allNodes[label]
		var childNodeCount int32
		binary.Read(r, binary.LittleEndian, &childNodeCount)
		for j := int32(0); j < childNodeCount; j++ {
			binary.Read(r, binary.LittleEndian, &label)
			childNode := allNodes[label]
			node.AddChild(childNode)
		}
	}
	return &dawg{root}
}
