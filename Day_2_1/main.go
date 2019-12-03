package main

import (
	"fmt"

	"github.com/aleechan/AdventOfCode2019/common"
)

func main() {
	input := common.ReadFile("input.txt")[0]
	mem := common.StringToIntArray(input)
	mem = common.RunProgram(mem)
	fmt.Println(mem)

}
