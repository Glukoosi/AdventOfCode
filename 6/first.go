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

	var orbitCount int
	for i := len(orbitMap) - 1; i >= 0; i-- {
		orbit := orbitMap[i]

		n := len(orbitMap) - 1
		for {
			if orbitMap[n][1] == orbit[0] {
				orbitCount++
				orbit = orbitMap[n]
				n = len(orbitMap)
			}
			if n == 0 {
				break
			}
			n--
		}
		orbitCount++
	}
	fmt.Println(orbitCount)
}
