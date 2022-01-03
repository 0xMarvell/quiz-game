package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

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
	fmt.Println(content)

}
