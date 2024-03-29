package main

import(

	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main(){
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var fuel_total int
	for scanner.Scan(){
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		fuel_total += mass / 3 - 2
	}

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	fmt.Println(fuel_total)

}
