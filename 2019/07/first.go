package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	var program []int
	for _, n := range strings.Split(scanner.Text(), ",") {
		number, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		program = append(program, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	options := []int{0, 1, 2, 3, 4}

	settings := permutations(options)

	var maxSignal int
	var inputValue int
	for _, setting := range settings {
		for _, i := range setting {
			inputValue = intCode(program, i, inputValue)
		}
		if inputValue > maxSignal {
			maxSignal = inputValue
		}
		inputValue = 0
	}
	fmt.Println(maxSignal)

}

func intInSlice(slice []int, integer int) bool {
	for n := range slice {
		if n == integer {
			return true
		}
	}
	return false
}

func intCode(program []int, inputIns int, inputAmp int) int {

	var output int
	var inputCount int
	var n int
	for {
		instruction := program[n]
		opcode := instruction % 100

		if opcode == 99 {
			break
		}

		var params [3]int

		instruction /= 100
		for i := 0; i <= 2; i++ {
			if instruction%10 == 0 {
				params[i] = program[n+i+1]
			} else {
				params[i] = n + i + 1
			}
			instruction /= 10
		}

		if opcode == 1 {
			program[params[2]] = program[params[0]] + program[params[1]]
			n += 4
		} else if opcode == 2 {
			program[params[2]] = program[params[0]] * program[params[1]]
			n += 4
		} else if opcode == 3 {
			if inputCount == 0 {
				program[params[0]] = inputIns
			} else if inputCount == 1 {
				program[params[0]] = inputAmp
			}
			inputCount++
			n += 2
		} else if opcode == 4 {
			output = program[params[0]]
			n += 2
		} else if opcode == 5 {
			if program[params[0]] != 0 {
				n = program[params[1]]
			} else {
				n += 3
			}
		} else if opcode == 6 {
			if program[params[0]] == 0 {
				n = program[params[1]]
			} else {
				n += 3
			}
		} else if opcode == 7 {
			if program[params[0]] < program[params[1]] {
				program[params[2]] = 1
			} else {
				program[params[2]] = 0
			}
			n += 4
		} else if opcode == 8 {
			if program[params[0]] == program[params[1]] {
				program[params[2]] = 1
			} else {
				program[params[2]] = 0
			}
			n += 4
		}
	}

	return output
}

// https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
