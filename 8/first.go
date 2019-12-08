package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	var image [][][]int
	var layer [][]int
	var row []int

	width := 25
	height := 6
	for i, n := range scanner.Text() {
		number := int(n - '0')
		row = append(row, number)
		if (i + 1) % width == 0 {
			layer = append(layer, row)
			row = nil
			if (i + 1) % height == 0 {
				image = append(image, layer)
				layer = nil
			}
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var answer int
	var lowestZeros int
	for _, l := range image{
		var zeros int
		var ones int
		var twos int
		for _, r := range l {
			zeros += intInSlice(r, 0)
			ones += intInSlice(r, 1)
			twos += intInSlice(r, 2)
		}
		if zeros < lowestZeros || lowestZeros == 0 {
			lowestZeros = zeros
			answer = ones * twos
		}
	}

	fmt.Println(answer)

}

func intInSlice(slice []int, integer int) int {
	var count int
	for _,n := range slice {
		if n == integer {
			count++
		}
	}
	return count
}


