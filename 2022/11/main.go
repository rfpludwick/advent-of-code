package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	flagShowHelp  bool
	flagTestMode  bool
	flagInputFile string
	flagRunPart   int
)

type Monkey struct {
	Items                   []uint64
	OperationLeft           string
	OperationOperator       rune
	OperationRight          string
	TestDivisor             uint64
	DivisorLCM              uint64
	TestTrueMonkeyReceiver  int
	TestFalseMonkeyReceiver int
	NumberItemsConsidered   int
}

func (m *Monkey) Operation(input uint64) uint64 {
	var (
		operationLeft  uint64
		operationRight uint64
	)

	switch m.OperationLeft {
	case "old":
		operationLeft = input
	default:
		operationLeft, _ = strconv.ParseUint(m.OperationLeft, 10, 64)
	}

	switch m.OperationRight {
	case "old":
		operationRight = input
	default:
		operationRight, _ = strconv.ParseUint(m.OperationRight, 10, 64)
	}

	switch m.OperationOperator {
	case '+':
		return operationLeft + operationRight
	case '*':
		return operationLeft * operationRight
	default:
		fmt.Printf("Monkey operation of '%s' is not supported\n", string(m.OperationOperator))

		os.Exit(1)
	}

	// We'll never get here
	return 0
}

func init() {
	flag.BoolVar(&flagShowHelp, "help", false, "Show this help")
	flag.BoolVar(&flagTestMode, "test", false, "Enable test mode")
	flag.StringVar(&flagInputFile, "input", "./input.txt", "Input file to use")
	flag.IntVar(&flagRunPart, "run-part", 1, "The part to run")

	flag.Parse()

	if flagShowHelp {
		flag.Usage()

		os.Exit(0)
	}

	if flagTestMode {
		flagInputFile = "./test-input.txt"
	}
}

func main() {
	file, err := os.Open(flagInputFile)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	switch flagRunPart {
	case 1:
		partOne(scanner)
	case 2:
		partTwo(scanner)
	default:
		log.Fatalf("Error: Part number %d is not supported", flagRunPart)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func initializeMonkeys(scanner *bufio.Scanner) []Monkey {
	monkeys := make([]Monkey, 0)

	currentMonkeyNumber := -1

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		// Note that the conditionals below are ordered by smaller to longer input string length checks
		// Also we're making an assumption here that our input monkeys are sequentially ordered from 0

		if line[0:6] == "Monkey" {
			currentMonkeyNumber, _ = strconv.Atoi(strings.Trim(line[7:], ":"))

			monkeys = append(monkeys, Monkey{})

			monkeys[currentMonkeyNumber].Items = make([]uint64, 0)
		} else if line[0:18] == "  Starting items: " {
			for _, linePart := range strings.Split(line[18:], ",") {
				item, _ := strconv.ParseUint(strings.Trim(linePart, " "), 10, 64)

				monkeys[currentMonkeyNumber].Items = append(monkeys[currentMonkeyNumber].Items, item)
			}
		} else if line[0:19] == "  Operation: new = " {
			lineParts := strings.Split(line[19:], " ")

			monkeys[currentMonkeyNumber].OperationLeft = lineParts[0]
			monkeys[currentMonkeyNumber].OperationOperator = rune(lineParts[1][0])
			monkeys[currentMonkeyNumber].OperationRight = lineParts[2]
		} else if line[0:21] == "  Test: divisible by " {
			monkeys[currentMonkeyNumber].TestDivisor, _ = strconv.ParseUint(line[21:], 10, 64)
		} else if line[0:29] == "    If true: throw to monkey " {
			monkeys[currentMonkeyNumber].TestTrueMonkeyReceiver, _ = strconv.Atoi(line[29:])
		} else if line[0:30] == "    If false: throw to monkey " {
			monkeys[currentMonkeyNumber].TestFalseMonkeyReceiver, _ = strconv.Atoi(line[30:])
		} else {
			fmt.Printf("Error processing input line '%s'\n", line)

			os.Exit(1)
		}
	}

	lcm := getMonkeysLcm(monkeys)

	for i := 0; i < len(monkeys); i++ {
		monkeys[i].DivisorLCM = lcm
	}

	return monkeys
}

// Following two functions generally from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func getMonkeysLcm(monkeys []Monkey) uint64 {
	// greatest common divisor (GCD) via Euclidean algorithm
	GCD := func(a, b uint64) uint64 {
		for b != 0 {
			t := b
			b = a % b
			a = t
		}
		return a
	}

	// find Least Common Multiple (LCM) via GCD
	var LCM func(...uint64) uint64

	LCM = func(integers ...uint64) uint64 {
		result := integers[0] * integers[1] / GCD(integers[0], integers[1])

		for i := 2; i < len(integers); i++ {
			result = LCM(result, integers[i])
		}

		return result
	}

	// My code
	allDivisors := make([]uint64, 0)

	for _, monkey := range monkeys {
		allDivisors = append(allDivisors, monkey.TestDivisor)
	}

	return LCM(allDivisors...)
}

func processRound(monkeys []Monkey, allowWorryRelief bool) {
	for i, monkey := range monkeys {
		for _, item := range monkey.Items {
			monkeys[i].NumberItemsConsidered++

			item = monkey.Operation(item)

			if allowWorryRelief {
				item /= 3
			}

			item %= monkey.DivisorLCM

			// if item%monkey.DivisorLCM == 0 {
			// 	item /= monkey.DivisorLCM
			// }

			var monkeyReceiver int

			if item%monkey.TestDivisor == 0 {
				monkeyReceiver = monkey.TestTrueMonkeyReceiver
			} else {
				monkeyReceiver = monkey.TestFalseMonkeyReceiver
			}

			monkeys[monkeyReceiver].Items = append(monkeys[monkeyReceiver].Items, item)
		}

		monkeys[i].Items = make([]uint64, 0)
	}
}
