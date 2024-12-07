package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func main() {
	partOne()
}

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

var dirs = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func partOne() {
	var maze [][]string
	mazeString := utils.ReadFile("maze.txt")
	mazeRows := strings.Split(mazeString, "\r\n")
	for _, row := range mazeRows {
		maze = append(maze, strings.Split(row, ""))
	}
	l, rl := len(maze), len(maze)
	visited := make([][]int, l)
	for i := range visited {
		visited[i] = make([]int, rl)
	}

	var watchMan []int
	for i, row := range maze {
		for j := range row {
			if maze[i][j] == "^" {
				watchMan = []int{i, j}
			}
		}
	}

	currentDirection := UP
	blockCounts := 0

	for !isOutOfBounds(watchMan, l, rl) {
		nextStepDir := dirs[currentDirection]
		nextStepCoords := []int{watchMan[0] + nextStepDir[0], watchMan[1] + nextStepDir[1]}

		if isOutOfBounds(nextStepCoords, l, rl) {
			blockCounts += 1 // block exit
			visited[watchMan[0]][watchMan[1]] = 1
			break
		}

		vis := visited[watchMan[0]][watchMan[1]]
		if vis == 1 {
			blockCounts += 1
		}

		if maze[nextStepCoords[0]][nextStepCoords[1]] == "#" {
			currentDirection = (currentDirection + 1) % 4
			continue
		}

		visited[watchMan[0]][watchMan[1]] = 1
		watchMan = nextStepCoords
	}

	visitedNum := 0

	for _, r := range visited {
		for _, v := range r {
			if v == 1 {
				visitedNum += 1
			}
		}
	}

	fmt.Println(visitedNum)
	fmt.Println(blockCounts)
}

func isOutOfBounds(arr []int, numRows, numCols int) bool {
	return (arr[0] < 0 ||
		arr[0] >= numCols ||
		arr[1] < 0 ||
		arr[1] >= numRows)
}

func checkIfBlockIsValid(visited [][]int, direction int, x, y int) bool {
	direction = (direction + 1) % 4
	nextStepDir := dirs[direction]
	nextStepCoords := []int{y + nextStepDir[0], x + nextStepDir[1]}

	return visited[nextStepCoords[0]][nextStepCoords[1]] == 1
}
