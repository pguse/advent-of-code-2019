package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("day05.txt")
	// This file only has one line
	code := strings.Split(lines[0], ",")
	// Create a integer slice from the string slice
	intCode := AtoiCode(code)
	output := intCodeComputer([]int{1}, intCode)
	for _, n := range output {
		fmt.Printf("%v ", n)
	}
}

// AtoiCode converts the array from strings to integers
func AtoiCode(code []string) []int {
	var intCode []int
	for _, s := range code {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Problem converting file info. to integers.")
		} else {
			intCode = append(intCode, i)
		}
	}
	return intCode
}

func intCodeComputer(input []int, code []int) []int {
	var c []int
	var DE, F, G int
	// Stores the intCode program in memory
	for _, n := range code {
		c = append(c, n)
	}

	// Runs the intCode program
	var out []int
	i := 0
	for {
		DE = c[i] % 100
		if DE == 99 {
			break
		}
		C := (c[i]%1000 - DE) / 100
		B := (c[i]%10000 - c[i]%1000) / 1000
		// Run opCode based on parameter mode
		if DE == 1 {
			if C == 0 {
				F = c[c[i+1]]
			} else {
				F = c[i+1]
			}
			if B == 0 {
				G = c[c[i+2]]
			} else {
				G = c[i+2]
			}
			c[c[i+3]] = F + G
			i = i + 4
		} else if DE == 2 {
			if C == 0 {
				F = c[c[i+1]]
			} else {
				F = c[i+1]
			}
			if B == 0 {
				G = c[c[i+2]]
			} else {
				G = c[i+2]
			}
			c[c[i+3]] = F * G
			i = i + 4
		} else if DE == 3 {
			c[c[i+1]] = input[0]
			i = i + 2
		} else if DE == 4 {
			out = append(out, c[c[i+1]])
			i = i + 2
		}
	}
	return out
}

// readFile creates a splice of strings representing each line of the file
func readLines(f string) []string {

	var lines []string

	// Open the file
	file, err := os.Open(f)
	if err != nil {
		fmt.Println(err)
		return lines
	}
	defer file.Close()

	// Scan the file line by line
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
