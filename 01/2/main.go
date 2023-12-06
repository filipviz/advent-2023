package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const fileName = "../input.txt"

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Could not open %s: %v\n", fileName, err)
		os.Exit(1)
	}
	defer file.Close()

	numStrings := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", //"0",
		"one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine", // No zero according to puzzle
	}

	strToDigit := make(map[string]string)
	strToDigit["0"] = "0"
	strToDigit["1"] = "1"
	strToDigit["2"] = "2"
	strToDigit["3"] = "3"
	strToDigit["4"] = "4"
	strToDigit["5"] = "5"
	strToDigit["6"] = "6"
	strToDigit["7"] = "7"
	strToDigit["8"] = "8"
	strToDigit["9"] = "9"
	strToDigit["one"] = "1"
	strToDigit["two"] = "2"
	strToDigit["three"] = "3"
	strToDigit["four"] = "4"
	strToDigit["five"] = "5"
	strToDigit["six"] = "6"
	strToDigit["seven"] = "7"
	strToDigit["eight"] = "8"
	strToDigit["nine"] = "9"

	firstBytes := make(map[byte][]string)
	lastBytes := make(map[byte][]string)
	for _, str := range numStrings {
		firstByte := str[0]
		_, exists := firstBytes[firstByte]
		if exists {
			firstBytes[firstByte] = append(firstBytes[str[0]], str)
		} else {
			firstBytes[firstByte] = []string{str}
		}

		lastByte := str[len(str)-1]
		_, exists = lastBytes[lastByte]
		if exists {
			lastBytes[lastByte] = append(lastBytes[lastByte], str)
		} else {
			lastBytes[lastByte] = []string{str}
		}
	}

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()

		var firstNum string
	firstLoop:
		for i := 0; i < len(t); i++ {
			foundStrs, exists := firstBytes[t[i]]
			if exists {
				for _, s := range foundStrs {
					if i+len(s)-1 < len(t) && s == t[i:i+len(s)] {
						firstNum = strToDigit[s]
						fmt.Printf("Found first match at index %d of %s: %s\n", i, t, firstNum)
						break firstLoop
					}
				}
			}
		}

		var lastNum string
	secondLoop:
		for i := len(t) - 1; i >= 0; i-- {
			foundStrs, exists := lastBytes[t[i]]
			if exists {
				for _, s := range foundStrs {
					if i-len(s)+1 >= 0 && s == t[i-len(s)+1:i+1] {
						lastNum = strToDigit[s]
						fmt.Printf("Found last match at index %d of %s: %s\n", i, t, lastNum)
						break secondLoop
					}
				}
			}
		}

		value, err := strconv.Atoi(firstNum + lastNum)
		if err != nil {
			fmt.Printf("Could not convert into string: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Combined to %d.\n", value)

		total += value
		fmt.Printf("New total: %d\n", total)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed while scanning file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Total: %d\n", total)
}
