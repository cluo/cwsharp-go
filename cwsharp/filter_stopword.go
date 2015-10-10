package cwsharp

// 停用词过滤

type StopwordFilter struct {
	t Tokenizer
	// 检查指定的Token是否忽略，不返回处理
	CheckIgnore func(Token) bool
}

type stopwordTokenIterator struct {
	p     *StopwordFilter
	inner TokenIterator
	cur   Token
}

func NewStopwordFilter(t Tokenizer) *StopwordFilter {
	return &StopwordFilter{t: t}
}

func (s *StopwordFilter) Traverse(r Reader) TokenIterator {
	if s.CheckIgnore == nil {
		s.CheckIgnore = defaultCheckIgnore
	}
	return &stopwordTokenIterator{
		p:     s,
		inner: s.t.Traverse(r),
	}
}

func (iter *stopwordTokenIterator) Next() bool {
	for iter.inner.Next() {
		t := iter.inner.Cur()
		if !iter.p.CheckIgnore(t) {
			iter.cur = t
			return true
		}
	}
	return false
}

func (iter *stopwordTokenIterator) Cur() Token {
	return iter.cur
}

func defaultCheckIgnore(t Token) bool {
	return false
}
