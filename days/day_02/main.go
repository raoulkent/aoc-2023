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

func parseLine(line string) (string, []bag) {
	// Regex to capture the game number only
	gameIDRegex := regexp.MustCompile(`Game\s+(\d+)`)
	// Regex to match color counts
	colorCountRegex := regexp.MustCompile(`(\d+)\s+(blue|red|green)`)

	// Extract game ID (number only)
	gameIDMatch := gameIDRegex.FindStringSubmatch(line)
	if len(gameIDMatch) < 2 {
		return "", nil // Return empty if game ID is not found
	}
	gameID := gameIDMatch[1] // The number part only

	// Split into sections by ";"
	parts := strings.SplitN(line, ":", 2)
	if len(parts) < 2 {
		return "", nil // Return empty if line is improperly formatted
	}
	sections := strings.Split(parts[1], ";")

	// Process each section individually
	var sectionBags []bag
	for _, section := range sections {
		sectionBag := bag{}

		// Find all matches for color counts
		matches := colorCountRegex.FindAllStringSubmatch(section, -1)
		for _, match := range matches {
			count, _ := strconv.Atoi(match[1])
			color := match[2]

			switch color {
			case "red":
				sectionBag.red += count
			case "blue":
				sectionBag.blue += count
			case "green":
				sectionBag.green += count
			}
		}

		// Append color counts for this section
		sectionBags = append(sectionBags, sectionBag)
	}

	return gameID, sectionBags
}

// Part1 solves the b part of the exercise
func Part1(input []string) string {
	start := time.Now()
	defer func() {
		fmt.Printf("Part 1 took: %v\n", time.Since(start))
	}()

	validGames := 0
	for _, line := range input {
		bag := defaultBag()
		gameID, sectionBags := parseLine(line)
		OK := true
		for _, sectionBag := range sectionBags {
			if !bag.isValidDraw(sectionBag) {
				OK = false
				break
			}
		}

		if !OK {
			continue
		}
		gamevalue, _ := strconv.Atoi(gameID)

		validGames += gamevalue
	}

	sum := strconv.Itoa(validGames)

	return sum
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
