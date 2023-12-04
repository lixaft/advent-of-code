package main

import "fmt"
import "os"
import "bufio"
import "regexp"


func main() {
	re := regexp.MustCompile(`[0-9]`)

	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, _ := reader.ReadLine()
		if line == nil {
			break
		}
		nums := re.FindAll(line, -1)
	}

}
