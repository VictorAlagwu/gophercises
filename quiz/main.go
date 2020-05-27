package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

//Problem :
type problem struct {
	q string
	a string
}
func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limit := flag.Int("limit", 30, "Get the limit for the question")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	problems := parseLines(lines)
   
	timer := time.NewTimer(time.Duration(*limit) * time.Second)
	correct := 0

	problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		//Make answer channel
		answerCh := make(chan string)
		 
		// Anonymous Function
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			//Assign answer to answer channel
			answerCh <- answer
		}()
		select {
			case <-timer.C: //Listen to timer channel
				fmt.Println()
				break problemloop
			case answer := <-answerCh: //Get answer sent to answer channel
				if answer == p.a {
					correct++
				}
		}
		
	}

	fmt.Printf("You were able to get %d correct answers, in %d questions \n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}