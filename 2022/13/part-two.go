package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func partTwo(scanner *bufio.Scanner) {
	decoderPackets := []string{
		"[[2]]",
		"[[6]]",
	}
	var decoderKey int

	packets := make([][]interface{}, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		packets = append(packets, parseLine(line))
	}

	for _, decoderPacket := range decoderPackets {
		packets = append(packets, parseLine(decoderPacket))
	}

	sort.Slice(packets, func(i int, j int) bool {
		return compareSides(packets[i], packets[j]) == COMPARE_PART_ONE_SUCCESS
	})

	for i, packet := range packets {
		for _, decoderPacket := range decoderPackets {
			if fmt.Sprint(packet) == decoderPacket {
				if decoderKey == 0 {
					decoderKey = i + 1
				} else {
					decoderKey *= i + 1
				}
			}
		}
	}

	fmt.Printf("The decoder key is %d\n", decoderKey)
}

func parseLine(line string) []interface{} {
	// Load up the JSON-like input
	var packet interface{}

	err := json.Unmarshal([]byte(line), &packet)

	if err != nil {
		fmt.Printf("Error unmarshaling JSON %s\n", err)

		os.Exit(1)
	}

	return packet.([]interface{})
}
