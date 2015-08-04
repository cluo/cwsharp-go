// Copyright (c) CWSharp. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package cwsharp

import (
	"hash/fnv"
)

func fnvHash(input []byte) uint32 {
	h := fnv.New32()
	h.Write(input)
	return h.Sum32()
}

func isLetterCase(c rune) bool {
	return (c >= 0x41 && c <= 0x5A) || (c >= 0x61 && c <= 0x7A)
}

func isUpperCase(c rune) bool {
	return c >= 0x41 && c <= 0x5A
}

func isLowercase(c rune) bool {
	return c >= 0x61 && c <= 0x7A
}

func isNumeralCase(c rune) bool {
	return c >= 0x30 && c <= 0x39
}

func isCjkCase(c rune) bool {
	return c >= 0x4e00 && c <= 0x9fa5
}

func isNull(c rune) bool {
	return c == 0
}
