package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func (c *Cave) AddFloor() {
	c.MaxY += 2

	for x := c.MinX; x <= c.MaxX; x++ {
		c.SetParticle(x, c.MaxY, PARTICLE_ROCK)
	}
}

func partTwo(scanner *bufio.Scanner) {
	cave := newCave()

	// Build the cave
	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, " -> ")

		previousX, previousY := 0, 0

		for i, linePart := range lineParts {
			xy := strings.Split(linePart, ",")

			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])

			if i == 0 { // Are we starting from the rock's starting point?
				cave.SetParticle(x, y, PARTICLE_ROCK)

				previousX, previousY = x, y
			} else {
				if previousX == x { // In this case, we're moving along y
					difference := y - previousY
					differenceAbsolute := int(math.Abs(float64(difference)))

					for differenceAbsolute != 0 {
						if difference < 1 {
							previousY--
						} else {
							previousY++
						}

						cave.SetParticle(previousX, previousY, PARTICLE_ROCK)

						differenceAbsolute--
					}
				} else { // Move along x
					difference := x - previousX
					differenceAbsolute := int(math.Abs(float64(difference)))

					for differenceAbsolute != 0 {
						if difference < 1 {
							previousX--
						} else {
							previousX++
						}

						cave.SetParticle(previousX, previousY, PARTICLE_ROCK)

						differenceAbsolute--
					}
				}
			}
		}
	}

	cave.AddFloor()

	for {
		if cave.Tick(false) {
			break
		}
	}

	cave.Render()

	fmt.Printf("There are %d units of sand in the cave\n", cave.GetSandUnitsCount())
}
