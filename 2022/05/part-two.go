package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	flagShowHelp  bool
	flagTestMode  bool
	flagInputFile string
)

func init() {
	flag.BoolVar(&flagShowHelp, "help", false, "Show this help")
	flag.BoolVar(&flagTestMode, "test", false, "Enable test mode")
	flag.StringVar(&flagInputFile, "input", "./input.txt", "Input file to use")
}

func main() {
	flag.Parse()

	if flagShowHelp {
		flag.Usage()

		os.Exit(0)
	}

	if flagTestMode {
		flagInputFile = "./test-input.txt"
	}

	file, err := os.Open(flagInputFile)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	crateStacks := parseInitialCrates(scanner)
	crateStacks = parseCrateMoves(scanner, crateStacks)
	printStackTops(crateStacks)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parseInitialCrates(scanner *bufio.Scanner) [][]rune {
	numberCrateStacks := 0
	crateStacks := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()

		// Check if we've reached the numerical base
		if line[1] == '1' {
			break
		}

		stackNumber := 0

		for {
			lineLength := len(line)

			// Line done processing?
			if lineLength == 0 {
				break
			}

			progressLine := func() {
				line = line[4:]

				stackNumber++
			}

			// Prep next stack
			if stackNumber >= numberCrateStacks {
				crateStacks = append(crateStacks, make([]rune, 0))

				numberCrateStacks++
			}

			// Skip this stack? Check for end of line as well...
			if lineLength > 3 && line[:4] == "    " {
				progressLine()

				continue
			}

			// Crate in stack
			if line[0] == '[' {
				crateStacks[stackNumber] = append(crateStacks[stackNumber], rune(line[1]))

				if lineLength == 3 { // Indicates the end of the line
					break
				}

				progressLine()

				continue
			}
		}
	}

	// Reverse the stacks so top is at the end of the data structure
	for s, stack := range crateStacks {
		crateStacks[s] = reverseStack(stack)
	}

	// Advance past the empty line
	scanner.Scan()

	return crateStacks
}

func parseCrateMoves(scanner *bufio.Scanner, crateStacks [][]rune) [][]rune {
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

func printStackTops(crateStacks [][]rune) {
	for _, stack := range crateStacks {
		fmt.Printf("%s", string(stack[len(stack)-1]))
	}

	fmt.Println("")
}

func reverseStack(crateStack []rune) []rune {
	for i, j := 0, len(crateStack)-1; i < j; i, j = i+1, j-1 {
		crateStack[i], crateStack[j] = crateStack[j], crateStack[i]
	}

	return crateStack
}
