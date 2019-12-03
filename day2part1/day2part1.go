package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	lines := readLines("day02.txt")

	// This file only has one line
	code := strings.Split(lines[0], ",")
	// Create a integer slice from the string slice
	var intCode []int

	for _, s := range code {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Problem converting file info. to integers.")
		} else {
			intCode = append(intCode, i)
		}
	}

	// Restore "1202 program alarm" state
	intCode[1] = 12
	intCode[2] = 2
	// Run the program
	for i := 0; intCode[i] != 99; i = i + 4 {
		//fmt.Printf("%v %v %v %v", intCode[i], intCode[i+1], intCode[i+2], intCode[i+3])
		//fmt.Println()
		if intCode[i] == 1 {
			intCode[intCode[i+3]] = intCode[intCode[i+1]] + intCode[intCode[i+2]]
		} else if intCode[i] == 2 {
			intCode[intCode[i+3]] = intCode[intCode[i+1]] * intCode[intCode[i+2]]
		}
	}

	fmt.Println("Intial Value:", intCode[0])
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
