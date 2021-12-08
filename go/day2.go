package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/KenethSandoval/advent-of-code/go/utils"
)

type inst struct {
	dir string
	val int
}

func main() {
	result := utils.ReadFile2("../inputzz/day2.txt")

	var ins []inst
	for _, s := range result {
		split := strings.Split(s, " ")
		v, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err.Error())
		}
		ins = append(ins, inst{split[0], v})
	}

	h := 0
	d := 0

	for _, i := range ins {
		switch i.dir {
		case "forward":
			h += i.val
		case "up":
			d -= i.val
		case "down":
			d += i.val
		default:
			log.Fatal(i)
		}
	}
	fmt.Println(h, d, h*d)
}
