package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func LinesInFile(fileName string) []string {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	chunk := ""
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			result = append(result, scanner.Text())
			chunk = ""

		} else {
			chunk += scanner.Text()
		}

	}
	return result
}

func GetParts(boardingPass string) (string, string) {
	firstPart := boardingPass[0:7]
	secondPart := boardingPass[7:]
	return firstPart, secondPart
}

func GetRowCode(rowCode string) int {
	minRow := 0
	maxRow := 8

	for _, bits := range rowCode {
		char := string(bits)

		if char == "L" {
			maxRow = (maxRow + minRow) / 2
		} else {
			minRow = (maxRow + minRow) / 2
		}
	}

	if minRow == 0 {
		return 0
	}
	return minRow
}

func GetSeatNumber(seatCode string) int {
	minRow := 0
	maxRow := 128

	for _, bits := range seatCode {
		char := string(bits)

		if char == "F" {
			maxRow = (maxRow + minRow) / 2
		} else {
			minRow = (maxRow + minRow) / 2
		}
	}

	if minRow == 0 {
		return 0
	}
	return minRow
}

// func GetRowNumber(rowNumber string) int {
// 	return 12
// }

func main() {
	seats := []int{}
	for _, line := range LinesInFile("./text.txt") {
		seatCode, rowCode := GetParts(line)
		seatNumber := GetSeatNumber(seatCode)
		rowID := GetRowCode(rowCode)

		fmt.Println("Seat Number: " + fmt.Sprint(seatNumber))
		fmt.Println("Row Number: " + fmt.Sprint(rowID))

		id := (seatNumber * 8) + rowID
		seats = append(seats, id)
	}
	sort.Ints(seats)

	fmt.Print(seats)

}
