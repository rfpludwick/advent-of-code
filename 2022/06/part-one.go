package main

import (
	"bufio"
	"fmt"
	"strings"
)

func partOne(scanner *bufio.Scanner) {
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
}
