package main // import "test/quiz"
import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvName := flag.String("csv", "problems.csv", "a csv file in the format of `question, answer`")
	timeLimit := flag.Int("limit", 30, "the time lmit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvName)
	if err != nil {
		exit(fmt.Sprintf("failed to open file %v", *csvName))

	}
	quiz := csv.NewReader(file)

	lines, err := quiz.ReadAll()
	if err != nil {
		exit("failed to parse the csv file")
	}
	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		answerChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("you got %d answers correct out of %d", correct, len(problems))
			return
		case answer := <-answerChan:
			if answer == p.answer {
				correct++
			}
		}

	}
	fmt.Printf("you got %d answers correct out of %d", correct, len(problems))

}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{line[0], line[1]}

	}
	return problems
}

type problem struct {
	question string
	answer   string
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
