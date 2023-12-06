package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const nums = "1234567890"
const fileName = "../input.txt"

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Could not open %s: %v\n", fileName, err)
		os.Exit(1)
	}
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()

		firstNum := strings.IndexAny(t, nums)
		lastNum := strings.LastIndexAny(t, nums)

		if firstNum == -1 {
			fmt.Printf("Could not find a number in line: %s\n", t)
			os.Exit(1)
		}

		value, err := strconv.Atoi(string(t[firstNum]) + string(t[lastNum]))
		if err != nil {
			fmt.Printf("Could not convert into string: %v\n", err)
			os.Exit(1)
		}

		total += value
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed while scanning file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Total: %d\n", total)
}
