package main

import (
	"aoc/utils"
	"fmt"
	"math"
)

type point struct {
	x int
	y int
}

func main() {
	city := utils.ReadFile("map.txt")
	cityMap := utils.Get2DArrayFromString(city, "")
	partOne(cityMap)
}

func partOne(cityMap [][]string) {
	var antenna = make(map[string][]point)

	for i, r := range cityMap {
		for j := range r {
			if cityMap[i][j] == "." {
				continue
			}

			if _, ok := antenna[cityMap[i][j]]; !ok {
				antenna[cityMap[i][j]] = make([]point, 0)
			}

			antenna[cityMap[i][j]] = append(antenna[cityMap[i][j]], point{x: j, y: i})

		}
	}

	var totalNodes = make(map[point]int)

	for _, v := range antenna {
		for _, n := range getAntinodes(v, len(cityMap[0]), len(cityMap)) {
			totalNodes[n] = 1
		}
	}

	fmt.Println(len(totalNodes))
}

func getAntinodes(antenna []point, xlim, ylim int) []point {
	var ret []point
	for i := range antenna {
		for j := i + 1; j < len(antenna); j++ {
			ret = append(ret, getAntinode(antenna[i], antenna[j], xlim, ylim)...)
		}
	}

	return ret
}

func getAntinode(p1, p2 point, xlim, ylim int) []point {
	var ret []point

	xDif := int(math.Abs(float64(p2.x) - float64(p1.x)))
	yDif := int(math.Abs(float64(p2.y) - float64(p1.y)))

	var newP1X, newP1Y int
	if p1.x > p2.x {
		newP1X = xDif
	} else {
		newP1X = -xDif
	}

	if p1.y > p2.y {
		newP1Y = yDif
	} else {
		newP1Y = -yDif
	}

	firstNode := point{x: p1.x + newP1X, y: p1.y + newP1Y}
	secondNode := point{x: p2.x - newP1X, y: p2.y - newP1Y}

	firstOob := pointOutOfBounds(firstNode, xlim, ylim)
	secondOob := pointOutOfBounds(secondNode, xlim, ylim)
	for !firstOob || !secondOob {
		if !firstOob {
			ret = append(ret, firstNode)
		}

		if !secondOob {
			ret = append(ret, secondNode)
		}

		firstNode = point{x: firstNode.x + newP1X, y: firstNode.y + newP1Y}
		secondNode = point{x: secondNode.x - newP1X, y: secondNode.y - newP1Y}

		firstOob = pointOutOfBounds(firstNode, xlim, ylim)
		secondOob = pointOutOfBounds(secondNode, xlim, ylim)
	}

	ret = append(ret, p1, p2)

	return ret
}

func pointOutOfBounds(p point, x, y int) bool {
	return !(p.x >= 0 && p.x < x &&
		p.y >= 0 && p.y < y)
}
