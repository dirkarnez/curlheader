package curlheader

import (
	"path/filepath"
	"testing"
)

func TestGetCurlHeader(t *testing.T) {
	t.Log("Hello World")
	// ex, _ := os.Executable()
	// exPath := filepath.Dir(ex)
	// file := filepath.Join(exPath, "seed.txt")
	file := filepath.Join(".", "seed.txt")
	t.Log(file)
	header, err := GetCurlHeader(file)
	if err != nil {
		t.FailNow()
	}

	if header.Get("authority") != "go.dev" {
		t.FailNow()
	}
	t.Log("Done")
}
