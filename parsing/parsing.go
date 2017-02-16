package parsing

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"strconv"
	"strings"

	"smartconn.cc/tosone/parsing/check"
	"smartconn.cc/tosone/parsing/handler"
	"smartconn.cc/tosone/parsing/model"
)

var selected []string
var parsingDir string

var currPage int
var totalPage int
var pages []string

var modelsOptions model.ModelsOptions

func init() {
	selected = []string{}
	parsingDir = ""
	currPage = 0
	totalPage = 0
	pages = []string{}
	modelsOptions = model.ModelsOptions{NetworkStatus: true}
}

// reset 重置
func reset() {
	selected = []string{}
}

// Start 开始 parsing
func Start(dir string, paragraphFlag string, manifestFile string) {
	parsingDir = dir
	if parsingDir == "" {
		fmt.Println("Parsing: Story path is a blank directory.")
		return
	}
	modelsOptions.Predir = dir
	manifestFileData, err := ioutil.ReadFile(path.Join(dir, manifestFile))
	if err != nil {
		fmt.Println("Parsing: Cannot found any thing in " + dir + ".")
		return
	}
	pages = strings.Split(string(manifestFileData), "\n")
	totalPage = len(pages)
	currPage = 0

	start, end := choosePoint(pages, paragraphFlag)
	parsing(start, end)
}

func parsing(start, end int) {
	handlers := handler.Index()
	pats := check.Pattern()
	checks := check.Check()
	if start < 0 || end < 0 || start > totalPage || end > totalPage {
		stopped()
		return
	}
	flag := true
	for index := start; index < end; index++ {
		temp := 0
		flag = true
		for mode, pat := range pats {
			temp++
			patResult := pat.FindAllStringSubmatch(pages[index], 1)
			if len(patResult) == 1 {
				flag = false
				log.Println("Line " + strconv.Itoa(index+1) + ": mode is " + mode + ".")
				request := checks[mode](patResult[0])
				if flagTest(request.Flag) {
					response := handlers[mode](request, modelsOptions)
					selected = append(selected, response.Sound...)
				} else {
					fmt.Println("Line " + strconv.Itoa(index+1) + " has been jumped.")
				}
				break
			}
		}
		if temp == len(pats) && flag {
			if pages[index] == "" {
				log.Println("Line " + strconv.Itoa(index+1) + ": mode is blank.")
			} else {
				log.Println("Line " + strconv.Itoa(index+1) + ": mode is none.")
			}
		}
	}
	stopped()
}

// flagTest 检查当前的 flag 是否在已经播放的列表中
func flagTest(flag string) bool {
	if flag == "" {
		return true
	}
	for _, v := range selected {
		if flag == v {
			return true
		}
	}
	return false
}

// Stop 结束 parsing
func Stop() {
	handler.Stop()
	stopped()
}

func stopped() {
	log.Println("parsing over.")
	return
}
