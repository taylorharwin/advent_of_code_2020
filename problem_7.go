package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LinesInFile(fileName string) []string {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}

func GetRule(secondPart string) map[string]int {
	rule := make(map[string]int)

	if secondPart == " no other bags." {
		return nil
	}

	parts := strings.Split(secondPart, ",")
	for _, part := range parts {
		part = strings.ReplaceAll(part, "bags.", "")
		part = strings.ReplaceAll(part, "bags", "")
		part = strings.ReplaceAll(part, "bag.", "")
		part = strings.ReplaceAll(part, "bag", "")
		part = strings.TrimSuffix(part, " ")

		segments := strings.Split(part, " ")

		i, err := strconv.Atoi(segments[1])
		if err != nil {
			fmt.Println(err)
		} else {
			color := strings.Join(segments[2:len(segments)], " ")

			rule[color] = i
		}

	}

	return rule
}

func SplitToRuleMap(rules []string) map[string]map[string]int {
	rulesMap := make(map[string]map[string]int)

	for _, rule := range rules {
		split := strings.Split(rule, " contain")
		firstPart := strings.Split(split[0], " bags")

		firstRule := firstPart[0]

		rulesMap[firstRule] = GetRule(split[1])

	}

	return rulesMap

}

func ValidateRulesMap(rulesMap map[string]map[string]int, color string) []string {
	validColors := []string{}
	for outerKey, rule := range rulesMap {
		for key, _ := range rule {
			if key == color {
				validColors = append(validColors, outerKey)
			}
		}
	}
	return validColors
}

func Unpack(rules map[string]map[string]int, colors []string, counter []string) []string {
	recOptions := counter
	if len(recOptions) == 0 {
		fmt.Print("reached end")
		return recOptions
	} else {
		new := Unpack(rules, colors, recOptions)
		recOptions = append(recOptions, new)
	}

	res := Unpack(r)

	for _, color := range colors {
		newColors := ValidateRulesMap(rules, color)
		if len(newColors) == 0 {
			return newColors
		} else {

		}

	}
	return recOptions
}

func main() {

	rulesMap := SplitToRuleMap(LinesInFile("./text.txt"))
	initialOptions := ValidateRulesMap(rulesMap, "shiny gold")
	newOptions := []string{}
	newOptions = Unpack(rulesMap, initialOptions, newOptions)

	fmt.Print(newOptions)

}
