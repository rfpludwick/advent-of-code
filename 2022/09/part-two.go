package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func partTwo(scanner *bufio.Scanner) {
	var (
		visits      map[string]bool
		knots       []Position
		leadingKnot int
	)

	visits = make(map[string]bool, 0)
	visits["0,0"] = true

	knots = make([]Position, 10)

	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, " ")

		direction := lineParts[0]
		distance, _ := strconv.Atoi(lineParts[1])

		for i := 0; i < distance; i++ {
			// Move the head first
			switch direction {
			case "U":
				knots[0].Y++
			case "D":
				knots[0].Y--
			case "L":
				knots[0].X--
			case "R":
				knots[0].X++
			}

			// Now move the remaining knots
			leadingKnot = 0

			for i := range knots {
				if i == 0 {
					continue
				}

				xDifference := knots[leadingKnot].X - knots[i].X
				yDifference := knots[leadingKnot].Y - knots[i].Y
				xDifferenceAbsolute := int64(math.Abs(float64(xDifference)))
				yDifferenceAbsolute := int64(math.Abs(float64(yDifference)))

				moveX := func() {
					if xDifference < 0 {
						knots[i].X--
					} else if xDifference > 0 {
						knots[i].X++
					}
				}

				moveY := func() {
					if yDifference < 0 {
						knots[i].Y--
					} else if yDifference > 0 {
						knots[i].Y++
					}
				}

				if xDifferenceAbsolute == 2 {
					moveX()

					if yDifferenceAbsolute == 1 { // Handle the diagonal move since total distance is 3
						moveY()
					}
				}

				if yDifferenceAbsolute == 2 {
					moveY()

					if xDifferenceAbsolute == 1 { // Handle the diagonal move since total distance is 3
						moveX()
					}
				}

				leadingKnot++
			}

			visits[fmt.Sprintf("%d,%d", knots[9].X, knots[9].Y)] = true
		}
	}

	fmt.Printf("Total tail visits is %d\n", len(visits))
}
