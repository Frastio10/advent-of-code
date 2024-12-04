package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() [][]int {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	str := strings.TrimSpace(string(dat))
	lines := strings.Split(str, "\n")

	result := make([][]int, 0, len(lines))

	for _, line := range lines {
		l := strings.Fields(line)

		nums := make([]int, 0, len(l))

		for _, num := range l {
			convertedNumber, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			nums = append(nums, convertedNumber)
		}

		result = append(result, nums)

	}

	return result

}

func isAscending(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] || arr[i]-arr[i-1] < 1 || arr[i]-arr[i-1] > 3 {
			return false
		}
	}
	return true
}

func isDescending(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] >= arr[i-1] || arr[i-1]-arr[i] < 1 || arr[i-1]-arr[i] > 3 {
			return false
		}
	}
	return true
}

func isSafe(report []int) bool {
	return isAscending(report) || isDescending(report)
}

func checkReportValidity(report []int) bool {
	if isSafe(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		// bruteforce lol
		level := make([]int, 0, len(report)-1)
		level = append(level, report[:i]...)
		level = append(level, report[i+1:]...)

		if isSafe(level) {
			return true
		}

	}

	return false
}

func main() {
	input := getInput()
	total := 0

	for _, report := range input {
		if isValid := checkReportValidity(report); isValid {
			total += 1
		}
	}

	fmt.Println(total)
}
