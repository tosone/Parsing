package single

import (
	"smartconn.cc/tosone/parsing/lib/audioMgr"
	"smartconn.cc/tosone/parsing/lib/fileresolve"
	"smartconn.cc/tosone/parsing/model"
)

var modeBool bool

func init() {
	modeBool = false
}

// Index 入口函数
func Index(res model.RegexResult, options model.ModelsOptions) model.ModeResponse {
	modeBool = true
	audiofile := fileresolve.Get(options.Predir, res.Con)
	audioMgr.Play(audiofile, 0)
	response := model.ModeResponse{Sound: []string{res.Con}, Backline: 1}
	return response
}

// Stop 停止当前模式
func Stop() {
	if modeBool {
		modeBool = false
		if audioMgr.IsBGMPlaying() {
			audioMgr.Stop()
		}
	}
}
