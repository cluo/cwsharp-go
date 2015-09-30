package cwsharp

var EOF rune = -1

type Reader interface {
	ReadRule() rune
	Peek() rune
	Seek(offset int)
	Pos() int
}
