package handler

import (
	"smartconn.cc/tosone/parsing/handler/choice"
	"smartconn.cc/tosone/parsing/handler/choiceComplex"
	"smartconn.cc/tosone/parsing/handler/comment"
	"smartconn.cc/tosone/parsing/handler/interaction"
	"smartconn.cc/tosone/parsing/handler/paragraph"
	"smartconn.cc/tosone/parsing/handler/single"
	"smartconn.cc/tosone/parsing/handler/wait"
	"smartconn.cc/tosone/parsing/model"
)

// Index 所有 handler 的索引
func Index() map[string]func(model.RegexResult, model.ModelsOptions) model.ModeResponse {
	return map[string]func(model.RegexResult, model.ModelsOptions) model.ModeResponse{
		"single":            single.Index,
		"singleFlag":        single.Index,
		"paragraph":         paragraph.Index,
		"choice":            choice.Index,
		"choiceFlag":        choice.Index,
		"interaction":       interaction.Index,
		"interactionFlag":   interaction.Index,
		"wait":              wait.Index,
		"waitFlag":          wait.Index,
		"choiceComplex":     choiceComplex.Index,
		"choiceComplexFlag": choiceComplex.Index,
		"comment":           comment.Index,
	}
}

// Stop 停止所有模式
func Stop() {
	single.Stop()
}
