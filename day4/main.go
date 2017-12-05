// http://adventofcode.com/2017/day/2
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	totalValid := readPassphraseList("./input.txt")

	fmt.Printf("Total Valid Passphrases: %d\n", totalValid)
}

func readPassphraseList(filepath string) int {
	totalValid := 0
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		valid := readLine(scanner.Text())
		if valid {
			totalValid++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return totalValid
}

func readLine(line string) bool {
	wordsMap := make(map[string]int, 0)

	words := strings.Fields(line)
	for _, word := range words {
		wordsMap[word]++
		if wordsMap[word] > 1 {
			return false
		}
	}

	return true
}
