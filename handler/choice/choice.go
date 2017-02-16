package choice

import (
	"encoding/json"
	"fmt"

	"log"

	"smartconn.cc/tosone/parsing/config"
	"smartconn.cc/tosone/parsing/lib/audioMgr"
	"smartconn.cc/tosone/parsing/lib/betterMatch"
	"smartconn.cc/tosone/parsing/model"
)

var keywords []string

var channel chan string

func init() {
	channel = make(chan string, 1)
}

// Index 简单选择模式的入口
func Index(res model.RegexResult, options model.ModelsOptions) model.ModeResponse {
	// 关键词列表
	keywords = []string{}
	// 解析分支部分
	var structContent model.ChoiceModel
	json.Unmarshal([]byte(res.Con), &structContent)

	for _, v := range structContent {
		keywords = append(keywords, v...)
	}
	betterMatch.SortByLen(keywords)

	go RecordChoice(2, false, 1, "")
	result := <-channel
	close(channel)
	fmt.Println(result, keywords, betterMatch.MatchWhole(result, keywords))
	response := model.ModeResponse{Sound: []string{}, Backline: 1}
	return response
}

// RecordChoice 选择录音
func RecordChoice(max int, once bool, times int, state string) {
	if times == 2 {
		if state == "" || state == "saidNothing" {
			audioMgr.Play(config.Audio["saidNothing"], 0)
		} else if state == "saidWrong" {
			audioMgr.Play(config.Audio["saidWrong"], 0)
		}
	}
	choiceText := ""
	res, err := audioMgr.RecordToText(max)
	if err != nil {
		log.Println(err)
		channel <- ""
	}
	choiceText = betterMatch.Match(res, keywords)
	if !((once == true && times == 1) || (once == false && times == 2)) {
		if res == "" || len(choiceText) == 0 {
			state := "saidNothing"
			if res != "" && len(choiceText) == 0 {
				state = "saidWrong"
			}
			RecordChoice(max, once, 2, state)
			return
		}
		channel <- choiceText
	}
	channel <- choiceText
}

// Stop 停止
func Stop() {
	close(channel)
}
