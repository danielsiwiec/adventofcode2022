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
		score += processLine(scanner.Text())
	}

	fmt.Println(score)
}

func processLine(text string) int32 {
	first, second := split(text)
	common := common(first, second)

	if common >= 97 && common <= 122 {
		return common - 96
	} else {
		return common - 38
	}
}

func common(first string, second string) int32 {
	for _, char := range first {
		contains := strings.Contains(second, string(char))
		if contains {
			return char
		}
	}
	return 0
}

func split(text string) (string, string) {
	length := len(text)
	first := text[0:(length/2)]
	second := text[(length/2): length]
	return first, second
}
