package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/aleechan/AdventOfCode2019/common"
)

func main() {
	file := common.ReadFile("input.txt")

	total := 0

	for _, line := range file {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		temp := int(math.Floor(float64(num)/float64(3))) - 2
		total = total + temp
	}
	fmt.Println(total)
}
