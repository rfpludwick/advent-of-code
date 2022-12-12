package main

import (
	"bufio"
	"fmt"
	"sort"
)

func partTwo(scanner *bufio.Scanner) {
	var (
		monkeys            []Monkey
		allItemsConsidered []int
		monkeyBusiness     int
	)

	monkeys = initializeMonkeys(scanner)

	for i := 0; i < 10000; i++ {
		processRound(monkeys, false)
	}

	// Gather the number of items considered
	for _, monkey := range monkeys {
		allItemsConsidered = append(allItemsConsidered, monkey.NumberItemsConsidered)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(allItemsConsidered)))

	monkeyBusiness = allItemsConsidered[0] * allItemsConsidered[1]

	fmt.Printf("Monkey business is %d\n", monkeyBusiness)
}
