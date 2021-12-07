package main

import (
	"fmt"
	"io"
	"os"
)


const file = "../inputzz/day1.txt"

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

func problem2() {
  result := readFile(file)

  data := result[0] + result[1] + result [2]
  fmt.Print(data)
}

func main() {
	result := readFile(file)

	firstNumber := result[0]
  count := 0

  for i := 1; i < len(result); i++ {
    if result [i] > firstNumber {
      count++
    }

    firstNumber = result[i]
  }

  problem2()
  fmt.Println(count)
}
