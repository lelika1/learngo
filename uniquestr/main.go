package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	uniqueStr := make(map[string]bool)
	for scanner.Scan() {
		str := scanner.Text()
		if _, ok := uniqueStr[str]; !ok {
			uniqueStr[str] = true
			fmt.Println(str)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
