package pinyin

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	tones = [][]rune{
		[]rune{'ā', 'ē', 'ī', 'ō', 'ū', 'ǖ', 'Ā', 'Ē', 'Ī', 'Ō', 'Ū', 'Ǖ'},
		[]rune{'á', 'é', 'í', 'ó', 'ú', 'ǘ', 'Á', 'É', 'Í', 'Ó', 'Ú', 'Ǘ'},
		[]rune{'ǎ', 'ě', 'ǐ', 'ǒ', 'ǔ', 'ǚ', 'Ǎ', 'Ě', 'Ǐ', 'Ǒ', 'Ǔ', 'Ǚ'},
		[]rune{'à', 'è', 'ì', 'ò', 'ù', 'ǜ', 'À', 'È', 'Ì', 'Ò', 'Ù', 'Ǜ'},
	}
	neutrals = []rune{'a', 'e', 'i', 'o', 'u', 'v', 'a', 'e', 'i', 'o', 'u', 'v'}
)

type Pinyin struct {
	// 从带声调的声母到对应的英文字符的映射
	tonesMap map[rune]rune

	// 从汉字到声调的映射
	numericTonesMap map[rune]int

	// 从汉字到拼音的映射（带声调）
	pinyinMap map[rune]string

	initialized bool
}

// 使用前必须初始化
func (py *Pinyin) Init(pinyinTablePath string) {
	if py.initialized {
		log.Fatal("不能重复初始化")
	}
	py.initialized = true

	py.tonesMap = make(map[rune]rune)
	py.numericTonesMap = make(map[rune]int)
	py.pinyinMap = make(map[rune]string)
	for i, runes := range tones {
		for j, tone := range runes {
			py.tonesMap[tone] = neutrals[j]
			py.numericTonesMap[tone] = i + 1
		}
	}

	f, err := os.Open(pinyinTablePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), "\t")
		if len(strs) < 2 {
			continue
		}
		i, err := strconv.ParseInt(strs[0], 16, 32)
		if err != nil {
			continue
		}
		py.pinyinMap[rune(i)] = strs[1]
	}
}

// 得到汉字的拼音，当withTone==true时返回带声调的拼音
// 当该字无法识别或者不是汉字时返回空字符串。
func (py *Pinyin) GetPinyin(hanzi rune, withTone bool) string {
	if !py.initialized {
		log.Fatal("尚未初始化")
	}

	if withTone {
		return py.pinyinMap[hanzi]
	}
	return py.getNeutral(py.pinyinMap[hanzi])
}

// 得到汉字的声调，1平, 2上, 3去, 4入, 0为轻声
func (py *Pinyin) GetNumericTone(hanzi rune) int {
	if !py.initialized {
		log.Fatal("尚未初始化")
	}

	tone := 0
	for _, c := range py.pinyinMap[hanzi] {
		newTone := py.numericTonesMap[c]
		if newTone != 0 {
			tone = newTone
		}
	}
	return tone
}

func (py *Pinyin) getNeutral(input string) string {
	if !py.initialized {
		log.Fatal("尚未初始化")
	}

	output := make([]rune, utf8.RuneCountInString(input))
	count := 0
	for _, tone := range input {
		neutral, found := py.tonesMap[tone]
		if found {
			output[count] = neutral
		} else {
			output[count] = tone
		}
		count++
	}
	return string(output)
}
