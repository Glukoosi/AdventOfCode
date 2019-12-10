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

	decodedImg := image[0]
	for _, l := range image[1:]{
		for ri,r := range l {
			for ni,n := range r {
				if n != 2 && decodedImg[ri][ni] == 2{
					decodedImg[ri][ni] = n
				}
			}
		}
	}

	for _, r := range decodedImg{
		fmt.Println(r)
	}
}



