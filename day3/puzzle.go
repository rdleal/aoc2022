package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	var sumOne, sumTwo int

	// auxiliary maps for part 2
	groupPriorityOf := make(map[rune]int)
	visiteds := make(map[rune]int)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 1; scanner.Scan(); i++ {
		// part one
		priorityOf := make(map[rune]int)
		alreadySummed := make(map[rune]bool)

		items := scanner.Text()
		mid := len(items) / 2
		firstCompartment := items[:mid]
		secondCompartment := items[mid:]

		for _, item := range firstCompartment {
			priorityOf[item] = itemPriority(item)
		}

		for _, item := range secondCompartment {
			p, ok := priorityOf[item]
			if ok && !alreadySummed[item] {
				sumOne += p
				alreadySummed[item] = true
			}
		}

		// part two
		alreadyVisited := make(map[rune]bool)
		for _, item := range items {
			groupPriorityOf[item] = itemPriority(item)

			if !alreadyVisited[item] {
				visiteds[item]++
				alreadyVisited[item] = true
			}
		}

		if isEndOfGroup(i) {
			for item, occurrences := range visiteds {
				if occurrences == 3 {
					sumTwo += groupPriorityOf[item]
				}
			}
			// resets auxiliary maps for next group
			groupPriorityOf = make(map[rune]int)
			visiteds = make(map[rune]int)
		}
	}

	fmt.Println(sumOne, sumTwo)
}

const (
	lowercaseStartsAt = 97
	uppercaseStartsAt = 65
)

func itemPriority(r rune) int {
	if unicode.IsLower(r) {
		return int(r-lowercaseStartsAt) + 1 // 1 to 26
	}

	if unicode.IsUpper(r) {
		return int(r-uppercaseStartsAt) + 27 // 27 to 52
	}

	return 0
}

func isEndOfGroup(i int) bool {
	return i%3 == 0
}
