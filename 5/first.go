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
			number, err := strconv.Atoi(input("Input: "))
			if err != nil {
				log.Fatal(err)
			}
			program[params[0]] = number
			n += 2
		} else if opcode == 4 {
			fmt.Println(program[params[0]])
			n += 2
		}
	}

}

func input(inputPrint string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(inputPrint)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	text = strings.TrimSuffix(text, "\n")
	return text

}
