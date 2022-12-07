package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func partTwo(scanner *bufio.Scanner) {
	scanner.Scan()

	line := []byte(scanner.Text())

	// Iterate through the string
	leftPosition := 0
	rightPosition := 14

	for {
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
}
