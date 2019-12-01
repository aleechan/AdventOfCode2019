package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/aleechan/AdventOfCode2019/common"
)

func calcFuel(mass int) int {
	fuel := int(math.Floor(float64(mass) / float64(3)))
	if fuel <= 2 {
		return 0
	}
	fuel = fuel - 2
	fmt.Println(fuel)
	return fuel + calcFuel(fuel)
}

func main() {
	file := common.ReadFile("input.txt")

	total := 0

	for _, line := range file {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		total = total + calcFuel(num)
	}
	fmt.Println(total)
}
