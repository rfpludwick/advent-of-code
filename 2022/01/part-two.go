package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sort"
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

	elvesCalories := []int{0}
	currentElf := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			currentElf++

			elvesCalories = append(elvesCalories, 0)

			continue
		}

		numeral, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		elvesCalories[currentElf] += numeral
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(elvesCalories)

	total := 0

	for i := 0; i < 3; i++ {
		total += elvesCalories[len(elvesCalories)-1]

		elvesCalories = elvesCalories[:len(elvesCalories)-1]
	}

	log.Printf("High three elves' total calories is %d", total)
}
