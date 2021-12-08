package main

import (
	"fmt"

	"github.com/KenethSandoval/advent-of-code/go/utils"
)

type inst struct {
	dir string
	val int
}

func main() {
	result := utils.ReadFile2("../inputzz/day2.txt")
	fmt.Println(result)
}
