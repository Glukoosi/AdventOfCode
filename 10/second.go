package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

type Asteroid struct {
	Y     int
	X     int
	Angle float64
}

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

	asteroids := getAngles(aMap, [2]int{16, 8})

	sort.Slice(asteroids, func(i, j int) bool {
		return asteroids[i].Angle < asteroids[j].Angle
	})

	var count int
	var lastAsteroidAngle float64

	asteroids = append(asteroids[101:], asteroids[:101]...)
	for _, asteroid := range asteroids {
		if lastAsteroidAngle != asteroid.Angle {
			count++
		}
		lastAsteroidAngle = asteroid.Angle

		if count == 200 {
			fmt.Println(asteroid.X*100 + asteroid.Y)
		}

	}

}

func getAngles(aMap [][]int, coordinate [2]int) []Asteroid {
	var asteroids []Asteroid
	for y := 0; y < len(aMap); y++ {
		for x := 0; x < len(aMap[0]); x++ {
			if aMap[y][x] == 1 {
				angle := math.Atan2(float64(y-coordinate[0]), float64(x-coordinate[1]))
				asteroids = append(asteroids, Asteroid{y, x, angle})
			}
		}
	}

	return asteroids
}

func float64InList(float float64, list []float64) bool {
	for _, i := range list {
		if float == i {
			return true
		}
	}
	return false
}
