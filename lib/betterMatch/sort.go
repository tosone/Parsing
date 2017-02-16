package betterMatch

import (
	"sort"
)

type sortStringByLen []string

func (a sortStringByLen) Len() int           { return len(a) }
func (a sortStringByLen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortStringByLen) Less(i, j int) bool { return len(a[i]) > len(a[j]) }

// SortByLen 将字符串以长度进行排序
func SortByLen(strs []string) []string { sort.Sort(sortStringByLen(strs)); return strs }

type similarityList struct {
	similarity int
	str        string
}

type sortStructByInt []similarityList

func (a sortStructByInt) Len() int           { return len(a) }
func (a sortStructByInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortStructByInt) Less(i, j int) bool { return a[i].similarity < a[j].similarity }
func sortSimilarity(strs []similarityList) []similarityList {
	sort.Sort(sortStructByInt(strs))
	return strs
}
