package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func main() {
	pointsVisited := make(map[int]map[point]bool)

	var head point
	tails := make([]*point, 9)
	for i := range tails {
		tails[i] = new(point)
		pointsVisited[i] = map[point]bool{
			point{}: true,
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		step, _ := strconv.Atoi(parts[1])

		for i := 1; i <= step; i++ {
			switch dir := parts[0]; dir {
			case "L": // left
				head.x -= 1
			case "R": // right
				head.x += 1
			case "U": // up
				head.y += 1
			case "D": // down
				head.y -= 1
			}

			trackKnotPosition(head, tails, 0, pointsVisited)
		}
	}

	fmt.Println(len(pointsVisited[0]))            // part 1
	fmt.Println(len(pointsVisited[len(tails)-1])) // part 2
}

// inspired by alsm solution: https://github.com/alsm/aoc2022/blob/master/day9/day9.go
var dirs = []point{
	{x: 1, y: 0},
	{x: 0, y: 1},
	{x: -1, y: 0},
	{x: 0, y: -1},
}

var dirsDiagonal = []point{
	{x: 1, y: 1},
	{x: -1, y: 1},
	{x: -1, y: -1},
	{x: 1, y: -1},
}

func trackKnotPosition(head point, tails []*point, i int, pointsVisited map[int]map[point]bool) {
	if len(tails) == 0 {
		return
	}

	tail := tails[0]
	d := dist(head, *tail)
	if d == 1 {
		return
	}

	for ; d > 1; d = dist(head, *tail) {
		var p point
		if head.x == tail.x || head.y == tail.y {
			p = dir(head, *tail, dirs)
		} else {
			p = dir(head, *tail, dirsDiagonal)
		}
		tail.x += p.x
		tail.y += p.y

		if visited := pointsVisited[i][*tail]; !visited {
			pointsVisited[i][*tail] = true
		}
	}

	trackKnotPosition(*tail, tails[1:], i+1, pointsVisited)
}

func dir(head, tail point, dirs []point) point {
	quadrant := math.Atan2(float64(head.y-tail.y), float64(head.x-tail.x))
	dir := int((quadrant*4)/(2*math.Pi)+4) % 4
	return dirs[dir]
}

func absDiff(x, y int) int {
	return int(math.Abs(float64(x - y)))
}

func dist(head, tail point) int {
	return int(math.Sqrt(math.Pow(float64(tail.x-head.x), 2) + math.Pow(float64(tail.y-head.y), 2)))
}
