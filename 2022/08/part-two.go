package main

import (
	"bufio"
	"fmt"
	"strconv"
)

type TreePartTwo struct {
	Height      int
	ScenicScore int64
}

func partTwo(scanner *bufio.Scanner) {
	// By convention when using single-char loop vars, i for height, j for width
	// This is to ease confusion when looking at different loops

	var (
		currentGridHeightIndex int
		gridWidth              int
		maxGridWidthIndex      int
		maxGridHeightIndex     int
		maxScenicScore         int64
		grid                   [][]TreePartTwo
	)

	currentGridHeightIndex = -1

	// Build the grid and calculate what we can along the way
	// We can figure out top row, left column, right column, left-of-tree comparisons, and top-of-tree comparisons
	// We have to figure out bottom row, right-of-tree comparisons, and bottom-of-tree comparisons later
	for scanner.Scan() {
		line := scanner.Text()

		currentGridHeightIndex++

		if gridWidth == 0 {
			gridWidth = len(line)
			maxGridWidthIndex = (gridWidth - 1)
		}

		grid = append(grid, make([]TreePartTwo, gridWidth))

		for j, char := range line {
			height, _ := strconv.Atoi(string(char))
			scenicScore := int64(0)

			if j > 0 && currentGridHeightIndex > 0 {
				// Check the previous columns' heights in this row to start
				for k := (j - 1); k >= 0; k-- {
					scenicScore++

					if height <= grid[currentGridHeightIndex][k].Height {
						break
					}
				}

				nextScenicScore := int64(0)

				// Check the previous rows' heights in this column to continue
				for k := (currentGridHeightIndex - 1); k >= 0; k-- {
					nextScenicScore++

					if height <= grid[k][j].Height {
						break
					}
				}

				scenicScore *= nextScenicScore
			}

			grid[currentGridHeightIndex][j] = TreePartTwo{
				Height:      height,
				ScenicScore: scenicScore,
			}
		}
	}

	maxGridHeightIndex = (len(grid) - 1)

	// Final nested loop to check remaining directions
	// Loops run from the bottom-up, right-to-left
	for i := maxGridHeightIndex; i >= 0; i-- {
		for j := maxGridWidthIndex; j > 0; j-- {
			// Handle the bottom row
			if i == maxGridHeightIndex {
				grid[i][j].ScenicScore = 0

				continue
			}

			// Check the previous columns' heights in this row to start
			nextScenicScore := int64(0)

			for k := (j + 1); k <= maxGridWidthIndex; k++ {
				nextScenicScore++

				if grid[i][j].Height <= grid[i][k].Height {
					break
				}
			}

			grid[i][j].ScenicScore *= nextScenicScore

			nextScenicScore = 0

			// Check the previous rows' heights in this column to continue
			for k := (i + 1); k <= maxGridHeightIndex; k++ {
				nextScenicScore++

				if grid[i][j].Height <= grid[k][j].Height {
					break
				}
			}

			grid[i][j].ScenicScore *= nextScenicScore

			if grid[i][j].ScenicScore > maxScenicScore {
				maxScenicScore = grid[i][j].ScenicScore
			}
		}
	}

	fmt.Printf("The maximum scenic score is %d\n", maxScenicScore)
}
