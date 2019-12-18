package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	counter := 0

	lines := readLines("day01.txt")

	for _, line := range lines {
		if n, err := strconv.Atoi(line); err == nil {
			counter = counter + (int(math.Floor(float64(n)/3)) - 2)
		} else {
			fmt.Println(line, "is not an integer.")
		}

	}

	fmt.Println("Fuel required:", counter)
}

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
