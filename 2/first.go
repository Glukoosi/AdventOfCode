package main

import(
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func main(){
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	var program []int
	for _, n := range strings.Split(scanner.Text(), ","){
		number, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		program = append(program, number)
	}

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	var n int
	program[1] = 12
	program[2] = 2

	for{
		if program[n] == 1 {
			program[program[n+3]] = program[program[n+1]] + program[program[n+2]]
		} else if program[n] == 2 {
			program[program[n+3]] = program[program[n+1]] * program[program[n+2]]
		} else if program[n] == 99 {
			fmt.Println(program[0])
			break
		}

		n += 4
	}

}

