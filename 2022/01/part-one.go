package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("High elf is %d with calories: %d\n", highElf, highElfCalories)
}
