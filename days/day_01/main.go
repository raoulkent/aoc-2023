package day_01

import (
	"fmt"
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

// Part1 solves the first part of the exercise
func Part1(input []string) int {
	start := time.Now()
	// Defer time tracking:
	defer func() {
		fmt.Printf("Part 1 took: %v\n", time.Since(start))
	}()

	sum := 0
	for _, line := range input {
		sum += sumFirstAndLastNumber(line)
	}
	return sum
}

// Function that reads a string, and returns the first character that is a number from the left
func getFirstNumber(line string) string {
	for _, char := range line {
		if char >= '0' && char <= '9' {
			return string(char)
		}
	}
	return ""
}

// Function that reads a string, and returns the first character that is a number from the right
func getLastNumber(line string) string {
	for i := len(line) - 1; i >= 0; i-- {
		char := line[i]
		if char >= '0' && char <= '9' {
			return string(char)
		}
	}
	return ""
}

// Function that sums the first and last number of a string
func sumFirstAndLastNumber(line string) int {
	numbers := getFirstNumber(line) + getLastNumber(line)
	value, _ := strconv.Atoi(numbers)
	return value
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
