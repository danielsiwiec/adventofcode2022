package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
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

	max := 0
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
			if max < current {
				max = current
			}
			current = 0
		}
	}

	fmt.Println(max)
}
