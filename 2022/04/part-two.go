package main

import (
	"bufio"
	"flag"
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

	numberAnyOverlaps := 0

	for scanner.Scan() {
		line := scanner.Text()

		sections := strings.Split(line, ",")
		leftBoundaries := strings.Split(sections[0], "-")
		rightBoundaries := strings.Split(sections[1], "-")

		leftMin, _ := strconv.Atoi(leftBoundaries[0])
		leftMax, _ := strconv.Atoi(leftBoundaries[1])
		rightMin, _ := strconv.Atoi(rightBoundaries[0])
		rightMax, _ := strconv.Atoi(rightBoundaries[1])

		if leftMax < rightMin { // Left lower than right
			continue
		}

		if rightMax < leftMin { // Right lower than left
			continue
		}

		numberAnyOverlaps++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Number of any overlaps is %d", numberAnyOverlaps)
}
