package parsing

import (
	"path/filepath"
	"testing"
)

func TestParsing(t *testing.T) {
	path, _ := filepath.Abs("../example/test/default-ac180c")
	Start(path, "456", ManifestFile)
}
