package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid := make([][]rune, 0)

	scanner := bufio.NewScanner(os.Stdin)

	var startRow, startCol int
	for row := 0; scanner.Scan(); row++ {
		letters := scanner.Text()

		grid = append(grid, make([]rune, len(letters)))
		for col, l := range letters {
			if l == 'S' {
				startRow, startCol = row, col
			}
			grid[row][col] = l
		}
	}

	grid[startRow][startCol] = 'a'

	fromStart := shortestPath(grid, startRow, startCol)

	fmt.Println(fromStart) // part 1

	minSteps := fromStart
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == 'a' {
				s := shortestPath(grid, row, col)
				if s != -1 && s < minSteps {
					minSteps = s
				}
			}
		}
	}

	fmt.Println(minSteps) // part 2
}

func shortestPath(grid [][]rune, startRow, startCol int) int {
	distTo := make([][]int, len(grid))
	visited := make([][]bool, len(grid))

	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
		distTo[i] = make([]int, len(grid[0]))
	}

	queue := [][2]int{
		{startRow, startCol},
	}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		row, col := cell[0], cell[1]

		if !visited[row][col] {
			visited[row][col] = true

			dirs := [][]int{
				{1, 0},  // right
				{0, -1}, // up
				{-1, 0}, // left
				{0, 1},  // down
			}

			for _, dir := range dirs {
				rowDir, colDir := dir[0], dir[1]

				nextRow := row + rowDir
				if 0 > nextRow || nextRow >= len(grid) {
					continue
				}

				nextCol := col + colDir
				if 0 > nextCol || nextCol >= len(grid[0]) {
					continue
				}

				if grid[row][col] == 'E' {
					return distTo[row][col]
				}

				nextVal := grid[nextRow][nextCol]
				if grid[nextRow][nextCol] == 'E' {
					nextVal = 'z'
				}
				if nextVal-grid[row][col] > 1 {
					continue
				}

				distTo[nextRow][nextCol] = distTo[row][col] + 1

				queue = append(queue, [2]int{nextRow, nextCol})
			}
		}
	}

	return -1
}
