package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	var numbers []int
	for _, text := range strings.Split(scanner.Text(), "-") {
		number, err := strconv.Atoi(text)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var count int
	for i := numbers[0]; i <= numbers[1]; i++ {
		if check_password(i) {
			count++
		}
	}

	fmt.Println(count)

}

func check_password(password int) bool {

	var flag int
	var counter int

	check := 999
	for password > 0 {
		digit := password % 10
		if digit > check {
			return false
		} else if digit == check {
			counter += 1
		} else {
			if counter == 1 {
				flag = 1
			}
			counter = 0
		}
		check = digit

		password /= 10
	}
	if flag == 1 || counter == 1 {
		return true
	}
	return false
}
