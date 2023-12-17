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

func run() error {
	if len(os.Args) < 2 {
		return errors.New("Please provide input file")
	}

	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		return err
	}

	result, err := solve(string(input))
	if err != nil {
		return err
	}
	fmt.Println(result)

	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
