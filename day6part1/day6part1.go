package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var orbit []string

	lines := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}
	//lines := readLines("day06.txt")

	planets := make(map[string]int)

	for _, p := range lines {
		orbit = strings.Split(p, ")")
		planets[orbit[0]] = 0
		fmt.Printf("%v : %v\n", orbit[0], orbit[1])
	}
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
