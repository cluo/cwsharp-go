// Copyright (c) CWSharp. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package cwsharp

type Tokenizer interface {
	Traverse(text string) TokenIterator
}

type TokenIterator func() (Token, TokenIterator)
