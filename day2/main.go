package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getGames() []string {
	return []string{
		"19 blue, 12 red; 19 blue, 2 green, 1 red; 13 red, 11 blue",
		"1 green, 1 blue, 1 red; 11 red, 3 blue; 1 blue, 18 red; 9 red, 1 green; 2 blue, 11 red, 1 green; 1 green, 2 blue, 10 red",
		"3 blue, 2 red, 6 green; 4 blue, 6 green, 1 red; 11 green, 12 blue; 2 red, 6 green, 4 blue; 4 green",
		"10 red, 5 green, 5 blue; 3 red, 3 blue, 6 green; 2 blue, 9 red, 6 green; 8 green, 10 red, 4 blue; 9 red, 2 green, 3 blue; 1 blue, 5 red, 15 green",
		"11 green, 7 blue; 5 green, 5 red, 1 blue; 1 green, 1 red, 4 blue; 1 red, 1 blue, 4 green; 4 blue, 1 red, 10 green; 5 red, 6 green",
		"1 green, 1 red, 11 blue; 1 blue, 2 green; 1 red, 5 green, 9 blue; 7 blue; 1 red, 2 green, 9 blue; 12 blue, 1 red, 2 green",
		"1 blue, 10 red, 7 green; 14 blue, 10 green; 12 red, 2 green; 16 red, 13 blue, 1 green; 12 green, 10 red, 3 blue; 9 red, 19 blue, 11 green",
		"3 blue, 1 green, 3 red; 4 blue, 10 red, 6 green; 1 green, 10 red, 9 blue; 9 blue, 7 red, 8 green; 8 green, 12 red, 8 blue; 6 blue, 1 green, 13 red",
		"10 green, 2 blue, 11 red; 2 green, 2 red; 6 blue, 8 red, 13 green",
		"8 red, 3 blue, 5 green; 5 green, 7 blue, 1 red; 3 red, 10 blue, 6 green; 2 red, 6 green, 7 blue; 3 blue, 11 red, 4 green; 8 red, 8 blue, 4 green",
		"14 green, 9 red; 3 blue, 6 green, 8 red; 14 green",
		"10 red, 5 blue, 1 green; 4 blue, 8 red; 5 blue, 1 green, 6 red; 14 red, 4 blue; 1 green, 11 red, 3 blue",
		"1 blue, 16 green, 1 red; 6 red, 2 blue, 5 green; 2 blue, 12 red, 10 green; 3 red, 4 blue, 13 green; 14 red, 4 blue, 12 green; 7 red, 2 green",
		"17 red, 11 blue, 3 green; 16 red, 3 blue, 8 green; 3 green, 9 red, 13 blue; 4 green, 15 red, 14 blue",
		"7 blue, 2 red, 2 green; 1 green, 5 red, 6 blue; 3 green, 6 red, 2 blue",
		"3 red, 3 green; 6 green, 4 red, 3 blue; 3 red, 4 blue; 4 blue, 2 red, 4 green",
		"6 red, 1 blue, 5 green; 3 red, 1 green, 12 blue; 13 green, 1 blue; 5 blue, 7 green, 6 red; 5 blue, 14 green, 2 red; 4 green, 6 red, 10 blue",
		"4 red, 8 blue; 8 blue, 4 red; 12 blue, 1 green",
		"1 blue, 15 green, 9 red; 1 red, 3 green; 4 blue, 2 green, 1 red",
		"7 blue, 4 green, 12 red; 1 red, 9 green, 8 blue; 4 blue, 2 green; 13 green, 8 blue; 3 red, 4 green, 1 blue; 6 green, 7 red, 3 blue",
		"9 green, 4 blue, 8 red; 5 blue; 7 red, 8 blue, 1 green",
		"3 green, 4 red; 6 red, 3 green; 4 red, 1 blue, 1 green; 11 red, 3 green, 1 blue; 7 red, 1 blue",
		"3 blue, 4 green; 3 green, 1 red; 1 red, 2 blue, 4 green",
		"2 blue, 3 green; 9 red, 4 green; 2 blue, 9 red; 2 green, 10 red, 1 blue; 1 blue, 1 red, 5 green",
		"8 green, 4 blue; 9 blue, 7 red; 5 green, 15 blue, 11 red; 11 green, 14 red, 10 blue",
		"3 blue; 2 red, 1 green; 2 red, 3 blue; 10 blue, 1 red, 3 green; 1 green, 2 red; 1 green, 6 blue",
		"1 green, 6 blue; 2 green, 1 red, 6 blue; 1 red, 2 blue, 1 green",
		"8 blue, 1 red, 5 green; 1 red; 3 green, 4 red, 2 blue; 4 green, 2 red, 4 blue; 5 blue, 3 red, 7 green",
		"2 green, 4 blue; 7 blue, 4 red, 10 green; 7 blue, 9 green; 14 green, 7 red, 5 blue",
		"19 green, 3 red; 19 green; 1 blue, 14 green; 2 blue, 5 green; 3 red, 19 green",
		"3 red, 1 green, 4 blue; 10 blue; 3 red, 4 green, 5 blue; 10 blue, 1 red, 6 green",
		"19 red, 1 green, 2 blue; 1 blue, 6 green, 13 red; 10 green, 9 red; 11 red, 2 blue, 6 green; 8 green, 5 red",
		"2 red, 8 blue, 2 green; 1 red, 3 green; 9 red, 9 blue, 1 green; 6 red, 1 green; 9 blue, 1 green, 8 red; 5 green, 10 red, 8 blue",
		"1 red, 6 blue, 2 green; 7 red; 14 red, 13 blue; 13 red, 12 blue; 1 green, 9 red, 13 blue; 2 green, 15 blue",
		"8 blue, 2 red, 3 green; 2 green, 2 red; 3 red, 6 blue, 2 green; 2 green, 6 blue; 1 green, 5 blue, 4 red; 3 green, 6 blue",
		"3 red, 5 blue, 10 green; 1 red, 1 green, 7 blue; 2 blue, 2 green, 1 red",
		"8 red, 7 green; 5 green, 1 blue, 6 red; 7 red, 6 blue, 11 green",
		"4 green, 10 red, 9 blue; 12 green, 2 blue, 2 red; 6 red, 6 blue, 9 green; 1 blue, 1 green, 6 red; 3 blue, 1 red, 5 green; 5 blue, 2 red, 12 green",
		"1 blue, 2 red; 7 blue, 2 green, 1 red; 7 blue, 11 green, 3 red; 8 blue, 13 green, 1 red; 6 green, 6 blue, 3 red",
		"8 green, 5 blue; 5 green, 1 blue, 10 red; 9 green, 3 blue; 3 green, 7 red; 2 green, 3 blue, 5 red",
		"7 green, 8 red; 3 blue, 15 green, 7 red; 2 red, 2 green, 4 blue; 10 green, 4 red, 5 blue; 3 red, 8 blue, 9 green; 7 red, 8 green",
		"6 blue, 12 green; 3 red, 1 green; 1 red, 12 green, 3 blue; 10 red, 9 green; 9 red, 4 green, 5 blue",
		"11 red, 6 green; 2 blue, 11 red; 3 red, 1 blue; 3 green, 11 red, 2 blue; 4 red, 5 green, 1 blue; 8 green, 2 blue, 17 red",
		"2 green, 9 blue, 3 red; 7 blue, 1 green, 4 red; 1 green",
		"1 green, 10 red; 5 red, 10 green, 1 blue; 11 red, 3 green, 2 blue; 2 blue, 3 green, 4 red; 7 green, 3 red, 2 blue; 1 blue, 10 red",
		"1 green, 4 blue, 7 red; 13 blue, 2 green, 9 red; 7 blue, 3 red, 1 green",
		"4 blue; 2 green, 2 red, 1 blue; 1 green, 1 red, 4 blue; 1 green, 2 red, 2 blue; 2 blue, 2 red",
		"5 green, 10 red; 7 red, 5 green; 1 green, 11 red; 12 red, 11 green; 11 red, 1 blue, 1 green",
		"2 green, 1 red, 1 blue; 1 blue, 2 red; 2 green, 1 red, 2 blue; 1 blue, 1 red, 1 green",
		"5 green, 2 blue; 4 green, 4 blue, 3 red; 1 red, 7 green, 3 blue",
		"9 green, 1 red, 2 blue; 7 red, 3 blue, 6 green; 5 green, 4 blue, 5 red",
		"2 green, 4 blue, 1 red; 2 blue, 2 red, 13 green; 8 blue, 3 green; 3 green, 4 blue, 2 red; 2 green",
		"3 red; 4 blue, 4 red; 2 blue, 2 red; 6 blue, 1 red, 2 green; 1 red, 1 green, 6 blue; 2 blue, 4 red",
		"3 blue, 3 green, 18 red; 4 blue, 18 red, 3 green; 7 blue, 4 green",
		"1 green, 2 red, 3 blue; 1 red, 4 blue, 1 green; 3 blue, 2 red; 2 blue, 1 green; 3 blue, 2 red; 1 blue, 1 green, 1 red",
		"12 green, 2 red, 1 blue; 11 green, 16 red, 13 blue; 7 red, 5 blue, 12 green; 4 blue, 16 red; 5 red, 1 blue, 3 green",
		"5 green, 17 blue, 11 red; 6 blue, 1 green; 1 green, 5 blue, 8 red; 9 green, 11 red, 1 blue; 9 green, 11 blue, 7 red; 8 green, 4 blue",
		"5 red, 10 blue, 6 green; 5 green, 11 blue, 5 red; 9 green; 4 red, 2 green",
		"2 red, 6 blue, 1 green; 1 green, 12 blue; 2 red",
		"6 blue, 10 green, 9 red; 8 red, 19 blue, 2 green; 16 red, 10 green, 12 blue; 13 red, 12 blue, 6 green",
		"12 green, 1 red, 3 blue; 3 red, 4 blue, 19 green; 1 blue, 7 green",
		"7 red, 6 blue, 8 green; 10 blue, 3 green, 17 red; 13 blue, 3 red, 10 green; 13 red, 5 blue, 9 green; 12 blue, 4 red; 10 red, 4 green",
		"19 green, 4 red; 5 blue, 4 red, 1 green; 4 red, 2 blue, 15 green; 5 green, 4 red, 5 blue",
		"6 red, 3 green; 6 green, 3 red, 3 blue; 3 blue, 8 red, 5 green; 3 blue, 7 red, 1 green; 1 blue, 6 red, 6 green",
		"1 green, 9 blue; 6 blue, 4 green, 6 red; 6 blue, 5 green; 3 red, 1 blue, 4 green",
		"1 blue, 2 red; 2 green, 1 blue; 2 red, 1 blue, 1 green; 1 blue, 1 green",
		"16 blue, 1 green; 1 blue, 2 green, 2 red; 1 red, 9 blue; 12 blue, 4 green, 1 red; 6 green, 11 blue, 3 red",
		"6 blue, 2 red, 1 green; 2 blue, 2 green; 1 green, 7 red, 15 blue; 14 blue, 12 green, 3 red; 13 green, 10 red, 6 blue; 2 green, 5 blue, 1 red",
		"2 red, 1 blue, 2 green; 1 blue, 7 green, 1 red; 3 blue, 1 red, 7 green; 2 red, 1 blue, 11 green",
		"2 green, 9 red, 3 blue; 12 blue, 1 green, 13 red; 6 red, 1 green, 5 blue; 1 red, 17 blue",
		"7 red, 5 green, 6 blue; 5 blue, 5 green; 7 green, 4 blue; 2 green, 4 blue, 8 red; 10 red, 8 green; 3 blue, 13 red, 7 green",
		"13 red, 17 green; 9 red, 20 green, 3 blue; 1 green, 3 blue, 8 red",
		"1 blue, 7 red, 2 green; 2 green, 1 blue, 8 red; 1 blue, 2 red; 4 red, 7 green; 4 red, 5 green; 3 green, 7 red",
		"2 green, 14 blue; 1 red, 1 blue, 7 green; 1 red, 8 green, 11 blue; 4 green, 12 blue; 1 green, 5 blue",
		"12 blue, 1 red; 1 red, 7 blue, 4 green; 4 blue, 6 green; 4 green, 3 blue, 1 red",
		"7 green, 5 red, 6 blue; 18 red, 1 green; 14 green, 4 red, 15 blue; 4 blue, 6 red",
		"2 blue, 2 green, 2 red; 2 blue, 1 red, 1 green; 2 green, 1 red; 6 blue, 4 green; 1 red, 1 blue, 6 green",
		"5 red, 16 blue, 12 green; 11 blue, 3 red, 2 green; 13 blue, 4 red",
		"9 red, 11 green, 6 blue; 1 red, 3 green; 7 blue, 7 red, 11 green; 8 red, 9 blue, 11 green; 7 red, 11 green, 4 blue",
		"7 green, 5 red, 2 blue; 1 blue, 7 green, 1 red; 2 red, 2 blue; 1 red, 4 blue, 12 green; 4 green, 2 blue",
		"5 blue, 2 green, 12 red; 2 green, 1 blue, 5 red; 3 blue, 13 red, 3 green; 3 green, 9 blue, 3 red; 10 blue, 4 red, 3 green",
		"11 blue, 1 red, 9 green; 11 green, 1 blue, 12 red; 13 red, 6 blue, 19 green",
		"6 red, 5 blue, 16 green; 4 green, 17 blue, 9 red; 15 red, 2 green, 9 blue",
		"19 green, 11 blue, 3 red; 1 blue, 18 green, 6 red; 17 blue, 5 green, 4 red; 18 blue, 7 green, 3 red",
		"3 green, 15 blue; 12 blue; 2 green, 1 red; 1 red, 9 blue, 1 green; 12 blue, 3 red, 1 green",
		"3 green, 4 blue, 5 red; 9 red, 4 green, 1 blue; 6 green, 1 blue, 8 red; 3 green, 2 blue, 5 red",
		"2 red, 8 blue, 5 green; 3 red, 5 blue, 10 green; 2 red, 3 green",
		"16 green, 13 red; 7 green, 1 blue, 2 red; 7 red, 12 green; 5 red, 7 green, 2 blue; 2 blue, 10 green, 7 red; 8 red, 16 green",
		"1 blue, 8 red; 2 green, 10 red, 12 blue; 13 green, 14 blue; 10 blue, 15 red, 13 green; 2 green, 5 red, 13 blue",
		"16 blue, 7 red, 4 green; 4 green, 6 red, 11 blue; 2 red, 8 blue, 2 green; 5 green, 8 red, 10 blue; 4 red, 2 green, 7 blue; 4 green, 5 blue, 5 red",
		"4 red, 4 green, 1 blue; 3 blue, 2 green; 6 blue, 4 green, 5 red; 2 red, 6 blue, 4 green; 6 blue, 1 green",
		"1 red, 3 green; 3 blue, 6 green; 5 blue, 1 red, 11 green; 1 red; 3 green, 13 blue",
		"1 red, 14 blue, 6 green; 10 blue, 6 red; 9 green, 15 red, 17 blue; 9 red, 1 green, 9 blue",
		"3 red, 14 green; 3 blue, 15 green, 3 red; 2 red, 15 green",
		"4 blue, 13 red; 5 blue, 1 green, 11 red; 3 green, 3 blue, 10 red; 13 red, 6 blue; 2 green, 5 blue; 3 green, 11 red",
		"7 blue, 1 green; 1 green, 4 blue; 1 green, 2 red, 5 blue; 1 red, 2 blue, 1 green; 1 blue",
		"15 green, 9 blue; 14 blue, 14 red, 2 green; 18 red, 12 blue, 2 green",
		"1 green, 9 red; 1 red, 2 green, 7 blue; 8 red, 1 blue; 6 red, 2 green; 1 green, 6 blue",
		"1 green, 2 red, 6 blue; 6 red, 1 green, 5 blue; 11 blue, 6 red; 11 red, 1 green; 1 green, 11 red, 9 blue",
		"12 green, 8 blue, 2 red; 7 blue, 14 red, 8 green; 14 red, 1 blue, 4 green",
	}
}

