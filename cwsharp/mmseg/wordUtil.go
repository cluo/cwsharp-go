package std

import (
	"bufio"
	"os"
	"strings"
)

type WordUtil struct {
	wordbag map[string]int32
}

func (this *WordUtil) Add(word string, frequency int32) {
	if len(word) == 0 {
		return
	}
	if this.wordbag == nil {
		this.wordbag = make(map[string]int32, 0)
	}
	word = strings.ToLower(word)
	this.wordbag[word] = frequency
}

func (this *WordUtil) Contains(word string) bool {
	_, found := this.wordbag[word]
	return found
}

//保存到文件，（dawg格式)
func (this *WordUtil) Save(file string) {
	dawg := buildDawg(this.wordbag)
	coder := dawgCoder{DawgFileVersion}
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	coder.Encode(w, dawg)
	w.Flush()
}

func (this *WordUtil) Load(file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	coder := dawgCoder{DawgFileVersion}
	dawg := coder.Decode(r)
	this.wordbag = dawg.MatchsPrefix("")
}
