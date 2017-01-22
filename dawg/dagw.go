package dawg

import (
	"bytes"
	"hash/fnv"
)

type Dawg struct {
	Version float32
	Root    *Node
}

type Node struct {
	char      rune
	childs    map[rune]*Node
	depth     int32
	frequency int32
	eow       bool
	parent    *Node
}

func (d *Dawg) Contains(word string) bool {
	if len(word) == 0 {
		return false
	}
	runes := []rune(word)
	nextNode := d.Root.Next(runes[0])
	i := 1
	for nextNode != nil && i < len(runes) {
		nextNode = nextNode.Next(runes[i])
		i = i + 1
	}

	return nextNode != nil && nextNode.eow
}

func (d *Dawg) MatchsPrefix(prefix string) map[string]int32 {
	result := make(map[string]int32)
	nextNode := d.Root
	runes := []rune(prefix)
	for i := 0; nextNode != nil && i < len(runes); i++ {
		nextNode = d.Root.Next(runes[i])
	}
	if nextNode == nil {
		return result
	}
	iterateNodesString(prefix, nextNode, result)
	return result
}

func iterateNodesString(commonPrefix string, node *Node, words map[string]int32) {
	for _, childNode := range node.childs {
		nextCommonPrefix := commonPrefix + string(childNode.char)
		//at end of word
		if childNode.eow {
			words[nextCommonPrefix] = childNode.frequency
		}
		iterateNodesString(nextCommonPrefix, childNode, words)
	}
}

func getDawgNodeId(node *Node) uint32 {
	var buffer bytes.Buffer
	var iter = node.descendants(true)
	for node := iter(); node != nil; node = iter() {
		buffer.WriteRune(node.char)
		if node.eow {
			buffer.WriteRune('1')
		} else {
			buffer.WriteRune('0')
		}
	}
	h := fnv.New32()
	h.Write(buffer.Bytes())
	return h.Sum32()
}

func (n *Node) addChild(node *Node) {
	if _, ok := n.childs[node.char]; ok {
		return
	}
	if node.parent == nil {
		node.parent = n
	}
	n.childs[node.char] = node
}

func (n *Node) removeChild(node *Node) bool {
	if _, ok := n.childs[node.char]; !ok {
		return false
	}
	delete(n.childs, node.char)
	return true
}

func (n *Node) descendants(self bool) func() *Node {
	pop := func(stack []*Node) ([]*Node, *Node) {
		if len(stack) <= 0 {
			return stack, nil
		} else {
			node := stack[0]
			return stack[1:], node
		}
	}

	var stack []*Node
	if self {
		stack = append(stack, n)
	}
	for _, node := range n.childs {
		stack = append(stack, node)
	}

	return func() *Node {
		var node *Node
		stack, node = pop(stack)
		if node == nil {
			return nil
		}
		for _, node := range node.childs {
			stack = append(stack, node)
		}
		return node
	}
}

func (n *Node) Next(char rune) *Node {
	return n.childs[char]
}

func (n *Node) HasChilds() bool {
	if len(n.childs) == 0 {
		return false
	}
	return true
}
