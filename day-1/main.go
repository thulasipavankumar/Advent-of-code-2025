package main

import (
	"strconv"

	"github.com/thulasipavankumar/Advent-of-code-2025/util"
)

const FILE_PATH = "input.txt"

func part2(lines []string) (count int) {
	var dail int = 50
	var mem int = 50

	for _, line := range lines {
		letter := line[0]
		number, _ := strconv.Atoi(line[1:])
		count += (number / 100)
		number %= 100
		println(line, number, count)
		if letter == 'L' {
			number *= -1
		}
		dail += number
		if dail < -99 {
			dail += 100
			count += (dail / 100)
			dail = (dail / 100) + 100

		} else if dail > 99 {
			dail = (dail / 100) - 100
			count += (dail / 100)
		}
	}
	count += part1(lines)
	return count

}
func part1(lines []string) (count int) {
	var number int
	var dail int = 50
	for _, line := range lines {
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

	}
	return

}
func main() {

	println(part2(util.GetFileContent(FILE_PATH)))
}
