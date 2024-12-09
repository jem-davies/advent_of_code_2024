package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	roomMap := parseInput("./input.txt")

	direction := "UP"
	directions := map[string][2]int{"UP": {-1, 0}, "DOWN": {1, 0}, "LEFT": {0, -1}, "RIGHT": {0, 1}}
	position := [2]int{52, 72}
	width := len(roomMap[0])
	height := len(roomMap)
	visitedPositions := [][2]int{}
	visitedPositions = append(visitedPositions, position)

	for {

		heightBound := position[0]+directions[direction][0] > height-1 || position[0]+directions[direction][0] < 0
		widthBound := position[1]+directions[direction][1] > width-1 || position[1]+directions[direction][1] < 0

		if heightBound || widthBound {
			fmt.Printf("Count: %v\n", len(visitedPositions))
			os.Exit(0)
		}

		inFront := roomMap[position[0]+directions[direction][0]][position[1]+directions[direction][1]]

		if inFront == "#" {
			direction = turn(direction)
			continue
		}

		position[0] = position[0] + directions[direction][0]
		position[1] = position[1] + directions[direction][1]

		if !visited(visitedPositions, position) {
			visitedPositions = append(visitedPositions, position)
		}
	}
}

func parseInput(filepath string) (roomMap [][]string) {
	roomMap = make([][]string, 0)

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		roomMap = append(roomMap, strings.Split(scanner.Text(), ""))
	}

	return roomMap
}

func turn(direction string) (newDirection string) {
	tempDirection := ""
	switch direction {
	case "UP":
		tempDirection = "RIGHT"
	case "DOWN":
		tempDirection = "LEFT"
	case "LEFT":
		tempDirection = "UP"
	case "RIGHT":
		tempDirection = "DOWN"
	}
	direction = tempDirection
	return direction

}

func visited(positions [][2]int, position [2]int) bool {
	for _, v := range positions {
		if v == position {
			return true
		}
	}
	return false
}
