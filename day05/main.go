package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//var crates = [][]string{
//	{"Z", "N"},
//	{"M", "C", "D"},
//	{"P"},
//}

var crates = [][]string{
	{"R", "N", "F", "V", "L", "J", "S", "M"},
	{"P", "N", "D", "Z", "F", "J", "W", "H"},
	{"W", "R", "C", "D", "G"},
	{"N", "B", "S"},
	{"M", "Z", "W", "P", "C", "B", "F", "N"},
	{"P", "R", "M", "W"},
	{"R", "T", "N", "G", "L", "S", "W"},
	{"Q", "T", "H", "F", "N", "B", "V"},
	{"L", "M", "H", "Z", "N", "F"},
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		move := strings.Split(scanner.Text(), " ")
		num, _ := strconv.Atoi(move[1])
		from, _ := strconv.Atoi(move[3])
		to, _ := strconv.Atoi(move[5])

		// Part one
		//for i := 0; i < num; i++ {
		//	crates[to-1] = append(crates[to-1], crates[from-1][len(crates[from-1])-1:]...)
		//	crates[from-1] = crates[from-1][:len(crates[from-1])-1]
		//}

		// Part two
		crates[to-1] = append(crates[to-1], crates[from-1][len(crates[from-1])-num:]...)
		crates[from-1] = crates[from-1][:len(crates[from-1])-num]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print("Top crates: ")
	for _, c := range crates {
		fmt.Print(c[len(c)-1])
	}

	fmt.Println()
}
