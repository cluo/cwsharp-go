package cwsharp

type stringReader struct {
	src []rune
	off int
}

func (r *stringReader) Init(src []rune) {
	r.src = src
	r.off = 0
}

func (r *stringReader) ReadRule() rune {
	if r.off == len(r.src) {
		return EOF
	}
	ch := r.src[r.off]
	r.off += 1
	return ch
}

func (r *stringReader) Peek() rune {
	if r.off == len(r.src) {
		return EOF
	}
	return r.src[r.off]
}

func (r *stringReader) Seek(offset int) {
	r.off = offset
}

func (r *stringReader) Pos() int {
	return r.off
}

func NewStringReader(src string) Reader {
	r := &stringReader{}
	r.Init([]rune(src))
	return r
}
