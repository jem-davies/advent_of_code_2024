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

	// vertical top -> bottom

	vTopBot := ""
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			vTopBot += wordMatrix[j][i]
		}
		vTopBot += " "
	}

	// horizontal left -> right

	hLeftRight := ""
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			hLeftRight = hLeftRight + wordMatrix[i][j]
		}
		hLeftRight = hLeftRight + " "
	}

	// diagonal topleft -> bottom right

	dTopLeftBotRight := ""
	for i := width - 1; i >= 0; i-- {
		dTopLeftBotRight += wordMatrix[0][i]
		j := 1
		for j <= height-1 && i+j <= width-1 {
			dTopLeftBotRight += wordMatrix[j][i+j]
			j++
		}
		dTopLeftBotRight += " "
	}

	for i := 1; i < height-1; i++ {
		dTopLeftBotRight += wordMatrix[i][0]
		j := 1
		for j <= height-1 && i+j <= width-1 {
			dTopLeftBotRight += wordMatrix[i+j][j]
			j++
		}
		dTopLeftBotRight += " "
	}

	// diagonal top right -> bottom left

	dTopRightBotLeft := ""

	for i := width - 1; i >= 0; i-- {
		dTopRightBotLeft += wordMatrix[0][i]
		j := 1
		for j <= height-1 && i-j >= 0 {
			dTopRightBotLeft += wordMatrix[j][i-j]
			j++
		}
		dTopRightBotLeft += " "
	}

	for i := 1; i < height; i++ {
		dTopRightBotLeft += wordMatrix[i][width-1]
		j := 1
		for j <= height-1 && i+j <= height-1 && width-1-j >= 0 {
			dTopRightBotLeft += wordMatrix[i+j][width-1-j]
			j++
		}
		dTopRightBotLeft += " "
	}

	superString := vTopBot + " " + hLeftRight + " " + dTopLeftBotRight + " " + dTopRightBotLeft
	a := strings.Count(superString, "XMAS")
	b := strings.Count(superString, "SAMX")

	fmt.Printf("Count: %v\n", a+b)
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
