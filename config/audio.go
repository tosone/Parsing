package config

// Audio 音频文件
var Audio map[string]string

func init() {
	Audio = audiofile()
}

func audiofile() map[string]string {
	return map[string]string{
		"saidNothing": "saidNothing",
		"saidWrong":   "saidWrong",
	}
}
