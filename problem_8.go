package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RulesInFile(fileName string) []map[string]int {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []map[string]int{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		partOne := parts[0]
		i, _ := strconv.Atoi(parts[1])

		var m = map[string]int{
			partOne: i,
		}

		result = append(result, m)
	}
	return result
}

func ProcessRule(rulesMap []map[string]int, rule map[string]int, accumulator int, index int) (map[string]int, int, int) {
	for instruction, number := range rule {
		if instruction == "jmp" {
			index = index + number
		}
		if instruction == "acc" {
			accumulator += number
			index = index + 1

		}
		if instruction == "nop" {
			index = index + 1
		}
	}

	return rulesMap[index], index, accumulator
}

func GetInfiniteLoop(rulesMap []map[string]int) int {
	newAccumulator := 0
	currentIndex := 0
	currentRule := rulesMap[currentIndex]

	seenRules := map[int]int{}

	for newAccumulator <= 500 {
		currentRule, currentIndex, newAccumulator = ProcessRule(rulesMap, currentRule, newAccumulator, currentIndex)

		seenRules[currentIndex]++

		if seenRules[currentIndex] > 1 {
			return newAccumulator
			break
		}
	}
	return 0
}

func main() {
	rulesMap := RulesInFile("./text.txt")

	for index := range rulesMap {
		copy := rulesMap
		for _, rule := range copy {
			for instruction, number := range rule {
				if instruction == "jmp" {
					copy[index] = map[string]int{"nop": number}
					newCount := GetInfiniteLoop(copy)
					if newCount == 0 {
						fmt.Print("Solved Problem!")
					}
				}
			}

		}

	}

}
