package main

import "fmt"
import "os"
import "bufio"
import "regexp"
import "strconv"

func main() {
	ret := 0

	re := regexp.MustCompile(`[0-9]`)

	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, _ := reader.ReadLine()
		if line == nil {
			break
		}
		nums := re.FindAllString(string(line), -1)
		num := nums[0] + nums[len(nums)-1]
		i, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		ret += i
	}

	fmt.Println(ret)
}
