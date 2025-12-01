package main

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		slog.Error("cannot read file", "error", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	puzzle1(scanner)

	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
	puzzle2(scanner)
}

func puzzle1(scanner *bufio.Scanner) {
	boundary := 100
	cursor := 50
	counter := 0

	for scanner.Scan() {
		line := scanner.Text()
		steps, _ := strconv.Atoi(line[1:])

		// direction left
		if line[0] == 'L' {
			steps = -steps
		}

		cursor = (cursor + steps) % boundary
		if cursor < 0 {
			cursor += boundary
		}

		if cursor == 0 {
			counter += 1
		}
	}

	if err := scanner.Err(); err != nil {
		slog.Error("scanning file failed", "error", err)
		os.Exit(1)
	}

	slog.Info("Puzzle 1", "password", counter)
}

func puzzle2(scanner *bufio.Scanner) {
	boundary := 99
	cursor := 50
	counter := 0

	for scanner.Scan() {
		line := scanner.Text()

		steps, _ := strconv.Atoi(line[1:])
		stepper := 1
		if line[0] == 'L' {
			stepper = -1
		}

		for range steps {
			cursor += stepper
			if cursor > boundary {
				cursor = 0
			} else if cursor < 0 {
				cursor = boundary
			}

			if cursor == 0 {
				counter += 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		slog.Error("scanning file failed", "error", err)
		os.Exit(1)
	}

	slog.Info("Puzzle 2", "password", counter)
}
