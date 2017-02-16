package choiceComplex

import (
	"encoding/json"

	"smartconn.cc/tosone/parsing/model"
)

var keywords []string

// Index 简单选择模式的入口
func Index(res model.RegexResult, options model.ModelsOptions) model.ModeResponse {
	var structContent model.ChoiceComplexModel
	json.Unmarshal([]byte(res.Con), &structContent)
	// fmt.Println(structContent)
	for _, v := range structContent {
		keywords = append(keywords, v.Keyword...)
	}
	// fmt.Println(keywords)
	response := model.ModeResponse{Sound: []string{}, Backline: 1}
	return response
}
