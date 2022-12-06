package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type Range struct {
	start int
	end   int
}

func (p Range) len() int {
	return p.end - p.start
}

func PairsContainsPairs(tee io.Reader) int {
	scanner := bufio.NewScanner(tee)
	scanner.Split(bufio.ScanLines)

	sum := 0

	for scanner.Scan() {
		var firstElf, secondElf Range
		_, err := fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &firstElf.start, &firstElf.end, &secondElf.start, &secondElf.end)

		if err != nil {
			return 0
		}

		if firstElf.len() < secondElf.len() {
			tmp := firstElf
			firstElf = secondElf
			secondElf = tmp
		}

		if firstElf.start <= secondElf.start && firstElf.end >= secondElf.end {
			sum++
		}
	}
	return sum
}

func main() {
	f, err := os.Open("../inputs/input-04")
	if err != nil {
		log.Fatal("failed to open file")
	}

	defer f.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(f, &buf)

	total := PairsContainsPairs(tee)
	fmt.Printf("Part 01: %d\n", total)

	//sum2 := OrganizeBadgets(&buf)
	//fmt.Printf("Part 02: %d\n", &)
}
