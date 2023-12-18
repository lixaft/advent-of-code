package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Play struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id    int
	plays []Play
}

func parse(line string) (Game, error) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return Game{}, errors.New("parse error")
	}

	id, err := strconv.Atoi(strings.TrimPrefix(parts[0], "Game "))
	if err != nil {
		return Game{}, err
	}

	var plays []Play

	for _, p := range strings.Split(parts[1], ";") {
		red, green, blue := 0, 0, 0

		for _, t := range strings.Split(p, ",") {
			tokens := strings.Fields(t)
			if len(tokens) != 2 {
				return Game{}, errors.New("unknown syntax")
			}

			color := tokens[1]
			value, err := strconv.Atoi(tokens[0])
			if err != nil {
				return Game{}, err
			}

			switch color {
			case "red":
				red = value
			case "green":
				green = value
			case "blue":
				blue = value
			default:
				return Game{}, errors.New("unable to parse color")
			}

		}
		plays = append(plays, Play{
			red:   red,
			green: green,
			blue:  blue,
		})
	}

	return Game{
		id:    id,
		plays: plays,
	}, nil
}

func solve(s string) (int, error) {
	ret := 0

	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		line := scanner.Text()
		game, err := parse(line)
		if err != nil {
			return 0, err
		}

		for _, play := range game.plays {
			if play.red > 12 || play.green > 13 || play.blue > 14 {
				goto next
			}
		}

		ret += game.id

	next:
		continue
	}

	return ret, nil
}

func run() int {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "error: please provide input file")
		return 1
	}

	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}

	result, err := solve(string(input))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}

	fmt.Println(result)

	return 0
}

func main() {
	code := run()
	os.Exit(code)

}
