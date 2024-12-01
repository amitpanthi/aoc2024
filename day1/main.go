package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("location.txt")
	if err != nil {
		fmt.Println("error opening location file")
		os.Exit(1)
	}

	rows := strings.Split(string(f), "\n")
	var dist1 []int
	var dist2 []int

	for _, r := range rows {
		cols := strings.Split(r, "   ")
		num1, _ := strconv.Atoi(cols[0])
		num2, _ := strconv.Atoi(strings.TrimSpace(cols[1]))

		dist1 = append(dist1, num1)
		dist2 = append(dist2, num2)
	}

	sort.Ints(dist1)
	sort.Ints(dist2)

	sum := 0.0

	for i := range dist1 {
		sum = sum + math.Abs(float64(dist2[i]-dist1[i]))
	}

	fmt.Println(int(sum)) // part 1

	rightMap := make(map[int]int)
	for _, d := range dist2 {
		if _, exists := rightMap[d]; exists {
			rightMap[d] = rightMap[d] + 1
		} else {
			rightMap[d] = 1
		}
	}

	distanceScore := 0
	leftMap := make(map[int]int)
	for _, d := range dist1 {
		if _, exists := leftMap[d]; !exists {
			rightOccurences := rightMap[d]
			distanceScore = distanceScore + (d * rightOccurences)
		} else {
			leftMap[d] = 0
		}
	}

	fmt.Println(distanceScore)
}
