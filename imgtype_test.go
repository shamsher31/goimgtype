package imgtype

import (
	"strings"
	"testing"

	"github.com/shamsher31/goimgext"
)

func TestGet(t *testing.T) {
	imgType, err := Get("golang.png")

	if err != nil {
		t.Fatal(err)
	}

	if isImgTypeFromExt(imgType) == false {
		t.Fatal("Not a valid image type")
	}
}

func isImgTypeFromExt(p string) bool {
	ext := imgext.Get()
	for i := 0; i < len(ext); i++ {
		if strings.Contains(ext[i], p[6:len(p)]) {
			return true
		}
	}
	return false
}
