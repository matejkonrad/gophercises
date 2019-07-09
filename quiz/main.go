package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Problem type is a quizz problem with a question and a correct answer
type Problem struct {
	id, question, answer string
}

// Answer is for keeping track if the user entered a correct value
type Answer struct {
	problem Problem
	correct bool
}

func parseCsv() []Problem {
	f, err := os.Open("problems.csv")
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(f))

	var problems []Problem
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		problems = append(problems, Problem{
			question: line[0],
			answer:   line[1]})
	}

	return problems
}

func readValue() string {
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	return name
}

func main() {
	problems := parseCsv()

	var answers []Answer

	for _, prb := range problems {
		fmt.Println("What is " + prb.question + "?")
		userAnswer := strings.TrimRight(readValue(), "\n")
		isCorrect := prb.answer == userAnswer
		answers = append(answers, Answer{problem: prb, correct: isCorrect})
	}

	fmt.Println(answers)
}
