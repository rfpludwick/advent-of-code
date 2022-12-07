package main

import (
	"bufio"
	"strconv"
	"strings"
)

func partTwo(scanner *bufio.Scanner) {
	crateStacks := parseInitialCrates(scanner)
	crateStacks = parseCrateMovesTwo(scanner, crateStacks)
	printStackTops(crateStacks)
}

func parseCrateMovesTwo(scanner *bufio.Scanner, crateStacks [][]rune) [][]rune {
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
		miniStack := make([]rune, 0)

		for i := 0; i < numberCrates; i++ {
			miniStack = append(miniStack, crateStacks[stackFrom][stackFromPosition])

			if stackFromPosition == 0 {
				crateStacks[stackFrom] = make([]rune, 0)
			} else {
				crateStacks[stackFrom] = crateStacks[stackFrom][:stackFromPosition]
			}

			stackFromPosition--
		}

		crateStacks[stackTo] = append(crateStacks[stackTo], reverseStack(miniStack)...)
	}

	return crateStacks
}

func reverseStack(crateStack []rune) []rune {
	for i, j := 0, len(crateStack)-1; i < j; i, j = i+1, j-1 {
		crateStack[i], crateStack[j] = crateStack[j], crateStack[i]
	}

	return crateStack
}
