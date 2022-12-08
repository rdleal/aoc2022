package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	trees := make([][]string, 0)
	for scanner.Scan() {
		row := scanner.Text()
		trees = append(trees, strings.Split(row, ""))
	}

	var maxScenicScore int

	visibleTrees := len(trees)*2 + (len(trees[0])-2)*2 // trees in edges are always visible

	for row := 1; row < len(trees)-1; row++ {
		for col := 1; col < len(trees[row])-1; col++ {
			visibleUp, viewingDistUp := isVisibleInRows(trees, row, col, -1)

			visibleDown, viewingDistDown := isVisibleInRows(trees, row, col, 1)

			visibleLeft, viewingDistLeft := isVisibleInCols(trees, row, col, -1)

			visibleRight, viewingDistRight := isVisibleInCols(trees, row, col, 1)

			if visibleUp || visibleDown || visibleLeft || visibleRight {
				visibleTrees++
			}

			scenicScore := viewingDistUp * viewingDistDown * viewingDistLeft * viewingDistRight

			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	fmt.Println(visibleTrees)   // answer for part 1
	fmt.Println(maxScenicScore) // answer for part 2
}

func isVisibleInRows(trees [][]string, row, col, dir int) (isVisible bool, viewingDist int) {
	treeHeight, _ := strconv.Atoi(trees[row][col])
	isVisible = true
	for row += dir; row >= 0 && row < len(trees); row += dir {
		viewingDist++
		if height, _ := strconv.Atoi(trees[row][col]); height >= treeHeight {
			isVisible = false
			break
		}
	}

	return isVisible, viewingDist
}

func isVisibleInCols(trees [][]string, row, col, dir int) (isVisible bool, viewingDist int) {
	treeHeight, _ := strconv.Atoi(trees[row][col])
	isVisible = true

	for col += dir; col >= 0 && col < len(trees[row]); col += dir {
		viewingDist++
		if height, _ := strconv.Atoi(trees[row][col]); height >= treeHeight {
			isVisible = false
			break
		}
	}

	return isVisible, viewingDist
}
