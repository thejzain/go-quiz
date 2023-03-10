package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "filename for the problems and answers")
	num := flag.Int("num", 12, "Number of questions")
	_ = num
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("The file %s couldnt read ", *csvFilename))
	}
	r := csv.NewReader(file)
	line, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse file")
	}
	problems := makeProblem(line)
	correct := 0
	for i, p := range problems {
		fmt.Printf("#%d Whats %s = ", (i + 1), p.q)
		var ans string
		fmt.Scanf("%s\n", &ans)
		if ans == p.a {
			correct++
		}
	}
	fmt.Printf("\nYour Scored %d out of %d", correct, len(line))
}

type problem struct {
	q string
	a string
}

func makeProblem(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}
