package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	DEFAULT_FILENAME           = "problems.csv"
	DEFAULT_TIME_LIMIT_SECONDS = 30.0
	DEFAULT_SHUFFLE            = false
)

type question struct {
	Text   string
	Answer string
}

func shuffleQuestions(questions []question) {
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
}

func loadQuestions(filename string, shuffle bool) []question {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filename, err)
	}

	var questions []question

	for _, row := range records {
		questions = append(questions, question{Text: row[0], Answer: strings.TrimSpace(row[1])})
	}

	if shuffle {
		shuffleQuestions(questions)
	}

	return questions
}

func displayQuestion(questionText string, quesNumber int) {
	fmt.Printf("Problem #%d:\t%s = ", quesNumber, questionText)
}

func runQuiz(questions []question, scoreChannel chan int, finishedChannel chan bool) {
	scanner := bufio.NewScanner(os.Stdin)
	correct := 0

	for i, question := range questions {
		displayQuestion(question.Text, i+1)

		scanner.Scan()
		response := scanner.Text()

		if question.Answer == response {
			correct += 1
			scoreChannel <- correct
		}
	}

	finishedChannel <- true
}

func startQuiz(questions []question, limitSeconds int) int {
	scoreChannel := make(chan int)
	finishedChannel := make(chan bool)

	go runQuiz(questions, scoreChannel, finishedChannel)

	timeout := time.After(time.Duration(limitSeconds) * time.Second)
	var correct int

	for {
		select {
		case correct = <-scoreChannel:
		case <-finishedChannel:
			return correct
		case <-timeout:
			return correct
		}
	}
}

func main() {
	filenamePtr := flag.String("csv", DEFAULT_FILENAME, "A CSV file in the format 'question,answer'")
	limitSecPtr := flag.Int("limit", DEFAULT_TIME_LIMIT_SECONDS, "The time limit for the quiz in seconds.")
	shouldShufflePtr := flag.Bool("shuffle", DEFAULT_SHUFFLE, "Should we shuffle questions before displaying them?")

	flag.Parse()

	questions := loadQuestions(*filenamePtr, *shouldShufflePtr)

	correct := startQuiz(questions, *limitSecPtr)
	fmt.Printf("\nYou scored %d out of %d.\n", correct, len(questions))
}
