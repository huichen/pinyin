从汉字得到拼音
===

* 可将四万多个常用汉字转化为拼音
* 可返回带声调的拼音
* 可得到汉字属于哪个四声

## 安装

    go get -u github.com/huichen/pinyin

## 用法

```go
var py pinyin.Pinyin

// 初始化，载入汉字拼音映射文件
py.Init("data/pinyin_table.txt")

// 返回汉字的拼音
// GetPinyin的第一个参数为单个汉字，第二个参数为是否返回带声调的拼音。
// 比如下面两行的输出分别为 "zhōng" 和 "zhong"
// 当该字无法识别或者不是汉字时返回空字符串。
fmt.Println(py.GetPinyin('中', true))
fmt.Println(py.GetPinyin('中', false))

// 返回汉字的声调（整数），0为轻声，1为平声，依此类推。
fmt.Println(py.GetNumericTone('中'))

// 下面的输出分别为 "lǜ" 和 "lv"
fmt.Println(py.GetPinyin('绿', true))
fmt.Println(py.GetPinyin('绿', false))
```

## 注意

本工具无法处理多音字，只能返回最常用的音，比如"和稀泥"和"和为贵"中的和字都会返回"he"。
