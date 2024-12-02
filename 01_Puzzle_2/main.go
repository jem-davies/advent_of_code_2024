package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	locationsA, locationsB := parseInput("./input.txt")

	simScore := 0
	for _, locationA := range locationsA {
		num := 0
		for _, locationB := range locationsB {

			if locationA == locationB {
				num += 1
			}
		}
		simScore += num * locationA
	}

	fmt.Printf("simScore: %v\n", simScore)
}

func parseInput(filePath string) (locationsA []int, locationsB []int) {
	locationsA = make([]int, 0)
	locationsB = make([]int, 0)

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var a int
	var b int
	for scanner.Scan() {
		line := scanner.Text()
		r := strings.Split(line, "   ")
		a, _ = strconv.Atoi(r[0])
		b, _ = strconv.Atoi(r[1])
		locationsA = append(locationsA, a)
		locationsB = append(locationsB, b)
	}

	return locationsA, locationsB
}
