// 工具类，生成MMSEG分词包采用的字典
package main

import (
	"bufio"
	"fmt"
	"github.com/zhengchun/cwsharp-go/cwsharp"
	"github.com/zhengchun/cwsharp-go/cwsharp/mmseg"
	"os"
	"strconv"
	"strings"
)

func main() {
	BuildDawgFile()
	//TEST
	TestStandard("cwsharp.dawg", "一次性交100元")
}

/*
	生成DAWG词典及测试
*/
func BuildDawgFile() {
	//词频文件
	var file_freq string = "cwsharp.freq"
	//字典文件
	var file_dict string = "cwsharp.dic"
	var dawg_savePath string = "cwsharp.dawg"
	var words = &mmseg.WordUtil{}
	f, err := os.Open(file_freq)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var line string = scanner.Text()
		if len(line) > 0 {
			array := strings.Split(line, " ")
			freq, err := strconv.ParseInt(array[1], 10, 32)
			if err != nil {
				panic(err)
			}
			words.Add(array[0], int32(freq))
		}
	}
	f.Close()
	f, err = os.Open(file_dict)
	if err != nil {
		panic(err)
	}
	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		var line string = scanner.Text()
		if len(line) > 0 {
			words.Add(line, 0)
		}
	}
	f.Close()
	words.Save(dawg_savePath)
	fmt.Println("BuildDawg[Success]")
}

//标准分词测试
func TestStandard(dawgFile string, text string) {
	fmt.Println("测试：" + text)
	tokenizer := mmseg.New(dawgFile)
	for iter := tokenizer.Traverse(cwsharp.NewStringReader(text)); iter.Next(); {
		fmt.Print(iter.Cur())
		fmt.Print(" / ")
	}
	fmt.Println()
}
