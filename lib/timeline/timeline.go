package timeline

import (
	"time"

	"smartconn.cc/tosone/parsing/model"
)

var currTime int
var breakFlag bool

func init() {
	currTime = 0
	breakFlag = false
}

func remove(s []model.Task, i int) []model.Task {
	return append(s[:i], s[i+1:]...)
}

// Start 开始任务
func Start(channel chan model.Task, tasks []model.Task) {
	currTime = 0
	breakFlag = false
	go func() {
		for len(tasks) != 0 || breakFlag == false {
			time.Sleep(time.Second)
			currTime++
			for k, v := range tasks {
				if Convert(v.Start) == currTime {
					tasks = append(tasks[:k], tasks[k+1:]...)
					channel <- v
				}
				if len(tasks) == 0 || breakFlag == true {
					breakFlag = true
					break
				}
			}
		}
	}()
}

// Stop 结束任务
func Stop() {
	breakFlag = true
}
