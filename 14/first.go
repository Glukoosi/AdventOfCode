package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type reaction struct {
	chemical string
	amount int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var listOfReactions [][]reaction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var reactions []reaction
		for _, recipe := range strings.Split(scanner.Text(), ", ") {
			//for _, parsed := range strings.Split(recipe[-1], " => ") {

			//}

			parsed := strings.Split(recipe, " ")
			n, chemical := parsed[0], parsed[1]
			number, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			reactions = append(reactions, reaction{chemical: chemical, amount: number})

			if strings.Contains(recipe, " => "){
				n, chemical := parsed[3], parsed[4]
				number, err := strconv.Atoi(n)
				if err != nil {
					log.Fatal(err)
				}
				reactions = append(reactions, reaction{chemical: chemical, amount: number})
			}
			listOfReactions = append(listOfReactions, reactions)
			reactions = nil
		}
	}

	fmt.Println(findReaction(listOfReactions, "FUEL"))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func findReaction(listOfReactions [][]reaction, chemical string) []reaction {

	var returnReaction []reaction
	for _, r := range listOfReactions {
		if r[len(r)-1].chemical == chemical {
			returnReaction = r
		}

	}
	return returnReaction
}

