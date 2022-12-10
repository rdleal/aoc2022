package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var cyclesCount, signalStrength int

	x := 1

	cycleValue := make(map[int]int)

	scanner := bufio.NewScanner(os.Stdin)
	for cycle := 1; cycle <= 240; cycle++ {
		signalStrength += cycleSignalStrength(cycle, x)

		if scanner.Scan() {
			parts := strings.Split(scanner.Text(), " ")

			switch op := parts[0]; op {
			case "addx":
				v, _ := strconv.Atoi(parts[1])
				cyclesCount += 2
				cycleValue[cyclesCount] = v
			case "noop":
				cyclesCount++
			}
		}

		drawCRT(cycle, x) // part 2

		if v, ok := cycleValue[cycle]; ok {
			x += v
		}
	}

	fmt.Println(signalStrength) // part 1
}

func cycleSignalStrength(cycle, x int) int {
	var signalStrength int
	switch cycle {
	case 20, 60, 100, 140, 180, 220:
		signalStrength = cycle * x
	}

	return signalStrength
}

const (
	litPx  = "#"
	darkPx = "."
)

func drawCRT(cycle, x int) {
	crtPos := (cycle - 1) % 40
	switch {
	case crtPos >= x-1 && x+1 >= crtPos:
		fmt.Print(litPx)
	default:
		fmt.Print(darkPx)
	}

	if cycle%40 == 0 {
		fmt.Print("\n")
	}
}
