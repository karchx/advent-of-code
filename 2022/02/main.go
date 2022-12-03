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

type Option int64
type Result int64

const (
	Rock Option = iota
	Paper
	Scissors
)

const (
	Lost Result = iota
	Draw
	Win
)

func parsePlayOpponent(input string) Option {
	switch input {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	}
	return 0
}

func parseResult(input string) Result {
	switch input {
	case "X":
		return Lost
	case "Y":
		return Draw
	case "Z":
		return Win
	}
	return 0
}

func (option Option) play(opponent Option) Result {
	if option == opponent {
		return Draw
	}

	if option.prev() == opponent {
		return Win
	}

	return Lost
}

func (option Option) getScore() int {
	switch option {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}
	return 0
}

func (result Result) getScore() int {
	switch result {
	case Lost:
		return 0
	case Draw:
		return 3
	case Win:
		return 6
	}
	return 0
}

func (option Option) next() Option {
	if option > 1 {
		return 0
	}
	return option + 1
}

func (option Option) prev() Option {
	if option < 1 {
		return 2
	}
	return option - 1
}

func (result Result) choosePlay(opponent Option) Option {
	if result == Draw {
		return opponent
	}

	if result == Win {
		return opponent.next()
	}

	if result == Lost {
		return opponent.prev()
	}

	return 0
}

func TotalScoreWithELfStrategy(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	sum := 0

	for scanner.Scan() {
		var opponent, player string
		_, err := fmt.Sscanf(scanner.Text(), "%s %s", &opponent, &player)
		if err != nil {
			return 0
		}

		optionOpponent := parsePlayOpponent(opponent)
		optionPlayer := parsePlayOpponent(opponent)

		result := optionPlayer.play(optionOpponent)

		sum += result.getScore() + optionPlayer.getScore()

	}
	return sum
}

func TotalScoreWithElfStrategyWinLose(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		var opponent, winLoseOrDraw string 
		_, err := fmt.Sscanf(scanner.Text(), "%s %s", &opponent, &winLoseOrDraw)
		if err != nil {
			return 0
		}

		opponentPlay := parsePlayOpponent(opponent)
		desiredResult := parseResult(winLoseOrDraw)

		playerPlay := desiredResult.choosePlay(opponentPlay)

		sum += desiredResult.getScore() + playerPlay.getScore()
	}

	return sum
}

func main() {
	f, err := os.Open("../inputs/input-test-02")
	if err != nil {
		log.Fatal("failed to open file")
	}

	defer f.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(f, &buf)
  got := TotalScoreWithELfStrategy(tee)
	fmt.Println(strconv.Itoa(got))
}
