package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	argv := os.Args[1:]
	if len(argv) != 2 {
		os.Exit(0)
	}

	x, er1 := strconv.Atoi(argv[0])
	y, er2 := strconv.Atoi(argv[1])

	if er1 != nil || er2 != nil {
		// Atoi error
		os.Exit(0)
	}

	corner := "o"
	h_edge := "-"
	v_edge := "|"
	inside := " "

	for j := 0; j < y; j += 1 {
		line := ""
		for i := 0; i < x; i += 1 {
			if i == 0 || i == x-1 {
				if j == 0 || j == y-1 {
					line += corner
				} else {
					line += v_edge
				}
			} else if j == 0 || j == y-1 {
				line += h_edge
			} else {
				line += inside
			}
		}
		fmt.Println(line)
	}
}