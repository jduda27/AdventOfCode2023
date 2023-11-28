# AdventOfCode2023 - Go
A Golang application for solving Advent of Code 2023 problems.
https://adventofcode.com/
## Goals
- Complete all of the daily Advent of Code Challenges
- Learn how to build a mix of different data structures and algorithms in Golang
- Besides very basic utility functions such as parsing input files and printing the results,
  do not copy over code from previous days.
- Create .md documentation for each day that explains the design process of each solution.
- Each day should be it's own branch.
- If there are multiple logical solutions to a problem, try using a solutino that has not been
  done so far in the project to practice different design patterns, data structures, and algorithms

## Useful Information
  ### Boiler Plate for Daily Project Files
  ``` Go
  import (
        "fmt"
  )
  
  func day1(inputFile string) {
    output := parseInputFileToString(inputFile) //Can also use parseInputFileToInt(inputFile)
	  //TODO: Implement the code to manipulate the input data to get the desired answer below

	  printAnswer({int day number}, "{answer unit type}", "{answer}", "{2nd aut}", "{2nd answer}")
  }
  ```
