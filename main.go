package main

/* EXERCISE DETAILS
This exercise is broken into two parts to help simplify the process of explaining it as well as to make
it easier to solve. The second part is harder than the first, so if you get stuck feel free to move on
to another problem then come back to part 2 later.

PART 1
Create a program that will read in a quiz provided via a CSV file (more details below) and will then
give the quiz to a user keeping track of how many questions they get right and how many they get
incorrect. Regardless of whether the answer is correct or wrong the next question should be asked
immediately afterwards.

The CSV file should default to problems.csv (example shown below), but the user should be able to
customize the filename via a flag.

The CSV file will be in a format like below, where the first column is a question and the second
column in the same row is the answer to that question.

5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7

You can assume that quizzes will be relatively short (< 100 questions) and will have single
word/number answers.
At the end of the quiz the program should output the total number of questions correct and how many
questions there were in total. Questions given invalid answers are considered incorrect.

NOTE: CSV files may have questions with commas in them. Eg: "what 2+2, sir?",4 is a valid row in a CSV.
I suggest you look into the CSV package in Go and don't try to write your own CSV parser.

PART 2
Adapt your program from part 1 to add a timer. The default time limit should be 30 seconds, but should
also be customizable via a flag.
Your quiz should stop as soon as the time limit has exceeded. That is, you shouldn't wait for the user
to answer one final questions but should ideally stop the quiz entirely even if you are currently
waiting on an answer from the end user.

Users should be asked to press enter (or some other key) before the timer starts, and then the questions
should be printed out to the screen one at a time until the user provides an answer. Regardless of
whether the answer is correct or wrong the next question should be asked.
At the end of the quiz the program should still output the total number of questions correct and how
many questions there were in total. Questions given invalid answers or unanswered are considered
incorrect.

BONUS
As a bonus exercises you can also...
Add string trimming and cleanup to help ensure that correct answers with extra whitespace,
capitalization, etc are not considered incorrect. Hint: Check out the strings package.
Add an option (a new flag) to shuffle the quiz order each time it is run. */

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
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

func shuffleQuiz(val []problem) []problem {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(val), func(x, y int) {
		val[x], val[y] = val[y], val[x]
	})
	return val
}

func exit(exitMessage string) {
	fmt.Println(exitMessage)
	os.Exit(1)
}

func main() {

	//Create all flags needed for the quiz game
	csvFilename := flag.String("csv", "problems.csv", "CSV file with text written in a 'Question, Answer' format")
	timeLimit := flag.Int("limit", 30, "Time allocated for the quiz (in seconds)")
	shuffle := flag.Bool("shuffle", true, "Shuffles the order of the questions displayed from the quiz file")
	flag.Parse()

	// Open the provided CSV file
	file, err := os.Open(*csvFilename)
	if err != nil { //if there is an error when opening CSV file (e.g. CSV file does not exist)
		msg := fmt.Sprintf("Failed to open the CSV file: %v\n", *csvFilename)
		exit(msg)
	}

	r := csv.NewReader(file)
	content, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	questions := parseContent(content)

	if *shuffle { // if shuffle flag has boolean value of "true"
		shuffleQuiz(questions)
	}

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
