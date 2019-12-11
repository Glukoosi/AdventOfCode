package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type tile struct{
	y int
	x int
	color int

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	var program [999999]int
	for i, n := range strings.Split(scanner.Text(), ",") {
		number, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		program[i] = number
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	inputChannel := make(chan int)
	outputChannel := make(chan int)
	closeChannel := make(chan bool)

	go intCode(&program, inputChannel, outputChannel, closeChannel)

	var grid []tile
	var coordinate [2]int
	input := 1
	var rotation int

	loop:
	for{
		var color int
		var dir int

		select{
		case <- closeChannel:
			printGrid(grid)
			break loop
		default:
			inputChannel <- input
			color = <-outputChannel
			dir = <-outputChannel

		var flag int
		for i,t := range grid {
			if t.y == coordinate[0] && t.x == coordinate[1] {
				grid[i].color = color
				flag = 1
			}
		}

		if flag == 0 || len(grid) == 0 {
			grid = append(grid, tile{y: coordinate[0], x: coordinate[1],color: color})
		}

		if dir == 0 {
			if rotation == 3{
				rotation = 0
			} else {
				rotation++
			}
		} else if dir == 1 {
			if rotation == 0{
				rotation = 3
			} else {
				rotation--
			}
		}

		if rotation == 0 { //up
			coordinate[0]++
		} else if rotation == 1 { //right {
			coordinate[1]++
		} else if rotation == 2 { //down {
			coordinate[0]--
		} else if rotation == 3 { //left {
			coordinate[1]--
		}

		input = 0
		for _, t := range grid {
			if t.y == coordinate[0] && t.x == coordinate[1] {
				input = t.color
			}
		}

		}
	}

}

func printGrid(grid []tile) {
	var minY int
	var minX int
	for _, t := range grid {
		if t.y < minY {
			minY = t.y
		}
		if t.x < minX{
			minX = t.x
		}
	}

	var gridArray [8][70]int
	for _, t := range grid {
		gridArray[t.y + -minY][t.x + -minX] = t.color
	}

	for _, r := range gridArray{
		fmt.Println(r)
	}
}

func intCode(program *[999999]int, inputChannel chan int, outputChannel chan int, closeChannel chan bool) {
	var n int
	var relBase int

	for {
		instruction := program[n]
		opcode := instruction % 100

		if opcode == 99 {
			closeChannel <- true
			break
		}

		var params [3]int

		instruction /= 100
		for i := 0; i <= 2; i++ {
			if instruction%10 == 0 {
				params[i] = program[n+i+1]
			} else if instruction%10 == 1 {
				params[i] = n + i + 1
			} else if instruction%10 == 2 {
				params[i] = program[n+i+1] + relBase
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
			program[params[0]] = <-inputChannel
			n += 2
		} else if opcode == 4 {
			outputChannel <- program[params[0]]
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
		} else if opcode == 9 {
			relBase += program[params[0]]
			n += 2
		}
	}

}

