package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/aleechan/AdventOfCode2019/common"
)

type point struct {
	x int
	y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	lines := common.ReadFile("input.txt")
	points := make(map[point]int)

	for i, line := range lines {
		dirs := strings.Split(line, ",")
		pos := point{x: 0, y: 0}

		for _, dir := range dirs {
			//fmt.Printf("%d %s\n", i, dir)
			d := rune(dir[0])
			steps, _ := strconv.Atoi(dir[1:len(dir)])

			switch d {
			case 'U':
				for j := 1; j <= steps; j++ {
					pos.y = pos.y + 1
					points[pos] = points[pos] + i + 1
				}
			case 'D':
				for j := 1; j <= steps; j++ {
					pos.y = pos.y - 1
					points[pos] = points[pos] + i + 1
				}
			case 'R':
				for j := 1; j <= steps; j++ {
					pos.x = pos.x + 1
					points[pos] = points[pos] + i + 1
				}
			case 'L':
				for j := 1; j <= steps; j++ {
					pos.x = pos.x - 1
					points[pos] = points[pos] + i + 1
				}
			default:
				fmt.Printf("Could not find direction for %d", d)
			}
		}
		fmt.Println(pos)
	}
	//fmt.Println(points)
	min := math.MaxInt32
	for point, value := range points {
		if value == 3 {
			dist := abs(point.x) + abs(point.y)
			if min > dist {
				min = dist
				fmt.Println(point)
			}
		}
	}
	fmt.Println(min)
}
