package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
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
	timeLimit := flag.Int("limit", 30, "Time allocated for the quiz (in seconds)")
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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, v := range questions {
		fmt.Printf("Question %d: %s = ", i+1, v.ques)
		answerChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(questions))
			return
		case answer := <-answerChan:
			if answer == v.ans {
				correct += 1
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(questions))

}
