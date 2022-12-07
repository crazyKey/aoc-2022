package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	Name        string
	Parent      *Directory
	Directories []*Directory
	Files       []File
}

type File struct {
	Name string
	Size int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	dir := Directory{Name: "/"}
	cur := &dir

	for scanner.Scan() {
		l := strings.Split(scanner.Text(), " ")

		switch l[0] {
		case "$":
			if l[1] == "cd" {
				if l[2] == ".." {
					cur = cur.Parent
					continue
				}

				for _, d := range cur.Directories {
					if d.Name == l[2] {
						cur = d
						break
					}
				}
			}
		case "dir":
			cur.Directories = append(cur.Directories, &Directory{Name: l[1], Parent: cur})
		default:
			s, _ := strconv.Atoi(l[0])
			cur.Files = append(cur.Files, File{Name: l[1], Size: s})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var total1 int
	total2 := 70000000
	toFree := 7052440 // 30000000 - (total2 - used)

	fmt.Println("Used:", find(&dir, &total1, &total2, &toFree))

	fmt.Println("Part 1:", total1)
	fmt.Println("Part 2:", total2)
}

func find(dir *Directory, i, j, m *int) int {
	var td int
	for _, d := range dir.Directories {
		td += find(d, i, j, m)
	}

	var tf int
	for _, f := range dir.Files {
		tf += f.Size
	}

	t := td + tf

	if t <= 100000 {
		*i += t
	}

	if t > *m && t < *j {
		*j = t
	}

	return t
}
