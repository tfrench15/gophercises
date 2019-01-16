package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	f, err := openFile()
	check(err)

	data, err := readFile(f)
	check(err)

	quiz := createProblems(data)
	playGame(quiz)
}

type problem struct {
	question string
	answer   string
}

func openFile() (*os.File, error) {
	return os.Open("problems.csv")
}

func readFile(file *os.File) ([][]string, error) {
	r := csv.NewReader(file)
	problems, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return problems, nil
}

func createProblems(problems [][]string) []problem {
	quiz := []problem{}
	for _, item := range problems {
		p := problem{
			question: item[0],
			answer:   item[1],
		}
		quiz = append(quiz, p)
	}
	return quiz
}

func playGame(quiz []problem) {
	correct := 0
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < len(quiz); i++ {
		fmt.Printf("Question: %s ", quiz[i].question)
		scanner.Scan()
		ans := scanner.Text()
		if ans == quiz[i].answer {
			correct++
		}
		continue
	}
	fmt.Printf("End of Quiz. Scored %d out of %d\n", correct, len(quiz))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
