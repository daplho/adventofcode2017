// http://adventofcode.com/2017/day/7
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	puzzleInput := readFile("./input.txt")

	registers := make(map[string]int)
	var allTimeMax int
	for _, instruction := range puzzleInput {
		currentRValue := compute(instruction, registers)
		if currentRValue > allTimeMax {
			allTimeMax = currentRValue
		}
	}

	fmt.Printf("Registers:\n %v\n", registers)
	fmt.Printf("Max Final Register Value: %d\n", findMax(registers))
	fmt.Printf("Max All Time Register Value: %d\n", allTimeMax)
}

func findMax(registers map[string]int) int {
	if len(registers) == 0 {
		return 0
	}

	var max int
	for _, rValue := range registers {
		if max < rValue {
			max = rValue
		}
	}

	return max
}

func compute(instruction []string, registers map[string]int) int {
	if len(instruction) != 7 {
		panic(fmt.Sprintf("ERROR: instruction '%v' has incorrect length", instruction))
	}

	rname := instruction[0]

	operation := instruction[1]
	amount, _ := strconv.Atoi(instruction[2])
	conditionRName := instruction[4]
	operator := instruction[5]
	conditionAmount, _ := strconv.Atoi(instruction[6])

	switch operator {
	case ">":
		if registers[conditionRName] > conditionAmount {
			registers[rname] = doOperation(registers[rname], amount, operation)
			return registers[rname]
		}
	case "<":
		if registers[conditionRName] < conditionAmount {
			registers[rname] = doOperation(registers[rname], amount, operation)
			return registers[rname]
		}
	case ">=":
		if registers[conditionRName] >= conditionAmount {
			registers[rname] = doOperation(registers[rname], amount, operation)
			return registers[rname]
		}
	case "<=":
		if registers[conditionRName] <= conditionAmount {
			registers[rname] = doOperation(registers[rname], amount, operation)
			return registers[rname]
		}
	case "==":
		if registers[conditionRName] == conditionAmount {
			registers[rname] = doOperation(registers[rname], amount, operation)
			return registers[rname]
		}
	case "!=":
		if registers[conditionRName] != conditionAmount {
			registers[rname] = doOperation(registers[rname], amount, operation)
			return registers[rname]
		}
	}

	return -1
}

func doOperation(first int, second int, operation string) int {
	if operation == "inc" {
		return first + second
	}

	if operation == "dec" {
		return first - second
	}

	return 0
}

func readFile(filepath string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([][]string, 0)
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), " "))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
