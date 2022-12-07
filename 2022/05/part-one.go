package main

import (
	"bufio"
	"strconv"
	"strings"
)

func partOne(scanner *bufio.Scanner) {
	crateStacks := parseInitialCrates(scanner)
	crateStacks = parseCrateMovesOne(scanner, crateStacks)
	printStackTops(crateStacks)
}

func parseCrateMovesOne(scanner *bufio.Scanner, crateStacks [][]rune) [][]rune {
	for scanner.Scan() {
		line := scanner.Text()

		// Split the move up into required integers
		line = line[5:]
		parts1 := strings.Split(line, " from ")
		parts2 := strings.Split(parts1[1], " to ")

		numberCrates, _ := strconv.Atoi(parts1[0])
		stackFrom, _ := strconv.Atoi(parts2[0])
		stackTo, _ := strconv.Atoi(parts2[1])

		stackFrom--
		stackTo--

		// Perform the move
		stackFromPosition := (len(crateStacks[stackFrom]) - 1)

		for i := 0; i < numberCrates; i++ {
			crateStacks[stackTo] = append(crateStacks[stackTo], crateStacks[stackFrom][stackFromPosition])

			if stackFromPosition == 0 {
				crateStacks[stackFrom] = make([]rune, 0)
			} else {
				crateStacks[stackFrom] = crateStacks[stackFrom][:stackFromPosition]
			}

			stackFromPosition--
		}
	}

	return crateStacks
}
