package main

import (
	"bufio"
	"fmt"
	"log"
	//"math"
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

	var moonsOrig []moon
	moonsOrig = make([]moon, len(moons))
	copy(moonsOrig, moons)

	var loopCount int
	var results [3]int
	for {

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

		loopCount++

		var countY int
		var countX int
		var countZ int
		for i4, _ := range moons {
			if moons[i4].yPos == moonsOrig[i4].yPos {
				countY++
			}
			if moons[i4].yVel == moonsOrig[i4].yVel {
				countY++
			}

			if moons[i4].xPos == moonsOrig[i4].xPos {
				countX++
			}
			if moons[i4].xVel == moonsOrig[i4].xVel {
				countX++
			}

			if moons[i4].zPos == moonsOrig[i4].zPos {
				countZ++
			}
			if moons[i4].zVel == moonsOrig[i4].zVel {
				countZ++
			}
		}

		if countY == 8 && results[0] == 0 {
			results[0] = loopCount
		}
		if countX == 8 && results[1] == 0 {
			results[1] = loopCount
		}
		if countZ == 8 && results[2] == 0 {
			results[2] = loopCount
		}

		if results[1] != 0 {
			break
		}

	}

	originalResult0 := results[0]
	originalResult1 := results[1]
	originalResult2 := results[2]

	for {
		if results[0] == results[1] {
			break
		}
		if results[0] < results[1] {
			results[0] = results[0] + originalResult0
		}
		if results[0] > results[1] {
			results[1] = results[1] + originalResult1
		}
	}

	for {
		if results[2] == results[1] {
			break
		}
		if results[2] < results[1] {
			results[2] = results[2] + originalResult2
		}
		if results[2] > results[1] {
			results[1] = results[1] + results[0]
		}
	}
	fmt.Println(results[2])
}
