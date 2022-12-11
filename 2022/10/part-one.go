package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type RegisterValue struct {
	Beginning int
	Ending    int
}

func partOne(scanner *bufio.Scanner) {
	var (
		previousCycleNumber      int
		registerHistory          []RegisterValue
		specialCycles            []int
		specialCyclesStrengthSum int64
	)

	previousCycleNumber = 0

	registerHistory = make([]RegisterValue, 0)
	registerHistory = append(registerHistory, RegisterValue{
		Beginning: 1,
		Ending:    1,
	})

	for scanner.Scan() {
		line := scanner.Text()

		carryLastCycle := func() {
			registerHistory = append(registerHistory, RegisterValue{
				Beginning: registerHistory[previousCycleNumber].Ending,
				Ending:    registerHistory[previousCycleNumber].Ending,
			})

			previousCycleNumber++
		}

		carryLastCycle()

		if line != "noop" {
			lineParts := strings.Split(line, " ")
			modifier, _ := strconv.Atoi(lineParts[1])

			carryLastCycle()

			registerHistory[previousCycleNumber].Ending += modifier
		}
	}

	registerHistory = registerHistory[1:] // Shave off the extra one on the beginning of the data structure

	specialCycles = []int{
		20,
		60,
		100,
		140,
		180,
		220,
	}

	for _, specialCycle := range specialCycles {
		specialCyclesStrengthSum += int64(specialCycle) * int64(registerHistory[specialCycle-1].Beginning)
	}

	fmt.Printf("Special cycles strength sum is %d\n", specialCyclesStrengthSum)
}
