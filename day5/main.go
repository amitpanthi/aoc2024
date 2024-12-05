package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	instructions, incorrect := partOne()
	partTwo(instructions, incorrect)
}

func partOne() (map[string][]string, [][]string) {
	instructions := processInstructions()
	manual := utils.ReadFile("manual.txt")
	manualArr := strings.Split(manual, "\n")
	var res [][]string
	var incorrect [][]string

	for _, m := range manualArr {
		valid := true
		m = strings.TrimSpace(m)
		mArr := strings.Split(m, ",")
		for i := range mArr {
			for j := i + 1; j < len(mArr)-1; j++ {
				if !utils.Contains(instructions[mArr[i]], mArr[j]) &&
					utils.Contains(instructions[mArr[j]], mArr[i]) {
					valid = false
					break
				}
			}

			if !valid {
				break
			}
		}

		if valid {
			res = append(res, mArr)
		} else {
			incorrect = append(incorrect, mArr)
		}
	}

	sum := 0
	for _, r := range res {
		middle, _ := strconv.Atoi(r[(len(r) / 2)])
		sum += middle
	}

	fmt.Println(sum)

	return instructions, incorrect
}

func partTwo(instructions map[string][]string, incorrect [][]string) {
	for _, s := range incorrect {
		fmt.Println(s)
		for i := range s {
			for j := range s {
				if utils.Contains(instructions[s[j]], s[i]) {
					temp := s[j]
					s[j] = s[i]
					s[i] = temp
				}
			}
		}
		fmt.Println(s)
	}

	sum := 0
	for _, r := range incorrect {
		middle, _ := strconv.Atoi(r[(len(r) / 2)])
		sum += middle
	}
	fmt.Println(sum)
}

func processInstructions() map[string][]string {
	var instructionMap = make(map[string][]string)
	instructions := utils.ReadFile("manual_instructions.txt")
	instructionArr := strings.Split(instructions, "\n")

	for _, inst := range instructionArr {
		instSplit := strings.Split(inst, "|")
		prev, next := strings.TrimSpace(instSplit[0]), strings.TrimSpace(instSplit[1])
		if _, ok := instructionMap[prev]; !ok {
			instructionMap[prev] = make([]string, 0)
		}

		instructionMap[prev] = append(instructionMap[prev], next)
	}

	return instructionMap
}
