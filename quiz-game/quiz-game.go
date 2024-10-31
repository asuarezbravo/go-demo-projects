package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Main function: parses command-line flags, reads the CSV file, and runs the quiz
func main() {
	// Parse flags for CSV file name and time limit
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	// Open the CSV file and handle any errors
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	defer file.Close() // Ensure the file is closed after usage

	// Parse CSV lines into a slice of problem structs
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)

	// Start the quiz with the parsed problems and the time limit
	correct := askQuestions(problems, *timeLimit)

	// Display the final score
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

// askQuestions iterates over the list of problems, prompts the user for answers, and applies a time limit.
// Returns the number of correct answers.
func askQuestions(problems []problem, timeLimit int) int {
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	correct := 0

	// Loop over each problem, prompting the user for an answer
problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)

		// Use a goroutine to read user input concurrently
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		// Select between receiving input or timeout signal
		select {
		case <-timer.C:
			fmt.Println("\nTime's up!") // Inform user when time limit is reached
			break problemloop           // Exit the loop if timer expires
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}

	return correct // Return the total number of correct answers
}

// parseLines converts CSV lines into a slice of problem structs with trimmed answers.
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]), // Trim whitespace for consistency in answers
		}
	}
	return ret
}

// problem struct defines a quiz question and answer
type problem struct {
	q string // Question
	a string // Answer
}

// exit prints an error message and exits the program
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
