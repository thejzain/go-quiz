package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problem.csv", "filename for the problems and answers")
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
	for i, j := range problems {
		fmt.Println(i, j)
	}
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
			a: line[1],
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}
