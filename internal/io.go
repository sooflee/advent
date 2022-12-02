package internal

import (
	"bufio"
	"os"
)

func Read(file string) []string {
	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
