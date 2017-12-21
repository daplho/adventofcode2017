package main

import (
	"bufio"
	"fmt"
	"github.com/junpengxiao/Stack"
	"log"
	"os"
)

func main() {
	puzzleInput := readFile("./input.txt")
	//fmt.Printf("Puzzle input: '%s'", puzzleInput)

	score, garbageCount := countGroups(puzzleInput)
	fmt.Printf("# of groups: %d\n", score)
	fmt.Printf("# of garbage: %d\n", garbageCount)
}

func countGroups(input string) (int, int) {
	garbage := false
	garbageCount := 0
	ignoreNext := false
	s := stack.NewStack(uint(len(input)))
	var sum int = 0
	for _, c := range input {
		switch c {
		case '{':
			if !garbage {
				s.Push(c)
			} else {
				if !ignoreNext {
					garbageCount++
				}
				ignoreNext = false
			}
		case '}':
			if !garbage {
				sum += int(s.Len())
				s.Pop()
			} else {
				if !ignoreNext {
					garbageCount++
				}
				ignoreNext = false
			}
		case '<':
			if !garbage {
				garbage = true
			} else {
				if !ignoreNext {
					garbageCount++
				}
				ignoreNext = false
			}
		case '>':
			if garbage {
				if !ignoreNext {
					garbage = false
				} else {
					ignoreNext = false
				}
			}
		case '!':
			if garbage {
				if !ignoreNext {
					ignoreNext = true
				} else {
					ignoreNext = false
				}
			}
		default:
			if garbage {
				if !ignoreNext {
					garbageCount++
				}
				ignoreNext = false
			}
		}
	}

	fmt.Printf("stack size: %d\n", s.Len())

	return sum, garbageCount
}

func readFile(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := ""
	for scanner.Scan() {
		line = line + scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return line
}
