package main

import (
	"bufio"
	"log"
	"strings"
)

func partOne(scanner *bufio.Scanner) {
	totalPriority := 0

	for scanner.Scan() {
		line := scanner.Text()
		midpoint := (len(line) / 2)

		lineLeft := make(map[rune]int)

		// Build the left side
		for _, char := range strings.Split(line[midpoint:], "") {
			runeChar := []rune(char)[0]

			_, ok := lineLeft[runeChar]

			if !ok {
				lineLeft[runeChar] = 0
			}

			lineLeft[runeChar]++
		}

		// Check the right side
		for _, char := range strings.Split(line[:midpoint], "") {
			runeChar := []rune(char)[0]

			_, ok := lineLeft[runeChar]

			if ok {
				runeValue := int(runeChar)

				if runeValue < 97 {
					runeValue -= 38
				} else {
					runeValue -= 96
				}

				totalPriority += runeValue

				break
			}
		}
	}

	log.Printf("Total priority is %d", totalPriority)
}
