package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	partOneAndTwo()
}

func partOneAndTwo() {
	s := utils.ReadFile("./reports.txt")

	reports := strings.Split(s, "\n")
	validCount := 0
	for _, report := range reports {
		if isValidReport(report) {
			validCount = validCount + 1
		}
	}

	isValidReport("1 2 7 8 9")

	fmt.Println(validCount)
}

func isValidReport(report string) bool {
	reportData := strings.Split(strings.TrimSpace(report), " ")
	inc := countUpOrDown(reportData, true)
	dec := countUpOrDown(reportData, false)

	if (inc > dec) && ((dec == 1) || (dec == 0 && inc+dec+2 == len(reportData))) {
		return isValid(reportData, true)
	} else if (dec > inc) && ((inc == 1) || (inc == 0 && inc+dec+2 == len(reportData))) {
		return isValid(reportData, false)
	} else {
		return false
	}
}

func isValid(a []string, increasing bool) bool {
	allowedMistakes := 1

	arr := make([]string, len(a))
	for i := range arr {
		arr[i] = a[i]
	}

	for i := 0; i < len(arr)-1; i++ {
		var first, second int

		if arr[i] == "s" {
			//skipped
			first, _ = strconv.Atoi(arr[i-1])
			second, _ = strconv.Atoi(arr[i+1])
		} else {
			first, _ = strconv.Atoi(arr[i])
			second, _ = strconv.Atoi(arr[i+1])
		}

		if !increasing {
			first, second = second, first
		}

		if (second-first) <= 0 || (second-first) > 3 {
			if allowedMistakes > 0 {
				allowedMistakes = allowedMistakes - 1
				arr[i] = "s"
			} else {
				return false
			}
		}
	}

	return true
}

func countUpOrDown(arr []string, up bool) int {
	var ret int
	for i := 0; i < len(arr)-1; i++ {
		first, _ := strconv.Atoi(arr[i])
		second, _ := strconv.Atoi(arr[i+1])

		if !up {
			first, second = second, first
		}

		if (second - first) > 0 {
			ret = ret + 1
		}
	}
	return ret
}