type Game struct {
	maxRed, maxGreen, maxBlue, power int
}

func ParseGame(gameString string) *Game {
	// example game: "12 green, 8 blue, 2 red; 7 blue, 14 red, 8 green; 14 red, 1 blue, 4 green"
	// we need to find the max of each color.  if a color is not present, it is 0

	maximums := make(map[string]int)
	maximums["red"] = 0
	maximums["green"] = 0
	maximums["blue"] = 0

	// split the string into the three parts
	parts := strings.Split(gameString, ";")
	// for each part, split it into the colors
	for _, part := range parts {
		// fmt.Println("part", part)
		colors := strings.Split(part, ",")
		for _, color := range colors {
			// fmt.Printf("color: %v\n", color)
			// for each color, split it into the number and the color
			// but we have to strip all leading and trailing whitespace first
			color = strings.TrimSpace(color)
			colorParts := strings.Split(color, " ")

			// for _, colorPart := range colorParts {
			// 	fmt.Printf("colorPart: %v\n", colorPart)
			// }
			number, _ := strconv.Atoi(colorParts[0])
			color := colorParts[1]
			if number > maximums[color] {
				maximums[color] = number
			}
		}

	}
	// get the maximum of these 2 options:
	// 1. the max of each color
	// 2. the product of the max of each color

	return &Game{
		maxRed:   maximums["red"],
		maxBlue:  maximums["blue"],
		maxGreen: maximums["green"],
		power:    int(math.Max(1, float64(maximums["red"])) * math.Max(float64(maximums["blue"]), 1) * math.Max(float64(maximums["green"]), 1)),
	}

}

func main() {

	minRed := 12
	minGreen := 13
	minBlue := 14

	games := getGames()

	sum := 0
	powerSum := 0
	for n, gameString := range games {
		game := ParseGame(gameString)
		powerSum += game.power

		if game.maxRed <= minRed && game.maxGreen <= minGreen && game.maxBlue <= minBlue {
			fmt.Println("\n" + strconv.Itoa(n+1))
			fmt.Printf("YES %v\n", gameString)
			sum += n + 1
		} else {
			fmt.Printf("NO %v\n", gameString)
		}
		// break
	}

	fmt.Printf("Part 1 answer: %v\n", sum)
	fmt.Printf("Part 2 answer: %v\n", powerSum)

}
