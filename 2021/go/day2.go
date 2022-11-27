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

func partOne(ins []inst) {
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

  fmt.Println("Part 1")
	fmt.Println(h, d, h*d)
}

func partTwo(ins []inst) {
	h := 0
	d := 0
	z := 0

	for _, i := range ins {
		switch i.dir {
		case "forward":
			h += i.val
			d += (i.val * z)
		case "up":
			z -= i.val
		case "down":
			z += i.val
		default:
			log.Fatal(i)
		}
	}

  fmt.Println("Part 2")
	fmt.Println(h, d, z, h*d)
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

  partOne(ins)
  partTwo(ins)
}
