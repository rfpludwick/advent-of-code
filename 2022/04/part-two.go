package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

func partTwo(scanner *bufio.Scanner) {
	numberAnyOverlaps := 0

	for scanner.Scan() {
		line := scanner.Text()

		sections := strings.Split(line, ",")
		leftBoundaries := strings.Split(sections[0], "-")
		rightBoundaries := strings.Split(sections[1], "-")

		leftMin, _ := strconv.Atoi(leftBoundaries[0])
		leftMax, _ := strconv.Atoi(leftBoundaries[1])
		rightMin, _ := strconv.Atoi(rightBoundaries[0])
		rightMax, _ := strconv.Atoi(rightBoundaries[1])

		if leftMax < rightMin { // Left lower than right
			continue
		}

		if rightMax < leftMin { // Right lower than left
			continue
		}

		numberAnyOverlaps++
	}

	log.Printf("Number of any overlaps is %d", numberAnyOverlaps)
}
