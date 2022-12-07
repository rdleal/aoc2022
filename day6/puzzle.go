package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		queue := scanner.Text()
		queuePointer := 0
		visiteds := make(map[byte]bool)

		var startOfPacketFound, startOfMessageFound bool

		i := queuePointer
		for !startOfPacketFound || !startOfMessageFound {
			b := queue[i]

			if visiteds[b] {
				visiteds = make(map[byte]bool)
				queuePointer++   // removes item from queue
				i = queuePointer // traceback to the next item in the queue
				continue
			}

			bytesRead := (i + 1) - queuePointer

			if !startOfPacketFound && bytesRead == 4 {
				fmt.Println("start-of-packet", i+1)
				startOfPacketFound = true
			}

			if !startOfMessageFound && bytesRead == 14 {
				fmt.Println("start-of-message", i+1)
				startOfMessageFound = true
			}

			visiteds[b] = true
			i++
		}
	}
}
