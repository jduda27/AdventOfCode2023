package main

import (
	"strconv"
)

func day1(inputFile string) {
	output := parseInputFileToString(inputFile)
	var sum int
	for _, line := range output {
		result := make([]string, 0, 2)
		for _, character := range line {
			switch {
			case character >= 48 && character <= 57:
				result = append(result, string(character))
				break
			case character == 111 || character == 79:
				result = append(result, "1")
			}
			// Add numeric characters to int from left to right
		}

		finalResult := result[0] + result[len(result)-1]
		final, err := strconv.Atoi(finalResult)
		if err != nil {
			panic(err)
		}
		sum += final

	}

	printAnswer(1, "calibration character", strconv.Itoa(sum), "", "")
}
