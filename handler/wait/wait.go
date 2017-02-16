package wait

import "smartconn.cc/tosone/parsing/model"
import "fmt"

// Index 简单选择模式的入口
func Index(res model.RegexResult, options model.ModelsOptions) model.ModeResponse {
	fmt.Println(res.Con)
	response := model.ModeResponse{Sound: []string{}, Backline: 1}
	return response
}
