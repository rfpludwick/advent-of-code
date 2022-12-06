package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
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

	line := []byte(scanner.Text())

	// Iterate through the string
	leftPosition := 0
	rightPosition := 14

	for i := 0; i < (len(line) - 13); i++ {
	loopStart:
		marker := line[leftPosition:rightPosition]

		for j := 0; j < 13; j++ {
			if bytes.Count(marker, marker[j:j+1]) > 1 {
				leftPosition++
				rightPosition++

				goto loopStart
			}
		}

		fmt.Printf("Buffer position is %d\n", rightPosition)

		break
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
