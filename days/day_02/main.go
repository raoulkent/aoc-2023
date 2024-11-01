package day_02

import (
	"fmt"
	"regexp"
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

type bag struct {
	red   int
	green int
	blue  int
}

func (b *bag) isValidDraw(second bag) bool {
	return b.red-second.red >= 0 && b.blue-second.blue >= 0 && b.green-second.green >= 0
}

func defaultBag() bag {
	return bag{red: 12, green: 13, blue: 14}
}

const (
	gameIDPattern     = `Game\s+(\d+)`
	colorCountPattern = `(\d+)\s+(blue|red|green)`
)

var (
	gameIDRegex     = regexp.MustCompile(gameIDPattern)
	colorCountRegex = regexp.MustCompile(colorCountPattern)
)

func extractGameID(line string) (string, error) {
	match := gameIDRegex.FindStringSubmatch(line)
	if len(match) < 2 {
		return "", fmt.Errorf("game ID not found")
	}
	return match[1], nil
}

func extractColorCounts(section string) ([]bag, error) {
	var bags []bag
	matches := colorCountRegex.FindAllStringSubmatch(section, -1)
	for _, match := range matches {
		count, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, err
		}
		color := match[2]
		b := bag{}
		switch color {
		case "red":
			b.red = count
		case "green":
			b.green = count
		case "blue":
			b.blue = count
		}
		bags = append(bags, b)
	}
	return bags, nil
}

func parseLine(line string) (string, []bag) {
	gameID, err := extractGameID(line)
	if err != nil {
		return "", nil
	}

	parts := strings.SplitN(line, ":", 2)
	if len(parts) < 2 {
		return "", nil
	}
	sections := strings.Split(parts[1], ";")

	var allBags []bag
	for _, section := range sections {
		bags, err := extractColorCounts(section)
		if err != nil {
			return "", nil
		}
		allBags = append(allBags, bags...)
	}

	return gameID, allBags
}

func isValidGame(sectionBags []bag) bool {
	bag := defaultBag()
	for _, sectionBag := range sectionBags {
		if !bag.isValidDraw(sectionBag) {
			return false
		}
	}
	return true
}

// Part1 solves the b part of the exercise
func Part1(input []string) string {
	start := time.Now()
	defer func() {
		fmt.Printf("Part 1 took: %v\n", time.Since(start))
	}()

	validGames := 0
	for _, line := range input {
		gameID, sectionBags := parseLine(line)
		if isValidGame(sectionBags) {
			gameValue, err := strconv.Atoi(gameID)
			if err != nil {
				fmt.Printf("Error converting game ID to integer: %v\n", err)
				continue
			}
			validGames += gameValue
		}
	}

	return strconv.Itoa(validGames)
}

func maxColorCounts(bags []bag) bag {
	maxBag := bag{red: 0, green: 0, blue: 0}
	for _, b := range bags {
		if b.red > maxBag.red {
			maxBag.red = b.red
		}
		if b.green > maxBag.green {
			maxBag.green = b.green
		}
		if b.blue > maxBag.blue {
			maxBag.blue = b.blue
		}
	}

	return maxBag
}

func bagPower(b bag) int {
	return b.red * b.green * b.blue
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	start := time.Now()
	defer func() {
		fmt.Printf("Part 1 took: %v\n", time.Since(start))
	}()

	powerSum := 0
	for _, line := range input {
		_, sectionBags := parseLine(line)
		maxBag := maxColorCounts(sectionBags)
		powerSum += bagPower(maxBag)
	}

	return strconv.Itoa(powerSum)
}
