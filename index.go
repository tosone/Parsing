package main

import "smartconn.cc/tosone/parsing/parsing"
import "path/filepath"
import "fmt"

func main() {
	path, _ := filepath.Abs("./example/test/default-ac180c")
	fmt.Println(path)
	parsing.Start(path, "456", parsing.ManifestFile)
	parsing.Stop()
}
