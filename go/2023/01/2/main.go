package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const REGEX = `(\d)`

func solve(s string) (int, error) {
	ret := 0

	translate := map[string]string{
		"one":   "one1one",
		"two":   "two2two",
		"three": "three3three",
		"four":  "four4four",
		"five":  "five5five",
		"six":   "six6six",
		"seven": "seven7seven",
		"eight": "eight8eight",
		"nine":  "nine9nine",
	}

	for k, v := range translate {
		s = strings.ReplaceAll(s, k, v)
	}

	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		line := scanner.Text()

		re := regexp.MustCompile(REGEX)
		nums := re.FindAllString(line, -1)

		if len(nums) < 1 {
			return 0, errors.New("Not enough numbers")
		}

		num, _ := strconv.Atoi(nums[0] + nums[len(nums)-1])
		ret += num
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
