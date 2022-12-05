package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type crateStack []byte

func (s *crateStack) PushBack(vals ...byte) {
	*s = append(*s, vals...)
}

func (s *crateStack) PushFront(vals ...byte) {
	*s = append(vals, (*s)...)
}

func (s *crateStack) PopFront() byte {
	v := (*s)[0]
	*s = (*s)[1:]
	return v
}

func (s *crateStack) PopFrontN(n int) []byte {
	vals := make([]byte, n)
	copy(vals, []byte((*s)[:n]))
	*s = (*s)[n:]

	return vals
}

func (s *crateStack) Front() byte {
	return (*s)[0]
}

func main() {
	stacksOne := make([]crateStack, 9)
	stacksTwo := make([]crateStack, 9)

	scannedCrates := false

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		if len(txt) == 0 || unicode.IsDigit(rune(txt[1])) {
			scannedCrates = true // if it's either an empty line or stack numbers, crates has been scanned.
			continue
		}

		if !scannedCrates {
			parseStackLine(txt, stacksOne, stacksTwo)
			continue
		}

		quantity, from, to := parseMove(txt)

		crateMover9000(quantity, from, to, stacksOne)

		crateMover9001(quantity, from, to, stacksTwo)
	}

	fmt.Println(topCratesOf(stacksOne), topCratesOf(stacksTwo))
}

func parseStackLine(txt string, stacksOne, stacksTwo []crateStack) {
	var stack int
	for i := 0; i < len(txt); {
		if txt[i] == '[' {
			stacksOne[stack].PushBack(txt[i+1])
			stacksTwo[stack].PushBack(txt[i+1])
		}

		stack++
		i += 4 // skips to the next crate in the line.
	}
}

func crateMover9000(quantity, from, to int, stacks []crateStack) {
	for i := 0; i < quantity; i++ {
		stacks[to-1].PushFront(stacks[from-1].PopFront())
	}
}

func crateMover9001(quantity, from, to int, stacks []crateStack) {
	stacks[to-1].PushFront(stacks[from-1].PopFrontN(quantity)...)
}

func parseMove(txt string) (quantity, from, to int) {
	fmt.Sscanf(txt, "move %d from %d to %d", &quantity, &from, &to)
	return
}

func topCratesOf(stacks []crateStack) string {
	var crates []byte
	for i := 0; i < len(stacks); i++ {
		if s := stacks[i]; len(s) > 0 {
			crates = append(crates, s.Front())
		}
	}

	return string(crates)
}
