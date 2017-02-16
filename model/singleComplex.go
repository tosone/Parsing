package model

// SingleComplexModel 复杂单句模式的数据结构
type SingleComplexModel struct {
	Sound    interface{}             `json:"sound"`
	Random   interface{}             `json:"random"`
	Backline int                     `json:"backline"`
	Loop     bool                    `json:"loop"`
	Last     string                  `json:"last"`
	Input    SingleComplexInputTask  `json:"input"`
	OutPut   SingleComplexOutputTask `json:"output"`
}

// SingleComplexInputTask 输入任务的 task 列表
type SingleComplexInputTask struct {
	Cmd    string                     `json:"cmd"`
	Start  string                     `json:"start"`
	Breaks string                     `json:"break"`
	Button map[string]map[string]Task `json:"button"`
}

// SingleComplexOutputTask 输出任务的 task 列表
type SingleComplexOutputTask struct {
	Sound []Task `json:"sound"`
	LED   []Task `json:"LED"`
}
