package comment

import (
	"log"

	"smartconn.cc/tosone/parsing/model"
)

// Index 简单选择模式的入口
func Index(res model.RegexResult, options model.ModelsOptions) model.ModeResponse {
	log.Println("curr mode is comment.")
	response := model.ModeResponse{Sound: []string{}, Backline: 1}
	return response
}
