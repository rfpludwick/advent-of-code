package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func partTwo(scanner *bufio.Scanner) {
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

	fmt.Println("Grid built")

	// Check
	var checkRangeMaxX, checkRangeMaxY int

	if flagTestMode {
		checkRangeMaxX = 20
		checkRangeMaxY = 20
	} else {
		checkRangeMaxX = 4000000
		checkRangeMaxY = 4000000
	}

	for y := 0; y <= checkRangeMaxY; y++ {
		mergedBoundaries := mergeBoundaries(y, grid.Pairs)

		sort.Slice(mergedBoundaries, func(i int, j int) bool {
			return mergedBoundaries[i][0] < mergedBoundaries[j][0]
		})

		// Find the gap
		x := 0

		for _, bounds := range mergedBoundaries {
			if bounds[0] < 0 {
				bounds[0] = 0
			}

			if bounds[0] > x {
				fmt.Printf("Open position found at %d, %d\n", y, x)
				fmt.Printf("Tuning frequency is %d\n", (4000000*x)+y)

				os.Exit(0)
			}

			if bounds[1] >= checkRangeMaxX {
				break
			}

			x = bounds[1] + 1
		}
	}
}

func mergeBoundaries(y int, gridPairs []GridPair) [][]int {
	merged1 := make([][]int, 0)

	for _, gridPair := range gridPairs {
		if _, ok := gridPair.Bounds[y]; ok {
			merged1 = append(merged1, gridPair.Bounds[y])
		}
	}

	startLength := len(merged1)
	var endLength int

	for {
		merged2 := make([][]int, 0)
		endLength = 0

		for _, bounds := range merged1 {
			if endLength == 0 {
				merged2 = append(merged2, bounds)
				endLength++

				continue
			}

			hasMerged := false

			for i, merge := range merged2 {
				if bounds[0] <= merge[0] && (bounds[1]+1) >= merge[0] { // Bounds overlap/touch on left
					if bounds[0] < merge[0] {
						merged2[i][0] = bounds[0]
					}

					if bounds[1] > merge[1] {
						merged2[i][1] = bounds[01]
					}

					hasMerged = true

					break
				}

				if (bounds[0]-1) <= merge[1] && bounds[1] >= merge[1] { // Bounds overlap/touch on right
					if bounds[0] < merge[0] {
						merged2[i][0] = bounds[0]
					}

					if bounds[1] > merge[1] {
						merged2[i][1] = bounds[01]
					}

					hasMerged = true

					break
				}

				if bounds[0] >= merge[0] && bounds[1] <= merge[1] { // Bounds contained entirely within
					hasMerged = true
				}
			}

			if !hasMerged {
				merged2 = append(merged2, bounds)
				endLength++
			}
		}

		if endLength == startLength {
			break
		}

		merged1 = merged2
		startLength = endLength
	}

	return merged1
}
