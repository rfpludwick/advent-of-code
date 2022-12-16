package main

import (
	"bufio"
	"flag"
	"log"
	"math"
	"os"
)

var (
	flagShowHelp  bool
	flagTestMode  bool
	flagInputFile string
	flagRunPart   int
)

const (
	GRID_POSITION_SENSOR = 1
	GRID_POSITION_BEACON = 2
)

type Grid struct {
	Pairs []GridPair
	MinX  int
	MaxX  int
	MinY  int
	MaxY  int
}

type GridPair struct {
	SensorX  int
	SensorY  int
	BeaconX  int
	BeaconY  int
	Distance int
	Bounds   map[int][]int
}

func newGrid() *Grid {
	return &Grid{
		MinX: math.MaxInt,
		MaxX: math.MinInt,
		MinY: math.MaxInt,
		MaxY: math.MinInt,
	}
}

func (g *Grid) AddPair(sensorX int, sensorY int, beaconX int, beaconY int) {
	distance := calculateGridDistance(sensorX, sensorY, beaconX, beaconY)

	minSensorX := sensorX - distance
	maxSensorX := sensorX + distance
	minBeaconX := beaconX - distance
	maxBeaconX := beaconX + distance
	minSensorY := sensorY - distance
	maxSensorY := sensorY + distance
	minBeaconY := beaconY - distance
	maxBeaconY := beaconY + distance

	if minSensorX < g.MinX {
		g.MinX = minSensorX
	}

	if minBeaconX < g.MinX {
		g.MinX = minBeaconX
	}

	if maxSensorX > g.MaxX {
		g.MaxX = maxSensorX
	}

	if maxBeaconX > g.MaxX {
		g.MaxX = maxBeaconX
	}

	if minSensorY < g.MinY {
		g.MinY = minSensorY
	}

	if minBeaconY < g.MinY {
		g.MinY = minBeaconY
	}

	if maxSensorY > g.MaxY {
		g.MaxY = maxSensorY
	}

	if maxBeaconY > g.MaxY {
		g.MaxY = maxBeaconY
	}

	gridPair := GridPair{
		SensorX:  sensorX,
		SensorY:  sensorY,
		BeaconX:  beaconX,
		BeaconY:  beaconY,
		Distance: distance,
		Bounds:   make(map[int][]int, 0),
	}
	gridPair.calculateBounds()

	g.Pairs = append(g.Pairs, gridPair)
}

func (gp *GridPair) calculateBounds() {
	leftX, y, rightX := gp.SensorX, gp.SensorY-gp.Distance, gp.SensorX

	for i := gp.Distance * 2; i >= 0; i-- {
		gp.Bounds[y] = []int{
			leftX,
			rightX,
		}

		y++

		if y <= gp.SensorY {
			leftX--
			rightX++
		} else {
			leftX++
			rightX--
		}
	}
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
