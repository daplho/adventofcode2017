// http://adventofcode.com/2017/day/4

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type ReadLine func(string) bool

func main() {
	totalValid := readPassphraseList("./input.txt", readLine)
	fmt.Printf("Total Valid Passphrases: %d\n", totalValid)

	totalValidAnagrams := readPassphraseList("./input.txt", readLineAnagram)
	fmt.Printf("Total Valid Anagram Passphrases: %d\n", totalValidAnagrams)
}

func readPassphraseList(filepath string, fn ReadLine) int {
	totalValid := 0
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		valid := fn(scanner.Text())
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

func readLineAnagram(line string) bool {
	wordsMap := make(map[string]int, 0)

	words := strings.Fields(line)
	for _, word := range words {
		wordKey := sortString(word)
		wordsMap[wordKey]++
		if wordsMap[wordKey] > 1 {
			return false
		}
	}

	return true
}

func sortString(word string) string {
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
