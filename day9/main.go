package main

import (
	"fmt"
	"github.com/junpengxiao/Stack"
)

func main() {
	puzzleInput := "{{<a!>},{<a!>},{<a!>},{<ab>}}"
	score := countGroups(puzzleInput)

	fmt.Printf("# of groups: %d\n", score)
}

func countGroups(input string) int {
	s := stack.NewStack(uint(len(input)))
	sum := 0
	for _, c := range input {
		switch c {
		case '{':
			s.Push(c)
		case '}':
			sum += s.Len()
			s.Pop()
		}
	}

	return sum
}
