package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) //get all the input in a buffer

	// Read the number of teachers
	fmt.Println("How many teachers do you have in your school?")
	nStr, _ := reader.ReadString('\n') // read input till first enter button
	nStr = strings.TrimSpace(nStr)
	n, err := strconv.Atoi(nStr) // turn input to integer
	if err != nil {
		fmt.Println("Invalid number of teachers")
		return
	}

	teacher := make([]string, n)
	t_score := make([]string, n)

	for i := 0; i < n; i++ {
		// Read teacher's name
		fmt.Printf("Tell me the name of teacher %d: ", i+1)
		name, _ := reader.ReadString('\n')
		teacher[i] = strings.TrimSpace(name)

		// Read scores
		fmt.Println("Enter the scores of students separated by spaces:")
		scoresLine, _ := reader.ReadString('\n')
		scores := strings.Fields(scoresLine)

		// Calculate sum of scores
		sum := 0
		for _, scoreStr := range scores {
			score, err := strconv.Atoi(scoreStr)
			if err != nil {
				fmt.Println("Invalid score:", scoreStr)
				return
			}
			sum += score
		}

		// Calculate average
		avg := sum / len(scores)

		// Determine rating
		switch {
		case avg >= 80:
			t_score[i] = "Excellent"
		case avg >= 60:
			t_score[i] = "Very Good"
		case avg >= 40:
			t_score[i] = "Good"
		default:
			t_score[i] = "Fair"
		}
	}

	// Print the results
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s\n", teacher[i], t_score[i])
	}
}
