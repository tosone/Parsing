package model

// InteractionModel 互动模式的模型
type InteractionModel struct {
	Cmd            string                     `json:"cmd"`
	InteractionBgm string                     `json:"InteractionBgm"`
	Button         map[string]map[string]Task `json:"button"`
	Nfc            []InteractionModelNfc      `json:"nfc"`
}

// InteractionModelNfc 互动模式的模型关于 nfc 详细参数
type InteractionModelNfc struct {
	Mode   string          `json:"mode"`
	MaxNfc int             `json:"maxNfc"`
	Nfc    interface{}     `json:"nfc"`
	Output map[string]Task `json:"output"`
}
