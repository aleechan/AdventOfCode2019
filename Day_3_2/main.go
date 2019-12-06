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

	line := lines[0]
	i := 1
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
				points[pos] = i
				i++
			}
		case 'D':
			for j := 1; j <= steps; j++ {
				pos.y = pos.y - 1
				points[pos] = i
				i++
			}
		case 'R':
			for j := 1; j <= steps; j++ {
				pos.x = pos.x + 1
				points[pos] = i
				i++
			}
		case 'L':
			for j := 1; j <= steps; j++ {
				pos.x = pos.x - 1
				points[pos] = i
				i++
			}
		default:
			fmt.Printf("Could not find direction for %d", d)
		}
	}
	fmt.Println(pos)

	line = lines[1]
	i = 1
	dirs = strings.Split(line, ",")
	pos = point{x: 0, y: 0}

	min := math.MaxInt32

	for _, dir := range dirs {
		//fmt.Printf("%d %s\n", i, dir)
		d := rune(dir[0])
		steps, _ := strconv.Atoi(dir[1:len(dir)])

		switch d {
		case 'U':
			for j := 1; j <= steps; j++ {
				pos.y = pos.y + 1
				if val, ok := points[pos]; ok {
					if i+val < min {
						min = i + val
						fmt.Println(min)
					}
				}
				i++
			}
		case 'D':
			for j := 1; j <= steps; j++ {
				pos.y = pos.y - 1
				if val, ok := points[pos]; ok {
					if i+val < min {
						min = i + val
						fmt.Println(min)
					}
				}
				i++
			}
		case 'R':
			for j := 1; j <= steps; j++ {
				pos.x = pos.x + 1
				if val, ok := points[pos]; ok {
					if i+val < min {
						min = i + val
						fmt.Println(min)
					}
				}
				i++
			}
		case 'L':
			for j := 1; j <= steps; j++ {
				pos.x = pos.x - 1
				if val, ok := points[pos]; ok {
					if i+val < min {
						min = i + val
						fmt.Println(min)
					}
				}
				i++
			}
		default:
			fmt.Printf("Could not find direction for %d", d)
		}
	}

	fmt.Println(min)
}
