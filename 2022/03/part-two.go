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
	rucksacksCollected := 0
	rucksacks := make([]map[rune]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		rucksack := make(map[rune]int)

		// Build the rucksack
		for _, char := range strings.Split(line, "") {
			runeChar := []rune(char)[0]

			_, ok := rucksack[runeChar]

			if !ok {
				rucksack[runeChar] = 0
			}

			rucksack[runeChar]++
		}

		rucksacks = append(rucksacks, rucksack)

		rucksacksCollected++

		if rucksacksCollected == 3 {
			// Iterate through the chars
			for _, runeChar := range createRunesSlice() {
				_, ok1 := rucksacks[0][runeChar]
				_, ok2 := rucksacks[1][runeChar]
				_, ok3 := rucksacks[2][runeChar]

				if ok1 && ok2 && ok3 {
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

			rucksacksCollected = 0
			rucksacks = make([]map[rune]int, 0)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Total priority is %d", totalPriority)
}

func createRunesSlice() []rune {
	runeChars := make([]rune, 0)

	for i := 65; i <= 91; i++ {
		runeChars = append(runeChars, rune(i))
	}

	for i := 97; i <= 123; i++ {
		runeChars = append(runeChars, rune(i))
	}

	return runeChars
}
