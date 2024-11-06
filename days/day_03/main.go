package day_03

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}

type coordinate struct {
	x int
	y int
}

type numberGroup struct {
	coordinates []coordinate
	value       int
}

func findSymbolPositions(s string) []int {
	re := regexp.MustCompile(`[^a-zA-Z0-9\s.,!?;:'"(){}\[\]<>]`)
	matches := re.FindAllStringIndex(s, -1)
	positions := make([]int, len(matches))
	for i, match := range matches {
		positions[i] = match[0]
	}
	return positions
}

func findSymbolCoordinates(s string, y int) []coordinate {
	positions := findSymbolPositions(s)
	coordinates := make([]coordinate, len(positions))
	for i, pos := range positions {
		coordinates[i] = coordinate{x: pos, y: y}
	}
	return coordinates
}

func findNumberGroupsPositions(s string) [][]int {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllStringIndex(s, -1)
	var groups [][]int
	for _, match := range matches {
		var group []int
		for i := match[0]; i < match[1]; i++ {
			group = append(group, i)
		}
		groups = append(groups, group)
	}
	return groups
}

func findNumberGroupCoordinates(s string, y int) [][]coordinate {
	groups := findNumberGroupsPositions(s)
	coordinates := make([][]coordinate, len(groups))
	for i, group := range groups {
		coordinates[i] = makeCoordinateList(group, y)
	}
	return coordinates
}

// isAdjacentTo checks if the coordinate c is adjacent to the coordinate d.
// Two coordinates are considered adjacent if they are next to each other
// horizontally, vertically, or diagonally, but not the same coordinate.
//
// The function returns true if the coordinates are adjacent, and false otherwise.
//
// Example:
//
//	c: (2, 3), d: (3, 3) -> true
//	c: (2, 3), d: (2, 2) -> true
//	c: (2, 3), d: (1, 1) -> false
//	c: (2, 3), d: (2, 3) -> false
func (c *coordinate) isAdjacentTo(d coordinate) bool {
	if c.x == d.x && c.y == d.y {
		return false
	}
	dx := abs(c.x - d.x)
	dy := abs(c.y - d.y)
	return dx <= 1 && dy <= 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (ng *numberGroup) isAdjacentToSymbol(symbolPositions []coordinate) bool {
	for _, coord := range ng.coordinates {
		for _, sym := range symbolPositions {
			if coord.isAdjacentTo(sym) {
				return true
			}
		}
	}
	return false
}

func makeCoordinateList(positions []int, y int) []coordinate {
	coordinates := []coordinate{}
	for _, pos := range positions {
		coordinates = append(coordinates, coordinate{x: pos, y: y})
	}
	return coordinates
}

func concatenateInts(ints []int) int {
	sum := 0
	for _, i := range ints {
		sum = sum*10 + i
	}
	return sum
}

func getValuesAtCoordinates(rows []string, coordinates []coordinate) []int {
	values := make([]int, len(coordinates))
	for i, coord := range coordinates {
		values[i] = int(rows[coord.y][coord.x] - '0')
	}
	return values
}

func sumOfNumberGroupsAdjacentToSymbols(rows []string) int {
	sum := 0

	symbolCoordinates := []coordinate{}
	numberGroups := []numberGroup{}

	for i, row := range rows {
		symbolCoordinates = append(symbolCoordinates, findSymbolCoordinates(row, i)...)
		numberGroupCoordinates := findNumberGroupCoordinates(row, i)
		for _, groupCoordinates := range numberGroupCoordinates {
			numberGroups = append(numberGroups, numberGroup{coordinates: groupCoordinates, value: concatenateInts(getValuesAtCoordinates(rows, groupCoordinates))})
		}
	}

	for _, ng := range numberGroups {
		if ng.isAdjacentToSymbol(symbolCoordinates) {
			sum += ng.value
		}
	}

	return sum
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	start := time.Now()
	defer func() {
		fmt.Printf("Part 1 took: %v\n", time.Since(start))
	}()

	return strconv.Itoa(sumOfNumberGroupsAdjacentToSymbols(input))
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
