package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"sort"
	"strconv"
)

func main() {
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		panic("No caller info")
	}

	f, err := os.Open(fmt.Sprintf("%s/%s", path.Dir(filename), "./input"))

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var calories []int
	current := 0

	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			int, err := strconv.Atoi(text)

			if err != nil {
				println(err.Error())
			} else {
				current += int
			}

		} else {
			calories = append(calories, current)
			current = 0
		}
	}

	sort.Ints(calories)
	result := 0
	for _, numb := range calories[len(calories)-3:] {
		result += numb
	}

	fmt.Println(result)
}
