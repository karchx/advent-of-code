package main

import (
	"fmt"
	"github.com/KenethSandoval/advent-of-code/utils"
)


const file = "../inputzz/day1.txt"

func problem2() {
  result := utils.ReadFile(file)

  data := result[0] + result[1] + result [2]
  fmt.Print(data)
}

func main() {
	result := utils.ReadFile(file)

	firstNumber := result[0]
	count := 0

  for i := 1; i < len(result); i++ {
    if result [i] > firstNumber {
      count++
    }

    firstNumber = result[i]
  }

  fmt.Println(count)
}
