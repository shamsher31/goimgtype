// Package imgtype return type of image based on mime type
package imgtype // import "github.com/shamsher31/goimgtype"

import (
	"log"
	"net/http"
	"os"
)

// Get returns the type of Image
func Get(p string) string {
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

	return filetype

}
