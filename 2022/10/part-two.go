package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func partTwo(scanner *bufio.Scanner) {
	var (
		rowCycleNumber       int
		spriteCenterPosition int
	)

	rowCycleNumber = -1
	spriteCenterPosition = 1

	for scanner.Scan() {
		line := scanner.Text()

		advanceCycle := func() {
			rowCycleNumber++

			if rowCycleNumber == 40 {
				fmt.Printf("\n")

				rowCycleNumber = 0
			}

			if spriteCenterPosition-1 <= rowCycleNumber && rowCycleNumber <= spriteCenterPosition+1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}

		advanceCycle()

		if line != "noop" {
			advanceCycle()

			lineParts := strings.Split(line, " ")
			modifier, _ := strconv.Atoi(lineParts[1])

			spriteCenterPosition += modifier
		}
	}
}
