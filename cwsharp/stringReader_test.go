// Copyright (c) CWSharp. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package cwsharp

import (
	"testing"
)

func TestRead(t *testing.T) {
	reader := newStringReader("Hello,World", true)
	char := reader.Read()
	if string(char) != "H" {
		t.Fail()
	}
	char = reader.Peek()
	if string(char) != "e" {
		t.Fail()
	}
	reader.Seek(0)
	for char = reader.Read(); char != 0; char = reader.Read() {
		t.Log(string(char))
	}
}

func TestReadCount(t *testing.T) {
	reader := newStringReader("Hello,World", false)
	str := string(reader.ReadCount(5)) //hello
	if str != "hello" {
		t.Fail()
	}
}
