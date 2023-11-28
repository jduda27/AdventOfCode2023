package main

import (
	"fmt"
	"reflect"
)

func day1(inputFile string) {
	output := parseInputFileToString(inputFile)
	fmt.Println(reflect.TypeOf(output))
	//TODO: @Josh - Implement the code to manipulate the input data to get the desired answer below

	printAnswer(1, "answerType", "answer", "", "")
}
