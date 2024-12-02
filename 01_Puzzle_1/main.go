package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	locationsA, locationsB := parseInput("./input.txt")

	slices.Sort(locationsA)
	slices.Sort(locationsB)

	distances := make([]int, 0)
	for i, location := range locationsA {
		distances = append(distances, abs(location-locationsB[i]))
	}

	answer := sum(distances)
	fmt.Printf("Total Distance: %v\n", answer)
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func sum(s []int) int {
	sum := 0
	for _, x := range s {
		sum += x
	}
	return sum
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
