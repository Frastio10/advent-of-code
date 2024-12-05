package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func getInput() string {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	str := strings.TrimSpace(string(dat))

	return str

}

// i dont like regex.
func main() {
	input := getInput()
	// cmd := regexp.MustCompile("(?:mul\\(\\d+,\\d+\\))") //  part 1
	cmd := regexp.MustCompile("(?:mul\\(\\d+,\\d+\\)|do\\(\\)|don't\\(\\))") //  part 2

	total := 0
	shouldDo := true

	for _, token := range cmd.FindAllString(input, -1) {
		if token[0:3] == "mul" {
			var n1, n2 int
			fmt.Sscanf(token, "mul(%d,%d)", &n1, &n2)
			if shouldDo {
				total += n1 * n2
			}
		} else {
			shouldDo = token == "do()"
		}
	}

	fmt.Println(total)
}
