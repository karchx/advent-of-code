package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
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

func ReadFile2(nameFile string) []string {
	file, err := os.Open(nameFile)
	var data []string

	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	//buffer := bufio.NewReader(file)
	//line, _, err := buffer.ReadLine()
	buffer := bufio.NewScanner(file)

	for buffer.Scan() {
		data = append(data, buffer.Text())
	}
	return data
}

func ReadFileToStringsSlice(f *os.File)[]string {
  scanner := bufio.NewScanner(f)
  slice := make([]string, 0, 100)

  for scanner.Scan() {
    slice = append(slice, scanner.Text())
  }
  return slice
}
