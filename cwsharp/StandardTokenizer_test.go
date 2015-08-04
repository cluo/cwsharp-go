// Copyright (c) CWSharp. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package cwsharp

import (
	"os"
	"path/filepath"
	"testing"
)

func InitDawgFile() string {
	words := map[string]int32{
		"研": 125550, "究": 99256, "生": 601668, "命": 87293, "起": 228817, "源": 48575,
		"一": 2542556, "次": 236297, "性": 199183, "交": 728005, "元": 100523, "百": 64888,
		"研究": 1, "生命": 1, "起源": 1, "研究生": 1, "一次": 100, "一次性": 1000, "性交": 10,
		"卡拉ok": 1, "卡拉": 1, "t恤": 1, "为首": 1, "首要": 1, "考虑": 1, "苹果": 1,
	}
	util := &WordUtil{}
	for word, freq := range words {
		util.Add(word, freq)
	}
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	file := dir + "\\test.dawg"
	util.Save(file)
	return file
}

func TestStandardTokenizer(t *testing.T) {
	file := InitDawgFile()
	tokenizer := NewStandardTokenizer(file, true)
	var text string
	t.Log("===================")
	text = "一次性交一百元" //歧义
	t.Log(text)
	iterator := tokenizer.Traverse(text)
	for token, next := iterator(); next != nil; token, next = iterator() {
		t.Log(token)
	}
	t.Log("===================")
	text = "研究生命起源"
	t.Log(text)
	iterator = tokenizer.Traverse(text)
	for token, next := iterator(); next != nil; token, next = iterator() {
		t.Log(token)
	}
	t.Log("===================")
	text = "卡拉OK"
	t.Log(text)
	iterator = tokenizer.Traverse(text)
	for token, next := iterator(); next != nil; token, next = iterator() {
		t.Log(token)
	}
	t.Log("===================")
	//no 任何内容
	text = ""
	t.Log(text)
	iterator = tokenizer.Traverse(text)
	for token, next := iterator(); next != nil; token, next = iterator() {
		t.Log(token)
	}
}
