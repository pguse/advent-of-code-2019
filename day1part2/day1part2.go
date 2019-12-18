package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	var fuel []int

	lines := readLines("day01.txt")

	for _, line := range lines {
		if n, err := strconv.Atoi(line); err == nil {
			// Determine the fuel requirements for each module
			counter := 0
			for {
				n = int(math.Floor(float64(n)/3)) - 2
				if n > 0 {
					counter = counter + n
				} else {
					break
				}
			}
			fuel = append(fuel, counter)
		} else {
			fmt.Println(line, "is not an integer.")
		}

	}

	// Determine the fuel required for all modules
	counter := 0
	for _, value := range fuel {
		counter = counter + value
	}

	fmt.Println("Fuel required:", counter)
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
