package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)


func offsetRune(char rune) int {
	if char > 64 && char < 92 {
		return int(char) - 65 + 27
	}

	if char > 96 && char < 123 {
		return int(char) - 97 + 1
	}

	return 0
}

func TotalPriorities(tee io.Reader) int {
	scanner := bufio.NewScanner(tee)
	scanner.Split(bufio.ScanLines)

	sum := 0

	for scanner.Scan() {
		backPacks := scanner.Text()

		firstPack := backPacks[:len(backPacks)/2]
		secondPack := backPacks[len(backPacks)/2:]

    for _, char := range firstPack {
      if strings.IndexRune(secondPack, char) < 0 {
        continue
      }
      sum += offsetRune(char)
      break
    }
	}

	return sum
}

func OrganizeBadgets(tee io.Reader) int {
  scanner := bufio.NewScanner(tee)
  scanner.Split(bufio.ScanLines)

  sum := 0
  var groups []string
  for scanner.Scan() {
    groups = append(groups, scanner.Text())

    if len(groups) != 3 {
      continue
    }

    var matchedFirst []rune

    for _, char := range groups[0] {
      if strings.IndexRune(groups[1], char) < 0 {
        continue
      }

      matchedFirst = append(matchedFirst, char)
    }

    for _, char := range groups[2] {
      if strings.IndexRune(string(matchedFirst), char) < 0 {
        continue
      }
      sum += offsetRune(char)
      break
    }

    groups = nil
  }
  return sum
}

func main() {
	f, err := os.Open("../inputs/input-03")
	if err != nil {
		log.Fatal("failed to open file")
	}

	defer f.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(f, &buf)

	sum := TotalPriorities(tee)
	fmt.Printf("Part 01: %d\n", sum)

	sum2 := OrganizeBadgets(&buf)
	fmt.Printf("Part 02: %d\n", sum2)
}
