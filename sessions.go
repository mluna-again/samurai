package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func loadSessions() ([][]string, error) {
	lines := []string{}
	counter := 0
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		line := reader.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if counter > 9 {
			break
		}
		lines = append(lines, line)
		counter++
	}
	if err := reader.Err(); err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}

	mapped := [][]string{}
	for _, line := range lines {
		cmps := strings.Split(line, separator)
		if len(cmps) < 2 {
			mapped = append(mapped, []string{line, ""})
			continue
		}

		right := cmps[1]
		if rightDates != "" && right != "" {
			epoch, err := strconv.Atoi(right)
			if err != nil {
				return nil, err
			}
			t := time.Unix(int64(epoch), 0)

			right = t.Format(rightDates)
		}
		mapped = append(mapped, []string{cmps[0], right})
	}

	return mapped, nil
}
