package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	wordsToLookForArr := os.Args[1:]
	wordsToLookFor := strings.Join(wordsToLookForArr, " ")

	fmt.Println("Will look for", wordsToLookFor)

	start := time.Now()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, wordsToLookFor) {
			fmt.Println("Found: ", line)
			fmt.Println("Took", time.Since(start))
		}
	}
}
