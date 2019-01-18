package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

var timeLimit = flag.Int("limit", 30, "sets the time limit for the quiz")

func main() {
	flag.Parse()
	f, err := openFile()
	check(err)

	data, err := readFile(f)
	check(err)

	quiz := newQuiz(data)
	tl := *timeLimit
	limit := time.Duration(tl) * time.Second
	quiz.play(limit)
}

type quiz struct {
	problems []problem
	score    int
}

type problem struct {
	question string
	answer   string
}

func (q *quiz) play(limit time.Duration) {
	scanner := bufio.NewScanner(os.Stdin)
	ch := make(chan int)
	go func() {
		for i := 0; i < len(q.problems); i++ {
			fmt.Printf("Question %d: %s ", i+1, q.problems[i].question)
			scanner.Scan()
			ans := scanner.Text()
			if ans == q.problems[i].answer {
				q.score++
			}
			continue
		}
		ch <- 1
	}()
	select {
	case <-ch:
		fmt.Printf("Quiz Complete: Scored %d out %d\n", q.score, len(q.problems))
		return
	case <-time.After(limit):
		fmt.Printf("Quiz Complete: Scored %d out %d\n", q.score, len(q.problems))
		return
	}
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

func newQuiz(problems [][]string) quiz {
	list := []problem{}
	for _, item := range problems {
		p := problem{
			question: item[0],
			answer:   item[1],
		}
		list = append(list, p)
	}
	q := quiz{
		problems: list,
		score:    0,
	}
	return q
}

// handle errs.
func check(err error) {
	if err != nil {
		panic(err)
	}
}
