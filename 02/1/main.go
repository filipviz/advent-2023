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

type Color int

const (
	Red Color = iota
	Green
	Blue
)

func main() {
	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("Could not open input.txt: ", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	// I probably don't need this since games are sequential, but let's use it just in case.
	gameNumberRegex, _ := regexp.Compile("Game ([0-9]+):")

	redReg, _ := regexp.Compile("([0-9]+) red")
	greenReg, _ := regexp.Compile("([0-9]+) green")
	blueReg, _ := regexp.Compile("([0-9]+) blue")

	maxCubes := make(map[Color]int)
	maxCubes[Red] = 12
	maxCubes[Green] = 13
	maxCubes[Blue] = 14

	total := 0
scanLoop:
	for scanner.Scan() {
		text := scanner.Text()

		match := gameNumberRegex.FindStringSubmatch(text)

		gameNum, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Println("Could not parse game number: ", err)
			os.Exit(1)
		}

		hands := strings.Split(text[len(match[0]):], ";")
		for _, hand := range hands {
			redMatch := redReg.FindStringSubmatch(hand)
			if redMatch != nil {
				redNum, _ := strconv.Atoi(redMatch[1])
				if redNum > maxCubes[Red] {
					continue scanLoop
				}
			}

			blueMatch := blueReg.FindStringSubmatch(hand)
			if blueMatch != nil {
				blueNum, _ := strconv.Atoi(blueMatch[1])
				if blueNum > maxCubes[Blue] {
					continue scanLoop
				}
			}

			greenMatch := greenReg.FindStringSubmatch(hand)
			if greenMatch != nil {
				greenNum, _ := strconv.Atoi(greenMatch[1])
				if greenNum > maxCubes[Green] {
					continue scanLoop
				}
			}
		}

		total += gameNum
	}

	if err = scanner.Err(); err != nil {
		fmt.Println("Encountered an error while scanning: ", err)
		os.Exit(1)
	}

	fmt.Println("Total: ", total)
}
