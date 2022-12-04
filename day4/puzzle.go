package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var fullyOverlappings, overlappings int // answers for part 1 and part 2, respectively.

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), ",")
		sectionOne := toSection(pairs[0])
		sectionTwo := toSection(pairs[1])

		if fullyOverlaps(sectionOne, sectionTwo) {
			fullyOverlappings++
		}

		if overlaps(sectionOne, sectionTwo) {
			overlappings++
		}
	}

	fmt.Println(fullyOverlappings, overlappings)
}

type section struct {
	startsAt, endsAt int
}

func toSection(pair string) section {
	parts := strings.Split(pair, "-")
	startsAt, _ := strconv.Atoi(parts[0])
	endsAt, _ := strconv.Atoi(parts[1])

	return section{startsAt, endsAt}
}

func fullyOverlaps(s1, s2 section) bool {
	return (s1.startsAt >= s2.startsAt && s1.endsAt <= s2.endsAt) || (s2.startsAt >= s1.startsAt && s2.endsAt <= s1.endsAt)
}

func overlaps(s1, s2 section) bool {
	return s1.startsAt <= s2.endsAt && s2.startsAt <= s1.endsAt
}
