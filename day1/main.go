package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Errorf("error opening the file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	fuelNeeded := 0
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("error reading mass out of file.")
		}
		fuelNeeded += calculateFuel(mass)
	}
	fmt.Println(fuelNeeded)
}

func calculateFuel(mass int) int {
	fuel := (mass / 3) - 2
	return fuel
}
