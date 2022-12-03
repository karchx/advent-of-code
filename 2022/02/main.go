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

type Move2 struct {
	opponentMove string
	win          int
	draw         int
	loss         int
}

const (
	win  = 6
	draw = 3
)

func TotalScoreWithELfStrategy(tee io.Reader) int {

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

func TotalScoreWithELfStrategyWinLose(tee io.Reader) int {
	rockScore := 1
	paperScore := 2
	scissorsScore := 3

	rock := Move2{"A", paperScore, rockScore, scissorsScore}
	paper := Move2{"B", scissorsScore, paperScore, rockScore}
	scissors := Move2{"C", rockScore, scissorsScore, paperScore}

	sum := 0

	input := bufio.NewScanner(tee)
	input.Split(bufio.ScanLines)

	// for i := 0; i < len(input)-1; i++ {
	for input.Scan() {
		moves := strings.Split(input.Text(), " ")

		opponentMove := moves[0]
		outCome := moves[1]

		var opponentMoveType Move2
		switch opponentMove {
		case "A":
			opponentMoveType = rock
		case "B":
			opponentMoveType = paper
		case "C":
			opponentMoveType = scissors
		default:
			fmt.Println("Uhh something is not right here, check your input.txt")
		}

		switch outCome {
		case "X":
			sum += opponentMoveType.loss
		case "Y":
			sum += opponentMoveType.draw + draw
		case "Z":
			sum += opponentMoveType.win + win
		default:
			fmt.Println("Uhh something is not right here, check your input.txt")
		}
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
	fmt.Printf("Part 01: %d\n", sum)

	sum2 := TotalScoreWithELfStrategyWinLose(&buf)
	fmt.Printf("Part 02: %d\n", sum2)
}
