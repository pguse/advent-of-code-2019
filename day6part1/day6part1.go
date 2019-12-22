package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var orbit []string

	//lines := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}
	lines := readLines("day06.txt")

	planets := make(map[string]string)

	// Add all planets to the map
	for _, p := range lines {
		orbit = strings.Split(p, ")")
		planets[orbit[1]] = orbit[0]
	}

	count := 0
	for key, _ := range planets {
		for {
			if planets[key] == "" {
				break
			}
			count++
			key = planets[key]
		}
	}
	fmt.Printf("Number of orbits = %v", count)
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
