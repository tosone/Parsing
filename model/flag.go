package model

// FlagModel 简单的 flag 模式
type FlagModel string

// ArrFlagModel 组合的 flag 模式
type ArrFlagModel []string

// ComplexFlagModel 复杂的 flag 模式
type ComplexFlagModel struct {
	Mode  string `json:"mode"`
	Unvar string `json:"unvar"`
}
