package common

import (
	"fmt"
	"strconv"
	"strings"
)

func StringToIntArray(dataString string) []int {
	data := strings.Split(dataString, ",")
	mem := []int{}
	for _, i := range data {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		mem = append(mem, j)
	}
	return mem
}

func RunProgram(mem []int) []int {
	//fmt.Println(mem)
	pc := 0
	for mem[pc] != 99 {
		if pc+3 >= len(mem) || mem[pc] >= len(mem) || mem[pc+1] >= len(mem) || mem[pc+2] >= len(mem) || mem[pc+3] >= len(mem) {
			fmt.Println("Array out of bounds error")
			return []int{0}
		}
		var op, a, b, out int
		op = mem[pc]
		a = mem[mem[pc+1]]
		b = mem[mem[pc+2]]
		out = mem[pc+3]
		if op == 1 {
			mem[out] = a + b
		} else if op == 2 {
			mem[out] = a * b
		} else {
			fmt.Printf("%d is not an op code\n", op)
			fmt.Printf("Op Code: %d, a: %d, b: %d, out: %d\n", op, a, b, out)
			return []int{0}
		}
		pc = pc + 4
	}
	//fmt.Printf("p1: %d, p2: %d, mem:%d\n", p1, p2, mem)
	return mem
}
