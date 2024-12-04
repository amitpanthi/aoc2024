package main

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strings"
)

var dirs = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

func main() {
	words := utils.ReadFile("words.txt")
	wordArr := strings.Split(words, "\n")
	letters := make([][]string, 0)
	for _, word := range wordArr {
		letters = append(letters, strings.Split(strings.TrimSpace(word), ""))
	}
	solve(letters)
}

func solve(letters [][]string) {
	//count = 3 for XMA. S takes care of itself
	//pick one dir, and keep going in that dir till you find XMAS, if not stop
	xmasCount := 0
	crossMasCount := 0

	for y, letterRow := range letters {
		for x := range letterRow {
			for _, dir := range dirs {
				xmasCount += getXmas(letters, x, y, dir, 3)
			}

			if x > 0 && x < len(letters[0])-1 && y > 0 && y < len(letters)-1 {
				if getCrossMas(letters, x, y) {
					crossMasCount += 1
				}
			}
		}
	}
	fmt.Println(xmasCount)
	fmt.Println(crossMasCount)
}

func getXmas(letters [][]string, x, y int, dir []int, count int) int {
	if count != 0 &&
		(x+dir[1] < 0 ||
			y+dir[0] < 0 ||
			x+dir[1] >= len(letters[0]) ||
			y+dir[0] >= len(letters)) {
		return 0
	}

	var nextExpectedLetter string
	currLetter := letters[y][x]

	if currLetter == "S" {
		if count == 0 {
			return 1
		}

		return 0
	}

	switch currLetter {
	case "X":
		nextExpectedLetter = "M"
	case "M":
		nextExpectedLetter = "A"
	case "A":
		nextExpectedLetter = "S"
	default:
		return 0
	}

	nextLetter := letters[y+dir[0]][x+dir[1]]

	if nextExpectedLetter != nextLetter {
		return 0
	}

	return getXmas(letters, x+dir[1], y+dir[0], dir, count-1)
}

func getCrossMas(letters [][]string, x, y int) bool {
	if letters[y][x] == "A" {
		lDiagString := []string{letters[y-1][x-1], letters[y][x], letters[y+1][x+1]}
		rDiagString := []string{letters[y-1][x+1], letters[y][x], letters[y+1][x-1]}
		sort.Strings(lDiagString)
		sort.Strings(rDiagString)

		return utils.IsSameArray(lDiagString, rDiagString) && utils.IsSameArray(lDiagString, []string{"A", "M", "S"})
	}

	return false
}
