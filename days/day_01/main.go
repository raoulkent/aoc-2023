package day_01

import (
	"fmt"
	"strconv"
	"strings"
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
func Part1(input []string) string {
	start := time.Now()
	// Defer time tracking:
	defer func() {
		fmt.Printf("Part 1 took: %v\n", time.Since(start))
	}()

	sum := 0
	for _, line := range input {
		sum += sumFirstAndLastNumber(line)
	}
	return strconv.Itoa(sum)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	start := time.Now()
	// Defer time tracking:
	defer func() {
		fmt.Printf("Part 2 took: %v\n", time.Since(start))
	}()

	sum := 0
	for _, line := range input {
		sum += combineFirstAndLastDigitOrWord(line)
	}
	return strconv.Itoa(sum)
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

var wordToNumber = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// Function that reads a string, and returns the first character or word that is a number from the left
func getFirstNumberIncludingWords(line string) string {
	for i := 0; i < len(line); i++ {
		for word, number := range wordToNumber {
			if strings.HasPrefix(line[i:], word) {
				return strconv.Itoa(number)
			}
		}
		if line[i] >= '0' && line[i] <= '9' {
			return string(line[i])
		}
	}
	return ""
}

// Function that reads a string, and returns the first character or word that is a number from the right
func getLastNumberIncludingWords(line string) string {
	for i := len(line) - 1; i >= 0; i-- {
		for word, number := range wordToNumber {
			if strings.HasSuffix(line[:i+1], word) {
				return strconv.Itoa(number)
			}
		}
		if line[i] >= '0' && line[i] <= '9' {
			return string(line[i])
		}
	}
	return ""
}

// Function that combines the first and last number or word of a string
func combineFirstAndLastDigitOrWord(line string) int {
	numbers := getFirstNumberIncludingWords(line) + getLastNumberIncludingWords(line)
	value, _ := strconv.Atoi(numbers)
	return value
}
