package main

import (
	"bufio"
	"fmt"
	"strconv"
)

type TreePartOne struct {
	Height  int
	Visible bool
}

func partOne(scanner *bufio.Scanner) {
	// By convention when using single-char loop vars, i for height, j for width
	// This is to ease confusion when looking at different loops

	var (
		currentGridHeightIndex int
		gridWidth              int
		maxGridWidthIndex      int
		numberVisible          int
		currentRowMaxHeight    int
		columnMaxHeights       []int
		grid                   [][]TreePartOne
	)

	currentGridHeightIndex = -1

	// Build the grid and calculate what we can along the way
	// We can figure out top row, left column, top-down comparisons, and left-to-right comparisons
	// We have to figure out bottom row, right column, right-to-left comparisons, and bottom-up comparisons later
	for scanner.Scan() {
		line := scanner.Text()

		currentGridHeightIndex++
		currentRowMaxHeight = -1

		if gridWidth == 0 {
			gridWidth = len(line)
			maxGridWidthIndex = (gridWidth - 1)

			columnMaxHeights = initColumnMaxHeights(gridWidth)
		}

		grid = append(grid, make([]TreePartOne, gridWidth))

		for j, char := range line {
			height, _ := strconv.Atoi(string(char))
			visible := false

			// Check if this tree is visible from the left (includes full left column)
			if height > currentRowMaxHeight {
				currentRowMaxHeight = height

				visible = true
			}

			// Check if this tree is visible from the top (includes full top row)
			if height > columnMaxHeights[j] {
				columnMaxHeights[j] = height

				visible = true
			}

			if visible {
				numberVisible++
			}

			grid[currentGridHeightIndex][j] = TreePartOne{
				Height:  height,
				Visible: visible,
			}
		}
	}

	// Final nested loop to check remaining directions
	// Loops run from the bottom-up, right-to-left
	columnMaxHeights = initColumnMaxHeights(gridWidth)

	for i := (len(grid) - 1); i > 0; i-- {
		rowMaxHeight := -1

		for j := maxGridWidthIndex; j > 0; j-- {
			// Check if this tree is visible from the right (includes full right column)
			if grid[i][j].Height > columnMaxHeights[j] {
				columnMaxHeights[j] = grid[i][j].Height

				if !grid[i][j].Visible {
					grid[i][j].Visible = true

					numberVisible++
				}
			}

			// Check if this tree is visible from the bottom (includes full bottom row)
			if grid[i][j].Height > rowMaxHeight {
				rowMaxHeight = grid[i][j].Height

				if !grid[i][j].Visible {
					grid[i][j].Visible = true

					numberVisible++
				}
			}
		}
	}

	fmt.Printf("The number of visible trees is %d\n", numberVisible)
}

func initColumnMaxHeights(width int) []int {
	columnMaxHeights := make([]int, width)

	for i := 0; i < width; i++ {
		columnMaxHeights[i] = -1
	}

	return columnMaxHeights
}
