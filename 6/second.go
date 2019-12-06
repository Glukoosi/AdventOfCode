package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var orbitMap [][]string
	for scanner.Scan() {
		var orbit []string
		for _, planet := range strings.Split(scanner.Text(), ")") {
			orbit = append(orbit, planet)
		}
		orbitMap = append(orbitMap, orbit)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var orbitMapSan [][]string
	var orbitMapYou [][]string
	for i := len(orbitMap) - 1; i >= 0; i-- {
		orbit := orbitMap[i]

		n := len(orbitMap) - 1
		for {
			if orbitMap[i][1] == "YOU" {
				if orbitMap[n][1] == orbit[0] {
					orbitMapYou = append(orbitMapYou, orbit)
					orbit = orbitMap[n]
					n = len(orbitMap)
				}
			} else if orbitMap[i][1] == "SAN" {
				if orbitMap[n][1] == orbit[0] {
					orbitMapSan = append(orbitMapSan, orbit)
					orbit = orbitMap[n]
					n = len(orbitMap)
				}
			}
			if n == 0 {
				break
			}
			n--
		}
	}

	var count int
	i := 1
	for {
		sanPop := orbitMapSan[len(orbitMapSan)-i]
		youPop := orbitMapYou[len(orbitMapYou)-i]
		if sanPop[0] == youPop[0] && sanPop[1] == youPop[1] {
			count++
		} else {
			break
		}
		i++

	}
	fmt.Println((len(orbitMapSan) - count - 1) + (len(orbitMapYou) - count - 1))
}
