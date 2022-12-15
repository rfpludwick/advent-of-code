package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

var (
	flagShowHelp  bool
	flagTestMode  bool
	flagInputFile string
	flagRunPart   int
)

const (
	PARTICLE_ROCK       = 1
	PARTICLE_SAND       = 2
	PARTICLE_SAND_ENTRY = 3
)

type Cave struct {
	Cavern     map[int]map[int]int
	MinX       int
	MaxX       int
	MinY       int
	MaxY       int
	SandEntryX int
	SandEntryY int
}

func newCave() *Cave {
	c := Cave{
		Cavern:     make(map[int]map[int]int, 0),
		MinX:       math.MaxInt,
		MaxX:       math.MinInt,
		MinY:       math.MaxInt,
		MaxY:       math.MinInt,
		SandEntryX: 500,
		SandEntryY: 0,
	}
	c.SetParticle(c.SandEntryX, c.SandEntryY, PARTICLE_SAND_ENTRY)

	return &c
}

func (c *Cave) SetParticle(x int, y int, particle int) {
	if _, ok := c.Cavern[x]; !ok {
		c.Cavern[x] = make(map[int]int)
	}

	c.Cavern[x][y] = particle

	if x < c.MinX {
		c.MinX = x
	}

	if x > c.MaxX {
		c.MaxX = x
	}

	if y < c.MinY {
		c.MinY = y
	}

	if y > c.MaxY {
		c.MaxY = y
	}
}

func (c *Cave) Tick(isAbyss bool) bool {
	x, y, possibleY := c.SandEntryX, c.SandEntryY, c.SandEntryY

	for {
		possibleXs := []int{
			x,
			x - 1,
			x + 1,
		}
		possibleY++
		settled := false

		if possibleY > c.MaxY {
			return true
		}

		for _, possibleX := range possibleXs {
			settled = false
			airFound := false

			if possibleX < c.MinX {
				if isAbyss {
					return true
				}

				c.MinX = possibleX
				c.SetParticle(c.MinX, c.MaxY, PARTICLE_ROCK)
			}

			if possibleX > c.MaxX {
				if isAbyss {
					return true
				}

				c.MaxX = possibleX
				c.SetParticle(c.MaxX, c.MaxY, PARTICLE_ROCK)
			}

			if _, ok := c.Cavern[possibleX]; !ok {
				airFound = true
			} else if _, ok2 := c.Cavern[possibleX][possibleY]; !ok2 {
				airFound = true
			}

			if airFound {
				x = possibleX
				y = possibleY

				break
			}

			settled = true
		}

		if settled {
			break
		}
	}

	c.SetParticle(x, y, PARTICLE_SAND)

	if x == c.SandEntryX && y == c.SandEntryY {
		return true
	}

	return false
}

func (c *Cave) Render() {
	// Render header
	leftSideWidth := len(fmt.Sprintf("%d", c.MaxY)) + 1
	leftPad := strings.Repeat(" ", leftSideWidth)

	for i := 0; i < len(fmt.Sprintf("%d", c.MaxX)); i++ {
		fmt.Print(leftPad)

		for x := c.MinX; x <= c.MaxX; x++ {
			// The following logic only renders:
			// Every 4th header
			// Min & max bounds
			// Position 500 since that is where the sand comes in
			if x%4 == 0 || x == c.MinX || x == c.MaxX || x == c.SandEntryX {
				fmt.Print(string(fmt.Sprintf("%3d", x)[i]))
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}

	fmt.Println()

	// Render body
	leftPadFormat := fmt.Sprintf("%s-%dd", "%", leftSideWidth)

	for y := c.MinY; y <= c.MaxY; y++ {
		fmt.Printf(leftPadFormat, y)

		for x := c.MinX; x <= c.MaxX; x++ {
			char := "."

			if _, ok := c.Cavern[x]; ok {
				if _, ok2 := c.Cavern[x][y]; ok2 {
					switch c.Cavern[x][y] {
					case PARTICLE_ROCK:
						char = "#"
					case PARTICLE_SAND:
						char = "o"
					case PARTICLE_SAND_ENTRY:
						char = "+"
					default:
						fmt.Printf("Somehow an unsupported particle type of %d is present\n", c.Cavern[x][y])
						os.Exit(1)
					}
				}
			}

			fmt.Print(char)
		}

		fmt.Println()
	}
}

func (c *Cave) GetSandUnitsCount() int {
	count := 0

	for x := c.MinX; x <= c.MaxX; x++ {
		for y := c.MinY; y <= c.MaxY; y++ {
			if c.Cavern[x][y] == PARTICLE_SAND {
				count++
			}
		}
	}

	return count
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
