package main

import (
	"bufio"
	"flag"
	"log"
	"os"
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

	totalPriority := 0

	for scanner.Scan() {
		line := scanner.Text()
		midpoint := (len(line) / 2)

		lineLeft := make(map[rune]int)

		// Build the left side
		for _, char := range strings.Split(line[midpoint:], "") {
			runeChar := []rune(char)[0]

			_, ok := lineLeft[runeChar]

			if !ok {
				lineLeft[runeChar] = 0
			}

			lineLeft[runeChar]++
		}

		// Check the right side
		for _, char := range strings.Split(line[:midpoint], "") {
			runeChar := []rune(char)[0]

			_, ok := lineLeft[runeChar]

			if ok {
				runeValue := int(runeChar)

				if runeValue < 97 {
					runeValue -= 38
				} else {
					runeValue -= 96
				}

				totalPriority += runeValue

				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Total priority is %d", totalPriority)
}
