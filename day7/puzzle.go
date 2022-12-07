package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var res int
	sizeDirs := make([]int, 0)
	rootDirSize := dfsFilesystem(scanner, &res, &sizeDirs)
	sizeDirs = append(sizeDirs, rootDirSize)

	fmt.Println(res) // part one answer

	freeUpSize := 30000000 - (70000000 - rootDirSize)

	sort.Ints(sizeDirs)
	for _, size := range sizeDirs {
		if size >= freeUpSize {
			fmt.Println(size) // part two answer
			break
		}
	}
}

func dfsFilesystem(scanner *bufio.Scanner, res *int, sizeDirs *[]int) int {
	var dirSize int
	for scanner.Scan() {
		s := scanner.Text()
		if isCommand(s) {
			cmd, subject := parseCommand(s)
			if subject == ".." {
				break
			}
			if cmd == "cd" {
				dirSize += dfsFilesystem(scanner, res, sizeDirs)
			}
			continue
		}

		parts := strings.Split(s, " ")
		if isDir(parts[0]) {
			continue
		}

		fileSize, _ := strconv.Atoi(parts[0])
		dirSize += fileSize
	}

	if dirSize <= 100000 {
		*res += dirSize
	}

	*sizeDirs = append(*sizeDirs, dirSize)

	return dirSize
}

func isCommand(s string) bool {
	return strings.HasPrefix(s, "$")
}

func parseCommand(s string) (cmd, subject string) {
	s = strings.TrimPrefix(s, "$ ")
	parts := strings.Split(s, " ")

	cmd = parts[0]
	if cmd == "ls" {
		return cmd, subject
	}

	return cmd, parts[1]
}

func isDir(s string) bool {
	return s == "dir"
}
