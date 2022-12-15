package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func partOne(scanner *bufio.Scanner) {
	currentPair := 1
	currentSide := 1
	var leftSide, rightSide interface{}
	correctOrderResult := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 { // We're moving on to a new pair
			currentPair++
			currentSide = 1

			continue
		}

		// Load up the JSON-like input to the correct side
		var err error

		if currentSide == 1 {
			err = json.Unmarshal([]byte(line), &leftSide)
		} else {
			err = json.Unmarshal([]byte(line), &rightSide)
		}

		if err != nil {
			fmt.Printf("Error unmarshaling JSON %s\n", err)

			os.Exit(1)
		}

		currentSide++

		// Process the pair
		if currentSide > 2 {
			if compareSides(leftSide.([]interface{}), rightSide.([]interface{})) == COMPARE_PART_ONE_SUCCESS {
				correctOrderResult += currentPair
			}
		}
	}

	fmt.Printf("The sum of the correctly-paired indices is %d\n", correctOrderResult)
}
