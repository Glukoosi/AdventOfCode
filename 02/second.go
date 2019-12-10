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


	for noun:=0;noun<100;noun++{
		for verb:=0;verb<100;verb++{
			tempProgram := make([]int, len(program))
			copy(tempProgram, program)
			answer := run_program(noun, verb, tempProgram)
			if answer == 19690720 {
				fmt.Println(100 * noun + verb)
			}
		}
	}
}

func run_program(noun int, verb int, program []int) int{

	var n int
	program[1] = noun
	program[2] = verb
	for{
		if program[n] == 1 {
			program[program[n+3]] = program[program[n+1]] + program[program[n+2]]
		} else if program[n] == 2 {
			program[program[n+3]] = program[program[n+1]] * program[program[n+2]]
		} else if program[n] == 99 {
			return program[0]
		}

		n += 4
	}
}
