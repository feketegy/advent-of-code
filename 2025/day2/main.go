package main

import (
	"bytes"
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
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		slog.Error("cannot read file", "error", err)
		os.Exit(1)
	}

	sections := bytes.SplitSeq(file, []byte(","))
	var reps uint64

	for sect := range sections {
		ids := bytes.Split(sect, []byte("-"))

		// clean the input
		s := strings.TrimRight(string(ids[0]), "\n")
		e := strings.TrimRight(string(ids[1]), "\n")

		start, _ := strconv.ParseUint(s, 10, 64)
		end, _ := strconv.ParseUint(e, 10, 64)

		for i := start; i <= end; i++ {
			si := strconv.FormatUint(i, 10)

			// odd digits can't repeat
			l := len(si)
			if l%2 != 0 {
				continue
			}

			if si[:l/2] == si[l/2:] {
				reps += i
			}
		}
	}

	slog.Info("Puzzle 1", "repeating sum", reps)
}

func puzzle2() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		slog.Error("cannot read file", "error", err)
		os.Exit(1)
	}

	sections := bytes.SplitSeq(file, []byte(","))
	var reps uint64

	for sect := range sections {
		ids := bytes.Split(sect, []byte("-"))

		// clean the input
		s := strings.TrimRight(string(ids[0]), "\n")
		e := strings.TrimRight(string(ids[1]), "\n")

		start, _ := strconv.ParseUint(s, 10, 64)
		end, _ := strconv.ParseUint(e, 10, 64)

		for i := start; i <= end; i++ {
			si := strconv.FormatUint(i, 10)

			rep := isRepeating(si)

			if rep == true {
				reps += i
			}
		}
	}

	slog.Info("Puzzle 2", "repeating sum", reps)

}

func isRepeating(s string) bool {
	n := len(s)

	for d := 1; d <= n/2; d++ {
		if n%d != 0 {
			continue
		}

		pattern := s[:d]
		repeatCount := n / d

		ok := true
		for i := 1; i < repeatCount; i++ {
			if s[i*d:(i+1)*d] != pattern {
				ok = false
				break
			}
		}

		if ok {
			return true
		}
	}

	return false
}
