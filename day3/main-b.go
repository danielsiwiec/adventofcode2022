package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)

	f, _ := os.Open(fmt.Sprintf("%s/%s", path.Dir(filename), "./input"))
	defer f.Close()

	scanner := bufio.NewScanner(f)
	score := int32(0)

	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()

		badge := findBadge(line1, line2, line3)
		score += scoreBadge(badge)
	}

	fmt.Println(score)
}

func scoreBadge(char int32) int32 {
	if char >= 97 && char <= 122 {
		return char - 96
	} else {
		return char - 38
	}
}

func findBadge(first string, second string, third string) int32 {
	for _, char := range first {
		containsSecond := strings.Contains(second, string(char))
		containsThird := strings.Contains(third, string(char))
		if containsSecond && containsThird {
			return char
		}
	}
	return 0
}