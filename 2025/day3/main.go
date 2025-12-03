package main

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func main() {
	puzzle1()
	puzzle2()
}

func puzzle1() {
	file, err := os.Open("./input.txt")
	if err != nil {
		slog.Error("cannot read file", "error", err)
		os.Exit(1)
	}

	defer file.Close()

	var total int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		total += calcHighestJoltage(line)
	}

	if err := scanner.Err(); err != nil {
		slog.Error("scanning file failed", "error", err)
		os.Exit(1)
	}

	slog.Info("Puzzle 1", "total joltage", total)
}

func puzzle2() {
	file, err := os.Open("./input.txt")
	if err != nil {
		slog.Error("cannot read file", "error", err)
		os.Exit(1)
	}

	defer file.Close()

	var total int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var bankTotal strings.Builder
		curIndex := 0

		for i := range 12 {
			r := getHighestVal(line[:len(line)-(11-i)], curIndex)

			bankTotal.WriteString(strconv.Itoa(r[0]))
			curIndex = r[1] + 1
		}

		t, _ := strconv.Atoi(bankTotal.String())
		total += t
	}

	if err := scanner.Err(); err != nil {
		slog.Error("scanning file failed", "error", err)
		os.Exit(1)
	}

	slog.Info("Puzzle 2", "total joltage", total)
}

func calcHighestJoltage(bank string) (ret int) {
	b1 := bank[:len(bank)-1]
	var digits [2]int

	for i := 0; i < len(b1); i++ {
		v1, _ := strconv.Atoi(string(b1[i]))

		if v1 <= digits[0] {
			continue
		}

		digits[0] = v1
		digits[1] = 0

		for j := i + 1; j < len(bank); j++ {
			v2, _ := strconv.Atoi(string(bank[j]))

			if v2 <= digits[1] {
				continue
			}

			digits[1] = v2
		}
	}

	return digits[0]*10 + digits[1]
}

// [0] = value, [1] = found at index
func getHighestVal(bank string, fromIndex int) (ret [2]int) {
	for i := fromIndex; i < len(bank); i++ {
		v, _ := strconv.Atoi(string(bank[i]))

		if v <= ret[0] {
			continue
		}

		ret[0] = v
		ret[1] = i
	}

	return
}
