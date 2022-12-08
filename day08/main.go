package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var grid [][]int

	for scanner.Scan() {
		var g []int

		for _, s := range scanner.Text() {
			i, _ := strconv.Atoi(string(s))
			g = append(g, i)
		}

		grid = append(grid, g)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	total1 := len(grid)*2 + len(grid[0])*2 - 4
	var total2 int

	// Part one
	for i := 1; i < len(grid)-1; i++ { // |
		for j := 1; j < len(grid[i])-1; j++ { // -
			var v []int
			for _, ii := range grid {
				v = append(v, ii[j])
			}

			a, ai := isHidden(v, i)
			b, bi := isHidden(grid[i], j)

			if !(a && b) {
				total1++
				if ai*bi > total2 {
					total2 = ai * bi
				}
			}
		}
	}

	fmt.Println("Part 1:", total1)
	fmt.Println("Part 2:", total2)
}

func isHidden(row []int, x int) (bool, int) {
	var li, ri int
	var l, r bool

	for i := x - 1; i >= 0; i-- {
		if !l {
			li++
		}
		if row[i] >= row[x] {
			l = true
		}
	}

	for i := x + 1; i < len(row); i++ {
		if !r {
			ri++
		}
		if row[i] >= row[x] {
			r = true
		}
	}

	return l && r, li * ri
}
