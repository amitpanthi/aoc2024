package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	calc := utils.ReadFile("calc.txt")
	partOne(calc)
}

func partOne(calc string) {
	calcs := getCalcMap(calc)

	ret := 0
	for target, nums := range calcs {
		if processCalc(target, nums[0], nums, 1, "*") ||
			processCalc(target, nums[0], nums, 1, "+") ||
			processCalc(target, nums[0], nums, 1, "||") {
			ret += target
		}
	}
	fmt.Println(ret)
}

func processCalc(target int, current int, nums []int, currIdx int, operator string) bool {
	if current == target && currIdx == len(nums) {
		return true
	}

	if current > target || currIdx >= len(nums) {
		return false
	}

	switch operator {
	case "||":
		currentStr := strconv.Itoa(current) + strconv.Itoa(nums[currIdx])
		current, _ = strconv.Atoi(currentStr)
	case "*":
		current = current * nums[currIdx]
	case "+":
		current = current + nums[currIdx]
	}

	return processCalc(target, current, nums, currIdx+1, "*") ||
		processCalc(target, current, nums, currIdx+1, "+") ||
		processCalc(target, current, nums, currIdx+1, "||")
}

func getCalcMap(calc string) map[int][]int {
	var cMap = make(map[int][]int)
	calcs := strings.Split(calc, "\r\n")
	for _, c := range calcs {
		var val []int
		cArr := strings.Split(c, ":")
		values := strings.Split(cArr[1], " ")

		for _, v := range values {
			num, _ := strconv.Atoi(v)
			val = append(val, num)
		}

		k, _ := strconv.Atoi(cArr[0])

		cMap[k] = val
	}

	return cMap
}
