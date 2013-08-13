package pinyin

import (
	"fmt"
	"testing"
)

func TestPinyin(t *testing.T) {
	var py Pinyin
	py.Init("data/pinyin_table.txt")
	fmt.Println(py.GetPinyin('中', true))
	fmt.Println(py.GetPinyin('中', false))
	fmt.Println(py.GetNumericTone('中'))
	fmt.Println(py.GetNumericTone('国'))

	fmt.Println(py.GetPinyin('和', true))

	fmt.Println(py.GetPinyin('绿', true))
	fmt.Println(py.GetPinyin('绿', false))
}
