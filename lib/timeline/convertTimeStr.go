package timeline

import "strings"
import "strconv"

// Convert 将字符串格式的时间转换成数字
func Convert(str string) int {
	if len(str) == 0 {
		return 0
	}
	arr := strings.Split(str, "'")
	if len(arr) == 1 {
		return str2int(arr[0])
	} else if len(arr) == 2 {
		if str2int(arr[1]) >= 60 {
			return str2int(arr[0])*60 + 60
		}
		return str2int(arr[0])*60 + str2int(arr[1])
	}
	return 0
}

func str2int(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}
