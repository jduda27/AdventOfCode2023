package main

import (
	"strconv"
)

func day1(inputFile string) {
	output := parseInputFileToString(inputFile)
	//TODO: @Josh - Implement the code to manipulate the input data to get the desired answer below
	var sum int
	for _, line := range output {
		result := make([]string, 0, 2)
		for _, character := range line {
			if character >= 30 && character <= 57 {
				result = append(result, string(character))
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
