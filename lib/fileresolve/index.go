package fileresolve

import (
	"path"
	"strings"
)

// Get 获取文件的绝对地址
func Get(predir, file string) string {
	if strings.HasPrefix(file, "/") {
		return file
	}
	return path.Join(predir, file)
}
