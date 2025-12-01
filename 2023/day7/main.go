package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type card struct {
	input string
	bid   int
	hand  int
}

var strenghts map[string]int = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

func main() {
	// Read file
	file, err := os.Open("./input.txt")
	if err != nil {
		slog.Error("cannot read file", "error", err)
		os.Exit(1)
	}

	defer file.Close()

	// var cards []card

	// Scan the file lines
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var c card

		spl := strings.Split(line, " ")
		bid, _ := strconv.Atoi(spl[1])

		c.input = spl[0]
		c.bid = bid

		fmt.Printf("\n%+v --- %+v", c.input, c.bid)

	}

	fmt.Printf("\n\n\n")
	if err := scanner.Err(); err != nil {
		slog.Error("scanning file failed", "error", err)
		os.Exit(1)
	}
}

func calcHand(input string) int {
	counter := map[int]

	for _, char := range input {

	}

}
