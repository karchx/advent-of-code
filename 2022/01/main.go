package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)


func main() {
  f, err := os.Open("../inputs/input-01")
  if err != nil {
    log.Fatal("failed to open file")
  }

  defer f.Close()

  var buf bytes.Buffer
  tee := io.TeeReader(f, &buf)

  p1, err := parseInputOne(tee)
  if err != nil {
    log.Fatal("failed get input")
  }
  fmt.Printf("p1: %d\n", p1)
}

func parseInputOne(tee io.Reader) (int, error) {
  return max(calories(tee)), nil
}

func calories(reader io.Reader) []int {
  scanner := bufio.NewScanner(reader)
  scanner.Split(bufio.ScanLines)

  var (
    total int
    calories []int
  )

  for scanner.Scan() {
    scanner.Bytes()
    txt := scanner.Text()
    if txt == "" {
      calories = append(calories, total)
      total = 0
      continue
    }

    cal, err := strconv.Atoi(txt)
    if err != nil {
      continue
    }

    total += cal
  }

  calories = append(calories, total)
  return calories
}

func max(array []int) int {
  var max int = array[0]

  for _, value := range array {
    if max < value {
      max = value
    }
  }

  return max
}
