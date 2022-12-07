package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

func partOne(scanner *bufio.Scanner) {
	numberFullOverlaps := 0

	for scanner.Scan() {
		line := scanner.Text()

		sections := strings.Split(line, ",")
		leftBoundaries := strings.Split(sections[0], "-")
		rightBoundaries := strings.Split(sections[1], "-")

		leftMin, _ := strconv.Atoi(leftBoundaries[0])
		leftMax, _ := strconv.Atoi(leftBoundaries[1])
		rightMin, _ := strconv.Atoi(rightBoundaries[0])
		rightMax, _ := strconv.Atoi(rightBoundaries[1])

		if leftMin >= rightMin && leftMax <= rightMax { // Check if left contained within right
			numberFullOverlaps++
		} else if rightMin >= leftMin && rightMax <= leftMax { // Check if right contained within left
			numberFullOverlaps++
		}
	}

	log.Printf("Number of full overlaps is %d", numberFullOverlaps)
}
