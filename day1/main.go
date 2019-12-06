package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getData(fileName string) []int {
	//Get data from file and append it to data array
	f, _ := os.Open(fileName)
	defer f.Close()

	fScanner := bufio.NewScanner(f)

	data := []int{}

	for fScanner.Scan() {
		lineValue, _ := strconv.Atoi(fScanner.Text())
		data = append(data, lineValue)
	}
	return data
}

func calculateFuel(data float64) int {
	return int(math.Floor(data/3) - 2)
}

func totalPerModule(data int) int {
	factor := float64(data / 3)
	if (factor - 2) <= 0 {
		return 0
	}
	return int(math.Floor(factor)-2) + totalPerModule(int(math.Floor(factor)-2))
}

func main() {
	//read data
	data := getData("input.txt")

	//Part 1 & 2--------------------
	var fuel int = 0
	var totalFuel int = 0
	for _, row := range data {
		fuel += calculateFuel(float64(row))
		totalFuel += totalPerModule(row)
	}

	//Print result
	fmt.Printf("Part 1: Total fuel is %d\n", fuel)
	fmt.Printf("Part 2: Total fuel including fuel mass: %d\n", totalFuel)
}
