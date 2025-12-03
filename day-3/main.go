package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thulasipavankumar/Advent-of-code-2025/util"
)

const FILE_PATH = "input.txt"

func part1(lines []string) (count int) {

	for _, line := range lines {
		count += getHighest2Digit(line, 2)
	}

	return
}
func getHighest2Digit(line string, digits int) (high int) {

	for i := 9; i > 0; i-- {

		for j := 9; j > 0; j-- {

			idx := strings.Index(line, strconv.Itoa(i))

			if idx != -1 {
				result := line[idx+1:]
				if strings.Contains(result, strconv.Itoa(j)) {
					high = i*10 + j
					fmt.Printf("%d %s \n", high, line)
					return
				}

			}

		}

	}
	return
}
func main() {
	var temp = "236125212541321122222227172221217116322222226214257223424222222181222612515252282425452323242222662"
	println(getHighest2Digit(temp))

	println(part1(util.GetFileContent(FILE_PATH)))
}
