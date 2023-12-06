package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const inputPath = "../input.txt"

func main() {
	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("Could not open input.txt: ", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	redReg, _ := regexp.Compile("([0-9]+) red")
	greenReg, _ := regexp.Compile("([0-9]+) green")
	blueReg, _ := regexp.Compile("([0-9]+) blue")

	total := 0
	for scanner.Scan() {
		text := scanner.Text()

		labelEnd := strings.Index(text, ":")
		hands := strings.Split(text[labelEnd+1:], ";")
		maxRed, maxBlue, maxGreen := 0, 0, 0
		for _, hand := range hands {

			redMatch := redReg.FindStringSubmatch(hand)
			if redMatch != nil {
				redNum, _ := strconv.Atoi(redMatch[1])
				if redNum > maxRed {
					maxRed = redNum
				}
			}

			blueMatch := blueReg.FindStringSubmatch(hand)
			if blueMatch != nil {
				blueNum, _ := strconv.Atoi(blueMatch[1])
				if blueNum > maxBlue {
					maxBlue = blueNum
				}
			}

			greenMatch := greenReg.FindStringSubmatch(hand)
			if greenMatch != nil {
				greenNum, _ := strconv.Atoi(greenMatch[1])
				if greenNum > maxGreen {
					maxGreen = greenNum
				}
			}
		}

		total += (maxRed * maxBlue * maxGreen)
	}

	if err = scanner.Err(); err != nil {
		fmt.Println("Encountered an error while scanning: ", err)
		os.Exit(1)
	}

	fmt.Println("Total: ", total)
}
