package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func partOne(scanner *bufio.Scanner) {
	// Build the grid
	grid := newGrid()

	for scanner.Scan() {
		line := scanner.Text()
		sensorBeaconParts := strings.Split(line[12:], ": closest beacon is at x=")
		sensorParts := strings.Split(sensorBeaconParts[0], ", y=")
		beaconParts := strings.Split(sensorBeaconParts[1], ", y=")

		sensorX, _ := strconv.Atoi(sensorParts[0])
		sensorY, _ := strconv.Atoi(sensorParts[1])
		beaconX, _ := strconv.Atoi(beaconParts[0])
		beaconY, _ := strconv.Atoi(beaconParts[1])

		grid.AddPair(sensorX, sensorY, beaconX, beaconY)
	}

	// Check the pertinent row
	var rowToCheck int

	if flagTestMode {
		rowToCheck = 10
	} else {
		rowToCheck = 2000000
	}

	numberPositions := 0

	for i := grid.MinX; i <= grid.MaxX; i++ {
		for _, gridPair := range grid.Pairs {
			if i == gridPair.BeaconX && rowToCheck == gridPair.BeaconY { // Don't count actual beacons
				break
			}

			if calculateGridDistance(i, rowToCheck, gridPair.SensorX, gridPair.SensorY) <= gridPair.Distance {
				numberPositions++

				break
			}
		}
	}

	fmt.Printf("The number of invalid positions is %d\n", numberPositions)
}

func calculateGridDistance(x1 int, y1 int, x2 int, y2 int) int {
	return int(math.Abs(float64(x1-x2))) + int(math.Abs(float64(y1-y2)))
}
