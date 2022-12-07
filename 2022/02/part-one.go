package main

import (
	"bufio"
	"log"
	"strings"
)

func partOne(scanner *bufio.Scanner) {
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

	log.Printf("Player's total score is %d", totalScore)
}
