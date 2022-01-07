package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	//"time"
)

type problem struct {
	ques string
	ans  string
}

func parseContent(lines [][]string) []problem {
	val := make([]problem, len(lines))
	for i, line := range lines {
		val[i] = problem{ques: line[0], ans: line[1]}
	}
	return val
}

func exit(exitMessage string) {
	fmt.Println(exitMessage)
	os.Exit(1)
}

func main() {

	//fmt.Println("Hello, quiz :)")
	csvFilename := flag.String("csv", "problems.csv", "a CSV file with text written in a 'Question, Answer' format")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		msg := fmt.Sprintf("Failed to open the CSV file: %v\n", *csvFilename)
		exit(msg)
	}

	r := csv.NewReader(file)
	content, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	questions := parseContent(content)

	correct := 0
	for i, v := range questions {
		fmt.Printf("Question %d: %s = \n", i+1, v.ques)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == v.ans {
			correct += 1
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(questions))

}
