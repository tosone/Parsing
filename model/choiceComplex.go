package model

// ChoiceComplexModel 复杂的选择模式
type ChoiceComplexModel []struct {
	Before             interface{} `json:"before"`
	OfflineBefore      interface{} `json:"offlineBefore"`
	Sound              interface{} `json:"sound"`
	Keyword            []string    `json:"keyword"`
	Vars               string      `json:"var"`
	PreSound           interface{} `json:"preSound"`
	Once               bool        `json:"once"`
	SaidNothing        interface{} `json:"saidNothing"`
	SaidWrong          interface{} `json:"saidWrong"`
	OfflineSaidNothing interface{} `json:"offlineSaidNothing"`
	OfflineSaidWrong   interface{} `json:"offlineSaidWrong"`
	MaxRecording       int         `json:"maxRecording"`
	LED                Task        `json:"LED"`
}
