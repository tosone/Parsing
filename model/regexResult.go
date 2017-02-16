package model

// RegexResult 正则结果
type RegexResult struct {
	Flag string
	Time string
	Con  string
	Def  string
}

// ModelsOptions 在模式中传入的参数
type ModelsOptions struct {
	NetworkStatus bool
	Predir        string
}
