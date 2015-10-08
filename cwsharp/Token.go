package cwsharp

type Token interface {
	Text() string
	Kind() Kind
}
