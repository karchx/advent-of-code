package main

import (
	"fmt"
	"strings"

	"github.com/KenethSandoval/advent-of-code/go/utils"
)

func main() {
	result := utils.ReadFile2("../inputzz/test.txt")

	for _, s := range result {
		split := strings.Split(s, "")
		fmt.Println(split[0])
	}
}
