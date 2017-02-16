package paragraph

import (
	"fmt"

	"smartconn.cc/tosone/parsing/model"
)

var m model.Paragraph

// Index 入口文件
func Index(res model.RegexResult, options model.ModelsOptions) model.ModeResponse {
	fmt.Println(res)
	response := model.ModeResponse{Sound: []string{}, Backline: 1}
	return response
}
