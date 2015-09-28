package cwsharp

// 提供简单的中文分词功能，识别单个中文，数字和英文
type SimpleTokenizer struct {
	Tokenizer
}

func (t *SimpleTokenizer) Traverse(r Reader) TokenIterator {
	return &TokenIteratorWrap{
		b: &SimpleTokenBreaker{r: r},
	}
}

type SimpleTokenBreaker struct {
	// 自定义字符串规则策略
	CheckRune func(cur rune, via []rune) bool
	Reader    Reader
	Breaker
}

func (s *SimpleTokenBreaker) Next() (t Token, err error) {
	if s.CheckRune == nil {

	}
	err = nil
	ch := s.Reader.ReadRule()
	if ch == EOF {
		err = DONE
		return
	}
	b := s.Reader.Pos()

	/*
		off := s.r.Pos()
		if isCjkCase(ch) {
			t.text = string(ch)
			t.kind = CJK
			return
		} else if isLetterCase(ch) {
			for ch = s.r.ReadRule(); ch != EOF; ch = s.r.ReadRule() {
				if isLetterCase(ch) || isNumeralCase(ch) {
					continue
				}
				if c, ok := _AlphaNumStops[ch]; ok && c {
					n := s.r.Peek()
					if _, ok = _AlphaNumStops[n]; ok || isLetterCase(n) || isNumeralCase(n) {
						continue
					}
				}
				s.r.Seek(s.r.Pos() - 1)
				break
			}
			l := s.r.Pos() - off
			s.r.Seek(off)
			t.text = string(s.r.ReadRules(l))
			t.kind = ALPHANUM
			return
		} else if isNumeralCase(ch) {
			mixed := false
			for ch = s.r.ReadRule(); ch != EOF; ch = s.r.ReadRule() {
				if isLetterCase(ch){
					mixed=true
					continue
				}else if isNumeralCase(ch){
					continue
				}
				if c,ok:=

			}
		}*/
	return
}
