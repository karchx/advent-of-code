package utils

import (
	"io"
	"os"
)

// FsReader read the file for input data
func FsReader(file string) (string, error) {
	f, err := os.Open(file)

	if err != nil {
		return "", err
	}

	defer f.Close()

	// var buf bytes.Buffer
	tee, _ := io.ReadAll(f)

	return  string(tee), nil
}
