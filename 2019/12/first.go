package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type moon struct {
	yPos int
	xPos int
	zPos int

	yVel int
	xVel int
	zVel int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var moons []moon
	for scanner.Scan() {
		var coordinates []int
		for _, n := range strings.Split(scanner.Text(), ",") {
			number, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			coordinates = append(coordinates, number)
		}
		moons = append(moons, moon{yPos: coordinates[1],
			xPos: coordinates[0],
			zPos: coordinates[2]})

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 1000; i++ {

		for i2, m := range moons {
			for _, m2 := range moons {
				if m != m2 {
					if m.yPos > m2.yPos {
						moons[i2].yVel--
					} else if m.yPos < m2.yPos {
						moons[i2].yVel++
					}
					if m.xPos > m2.xPos {
						moons[i2].xVel--
					} else if m.xPos < m2.xPos {
						moons[i2].xVel++
					}
					if m.zPos > m2.zPos {
						moons[i2].zVel--
					} else if m.zPos < m2.zPos {
						moons[i2].zVel++
					}
				}
			}
		}

		for i3, m := range moons {
			moons[i3].yPos += m.yVel
			moons[i3].xPos += m.xVel
			moons[i3].zPos += m.zVel
		}
	}

	var answer float64
	for _, m := range moons {
		pos := math.Abs(float64(m.yPos)) + math.Abs(float64(m.xPos)) + math.Abs(float64(m.zPos))
		vel := math.Abs(float64(m.yVel)) + math.Abs(float64(m.xVel)) + math.Abs(float64(m.zVel))
		total := pos * vel
		answer += total
	}

	fmt.Println(answer)
}
