package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var okLevelJumps = map[int]int{1: 0, 2: 0, 3: 0}

func main() {
	report := parseInput("./input.txt")

	puzzleTwo := true
	safe := 0
	if puzzleTwo {
		for _, line := range report {
			for i := range line {
				modLine := removeElementFromSlice(line, i)
				fmt.Printf("%v\n", modLine)
				if checkIncreaseOrDecrease(modLine) && checkLevelDifferences(modLine) {
					safe += 1
					fmt.Println("ok")
					break
				}
			}
		}

	} else {

		for _, line := range report {
			if checkIncreaseOrDecrease(line) && checkLevelDifferences(line) {
				safe += 1
			}
		}
	}
	fmt.Printf("Number of safe reports : %v\n", safe)
}

func parseInput(filepath string) (report [][]int) {
	report = make([][]int, 0)

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		r := strings.Split(line, " ")
		lvls := make([]int, 0)

		for _, lvl := range r {
			l, _ := strconv.Atoi(lvl)
			lvls = append(lvls, l)
		}

		report = append(report, lvls)
	}
	return report
}

func checkIncreaseOrDecrease(input []int) (ok bool) {
	var isDecreasing bool = false
	if input[0] > input[1] {
		isDecreasing = true
	}

	if isDecreasing {
		for i := 0; i < len(input)-1; i++ {
			if input[i] < input[i+1] {
				return false
			}
		}
	} else {
		for i := 0; i < len(input)-1; i++ {
			if input[i] > input[i+1] {
				return false
			}
		}
	}
	return true
}

func checkLevelDifferences(input []int) (ok bool) {
	var diff int
	for i := 0; i < len(input)-1; i++ {
		diff = abs(input[i] - input[i+1])
		if _, ok := okLevelJumps[diff]; !ok {
			return false
		}
	}
	return true
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func removeElementFromSlice(s []int, index int) (rs []int) {
	rs = make([]int, 0)
	rs = append(rs, s[:index]...)
	return append(rs, s[index+1:]...)
}
