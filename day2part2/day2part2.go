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

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if output(i, j, intCode) == 19690720 {
				fmt.Printf("noun: %v verb: % v  answer: %v", i, j, 100*i+j)
				break
			}
		}
	}
}

func output(noun int, verb int, code []int) int {
	var c []int
	for _, n := range code {
		c = append(c, n)
	}
	c[1] = noun
	c[2] = verb
	// Run the program
	for i := 0; c[i] != 99; i = i + 4 {
		if c[i] == 1 {
			c[c[i+3]] = c[c[i+1]] + c[c[i+2]]
		} else if c[i] == 2 {
			c[c[i+3]] = c[c[i+1]] * c[c[i+2]]
		}
	}

	return c[0]

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
