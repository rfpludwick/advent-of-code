package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
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

	highElf := 0
	highElfCalories := 0
	currentElf := 1
	currentElfCalories := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			if currentElfCalories > highElfCalories {
				highElf = currentElf
				highElfCalories = currentElfCalories
			}

			currentElf++
			currentElfCalories = 0

			continue
		}

		numeral, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		currentElfCalories += numeral
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("High elf is #%d with calories: %d", (highElf + 1), highElfCalories)
}
