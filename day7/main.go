// http://adventofcode.com/2017/day/7
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	puzzleInput := readFile("./input.txt")

	level := 1
	nodesByLevel := make([]map[string][]string, 0)
	nodesByLevel = append(nodesByLevel, puzzleInput)
	for {
		nodesByLevel = append(nodesByLevel, make(map[string][]string))

		prevLevel := makeMapCopy(nodesByLevel[level-1])
		for key, subkeys := range prevLevel {
			// if subkeys != nil && len(subkeys) > 0 {
			if prevLevelContainsAnyKey(prevLevel, subkeys) {
				nodesByLevel[level][key] = subkeys
				delete(nodesByLevel[level-1], key)
			}
			// }
		}

		if len(nodesByLevel[level]) == 0 {
			break
		}
		level++
	}

	for k, v := range nodesByLevel {
		fmt.Printf("%v -> %v\n", k, v)
	}
	fmt.Printf("Number of Levels: %v\n", level-1)
}

func prevLevelContainsAnyKey(prevLevel map[string][]string, keys []string) bool {
	for _, key := range keys {
		if _, ok := prevLevel[key]; ok {
			return true
		}
	}

	return false
}

func makeMapCopy(src map[string][]string) map[string][]string {
	dst := make(map[string][]string)

	for k, v := range src {
		dst[k] = v
	}

	return dst
}

func readFile(filepath string) map[string][]string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nodes := make(map[string][]string, 0)
	for scanner.Scan() {
		key, subkeys := parseLine(scanner.Text())
		nodes[key] = subkeys
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nodes
}

func parseLine(line string) (string, []string) {
	lineTokens := strings.Split(line, "->")

	if len(lineTokens) == 1 {
		keyWithWeightTokens := strings.Split(lineTokens[0], " ")
		key := strings.Trim(keyWithWeightTokens[0], " ")
		return key, nil
	}

	keyWithWeight, subkeysString := strings.Trim(lineTokens[0], " "), strings.Trim(lineTokens[1], " ")
	keyWithWeightTokens := strings.Split(keyWithWeight, " ")
	key := strings.Trim(keyWithWeightTokens[0], " ")

	return key, strings.Split(subkeysString, ", ")
}
