package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var rucksacks []string
	var total1 int32
	var total2 int32
	var i int

	for scanner.Scan() {
		s := scanner.Text()
		shared1 := make(map[int32]bool)
		c1, c2 := s[:len(s)/2], s[len(s)/2:]

		// Part one
		for _, c := range c1 {
			if strings.Contains(c2, string(c)) {
				shared1[c] = true
			}
		}

		for k := range shared1 {
			total1 += getPriorVal(k)
		}

		// Part two
		rucksacks = append(rucksacks, s)
		i++

		if i == 3 {
			i = 0
			for _, r := range rucksacks[0] {
				if strings.Contains(rucksacks[1], string(r)) && strings.Contains(rucksacks[2], string(r)) {
					total2 += getPriorVal(r)
					break
				}
			}

			rucksacks = []string{}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", total1)
	fmt.Println("Part 2:", total2)
}

func getPriorVal(k int32) int32 {
	if k-96 < 0 {
		k += 58
	}
	return k - 96
}
