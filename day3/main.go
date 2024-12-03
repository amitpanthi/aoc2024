package main

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	programText := utils.ReadFile("program.txt")
	partOne(programText)
	partTwo(programText)
}

func partOne(programText string) {
	r := regexp.MustCompile(`mul\(\d*,\d*\)`)
	validMuls := r.FindAllString(programText, -1)

	mulRegex := regexp.MustCompile(`\d+`)
	total := 0
	for _, v := range validMuls {
		nums := mulRegex.FindAllString(v, 2)
		first, _ := strconv.Atoi(nums[0])
		second, _ := strconv.Atoi(nums[1])

		total = total + (first * second)
	}

	fmt.Println(total)
}

func partTwo(programText string) {
	negRegex := regexp.MustCompile(`don\'t\(\)`)
	posRegex := regexp.MustCompile(`do\(\)`)

	//start looking at each negIdx, and find the nearest posIdx after that.
	negIdxs := getEndingIdxOfWord(negRegex.FindAllIndex([]byte(programText), -1))
	posIdxs := getEndingIdxOfWord(posRegex.FindAllIndex([]byte(programText), -1))

	validStartAndEndIdx := getValidStartAndEndIdxs(posIdxs, negIdxs)
	finalProgram := getFinalProgramData(programText, validStartAndEndIdx)
	partOne(finalProgram)
}

func getEndingIdxOfWord(idxs [][]int) []int {
	var ret []int
	for _, idx := range idxs {
		ret = append(ret, idx[1])
	}

	return ret
}

func getValidStartAndEndIdxs(pos []int, neg []int) [][]int {
	var ret [][]int
	var posIter = 0

	for i, n := range neg {
		for {
			//skip repeated dont's

			if posIter >= len(pos) {
				ret = append(ret, []int{n, -1})
				break
			}

			if pos[posIter] > n {
				if i > 0 && ret[len(ret)-1][1] == pos[posIter] {
					break
				}
				ret = append(ret, []int{n, pos[posIter]})
				break
			}

			posIter++
		}

		if posIter >= len(pos) {
			break
		}
	}

	return ret
}

func getFinalProgramData(program string, invalidIdx [][]int) string {
	currentIdx := 0
	ret := ""
	for _, idx := range invalidIdx {
		ret = ret + program[currentIdx:idx[0]]
		currentIdx = idx[1]
	}

	if currentIdx < len(program) && invalidIdx[len(invalidIdx)-1][1] != -1 {
		ret = ret + program[invalidIdx[len(invalidIdx)-1][1]:]
	}

	return ret
}
