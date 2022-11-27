package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/KenethSandoval/advent-of-code/go/utils"
)

func main() {
  f, err := os.Open("../inputzz/day3.txt")
  if err != nil {
    panic(err)
  }

  bits := utils.ReadFileToStringsSlice(f)
  fmt.Printf("%d\n", firstPart(bits))
}

func firstPart (bitLines []string) uint {
  counts := make(map[int]map[rune]int)
  dataSize := len(bitLines[0])
  for i := 0; i < dataSize; i++ {
    counts[i] = make(map[rune]int)
  }

  for _, line := range bitLines {
    for j, bit := range line {
      counts[j][bit]++
    }
  }

  var gb, eb strings.Builder
  for i := 0; i < dataSize; i++ {
    if counts[i]['0'] > counts[i]['1'] {
      gb.WriteRune('0')
      eb.WriteRune('1')
    } else {
      gb.WriteRune('1')
      eb.WriteRune('0')
    }
  }
  gamma, err := strconv.ParseInt(gb.String(), 2, 64)
  if err != nil {
    panic(err)
  }
  epsilon, err := strconv.ParseInt(eb.String(), 2, 64)
  if err != nil{
    panic(err)
  }

  return uint(gamma) * uint(epsilon)
}
