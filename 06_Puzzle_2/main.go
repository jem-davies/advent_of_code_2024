package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	roomMap := parseInput("./input.txt")

	/*
		UP 		0
		DOWN 	1
		LEFT 	2
		RIGHT 	3
	*/

	direction := 0
	directions := map[int][2]int{0: {-1, 0}, 1: {1, 0}, 2: {0, -1}, 3: {0, 1}}
	position := [3]int{52, 72, 0}
	width := len(roomMap[0])
	height := len(roomMap)
	visitedPositionsWithDirection := [][3]int{}
	visitedPositionsWithDirection = append(visitedPositionsWithDirection, position)

	for {

		heightBound := position[0]+directions[direction][0] > height-1 ||
			position[0]+directions[direction][0] < 0

		widthBound := position[1]+directions[direction][1] > width-1 ||
			position[1]+directions[direction][1] < 0

		if heightBound || widthBound {
			break
		}

		inFront := roomMap[position[0]+directions[direction][0]][position[1]+directions[direction][1]]

		if inFront == "#" {
			direction = turn(direction)
			continue
		}

		position[0] = position[0] + directions[direction][0]
		position[1] = position[1] + directions[direction][1]
		position[2] = direction

		if !visited(visitedPositionsWithDirection, position) {
			visitedPositionsWithDirection = append(visitedPositionsWithDirection, position)
		}
	}

	loopCount := 0

outerLoop:
	for _, v := range visitedPositionsWithDirection {
		if v == [3]int{52, 72, 0} {
			continue
		}

		visitedPositionsWithDirectionCopy := [][3]int{}
		position = [3]int{52, 72, 0}
		direction = 0
		newRoomMap := deepCopySlice(roomMap)
		newRoomMap[v[0]][v[1]] = "#"

	innerLoop:
		for {
			heightBound := position[0]+directions[direction][0] > height-1 || position[0]+directions[direction][0] < 0

			widthBound := position[1]+directions[direction][1] > width-1 || position[1]+directions[direction][1] < 0

			if heightBound || widthBound {
				continue outerLoop
			}

			inFront := newRoomMap[position[0]+directions[direction][0]][position[1]+directions[direction][1]]

			if inFront == "#" {
				direction = turn(direction)
				position[2] = direction
				continue innerLoop
			}

			position[0] = position[0] + directions[direction][0]
			position[1] = position[1] + directions[direction][1]
			position[2] = direction

			if looped(visitedPositionsWithDirectionCopy, position) {
				loopCount += 1
				continue outerLoop
			}
			visitedPositionsWithDirectionCopy = append(visitedPositionsWithDirectionCopy, position)

		}
	}
	timeElapsed := time.Since(start)
	fmt.Printf("The program took %s\n", timeElapsed)

	fmt.Printf("loopCount: %v\n", loopCount)
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

func turn(direction int) (newDirection int) {
	tempDirection := 0
	switch direction {
	case 0:
		tempDirection = 3
	case 1:
		tempDirection = 2
	case 2:
		tempDirection = 0
	case 3:
		tempDirection = 1
	}
	direction = tempDirection
	return direction
}

func visited(positions [][3]int, position [3]int) bool {
	for _, v := range positions {
		if v[0] == position[0] && v[1] == position[1] {
			return true
		}
	}
	return false
}

func looped(positions [][3]int, position [3]int) bool {
	for _, v := range positions {
		if v == position {
			return true
		}
	}
	return false
}

func deepCopySlice(original [][]string) [][]string {
	copySlice := make([][]string, len(original))

	for i := range original {
		if original[i] != nil {
			copySlice[i] = make([]string, len(original[i]))
			copy(copySlice[i], original[i])
		}
	}

	return copySlice
}
