package day_04

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"time"

	. "github.com/samber/lo"
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

func captureDigitGroups(line string) (left string, right string) {
	pattern := `(\d+(?:\s+\d+)*)\s*\|\s*(\d+(?:\s+\d+)*)`

	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(line)

	left = matches[1]
	right = matches[2]

	return left, right
}

func captureDigits(s string) []int {
	pattern := `\d+`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(s, -1)

	digits := make([]int, 0)
	for _, match := range matches {
		d, _ := strconv.Atoi(match)
		digits = append(digits, d)
	}

	return digits
}

func power2MatchedLine(line string) int {
	matches := 0

	leftGroup, rightGroup := captureDigitGroups(line)
	leftDigits := captureDigits(leftGroup)
	rightDigits := captureDigits(rightGroup)

	for _, leftDigit := range leftDigits {
		for _, rightDigit := range rightDigits {
			if leftDigit == rightDigit {
				matches++
			}
		}
	}

	if matches > 0 {
		pow := float64(matches)
		score := int(math.Pow(2, pow-1))

		return score
	}

	return 0
}

func sumOfPower2MatchedLines(input []string) int {
	sum := 0
	for _, line := range input {
		sum += power2MatchedLine(line)
	}

	return sum
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	start := time.Now()
	defer func() {
		fmt.Printf("Part 1 took: %v\n", time.Since(start))
	}()

	return strconv.Itoa(sumOfPower2MatchedLines(input))
}

func getMatchingNumbers(s string) int {
	leftGroup, rightGroup := captureDigitGroups(s)
	leftDigits := captureDigits(leftGroup)
	rightDigits := captureDigits(rightGroup)

	matches := len(Intersect(leftDigits, rightDigits))

	return matches
}

func totalScratchCards(input []string) int {
	total := 0
	additionalCards := make(map[int]int)

	for i, line := range input {
		// Find the number of additional cards for this line, and add the base card (1)
		cardsOfThisLine := additionalCards[i]

		// Find the number of matching numbers
		matches := getMatchingNumbers(line)

		// For each of the cards of this line, add an additional card for the subsequent [matches] lines
		for j := 1; j <= matches; j++ {
			additionalCards[i+j] += cardsOfThisLine + 1
		}

		// Add the number of cards of this line to the total
		total += cardsOfThisLine + 1
	}

	return total
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	start := time.Now()
	defer func() {
		fmt.Printf("Part 2 took: %v\n", time.Since(start))
	}()

	return strconv.Itoa(totalScratchCards(input))
}
