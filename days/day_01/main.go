package day_01

import (
	"fmt"
	"strconv"
	"sync"
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

	// Channel to send the input to
	inputChan := make(chan string, len(input))

	// Channel to send the result to
	resultChan := make(chan int)

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Start the goroutine to sum the integers
	wg.Add(1)
	go sumFirstAndLastNumber(inputChan, &wg, resultChan)

	// Send ints to the channel
	go func() {
		for _, line := range input {
			inputChan <- line
		}
		close(inputChan)
	}()

	// Wait for the goroutines to finish
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Receive the result from the result channel
	result := <-resultChan

	return strconv.Itoa(result)
}

// Function that takes a string, and returns the first number from the left
// and the first character that is a number from the right as a number.
// Example: "abc123" -> "13"
func sumFirstAndLastNumber(channel <-chan string, wg *sync.WaitGroup, resultChan chan<- int) {
	defer wg.Done()
	sum := 0

	for line := range channel {
		numbers := getFirstNumber(line) + getLastNumber(line)

		value, _ := strconv.Atoi(numbers)
		sum += value
	}

	resultChan <- sum
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

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
