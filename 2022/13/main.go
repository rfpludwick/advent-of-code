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

const (
	COMPARE_PART_ONE_SUCCESS  = 1
	COMPARE_PART_ONE_FAIL     = 2
	COMPARE_PART_ONE_CONTINUE = 3
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

func compareSides(left []interface{}, right []interface{}) int {
	leftLength, rightLength := len(left), len(right)
	var maxLength int

	if leftLength > rightLength {
		maxLength = leftLength
	} else {
		maxLength = rightLength
	}

	for i := 0; i < maxLength; i++ {
		if i >= rightLength {
			return COMPARE_PART_ONE_FAIL
		}

		if i >= leftLength {
			return COMPARE_PART_ONE_SUCCESS
		}

		var leftIsFloat, rightIsFloat bool
		var leftFloat, rightFloat float64
		var leftSlice, rightSlice []interface{}

		switch value := left[i].(type) {
		case float64:
			leftIsFloat = true
			leftFloat = value
		case []interface{}:
			leftIsFloat = false
			leftSlice = value
		default:
			fmt.Println("Unable to handle left side type")

			os.Exit(1)
		}

		switch value := right[i].(type) {
		case float64:
			rightIsFloat = true
			rightFloat = value
		case []interface{}:
			rightIsFloat = false
			rightSlice = value
		default:
			fmt.Println("Unable to handle right side type")

			os.Exit(1)
		}

		if leftIsFloat && rightIsFloat {
			if leftFloat < rightFloat {
				return COMPARE_PART_ONE_SUCCESS
			}

			if leftFloat > rightFloat {
				return COMPARE_PART_ONE_FAIL
			}

			continue
		}

		if leftIsFloat && !rightIsFloat { // Convert left to a slice
			leftSlice = append(leftSlice, left[i])
		} else if !leftIsFloat && rightIsFloat { // Convert right to a slice
			rightSlice = append(rightSlice, right[i])
		}

		nestedComparison := compareSides(leftSlice, rightSlice)

		if nestedComparison != COMPARE_PART_ONE_CONTINUE {
			return nestedComparison
		}
	}

	return COMPARE_PART_ONE_CONTINUE
}
