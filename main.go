package main

import (
	"flag"
	"fmt"

	"github.com/sooflee/advent/problems"
)

// args := flag.Args()
// file := args[0]

func main() {

	day := flag.Int("day", 1, "day value, default 1")
	part := flag.Int("part", 1, "a part value, default 1")
	flag.Parse()

	answer, _ := problems.Call(problems.Dailies, problems.GetFunc(*day), *part, "problems/"+problems.GetFile(*day))
	fmt.Println(answer[0])

}
