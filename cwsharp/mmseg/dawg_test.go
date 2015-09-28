package std

import (
	"bufio"
	"os"
	"path/filepath"
)

import (
	"testing"
)

func TestAddChild(t *testing.T) {
	rootNode := newDawgNode('a')
	rootNode.AddChild(newDawgNode('a')) //aa
	rootNode.AddChild(newDawgNode('b')) //ab
	//        a
	//      a    b
	if len(rootNode.childs) != 2 {
		t.Fail()
	}
}

func TestRemoveChild(t *testing.T) {
	rootNode := newDawgNode('a')
	childNode := newDawgNode('b')
	rootNode.AddChild(childNode)
	rootNode.RemoveChild(childNode)
	if len(rootNode.childs) > 0 {
		t.Fail()
	}
}

func TestDescendants(t *testing.T) {
	//hello > h, e, l, l, o
	s := []rune{'h', 'e', 'l', 'l', 'o'}
	rootNode := newDawgNode(' ')
	childNode := rootNode
	for _, c := range s {
		node := newDawgNode(c)
		childNode = childNode.AddChild(node)
	}
	childNode.eow = true
	//hei h,a,i
	s = []rune{'h', 'e', 'i'}
	childNode = rootNode
	for _, c := range s {
		node := newDawgNode(c)
		childNode = childNode.AddChild(node)
	}
	childNode.eow = true
	if len(rootNode.childs['h'].childs['e'].childs) != 2 {
		t.Fail()
	}
	for node, next := rootNode.Descendants(false)(); next != nil; node, next = next() {
		t.Logf("%s", node)
	}
}

func CreateDawg() *dawg {
	wordBag := map[string]int32{
		"中": 1, "国": 1, "世": 1, "博": 1, "会": 1, "人": 1,
		"世界": 1, "世博": 1, "世博会": 1, "中国": 1, "中国人": 1, "做一天和尚撞一天钟": 1,
	}
	dawg := buildDawg(wordBag)
	return dawg
}

func TestContains(t *testing.T) {
	dawg := CreateDawg()
	node := dawg.root.childs['世']
	if node == nil {
		t.Error("No any dawg-node.")
	}
	node = node.childs['界']
	if node == nil {
		t.Error("No any dawg-node.")
	}
	found := dawg.Contains("世博")
	if !found {
		t.Error("Doesn't exits this word.")
	}
}

func TestMatchsPrefix(t *testing.T) {
	dawg := CreateDawg()
	result := dawg.MatchsPrefix("中")
	//中国
	//中国人
	t.Log("prefix:中")
	count := len(result)
	if count != 2 {
		t.Fail()
	}
	for word, freq := range result {
		t.Logf("%s - %d", word, freq)
	}
}

func TestDawg(t *testing.T) {
	dawg := CreateDawg()
	rootNode := dawg.root
	for node, next := rootNode.Descendants(false)(); next != nil; node, next = next() {
		t.Logf("%s", node)
	}
}

func TestDawgCoder(t *testing.T) {
	coder := dawgCoder{version: 1.0}
	file := "d:\\go-dawg.dawg"
	f, _ := os.Create(file)
	dawg := CreateDawg()
	//test Encoding
	w := bufio.NewWriter(f)
	coder.Encode(w, dawg)
	w.Flush()
	f.Close()
	t.Log("write a dawg to file")
	f, _ = os.Open(file)
	//test decoding
	r := bufio.NewReader(f)
	newdawg := coder.Decode(r)
	t.Log("read from file")
	//node:=newdawg.root.childs['世']
	for _, node := range newdawg.root.childs {
		t.Log(node.String())
	}
	found := newdawg.Contains("世界")
	t.Log("query:世界")
	if !found {
		t.Error("No '世界' in this dawg. ")
	}
	found = newdawg.Contains("世博会")
	if !found {
		t.Fail()
	}
	t.Log("query:世博会")
	f.Close()
}

func TestWriteAndRead(t *testing.T) {
	util := &WordUtil{}
	util.Add("中国", 1)
	util.Add("世博", 1)
	util.Add("c#", 1)
	util.Add("java", 1)
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	file := dir + "\\test.dawg"
	util.Save(file)
	t.Log("保存到文件:" + file)
	util.Load(file)
	t.Log("加载文件:" + file)

	found := util.Contains("c#")
	if !found {
		t.Error("找不到'c#'词组")
	}
}
