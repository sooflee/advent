package problems

import (
	"bufio"
	"os"
)

func Day02(part int, file string) int {

	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	if part == 1 {

		for scanner.Scan() {
			scanner.Text()

		}
		return 0

	} else {
		return 0
	}

}
