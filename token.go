package cwsharp

// Kind is a token type.
type Kind interface{}

// Token represents a word text and with its kind of type.
type Token struct {
	Text string
	Kind Kind
}
