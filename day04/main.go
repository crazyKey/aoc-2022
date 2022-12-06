package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var total1 int
	var total2 int

	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), ",")
		elf1 := strings.Split(pair[0], "-")
		elf2 := strings.Split(pair[1], "-")

		i1l, _ := strconv.Atoi(elf1[0])
		i1h, _ := strconv.Atoi(elf1[1])
		i2l, _ := strconv.Atoi(elf2[0])
		i2h, _ := strconv.Atoi(elf2[1])

		// Part one
		if i1l <= i2l && i2h <= i1h || i2l <= i1l && i1h <= i2h {
			total1++
		}

		// Part two
		if i1l <= i2l && i2l <= i1h || i1l <= i2h && i2h <= i1h || i2l <= i1l && i1l <= i2h || i2l <= i1h && i1h <= i2h {
			total2++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", total1)
	fmt.Println("Part 2:", total2)
}
