package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(nameFile string) []string {
	var result []string

	lengthBuffer := 5

	file, err := os.Open(nameFile)
	if err != nil {
		fmt.Printf("Failed open file %s: %v", nameFile, err)
	}

	defer file.Close()

	bufer := make([]byte, lengthBuffer)

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

func main() {
	file := "../../inputzz/day1.txt"

	result := readFile(file)

	firstNumber := result[0]
	var count int = 0

	for _, number := range result {
		if number > firstNumber {
			firstNumber = number
			count++
		} else {
			firstNumber = result[0]
		}
	}
	//fmt.Print(count)
}
