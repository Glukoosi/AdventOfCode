package main

import(
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

type coordinate struct {
	x	int
	y	int
}

func main(){
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var paths [][]string
	for scanner.Scan(){ 
		var path []string
		for _, text := range strings.Split(scanner.Text(), ","){ path = append(path, text)
		}
		paths = append(paths, path)
	}

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	var result int

	var c1 coordinate
	var steps1 int
	for _, n1 := range paths[0]{

		dir1 := string(n1[0])
		number1, err := strconv.Atoi(n1[1:])
		if err != nil {
			log.Fatal(err)
		}

		for number1 > 0 {
			if dir1 == "U" {
				c1.y++
			} else if dir1 == "D" {
				c1.y--
			} else if dir1 == "R" {
				c1.x++
			} else if dir1 == "L" {
				c1.x--
			}
			number1--
			steps1++

			var c2 coordinate
			var steps2 int
			for _, n2 := range paths[1]{
				dir2 := string(n2[0])
				number2, err := strconv.Atoi(n2[1:])
				if err != nil {
					log.Fatal(err)
				}

				for number2 > 0 {
					if dir2 == "U" {
						c2.y++
					} else if dir2 == "D" {
						c2.y--
					} else if dir2 == "R" {
						c2.x++
					} else if dir2 == "L" {
						c2.x--
					}
					number2--
					steps2++

					if c1 == c2 {
						steps := steps1 + steps2
						if steps < result || result == 0{
							result = steps
						}
					}
				}
			}
		}
	}

	fmt.Println(result)
}
