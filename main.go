package main

import (
	"flag"
	"fmt"

	"github.com/sooflee/advent/internal"
	"github.com/sooflee/advent/problems"
)

// args := flag.Args()
// file := args[0]

func main() {

	day := flag.Int("day", 1, "day value, default 1")
	part := flag.Int("part", 1, "a part value, default 1")
	flag.Parse()

	answer, _ := internal.Call(problems.Dailies, internal.GetFunc(*day), *part, "problems/"+internal.GetFile(*day))
	fmt.Println(answer[0])

}
