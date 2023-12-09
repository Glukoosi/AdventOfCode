package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var aMap [][]int
	for scanner.Scan() {
		var row []int
		for _, position := range scanner.Text() {
			if string(position) == "#" {
				row = append(row, 1)
			} else if string(position) == "." {
				row = append(row, 0)
			}
		}
		aMap = append(aMap, row)
		row = nil
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var mostAngles int
	var coordinate [2]int
	for y := 0; y < len(aMap); y++ {
		for x := 0; x < len(aMap[0]); x++ {
			if aMap[y][x] == 1 {
				angle := countSights(aMap, [2]int{y, x})
				if angle > mostAngles {
					mostAngles = angle
					coordinate = [2]int{y, x}
				}
			}
		}
	}

	fmt.Println(mostAngles, coordinate)

}

func countSights(aMap [][]int, coordinate [2]int) int {
	var angles []float64
	for y := 0; y < len(aMap); y++ {
		for x := 0; x < len(aMap[0]); x++ {
			if aMap[y][x] == 1 {
				angle := math.Atan2(float64(y-coordinate[0]), float64(x-coordinate[1]))
				if !float64InList(angle, angles) {
					angles = append(angles, angle)
				}
			}
		}
	}

	return len(angles)
}

func float64InList(float float64, list []float64) bool {
	for _, i := range list {
		if float == i {
			return true
		}
	}
	return false
}
