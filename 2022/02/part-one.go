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

	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		switch split[1] {
		case "X": // Player plays rock
			totalScore++

			switch split[0] {
			case "A": // Elf plays rock
				totalScore += 3
			case "C": // Elf plays scissors
				totalScore += 6
			}
		case "Y": // Player plays paper
			totalScore += 2

			switch split[0] {
			case "A": // Elf plays rock
				totalScore += 6
			case "B": // Elf plays paper
				totalScore += 3
			}
		case "Z": // Player plays scissors
			totalScore += 3

			switch split[0] {
			case "B": // Elf plays paper
				totalScore += 6
			case "C": // Elf plays scissors
				totalScore += 3
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Player's total score is %d", totalScore)
}
