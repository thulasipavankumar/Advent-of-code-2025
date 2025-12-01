package main

import (
	"fmt"
	"strconv"

	"github.com/thulasipavankumar/Advent-of-code-2025/util"
)

const FILE_PATH = "input.txt"

func main() {
	var number int
	var dail int = 50
	var count int = 0
	for _, line := range util.GetFileContent(FILE_PATH) {
		letter := line[0]
		number, _ = strconv.Atoi(line[1:])
		number %= 100
		if letter == 'L' {
			number *= -1
		}
		dail += number
		dail %= 100
		if dail < 0 {
			dail += 100
		} else if dail > 99 {
			dail -= 100
		}
		if dail == 0 {
			count++
		}

		fmt.Println("dail: ", dail)

	}
	println(count)
}
