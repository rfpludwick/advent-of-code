package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	flagShowHelp  bool
	flagTestMode  bool
	flagInputFile string
	flagRunPart   int
)

func init() {
	flag.BoolVar(&flagShowHelp, "help", false, "Show this help")
	flag.BoolVar(&flagTestMode, "test", false, "Enable test mode")
	flag.StringVar(&flagInputFile, "input", "./input.txt", "Input file to use")
	flag.IntVar(&flagRunPart, "run-part", 1, "The part to run")

	flag.Parse()

	if flagShowHelp {
		flag.Usage()

		os.Exit(0)
	}

	if flagTestMode {
		flagInputFile = "./test-input.txt"
	}
}

func main() {
	file, err := os.Open(flagInputFile)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	switch flagRunPart {
	case 1:
		partOne(scanner)
	case 2:
		partTwo(scanner)
	default:
		log.Fatalf("Error: Part number %d is not supported", flagRunPart)
	}

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

	// Invert the stacks so top is at the end of the data structure
	for s, stack := range crateStacks {
		for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
			crateStacks[s][i], crateStacks[s][j] = crateStacks[s][j], crateStacks[s][i]
		}
	}

	// Advance past the empty line
	scanner.Scan()

	return crateStacks
}

func printStackTops(crateStacks [][]rune) {
	for _, stack := range crateStacks {
		fmt.Printf("%s", string(stack[len(stack)-1]))
	}

	fmt.Println("")
}
