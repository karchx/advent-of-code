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

type Move struct {
	opponentMove string
	yourMove     string
	win          string
	loss         string
	score        int
}

func TotalScoreWithELfStrategy(tee io.Reader) int {

	win := 6
	draw := 3

	rock := Move{"A", "X", "C", "B", 1}
	paper := Move{"B", "Y", "A", "C", 2}
	scissors := Move{"C", "Z", "B", "A", 3}

	sum := 0

	input := bufio.NewScanner(tee)
	input.Split(bufio.ScanLines)

	// for i := 0; i < len(input)-1; i++ {
	for input.Scan() {
		moves := strings.Split(input.Text(), " ")

		opponentMove := moves[0]
		yourMove := moves[1]

		var yourMoveType Move
		switch yourMove {
		case "X":
			yourMoveType = rock
		case "Y":
			yourMoveType = paper
		case "Z":
			yourMoveType = scissors
		default:
			fmt.Println("Uhh something is not right here, check your input.txt")
		}

		switch opponentMove {
		case yourMoveType.opponentMove:
			sum += draw
		case yourMoveType.win:
			sum += win
		case yourMoveType.loss:
			sum += 0
		default:
			fmt.Println("Uhh something is not right here, check your input.txt")
		}

		sum += yourMoveType.score
	}

	return sum
}

func main() {
	f, err := os.Open("../inputs/input-02")
	if err != nil {
		log.Fatal("failed to open file")
	}

	defer f.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(f, &buf)
	sum := TotalScoreWithELfStrategy(tee)
	fmt.Println(sum)
}
