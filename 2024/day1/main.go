package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInput() ([]int, []int) {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	str := strings.TrimSpace(string(dat))
	lines := strings.Split(str, "\n")

	left := make([]int, 0, len(lines))
	right := make([]int, 0, len(lines))

	for _, line := range lines {
		unsorted := strings.Fields(line)
		l, err := strconv.Atoi(unsorted[0])
		if err != nil {
			panic(err)
		}

		r, err := strconv.Atoi(unsorted[1])
		if err != nil {
			panic(err)
		}

		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}

func main() {
	left, right := getInput()

	sort.Ints(left)
	sort.Ints(right)

	total := 0

	for i, v := range left {
		pairSum := v - right[i]
		total += int(math.Abs(float64(pairSum)))
	}

	fmt.Println(total)
}
