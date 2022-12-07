package main

import (
	"bufio"
	"log"
	"strconv"
)

func partOne(scanner *bufio.Scanner) {
	highElf := 0
	highElfCalories := 0
	currentElf := 1
	currentElfCalories := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			if currentElfCalories > highElfCalories {
				highElf = currentElf
				highElfCalories = currentElfCalories
			}

			currentElf++
			currentElfCalories = 0

			continue
		}

		numeral, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		currentElfCalories += numeral
	}

	log.Printf("High elf is #%d with calories: %d", (highElf + 1), highElfCalories)
}
