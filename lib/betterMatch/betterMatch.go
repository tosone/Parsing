package betterMatch

import (
	"strings"

	"smartconn.cc/tosone/parsing/lib/convertPinyin"
)

var similarityNum int

func init() {
	similarityNum = 3
}

// SetSimilarity 设置相似度
func SetSimilarity(s int) {
	similarityNum = s
}

// Match 匹配发音相似的字符串，但仅仅会返回第一的相似的字符串
func Match(origin string, match []string) (res string) {
	if len(origin) == 0 || len(match) == 0 {
		return ""
	}
	match = SortByLen(match)
	for _, v := range match {
		matchStr, _ := matchOne(origin, v)
		if len(matchStr) != 0 {
			return v
		}
	}
	return ""
}

// MatchWhole 匹配发音相似的字符串，但会返回所有相似的字符串，所有字符串以相似度排序
func MatchWhole(origin string, match []string) (res []string) {
	if len(origin) == 0 || len(match) == 0 {
		return res
	}
	var list []similarityList
	match = SortByLen(match)
	for _, v := range match {
		matchStr, matchVal := matchOne(origin, v)
		if len(matchStr) != 0 {
			list = append(list, similarityList{matchVal, matchStr})
		}
	}
	sortSimilarity(list)
	for _, v := range list {
		res = append(res, v.str)
	}
	return res
}

func matchOne(origin, match string) (string, int) {
	if strings.Index(origin, match) != -1 {
		return match, 0
	}
	originPinyin := convertPinyin.Convert(origin)
	matchPinyin := convertPinyin.Convert(match)
	if len(originPinyin) < len(matchPinyin) {
		return "", 1000
	}
	offset := abs(len(originPinyin) - len(matchPinyin))
	for i := 0; i < (offset + 1); i++ {
		similarity := 0
		for k := range matchPinyin {
			similarity += simpleMatch(matchPinyin[k], originPinyin[k+i])
		}
		if similarity < similarityNum {
			return match, similarity
		}
	}
	return "", 1000
}

func simpleMatch(str1, str2 string) int {
	similarity := 0
	str1Arr := strings.Split(str1, "")
	str2Arr := strings.Split(str2, "")
	travel1 := str1Arr // 以谁为遍历目标遍历
	travel2 := str2Arr // 长度长的为第一遍历字串
	if len(str1Arr) > len(str2Arr) {
		travel1 = str2Arr
		travel2 = str1Arr
	}
	for k := range travel1 {
		if travel1[k] != travel2[k] {
			similarity++
			if k != (len(travel2) - 1) {
				similarity += simpleMatch(strings.Join(travel2[k+1:], ""), strings.Join(travel1[k:], ""))
				break
			}
		}
	}
	return similarity
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
