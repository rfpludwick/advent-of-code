package main

import (
	"bufio"
	"flag"
	"fmt"
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
	scanner.Scan()

	line := scanner.Text()

	// Iterate through the string
	leftPosition := 0
	rightPosition := 4

	for i := 0; i < (len(line) - 3); i++ {
		marker := line[leftPosition:rightPosition]

		if (strings.Count(marker, string(marker[0])) + strings.Count(marker, string(marker[1])) + strings.Count(marker, string(marker[2])) + strings.Count(marker, string(marker[3]))) == 4 {
			fmt.Printf("Buffer position is %d\n", rightPosition)

			break
		}

		leftPosition++
		rightPosition++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
