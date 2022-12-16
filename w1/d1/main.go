package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var (
		maximumCalories int
		caloriesPerElf  int
		calorieTracker  []int
		err             error
		f               *os.File
	)

	f, err = os.Open("inputs.txt")
	if err != nil {
		fmt.Printf("file open failed: %v\n", err)
		os.Exit(1)
	}

	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		scannedNumber, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			calorieTracker = append(calorieTracker, caloriesPerElf)
			caloriesPerElf = 0
		}
		caloriesPerElf += scannedNumber
	}
	if fileScanner.Err() != nil {
		fmt.Printf("file scanner failed: %v\n", fileScanner.Err())
	}

	// EOF
	calorieTracker = append(calorieTracker, caloriesPerElf)
	maximumCalories = findMax(calorieTracker)
	fmt.Printf("Fattest ELF: %d\n", maximumCalories)
	fmt.Printf("Top 3 total calories: %d\n", countTopThreeEaters(calorieTracker))
}

func findMax(l []int) int {
	var m, t int
	for _, v := range l {
		if v > t {
			t = v
			m = t
		}
	}
	return m
}

func countTopThreeEaters(l []int) int {
	var topThreeCalories, length int
	length = len(l)

	sort.Ints(l)

	for x := 1; x <= 3; x++ {
		topThreeCalories += l[length-x]
	}
	return topThreeCalories
}
