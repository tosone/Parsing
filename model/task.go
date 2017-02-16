package model

// TaskOption 任务的参数信息
type TaskOption struct {
	Color string `json:"colors"`
}

// Task 任务信息
type Task struct {
	Type   string     `json:"type"`
	Start  string     `json:"start"`
	Sound  string     `json:"sound"`
	Mode   string     `json:"mode"`
	LED    string     `json:"LED"`
	Option TaskOption `json:"option"`
}
