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

	score := countGroups(puzzleInput)
	fmt.Printf("# of groups: %d\n", score)
}

func countGroups(input string) int64 {
	garbage := false
	ignoreNext := false
	s := stack.NewStack(uint(len(input)))
	var sum int64 = 0
	for _, c := range input {
		switch c {
		case '{':
			if !garbage && !ignoreNext {
				s.Push(c)
			}
		case '}':
			if !garbage && !ignoreNext {
				sum += int64(s.Len())
				s.Pop()
			}
		case '<':
			if !ignoreNext {
				garbage = true
			}
		case '>':
			if !ignoreNext {
				garbage = false
			}
		case '!':
			if !ignoreNext {
				ignoreNext = true
			} else {
				ignoreNext = false
			}
		default:
			ignoreNext = false
		}
	}

	fmt.Printf("stack size: %d\n", s.Len())

	return sum
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
