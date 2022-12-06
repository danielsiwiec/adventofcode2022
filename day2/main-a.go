package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)

	f, _ := os.Open(fmt.Sprintf("%s/%s", path.Dir(filename), "./input"))
	defer f.Close()

	scanner := bufio.NewScanner(f)

	score := 0

	score_map := map[string]int{
		"A X": 4,
		"B X": 1,
		"C X": 7,
		"A Y": 8,
		"B Y": 5,
		"C Y": 2,
		"A Z": 3,
		"B Z": 9,
		"C Z": 6,
	}

	for scanner.Scan() {
		score += score_map[scanner.Text()]
	}

	fmt.Println(score)
}
