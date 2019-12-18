package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Point represents a point in 2-space
type Point struct {
	x     int
	y     int
	steps int
}

func main() {

	lines := readLines("day03.txt")
	wire1 := strings.Split(lines[0], ",")
	wire2 := strings.Split(lines[1], ",")
	//wire1 := []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
	//wire2 := []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}
	//wire1 := []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}
	//wire2 := []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}
	points1 := createWirePos(wire1)
	points2 := createWirePos(wire2)

	// Create a slice of common points on both wires
	var common []Point
	for _, p1 := range points1 {
		for _, p2 := range points2 {
			if p1.x == p2.x && p1.y == p2.y {
				common = append(common, Point{p1.x, p1.y, p1.steps + p2.steps})
			}
		}
	}

	// Calculate the minimum number of steps to a common point
	var min int
	if len(common) != 0 {
		min = common[0].steps
		for _, p := range common {
			if min > p.steps {
				min = p.steps
			}
		}
	}
	fmt.Println("Minimum Number of Steps to Common Point: ", min)

}

// containsPoint determines if a point is already in a Point slice
func containsPoint(points []Point, p Point) bool {
	result := false
	for _, point := range points {
		if point.x == p.x && point.y == p.y {
			result = true
			break
		}
	}
	return result
}

// createWirePos creates a slice of points where the wire is located
func createWirePos(w []string) []Point {
	p := Point{0, 0, 0}
	var points []Point
	for _, v := range w {
		runes := []rune(v)
		step, _ := (strconv.Atoi(string(runes[1:])))
		switch runes[0] {
		case 'R':
			for i := 0; i < step; i++ {
				p.x++
				p.steps++
				if !containsPoint(points, p) {
					points = append(points, p)
				}
			}
		case 'L':
			for i := 0; i < step; i++ {
				p.x--
				p.steps++
				if !containsPoint(points, p) {
					points = append(points, p)
				}
			}
		case 'U':
			for i := 0; i < step; i++ {
				p.y++
				p.steps++
				if !containsPoint(points, p) {
					points = append(points, p)
				}
			}
		case 'D':
			for i := 0; i < step; i++ {
				p.y--
				p.steps++
				if !containsPoint(points, p) {
					points = append(points, p)
				}
			}
		}
	}
	return points
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
