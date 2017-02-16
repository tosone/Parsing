package check

import (
	"regexp"

	"smartconn.cc/tosone/parsing/model"
)

// Pattern 模式
func Pattern() map[string]*regexp.Regexp {
	return map[string]*regexp.Regexp{
		"comment":           regexp.MustCompile(`^###.*$`),
		"bgmStart":          regexp.MustCompile(`^\$bgm\s+(\[.*\])$`),
		"bgmStop":           regexp.MustCompile(`^\$bgm\s+off$`),
		"paragraph":         regexp.MustCompile(`^\$paragraph\s+(.*)$`),
		"shareFlag":         regexp.MustCompile(`^([a-zA-Z0-9\-_\/]+\.(wav|ogg|amr|mp3)|\[.*\])\s+\$share$`),
		"wait":              regexp.MustCompile(`^\$wait\s+(\d+)\s+(\[\{.*\}\])$`),
		"waitFlag":          regexp.MustCompile(`^\"?([a-zA-Z0-9\-_\/]+\.(wav|ogg|amr|mp3))\"?\s+\$wait\s+(\d+)\s+(\[\{.*\}\])$`),
		"single":            regexp.MustCompile(`^\"?([a-zA-Z0-9\-_\/]+\.(wav|ogg|amr|mp3))\"?$`),
		"singleFlag":        regexp.MustCompile(`^([a-zA-Z0-9\-_\/]+\.(wav|ogg|amr|mp3)|\[.*\])\s+([a-zA-Z0-9\-_\/]+\.(wav|ogg|amr|mp3))$`),
		"singleComplex":     regexp.MustCompile(`^(\{.*sound.*\})$`),
		"singleComplexFlag": regexp.MustCompile(`^([a-zA-Z0-9\-_\/]+\.(wav|ogg|amr|mp3)|\[.*\])\s+(\{.*sound.*\})$`),
		"choice":            regexp.MustCompile(`^(\{.*\})\s+([a-zA-Z0-9\-_\/]+\.(wav|ogg|amr|mp3))$`),
		"choiceFlag":        regexp.MustCompile(`^([a-zA-Z0-9\-_\/]+\.(wav|ogg|amr|mp3)|\[.*\])\s+(\{.*\})\s+([a-zA-Z0-9\-_\/]+\.(wav|ogg|amr|mp3))$`),
		"choiceComplex":     regexp.MustCompile(`^(\[\{.*keywords?.*\}\])$`),
		"choiceComplexFlag": regexp.MustCompile(`^([a-zA-Z0-9\-_\/]+\.(wav|ogg|amr|mp3)|\[.*\])\s+(\[\{.*keywords?.*\}\])$`),
		"interaction":       regexp.MustCompile(`^(\[\{.*button.*\}\])$`),
		"interactionFlag":   regexp.MustCompile(`^([a-zA-Z0-9\-_\/]+\.(wav|ogg|amr|mp3)|\[.*\])\s+(\[\{.*\}\])$`),
	}
}

// Check 分析正则结果
func Check() map[string]func([]string) model.RegexResult {
	return map[string]func([]string) model.RegexResult{
		"comment": func(results []string) model.RegexResult {
			return model.RegexResult{"", "", results[0], ""}
		},
		"bgmStart": func(results []string) (res model.RegexResult) {
			return model.RegexResult{"", "", results[1], ""}
		},
		"bgmStop": func(results []string) model.RegexResult {
			return model.RegexResult{"", "", results[1], ""}
		},
		"paragraph": func(results []string) model.RegexResult { // need to rewrite
			return model.RegexResult{"", "", results[1], ""}
		},
		"share": func(results []string) model.RegexResult {
			return model.RegexResult{"", "", results[0], ""}
		},
		"shareFlag": func(results []string) model.RegexResult {
			return model.RegexResult{results[1], "", "", ""}
		},
		"wait": func(results []string) model.RegexResult {
			return model.RegexResult{"", results[1], results[2], ""}
		},
		"waitFlag": func(results []string) model.RegexResult {
			return model.RegexResult{results[1], results[3], results[4], ""}
		},
		"single": func(results []string) model.RegexResult {
			return model.RegexResult{"", "", results[1], ""}
		},
		"singleFlag": func(results []string) model.RegexResult {
			return model.RegexResult{results[1], "", results[3], ""}
		},
		"singleComplex": func(results []string) model.RegexResult {
			return model.RegexResult{"", "", results[1], ""}
		},
		"singleComplexFlag": func(results []string) model.RegexResult {
			return model.RegexResult{results[1], "", results[3], ""}
		},
		"choice": func(results []string) model.RegexResult {
			return model.RegexResult{"", "", results[1], results[2]}
		},
		"choiceFlag": func(results []string) model.RegexResult {
			return model.RegexResult{results[1], "", results[3], results[4]}
		},
		"choiceComplex": func(results []string) model.RegexResult {
			return model.RegexResult{"", "", results[1], ""}
		},
		"choiceComplexFlag": func(results []string) model.RegexResult {
			return model.RegexResult{results[1], "", results[3], ""}
		},
		"interaction": func(results []string) model.RegexResult {
			return model.RegexResult{"", "", results[1], ""}
		},
		"interactionFlag": func(results []string) model.RegexResult {
			return model.RegexResult{results[1], "", results[3], ""}
		},
	}
}
