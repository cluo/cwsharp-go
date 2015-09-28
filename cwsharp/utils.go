package cwsharp

func isLetterCase(c rune) bool {
	return (c >= 0x41 && c <= 0x5A) || (c >= 0x61 && c <= 0x7A)
}

func isLetterUpperCase(c rune) bool {
	return c >= 0x41 && c <= 0x5A
}

func isLetterLowerCase(c rune) bool {
	return c >= 0x61 && c <= 0x7A
}

func isNumeralCase(c rune) bool {
	return c >= 0x30 && c <= 0x39
}

func isCjkCase(c rune) bool {
	return c >= 0x4e00 && c <= 0x9fa5
}
