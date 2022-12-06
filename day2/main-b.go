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
		"A X": 3,
		"B X": 1,
		"C X": 2,
		"A Y": 4,
		"B Y": 5,
		"C Y": 6,
		"A Z": 8,
		"B Z": 9,
		"C Z": 7,
	}

	for scanner.Scan() {
		score += score_map[scanner.Text()]
	}

	fmt.Println(score)
}
