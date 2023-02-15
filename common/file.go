package common

import (
	"os"
)

func File() {
	r, _ := os.OpenFile(Filename, os.O_CREATE, 0644)
	defer func() { r.Close() }()
	_, err := r.WriteString(Webshells)
	if err != nil {
		return
	}

}
