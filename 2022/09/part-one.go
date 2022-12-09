package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func partOne(scanner *bufio.Scanner) {
	var (
		visits       map[string]bool
		headPosition Position
		tailPosition Position
	)

	visits = make(map[string]bool, 0)

	visits["0,0"] = true

	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, " ")

		direction := lineParts[0]
		distance, _ := strconv.Atoi(lineParts[1])

		for i := 0; i < distance; i++ {
			switch direction {
			case "U":
				headPosition.Y++
			case "D":
				headPosition.Y--
			case "L":
				headPosition.X--
			case "R":
				headPosition.X++
			}

			xDifference := headPosition.X - tailPosition.X
			yDifference := headPosition.Y - tailPosition.Y
			xDifferenceAbsolute := int64(math.Abs(float64(xDifference)))
			yDifferenceAbsolute := int64(math.Abs(float64(yDifference)))

			moveX := func() {
				if xDifference < 0 {
					tailPosition.X--
				} else if xDifference > 0 {
					tailPosition.X++
				}
			}

			moveY := func() {
				if yDifference < 0 {
					tailPosition.Y--
				} else if yDifference > 0 {
					tailPosition.Y++
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

			visits[fmt.Sprintf("%d,%d", tailPosition.X, tailPosition.Y)] = true
		}
	}

	fmt.Printf("Total tail visits is %d\n", len(visits))
}
