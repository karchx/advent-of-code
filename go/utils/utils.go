package utils

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(nameFile string) []string {
	var result []string

	lengthBufer := 5
	file, err := os.Open(nameFile)
	if err != nil {
		fmt.Printf("Failed open file %s: %v", nameFile, err)
	}

	defer file.Close()

	bufer := make([]byte, lengthBufer)

	for {
		bytesRead, err := file.Read(bufer)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Failed read content: %v", err)
			}
			break
		}
		fragment := string(bufer[:bytesRead])
		result = append(result, fragment)
	}
	return result
}
