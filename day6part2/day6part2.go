package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var orbit []string

	//lines := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"}
	lines := readLines("day06.txt")

	planets := make(map[string]string)

	// Add all planets to the map
	for _, p := range lines {
		orbit = strings.Split(p, ")")
		planets[orbit[1]] = orbit[0]
	}

	// Store all the planets that YOU indirectly/directly orbits
	var YouPlanets []string

	for k := "YOU"; k != ""; {
		YouPlanets = append(YouPlanets, planets[k])
		k = planets[k]
		//fmt.Printf("%v ", k)
	}
	//fmt.Println()

	// Store all the planets that SAN indirectly/directly orbits
	var SanPlanets []string

	for k := "SAN"; k != ""; {
		SanPlanets = append(SanPlanets, planets[k])
		k = planets[k]
		//fmt.Printf("%v ", k)
	}
	//fmt.Println()

	// Find the shortest path between YOU and SAN
	min := false
	for i, yp := range YouPlanets {
		for j, sp := range SanPlanets {
			if yp == sp {
				fmt.Printf("Minimum # of orbital transfers: %v", i+j)
				min = true
				break
			}
		}
		if min {
			break
		}
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
