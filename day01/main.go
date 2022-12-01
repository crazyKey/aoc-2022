package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	one()
}

func one() {
	// Open file
	f, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var elves []int
	var current int

	// Read line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()

		// Append to elves or increment current
		if t == "" {
			elves = append(elves, current)
			current = 0
		} else {
			i, err := strconv.Atoi(t)
			if err != nil {
				log.Fatal(err)
			}
			current += i
		}

	}

	// Scanner errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Print sum of top 3
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	fmt.Println(elves[0] + elves[1] + elves[2])
}
