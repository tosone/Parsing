package parsing

import check "smartconn.cc/tosone/parsing/check"

func init() {

}

func choosePoint(pages []string, paragraphFlag string) (start, end int) {
	pats := check.Pattern()
	checks := check.Check()
	start, end = -1, -1
	for line, page := range pages {
		for mode, pat := range pats {
			patResult := pat.FindAllStringSubmatch(page, 1)
			if len(patResult) == 1 {
				if mode == "paragraph" && checks[mode](patResult[0]).Con == paragraphFlag {
					start = line
				} else if start != -1 && mode == "paragraph" && checks[mode](patResult[0]).Con != paragraphFlag {
					return start, line
				}
			}
		}
	}
	return start, len(pages)
}
