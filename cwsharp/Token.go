package cwsharp

type TokenType int32

type Token interface {
	Text() string
	Kind() TokenType
}
