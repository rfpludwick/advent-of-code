package main

import (
	"bufio"
	"log"
	"strings"
)

func partTwo(scanner *bufio.Scanner) {
	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		switch split[0] {
		case "A": // Elf plays rock
			switch split[1] {
			case "X": // Player to lose with scissors
				totalScore += 3
			case "Y": // Player to tie with rock
				totalScore += (1 + 3)
			case "Z": // Player to win with paper
				totalScore += (2 + 6)
			}
		case "B": // Elf plays paper
			switch split[1] {
			case "X": // Player to lose with rock
				totalScore += 1
			case "Y": // Player to tie with paper
				totalScore += (2 + 3)
			case "Z": // Player to win with scissors
				totalScore += (3 + 6)
			}
		case "C": // Elf plays scissors
			switch split[1] {
			case "X": // Player to lose with paper
				totalScore += 2
			case "Y": // Player to tie with scissors
				totalScore += (3 + 3)
			case "Z": // Player to win with rock
				totalScore += (1 + 6)
			}
		}
	}

	log.Printf("Player's total score is %d", totalScore)
}
