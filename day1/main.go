package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

//Rerad textfile to data
func getData(fileName string) ([]int, error) {
	//Get data from file
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fScanner := bufio.NewScanner(f)

	var data []int

	for fScanner.Scan() {
		lineValue, err := strconv.Atoi(fScanner.Text())
		if err != nil {
			return nil, fmt.Errorf("Cant parse %s: %v", fScanner.Text(), err)
		}
		data = append(data, lineValue)
	}

	if err = fScanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func calculateFuel(mass int) int {
	return mass/3 - 2
}

func totalPerModule(mass int) int {
	factor := mass/3 - 2
	if factor <= 0 {
		return 0
	}
	return factor + totalPerModule(factor)
}

func main() {
	//read data
	data, err := getData("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	//Part 1 & 2--------------------
	var fuel int = 0
	var totalFuel int = 0

	for _, c := range data {
		fuel += calculateFuel(c)
		totalFuel += totalPerModule(c)
	}

	//Print result
	fmt.Printf("Part 1: Total fuel is %d\n", fuel)
	fmt.Printf("Part 2: Total fuel including fuel mass: %d\n", totalFuel)
}
