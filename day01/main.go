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
	f, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var elves []int
	var current int

	scanner := bufio.NewScanner(f)

	// Known bug - last elf isn't counted as there's no "" in input1.txt after it
	for scanner.Scan() {
		if scanner.Text() == "" {
			elves = append(elves, current)
			current = 0
			continue
		}

		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		current += i
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	fmt.Println("Part 1:", elves[0])
	fmt.Println("Part 2:", elves[0]+elves[1]+elves[2])
}
