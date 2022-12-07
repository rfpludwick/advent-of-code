package main

import (
	"bufio"
	"log"
	"sort"
	"strconv"
)

func partTwo(scanner *bufio.Scanner) {
	elvesCalories := []int{0}
	currentElf := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			currentElf++

			elvesCalories = append(elvesCalories, 0)

			continue
		}

		numeral, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		elvesCalories[currentElf] += numeral
	}

	sort.Ints(elvesCalories)

	total := 0

	for i := 0; i < 3; i++ {
		total += elvesCalories[len(elvesCalories)-1]

		elvesCalories = elvesCalories[:len(elvesCalories)-1]
	}

	log.Printf("High three elves' total calories is %d", total)
}
