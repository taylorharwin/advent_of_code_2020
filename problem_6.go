package main

import (
	"bufio"
	"fmt"
	"os"
)

func LinesInFile(fileName string) [][]string {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := [][]string{}
	// Use Scan.
	var chunk []string
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			result = append(result, chunk)
			chunk = nil

		} else {
			chunk = append(chunk, scanner.Text())
		}

	}
	return result
}

func CountAllYes(group []string) int {

	fmt.Println(group)
	answersMap := make(map[string]int)
	peopleCount := len(group)
	answerCount := 0

	for _, answer := range group {
		answerString := string(answer)

		for _, char := range answerString {
			a := string(char)
			if count, ok := answersMap[a]; ok {
				answersMap[a] = count + 1
			} else {
				answersMap[a] = 1
			}
		}

		for _, count := range answersMap {
			if count == peopleCount {
				answerCount++
			}
		}
	}

	return answerCount
}
func main() {
	numYes := 0
	for _, group := range LinesInFile("./text.txt") {
		countAllYesAnswers := CountAllYes(group)
		fmt.Println(countAllYesAnswers)
		numYes += countAllYesAnswers
	}

	fmt.Println(fmt.Sprintf("Total Group Yes Answers %v", numYes))

}
