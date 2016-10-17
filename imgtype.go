// Package imgtype return type of image based on mime type
package imgtype // import "github.com/shamsher31/goimgtype"

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/shamsher31/goimgext"
)

// Get returns the type of Image
func Get(p string) (string, error) {
	file, err := os.Open(p)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	buff := make([]byte, 512)

	_, err = file.Read(buff)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	filetype := http.DetectContentType(buff)

	ext := imgext.Get()

	for i := 0; i < len(ext); i++ {
		if strings.Contains(ext[i], filetype[6:len(filetype)]) {
			return filetype, nil
		}
	}

	return "", errors.New("Invalid image type")

}
