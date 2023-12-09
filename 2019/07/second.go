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

	options := []int{9, 8, 7, 6, 5}

	settings := permutations(options)

	var maxThruster int
	for _, setting := range settings {
		var channels []chan int
		for i := 0; i < 5; i++ {
			channels = append(channels, make(chan int))
		}
		channels = append(channels, channels[0])
		channels = append(channels, make(chan int))
		for count, i := range setting {
			b := append(program[:0:0], program...)
			if count == 4 {
				go intCode(b, i, channels[count:count+3], count)
			} else {
				go intCode(b, i, channels[count:count+2], count)
			}
		}
		result := <-channels[6]
		if result > maxThruster {
			maxThruster = result

		}
	}

	fmt.Println(maxThruster)
}

func intInSlice(slice []int, integer int) bool {
	for n := range slice {
		if n == integer {
			return true
		}
	}
	return false
}

func intCode(program []int, inputIns int, channels []chan int, number int) {

	var n int
	var inputCount int
	for {
		instruction := program[n]
		opcode := instruction % 100

		var params [3]int

		instruction /= 100
		if len(program)-n <= 1 {
		} else if len(program)-n <= 3 {
			for i := 0; i <= 1; i++ {
				if instruction%10 == 0 {
					params[i] = program[n+i+1]
				} else {
					params[i] = n + i + 1
				}
				instruction /= 10
			}
		} else {
			for i := 0; i <= 2; i++ {
				if instruction%10 == 0 {
					params[i] = program[n+i+1]
				} else {
					params[i] = n + i + 1
				}
				instruction /= 10
			}
		}

		if opcode == 99 {
			if number == 4 {
				last := <-channels[0]
				channels[2] <- last
			}
			break
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
			} else if inputCount == 1 && number == 4 {
				program[params[0]] = 0
			} else {
				program[params[0]] = <-channels[0]
			}
			inputCount++
			n += 2
		} else if opcode == 4 {
			channels[1] <- program[params[0]]
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
