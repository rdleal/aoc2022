package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	pairs := [][2]string{}
	for i, j := 0, 0; scanner.Scan(); i++ {
		txt := scanner.Text()
		if len(pairs) < j+1 {
			pairs = append(pairs, [2]string{})
		}

		if txt == "" {
			j++
			continue
		}

		pairs[j][(i-j)%2] = txt
	}

	packets := []string{
		"[[2]]",
		"[[6]]",
	}
	var indices int
	for i, pair := range pairs {
		ordered, _ := packetsInOrder(pair[0], pair[1])
		if ordered {
			indices += (i + 1)
		}

		packets = append(packets, pair[0], pair[1])
	}

	fmt.Println(indices) // part 1

	sort.Slice(packets, func(i, j int) bool {
		ordered, _ := packetsInOrder(packets[i], packets[j])
		return ordered
	})

	var indice2, indice6 int
	for i, packet := range packets {
		switch packet {
		case "[[2]]":
			indice2 = i + 1
		case "[[6]]":
			indice6 = i + 1
		}
	}

	fmt.Println(indice2 * indice6) // part 2
}

func packetsInOrder(left, right string) (inOrder, resolved bool) {
	var i, j int
	for i < len(left) && j < len(right) {
		l, r := rune(left[i]), rune(right[j])
		switch {
		case isListOpen(l):
			switch {
			case unicode.IsNumber(r):
				rn, sz := toInt(right[j:])
				j += sz - 1

				if inOrder, resolved = packetsInOrder(left[i:], "["+strconv.Itoa(rn)+"]"); resolved {
					return inOrder, resolved
				}
			case isListClose(r):
				return false, true
			}
		case unicode.IsNumber(l):
			ln, sz := toInt(left[i:])
			i += sz - 1

			switch {
			case unicode.IsNumber(r):
				rn, sz := toInt(right[j:])
				j += sz - 1

				if ln < rn {
					return true, true
				}

				if ln > rn {
					return false, true
				}

			case isListOpen(r):
				if inOrder, resolved = packetsInOrder("["+strconv.Itoa(ln)+"]", right[j:]); resolved {
					return inOrder, resolved
				}
			case isListClose(r):
				return false, true
			}
		case isListClose(l):
			if !isListClose(r) {
				return true, true
			}

		default:
			if isListClose(r) {
				return false, true
			}
		}

		j++
		i++
	}

	return false, false
}

func isListOpen(r rune) bool {
	return r == '['
}

func isListClose(r rune) bool {
	return r == ']'
}

func toInt(s string) (n, size int) {
	for i := 0; unicode.IsNumber(rune(s[i])); i++ {
		n = (int(rune(s[i])-'0') + (i * 10 * n))
		size++
	}

	return
}
