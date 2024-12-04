package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	wordMatrix := parseInput("./input.txt")
	height := len(wordMatrix)
	width := len(wordMatrix[0])

	sum := 0

	for i := 0; i < height-2; i++ {
		for j := 0; j < width-2; j++ {

			tmp := getThreeByThree(i, j, wordMatrix)
			if check(tmp) {
				sum += 1
			}
		}
	}
	fmt.Printf("%v\n", sum)
}

func parseInput(filepath string) (wordMatrix [][]string) {
	wordMatrix = make([][]string, 0)

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordMatrix = append(wordMatrix, strings.Split(scanner.Text(), ""))
	}

	return wordMatrix
}

func getThreeByThree(i int, j int, input [][]string) [][]string {
	return [][]string{
		{input[i][j], input[i][j+1], input[i][j+2]},
		{input[i+1][j], input[i+1][j+1], input[i+1][j+2]},
		{input[i+2][j], input[i+2][j+1], input[i+2][j+2]},
	}
}

func check(input [][]string) bool {
	ff := [][]string{{"M", "?", "S"}, {"?", "A", "?"}, {"M", "?", "S"}}
	fb := [][]string{{"M", "?", "M"}, {"?", "A", "?"}, {"S", "?", "S"}}
	bf := [][]string{{"S", "?", "S"}, {"?", "A", "?"}, {"M", "?", "M"}}
	bb := [][]string{{"S", "?", "M"}, {"?", "A", "?"}, {"S", "?", "M"}}

	if input[0][0] == ff[0][0] && input[0][2] == ff[0][2] && input[1][1] == "A" && input[2][0] == ff[2][0] && input[2][2] == ff[2][2] {
		return true
	}
	if input[0][0] == fb[0][0] && input[0][2] == fb[0][2] && input[1][1] == "A" && input[2][0] == fb[2][0] && input[2][2] == fb[2][2] {
		return true
	}
	if input[0][0] == bf[0][0] && input[0][2] == bf[0][2] && input[1][1] == "A" && input[2][0] == bf[2][0] && input[2][2] == bf[2][2] {
		return true
	}
	if input[0][0] == bb[0][0] && input[0][2] == bb[0][2] && input[1][1] == "A" && input[2][0] == bb[2][0] && input[2][2] == bb[2][2] {
		return true
	}
	return false
}
