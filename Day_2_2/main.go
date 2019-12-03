package main

import (
	"fmt"

	"github.com/aleechan/AdventOfCode2019/common"
)

func runProgram(mem []int, p1, p2 int) []int {
	mem[1] = p1
	mem[2] = p2
	return common.RunProgram(mem)
}

func main() {
	const target = 19690720

	input := common.ReadFile("input.txt")[0]
	result := 0
	p1 := 0
	p2 := 0
	var mem []int
	for p1 = 0; result != target; p1++ {
		for p2 = 0; result != target && p2 < 100; p2++ {
			mem = runProgram(common.StringToIntArray(input), p1, p2)
			result = mem[0]
			fmt.Printf("p1: %d, p2: %d, result:%d\n", p1, p2, mem[0])
		}
	}
}
