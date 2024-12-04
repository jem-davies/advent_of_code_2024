package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	sum := 0
	lines := parseInput("input.txt")
	for _, line := range lines {
		matches := r.FindAllString(line, -1)
		for _, match := range matches {
			sum += executeMul(match)
		}
	}
	fmt.Printf("sum: %v\n", sum)
}

func parseInput(filepath string) (lines []string) {
	lines = make([]string, 0)

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func executeMul(mul string) int {
	mul = mul[4:]
	mul = mul[:len(mul)-1]
	x := strings.Split(mul, ",")
	a, _ := strconv.Atoi(x[0])
	b, _ := strconv.Atoi(x[1])

	return a * b
}
