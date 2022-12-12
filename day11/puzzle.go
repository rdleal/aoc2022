package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items          []int
	operation      string
	divisibleBy    int
	toMonkey       map[bool]int
	inspectedTimes int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	monkeys1, lcm := parseMonkeys(scanner)

	// copies monkeys for part 2
	monkeys2 := make([]*monkey, len(monkeys1))
	for i, monkey := range monkeys1 {
		m := *monkey
		monkeys2[i] = &m
	}

	monkeysRounds(monkeys1, 20, 0)
	monkeysRounds(monkeys2, 10000, lcm)

	sortMonkeys(monkeys1)
	sortMonkeys(monkeys2)

	fmt.Println(monkeys1[0].inspectedTimes * monkeys1[1].inspectedTimes)
	fmt.Println(monkeys2[0].inspectedTimes * monkeys2[1].inspectedTimes)
}

func parseMonkeys(scanner *bufio.Scanner) (monkeys []*monkey, lcm int) {
	lcm = 1
	monkeys = make([]*monkey, 0)
	var curMonkey *monkey
	for scanner.Scan() {
		txt := strings.TrimSpace(scanner.Text())
		if _, _, found := strings.Cut(txt, "Monkey "); found {
			curMonkey = &monkey{
				toMonkey: make(map[bool]int),
			}
			monkeys = append(monkeys, curMonkey)
		}

		if _, items, found := strings.Cut(txt, "Starting items: "); found {
			parts := strings.Split(items, ", ")

			startingItems := make([]int, len(parts))
			for i := range parts {
				startingItems[i], _ = strconv.Atoi(parts[i])
			}

			curMonkey.items = startingItems
		}

		if _, operation, found := strings.Cut(txt, "Operation: new = "); found {
			curMonkey.operation = operation
		}

		if _, num, found := strings.Cut(txt, "Test: divisible by "); found {
			n, _ := strconv.Atoi(num)
			curMonkey.divisibleBy = n
			// all dividors will always be prime in this aoc challenge,
			// so the lcm of all divisors is the product of all of them.
			// For more on this, see https://aoc.just2good.co.uk/2022/11
			lcm *= n
		}

		if _, m, found := strings.Cut(txt, "If true: throw to monkey "); found {
			toMonkey, _ := strconv.Atoi(m)
			curMonkey.toMonkey[true] = toMonkey

		}

		if _, m, found := strings.Cut(txt, "If false: throw to monkey "); found {
			toMonkey, _ := strconv.Atoi(m)
			curMonkey.toMonkey[false] = toMonkey
		}
	}

	return monkeys, lcm
}

func monkeysRounds(monkeys []*monkey, roundsNum int, lcm int) {
	for i := 0; i < roundsNum; i++ {
		for j := range monkeys {
			monkey := monkeys[j]

			parts := strings.Split(monkey.operation, " ")
			for len(monkey.items) > 0 {
				item := monkey.items[0]
				monkey.items = monkey.items[1:]

				var operand1, operand2 int
				if parts[0] == "old" {
					operand1 = item
				} else {
					operand1, _ = strconv.Atoi(parts[0])
				}

				if parts[2] == "old" {
					operand2 = item
				} else {
					operand2, _ = strconv.Atoi(parts[2])
				}

				var newValue int
				switch parts[1] {
				case "+":
					newValue = operand1 + operand2
				case "*":
					newValue = operand1 * operand2
				}

				if lcm == 0 {
					newValue /= 3 // part 1
				} else {
					newValue %= lcm // part 2
				}

				isDivisible := (newValue % monkey.divisibleBy) == 0
				toMonkey := monkey.toMonkey[isDivisible]

				monkeys[toMonkey].items = append(monkeys[toMonkey].items, newValue)

				monkey.inspectedTimes++
			}
		}
	}
}

func sortMonkeys(m []*monkey) {
	sort.Slice(m, func(i, j int) bool { return m[i].inspectedTimes > m[j].inspectedTimes })
}
