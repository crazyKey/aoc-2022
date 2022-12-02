package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Pt.1 - XYZ as shapes
// A X Rock
// B Y Paper
// C Z Scissors

// Pt.2 - XYZ as instruction
// X Lose
// Y Draw
// Z Win

var (
	// Value of shapes
	shapes = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	// Scores for shape combinations
	combinations = map[string]map[string]int{
		"X": {"A": 3, "B": 0, "C": 6},
		"Y": {"A": 6, "B": 3, "C": 0},
		"Z": {"A": 0, "B": 6, "C": 3},
	}

	// Instructions for how to lose/draw/win
	instructions = map[string]map[string]string{
		"A": {"X": "Z", "Y": "X", "Z": "Y"},
		"B": {"X": "X", "Y": "Y", "Z": "Z"},
		"C": {"X": "Y", "Y": "Z", "Z": "X"},
	}
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
		// s[0] is elf shape and s[1] is my shape
		s := strings.Split(scanner.Text(), " ")

		// Part one
		total1 += combinations[s[1]][s[0]]
		total1 += shapes[s[1]]

		// Part two
		i := instructions[s[0]][s[1]]
		total2 += combinations[i][s[0]]
		total2 += shapes[i]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", total1)
	fmt.Println("Part 2:", total2)
}
