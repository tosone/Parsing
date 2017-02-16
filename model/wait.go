package model

// WaitModel 等待模式模型
type WaitModel []struct {
	Cmd    string                     `json:"cmd"`
	Button map[string]map[string]Task `json:"button"`
}

// WaitModelTime 等待模式的等待时间
type WaitModelTime int
