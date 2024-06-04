package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func parseInputFileToString(filePath string) []string {
	var output []string
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			output = append(output, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return output
}

func parseInputFileToInt(filePath string) []int {
	var output []int
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			number, err1 := strconv.Atoi(scanner.Text())
			if err1 != nil {
				panic(err1)
			}
			output = append(output, number)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return output
}

/**
 *
 */
func printAnswer(dayNum int, answerType1 string, answer1 string, answerType2 string, answer2 string) {
	fmt.Printf("Advent of Code Day %d\n---------------------\n", dayNum)
	if answerType1 != "" && answer1 != "" {
		fmt.Printf("%s: %s\n", answerType1, answer1)
	}
	if answerType2 != "" && answer2 != "" {
		fmt.Printf("\n%s: %s", answerType2, answer2)
	}
}

func executeChallenge(challengeNum int, isTest bool) {
	testPath := "Media/txt/test-inputs/"
	path := "Media/txt/"
	switch challengeNum {
	case 1:
		if isTest {
			day1(testPath + "day1.txt")
		} else {
			day1(path + "day" + strconv.Itoa(challengeNum) + ".txt")
		}
	case 2:
		//Put code for execution here
		fmt.Println("Challenge 2")
	default:
		fmt.Println("Challenge Number does not exist.")
	}
}
