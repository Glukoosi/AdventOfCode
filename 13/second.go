package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type tile struct{
	y int
	x int
	tileId int

}

var grid []tile

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

	program[0] = 2

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	outputChannel := make(chan int)
	closeChannel := make(chan bool)

	go intCode(&program, outputChannel, closeChannel)

	var coordinate [2]int
	var tileId int
	var score int

	loop:
	for{
		select{
		case <- closeChannel:
			fmt.Println(score)
			break loop
		case c := <-outputChannel:
			coordinate[1] = c
			coordinate[0] = <-outputChannel
			tileId = <-outputChannel

			if coordinate == [2]int{0, -1}{
				score = tileId
			} else {

				var flag int
				for i,t := range grid {
					if t.y == coordinate[0] && t.x == coordinate[1] {
						grid[i].tileId = tileId
						flag = 1
					}
				}

				if flag == 0 || len(grid) == 0 {
					grid = append(grid, tile{y: coordinate[0], x: coordinate[1],tileId: tileId})
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

	var gridArray [46][46]int
    for _, t := range grid {
        gridArray[t.y + -minY][t.x + -minX] = t.tileId
    }

    for _, r := range gridArray{
        fmt.Println(r)
    }
}

func intCode(program *[999999]int, outputChannel chan int, closeChannel chan bool) {
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
			program[params[0]] = input()
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

func input() int {
	var ballX int
	var output int

	time.Sleep(1 * time.Millisecond)
	for _,t := range grid {
		if t.tileId == 4 {
			ballX = t.x
		}
	}
	for _,t := range grid {
		if t.tileId == 3 {
			if ballX < t.x{
				output = -1
			} else if ballX > t.x{
				output = 1
			} else if ballX == t.x{
				output = 0
			}
		}
	}

	return output
}
